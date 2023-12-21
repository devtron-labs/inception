package providers

import (
	"github.com/devtron-labs/inception/pkg/providerIdentifier/bean"
	"github.com/go-logr/logr"
	"net/http"
	"os"
	"strings"
)

type IdentifyGoogle struct {
	Log logr.Logger
}

func (impl *IdentifyGoogle) Identify() (string, error) {
	data, err := os.ReadFile(bean.GoogleSysFile)
	if err != nil {
		impl.Log.Error(err, "error while reading file", "error", err)
		return bean.Unknown, err
	}
	if strings.Contains(string(data), bean.GoogleIdentifierString) {
		return bean.Google, nil
	}
	return bean.Unknown, nil
}

func (impl *IdentifyGoogle) IdentifyViaMetadataServer(detected chan<- string) {
	req, err := http.NewRequest("GET", bean.GoogleMetadataServer, nil)
	if err != nil {
		impl.Log.Error(err, "error while creating new request", "error", err)
		detected <- bean.Unknown
		return
	}
	req.Header.Set("Metadata-Flavor", "Google")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		impl.Log.Error(err, "error while requesting", "error", err, "request", req)
		detected <- bean.Unknown
		return
	}
	if resp.StatusCode == http.StatusOK {
		detected <- bean.Google
	} else {
		detected <- bean.Unknown
		return
	}
}
