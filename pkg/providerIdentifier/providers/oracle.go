package providers

import (
	"encoding/json"
	"github.com/devtron-labs/inception/pkg/providerIdentifier/bean"
	"github.com/go-logr/logr"
	"io/ioutil"
	"net/http"
	"strings"
)

type oracleMetadataResponse struct {
	OkeTM string `json:"oke-tm"`
}

type IdentifyOracle struct {
	Log logr.Logger
}

func (impl *IdentifyOracle) Identify() (string, error) {
	data, err := ioutil.ReadFile(bean.OracleSysFile)
	if err != nil {
		impl.Log.Error(err, "error while reading file", "error", err)
		return bean.Unknown, err
	}
	if strings.Contains(string(data), bean.OracleIdentifierString) {
		return bean.Oracle, nil
	}
	return bean.Unknown, nil
}

func (impl *IdentifyOracle) IdentifyViaMetadataServer(detected chan<- string) {
	r := oracleMetadataResponse{}
	req, err := http.NewRequest("GET", bean.OracleMetadataServer, nil)
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
		err = json.Unmarshal(body, &r)
		if err != nil {
			impl.Log.Error(err, "error while unmarshaling json", "error", err, "body", body)
			detected <- bean.Unknown
			return
		}
		if strings.Contains(r.OkeTM, "oke") {
			detected <- bean.Oracle
		}
	} else {
		detected <- bean.Unknown
		return
	}
}
