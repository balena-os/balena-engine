package registry // import "github.com/docker/docker/registry"

import (
	"fmt"
	"net/url"
	"strings"

	registrytypes "github.com/docker/docker/api/types/registry"
	"github.com/docker/go-connections/tlsconfig"
)

func (s *DefaultService) lookupV2Endpoints(reference string) (endpoints []APIEndpoint, err error) {
	tlsConfig := tlsconfig.ServerDefault()

	// extraxt the hostname from the reference
	refURL := reference
	if !strings.HasPrefix(refURL, "http://") && !strings.HasPrefix(refURL, "https://") {
		refURL = "https://" + refURL
	}
	u, err := url.Parse(refURL)
	if err != nil {
		return nil, fmt.Errorf("SUSE PATCH [lookupV2Endpoints]: error parsing reference %s: %s", reference, err)
	}
	hostname := u.Host // hostname + port (if present)
	if hostname == "" {
		return nil, fmt.Errorf("SUSE PATCH [lookupV2Endpoints]: cannot determine hostname of reference %s", reference)
	}

	// create endpoints for official and configured registries
	official := false
	if hostname == "docker.io" {
		official = true
	}
	reg := s.config.FindRegistry(reference)

	if reg != nil || official {
		if reg == nil {
			reg = &registrytypes.Registry{}
		}
		// if present, add mirrors prior to the registry
		for _, mirror := range reg.Mirrors {
			mURL := mirror.URL.URL()
			mirrorTLSConfig, err := s.tlsConfigForMirror(&mURL)
			if err != nil {
				return nil, fmt.Errorf("SUSE PATCH [lookupV2Endpoints]: %s", err)
			}
			endpoints = append(endpoints, APIEndpoint{
				URL:          &mURL,
				Version:      APIVersion2,
				Mirror:       true,
				TrimHostname: true,
				TLSConfig:    mirrorTLSConfig,
			})
		}
		// add the registry
		var endpointURL *url.URL
		if official {
			endpointURL = DefaultV2Registry
		} else {
			endpointURL = &url.URL{
				Scheme: reg.URL.Scheme(),
				Host:   reg.URL.Host(),
			}
		}
		endpoints = append(endpoints, APIEndpoint{
			URL:          endpointURL,
			Version:      APIVersion2,
			Official:     official,
			TrimHostname: true,
			TLSConfig:    tlsConfig,
		})

		return endpoints, nil
	}

	ana := allowNondistributableArtifacts(s.config, hostname)

	tlsConfig, err = s.tlsConfig(hostname)
	if err != nil {
		return nil, fmt.Errorf("SUSE PATCH [lookupV2Enpoints]: %s", err)
	}

	endpoints = []APIEndpoint{
		{
			URL: &url.URL{
				Scheme: "https",
				Host:   hostname,
			},
			Version:                        APIVersion2,
			AllowNondistributableArtifacts: ana,
			TrimHostname:                   true,
			TLSConfig:                      tlsConfig,
		},
	}

	if tlsConfig.InsecureSkipVerify {
		endpoints = append(endpoints, APIEndpoint{
			URL: &url.URL{
				Scheme: "http",
				Host:   hostname,
			},
			Version:                        APIVersion2,
			AllowNondistributableArtifacts: ana,
			TrimHostname:                   true,
			// used to check if supposed to be secure via InsecureSkipVerify
			TLSConfig: tlsConfig,
		})
	}

	return endpoints, nil
}
