module github.com/devtron-labs/inception

go 1.13

require (
	github.com/antlr/antlr4 v0.0.0-20200820155224-be881fa6b91d
	github.com/argoproj/gitops-engine v0.3.2
	github.com/ghodss/yaml v1.0.0
	github.com/go-logr/logr v0.3.0
	github.com/go-logr/zapr v0.4.0 // indirect
	github.com/google/go-cmp v0.5.9
	github.com/onsi/ginkgo v1.12.1
	github.com/onsi/gomega v1.10.1
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/posthog/posthog-go v0.0.0-20210616140600-8bebac98515e
	github.com/sirupsen/logrus v1.6.0
	github.com/tidwall/gjson v1.9.3
	github.com/tidwall/sjson v1.1.1
	google.golang.org/grpc v1.56.3
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v1.5.2
	k8s.io/klog v1.0.0
	k8s.io/kubectl v0.20.7
	k8s.io/utils v0.0.0-20201110183641-67b214c5f920
	sigs.k8s.io/controller-runtime v0.6.2
	sigs.k8s.io/yaml v1.2.0
)

replace (
	github.com/go-logr/zapr => github.com/go-logr/zapr v0.3.0
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20201216223049-8b5274cf687f
	k8s.io/api => k8s.io/api v0.20.7
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.7
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.7
	k8s.io/apiserver => k8s.io/apiserver v0.20.7
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.7
	k8s.io/client-go => k8s.io/client-go v0.20.7
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.7
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.7
	k8s.io/code-generator => k8s.io/code-generator v0.20.7
	k8s.io/component-base => k8s.io/component-base v0.20.7
	k8s.io/component-helpers => k8s.io/component-helpers v0.20.7
	k8s.io/controller-manager => k8s.io/controller-manager v0.20.7
	k8s.io/cri-api => k8s.io/cri-api v0.20.7
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.7
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.7
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.7
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.7
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.7
	k8s.io/kubectl => k8s.io/kubectl v0.20.7
	k8s.io/kubelet => k8s.io/kubelet v0.20.7
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.7
	k8s.io/metrics => k8s.io/metrics v0.20.7
	k8s.io/mount-utils => k8s.io/mount-utils v0.20.7
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.8
//github.com/devtron-labs/inception/pkg => /Users/nishant/go/src/github.com/devtron-labs/inception/pkg
)
