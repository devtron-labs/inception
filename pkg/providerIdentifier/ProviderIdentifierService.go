package providerIdentifier

import (
	"github.com/devtron-labs/inception/pkg/providerIdentifier/bean"
	"github.com/devtron-labs/inception/pkg/providerIdentifier/providers"
	"github.com/go-logr/logr"
)

type Identifier interface {
	Identify() (string, error)
	IdentifyViaMetadataServer(detected chan<- string)
}

type ProviderIdentifierService interface {
	IdentifyProvider() (string, error)
}

type ProviderIdentifierServiceImpl struct {
	Log logr.Logger
}

func NewProviderIdentifierServiceImpl(
	Log logr.Logger) *ProviderIdentifierServiceImpl {
	providerIdentifierService := &ProviderIdentifierServiceImpl{
		Log: Log,
	}
	return providerIdentifierService
}

func (impl *ProviderIdentifierServiceImpl) IdentifyProvider() (string, error) {
	identifiers := []Identifier{
		&providers.IdentifyAlibaba{Log: impl.Log},
		&providers.IdentifyAmazon{Log: impl.Log},
		&providers.IdentifyAzure{Log: impl.Log},
		&providers.IdentifyDigitalOcean{Log: impl.Log},
		&providers.IdentifyGoogle{Log: impl.Log},
		&providers.IdentifyOracle{Log: impl.Log},
	}

	identifiedProv := bean.Unknown
	var err error
	for _, prov := range identifiers {
		identifiedProv, err = prov.Identify()
		if err != nil {
			continue
		}
		if identifiedProv != bean.Unknown {
			return identifiedProv, nil
		}
	}

	detected := make(chan string)
	defer close(detected)

	provs := make([]func(chan<- string), 0)
	for _, prov := range identifiers {
		provs = append(provs, prov.IdentifyViaMetadataServer)
	}

	for _, functions := range provs {
		go functions(detected)
	}
	for range provs {
		d := <-detected
		if d != bean.Unknown {
			return d, nil
		}
	}
	return bean.Unknown, nil
}
