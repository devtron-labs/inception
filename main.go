/*
Copyright 2020 Devtron Labs Pvt Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/base64"
	"encoding/json"
	"flag"
	"github.com/devtron-labs/inception/pkg/language"
	"github.com/posthog/posthog-go"
	"io/ioutil"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	"net/http"
	"os"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"time"

	installerv1alpha1 "github.com/devtron-labs/inception/api/v1alpha1"
	"github.com/devtron-labs/inception/controllers"
	// +kubebuilder:scaffold:imports
	"github.com/patrickmn/go-cache"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	utilruntime.Must(clientgoscheme.AddToScheme(scheme))

	utilruntime.Must(installerv1alpha1.AddToScheme(scheme))
	// +kubebuilder:scaffold:scheme
}

var (
	PosthogApiKey           string = ""
	PosthogEndpoint         string = "https://app.posthog.com"
	CacheExpiry             int    = 720
	TelemetryApiKeyEndpoint string = "aHR0cHM6Ly90ZWxlbWV0cnkuZGV2dHJvbi5haS9kZXZ0cm9uL3RlbGVtZXRyeS9hcGlrZXk="
)

func main() {
	var metricsAddr string
	var enableLeaderElection bool
	flag.StringVar(&metricsAddr, "metrics-addr", ":8080", "The address the metric endpoint binds to.")
	flag.BoolVar(&enableLeaderElection, "enable-leader-election", false,
		"Enable leader election for controller manager. "+
			"Enabling this will ensure there is only one active controller manager.")
	flag.Parse()

	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: metricsAddr,
		Port:               9443,
		LeaderElection:     enableLeaderElection,
		LeaderElectionID:   "2b3710c5.devtron.ai",
	})
	if err != nil {
		setupLog.Error(err, "unable to start manager")
		os.Exit(1)
	}

	if len(PosthogApiKey) == 0 {
		_, apiKey, err := getPosthogApiKey(TelemetryApiKeyEndpoint)
		if err != nil {
			setupLog.Error(err, "exception caught while getting api key")
		}
		PosthogApiKey = apiKey
	}
	client, _ := posthog.NewWithConfig(PosthogApiKey, posthog.Config{Endpoint: PosthogEndpoint})
	//defer client.Close()
	d := time.Duration(CacheExpiry)
	c := cache.New(d*time.Minute, 1440*time.Minute)

	if err = (&controllers.InstallerReconciler{
		Client:        mgr.GetClient(),
		Log:           ctrl.Log.WithName("controllers").WithName("Installer"),
		Scheme:        mgr.GetScheme(),
		Mapper:        language.NewMapperFactory(),
		PosthogClient: client,
		Cache:         c,
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "unable to create controller", "controller", "Installer")
		os.Exit(1)
	}
	// +kubebuilder:scaffold:builder

	setupLog.Info("starting manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "problem running manager")
		os.Exit(1)
	}
}

func getPosthogApiKey(encodedPosthogApiKeyUrl string) (string, string, error) {
	dncodedPosthogApiKeyUrl, err := base64.StdEncoding.DecodeString(encodedPosthogApiKeyUrl)
	if err != nil {
		return "", "", err
	}
	apiKeyUrl := string(dncodedPosthogApiKeyUrl)
	req, err := http.NewRequest(http.MethodGet, apiKeyUrl, nil)
	if err != nil {
		return "", "", err
	}
	//var client *http.Client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", "", err
	}
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", "", err
		}
		var apiRes map[string]interface{}
		err = json.Unmarshal(resBody, &apiRes)
		if err != nil {
			return "", "", err
		}
		encodedApiKey := apiRes["result"].(string)
		apiKey, err := base64.StdEncoding.DecodeString(encodedApiKey)
		if err != nil {
			return "", "", err
		}
		return encodedApiKey, string(apiKey), err
	}
	return "", "", err
}
