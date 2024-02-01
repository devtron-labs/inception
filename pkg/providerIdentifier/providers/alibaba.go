package providers

import (
	"github.com/devtron-labs/inception/pkg/providerIdentifier/bean"
	"github.com/go-logr/logr"
	"io/ioutil"
	"net/http"
	"strings"
)

type IdentifyAlibaba struct {
	Log logr.Logger
}

func (impl *IdentifyAlibaba) Identify() (string, error) {
	data, err := ioutil.ReadFile(bean.AlibabaSysFile)
	if err != nil {
		impl.Log.Error(err, "error while reading file", "error", err)
		return bean.Unknown, err
	}
	if strings.Contains(string(data), bean.AlibabaIdentifierString) {
		return bean.Alibaba, nil
	}
	return bean.Unknown, nil
}

func (impl *IdentifyAlibaba) IdentifyViaMetadataServer(detected chan<- string) {
	req, err := http.NewRequest("GET", bean.AlibabaMetadataServer, nil)
	if err != nil {
		impl.Log.Error(err, "error while creating new request", "error", err)
		detected <- bean.Unknown
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		impl.Log.Error(err, "error while requesting", "error", err, "request", req)
		detected <- bean.Unknown
		return
	}
	if resp.StatusCode == http.StatusOK {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			impl.Log.Error(err, "error while reading response body", "error", err, "respBody", resp.Body)
			detected <- bean.Unknown
			return
		}
		if strings.HasPrefix(string(body), "ecs.") {
			detected <- bean.Alibaba
			return
		}
	} else {
		detected <- bean.Unknown
		return
	}
}
