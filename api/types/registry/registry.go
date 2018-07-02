package registry // import "github.com/docker/docker/api/types/registry"

import (
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/opencontainers/image-spec/specs-go/v1"
)

// ServiceConfig stores daemon registry services configuration.
type ServiceConfig struct {
	AllowNondistributableArtifactsCIDRs     []*NetIPNet
	AllowNondistributableArtifactsHostnames []string
	InsecureRegistryCIDRs                   []*NetIPNet           `json:"InsecureRegistryCIDRs"`
	IndexConfigs                            map[string]*IndexInfo `json:"IndexConfigs"`
	Mirrors                                 []string
	Registries                              map[string]Registry
}

// Registry holds information for a registry and its mirrors.
type Registry struct {
	// Prefix is used for the lookup of endpoints, where the given registry
	// is selected when its Prefix is a prefix of the passed reference, for
	// instance, Prefix:"docker.io/opensuse" will match a `docker pull
	// opensuse:tumleweed`.
	URL RegURL `json:"Prefix"`
	// The mirrors will be selected prior to the registry during lookup of
	// endpoints.
	Mirrors []Mirror `json:"Mirrors,omitempty"`
}

// NewRegistry returns a Registry and interprets input as a URL.
func NewRegistry(input string) (Registry, error) {
	reg := Registry{}
	err := reg.URL.Parse(input)
	return reg, err
}

// AddMirror interprets input as a URL and adds it as a new mirror.
func (r *Registry) AddMirror(input string) error {
	mir, err := NewMirror(input)
	if err != nil {
		return err
	}
	r.Mirrors = append(r.Mirrors, mir)
	return nil
}

// ContainsMirror returns true if the URL of any mirror equals input.
func (r *Registry) ContainsMirror(input string) bool {
	for _, m := range r.Mirrors {
		if m.URL.String() == input {
			return true
		}
	}
	return false
}

// Mirror holds information for a given registry mirror.
type Mirror struct {
	// The URL of the mirror.
	URL RegURL `json:"URL,omitempty"`
}

// NewMirror returns a Registry and interprets input as a URL.
func NewMirror(input string) (Mirror, error) {
	mir := Mirror{}
	err := mir.URL.Parse(input)
	return mir, err
}

// RegURL is a wrapper for url.URL to unmarshal it from the JSON config and to
// make it an embedded type for its users.
type RegURL struct {
	// rURL is a simple url.URL.  Notice it is no pointer to avoid potential
	// null pointer dereferences.
	rURL url.URL
}

// UnmarshalJSON unmarshals the byte array into the RegURL pointer.
func (r *RegURL) UnmarshalJSON(b []byte) error {
	var input string
	if err := json.Unmarshal(b, &input); err != nil {
		return err
	}
	return r.Parse(input)
}

// MarshalJSON marshals the RegURL.
func (r *RegURL) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// Parse parses input as a URL.
func (r *RegURL) Parse(input string) error {
	input = strings.ToLower(input)
	uri, err := url.Parse(input)
	if err == nil {
		r.rURL = *uri
	} else {
		return err
	}
	// default to https if no URI scheme is specified
	if uri.Scheme == "" {
		// we have to parse again to update all associated data
		return r.Parse("https://" + input)
	}

	// sanity checks
	if uri.Scheme != "http" && uri.Scheme != "https" {
		return fmt.Errorf("invalid url: unsupported scheme %q in %q", uri.Scheme, uri)
	}
	if uri.Host == "" {
		return fmt.Errorf("invalid url: unspecified hostname in %s", uri)
	}
	if uri.User != nil {
		// strip password from output
		uri.User = url.UserPassword(uri.User.Username(), "xxxxx")
		return fmt.Errorf("invalid url: username/password not allowed in URI %q", uri)
	}

	return nil
}

// Host returns the host:port of the URL.
func (r *RegURL) Host() string {
	return r.rURL.Host
}

// Prefix returns the host:port/path of the URL.
func (r *RegURL) Prefix() string {
	return r.rURL.Host + r.rURL.Path
}

// IsOfficial returns true if the URL points to an official "docker.io" host.
func (r *RegURL) IsOfficial() bool {
	return r.rURL.Hostname() == "docker.io"
}

// IsSecure returns true if the URI scheme of the URL is "https".
func (r *RegURL) IsSecure() bool {
	return r.Scheme() == "https"
}

// Scheme returns the URI scheme.
func (r *RegURL) Scheme() string {
	return r.rURL.Scheme
}

// URL return URL of the RegURL.
func (r *RegURL) URL() url.URL {
	return r.rURL
}

// String return URL as a string.
func (r *RegURL) String() string {
	return r.rURL.String()
}

// NetIPNet is the net.IPNet type, which can be marshalled and
// unmarshalled to JSON
type NetIPNet net.IPNet

// String returns the CIDR notation of ipnet
func (ipnet *NetIPNet) String() string {
	return (*net.IPNet)(ipnet).String()
}

// MarshalJSON returns the JSON representation of the IPNet
func (ipnet *NetIPNet) MarshalJSON() ([]byte, error) {
	return json.Marshal((*net.IPNet)(ipnet).String())
}

// UnmarshalJSON sets the IPNet from a byte array of JSON
func (ipnet *NetIPNet) UnmarshalJSON(b []byte) (err error) {
	var ipnetStr string
	if err = json.Unmarshal(b, &ipnetStr); err == nil {
		var cidr *net.IPNet
		if _, cidr, err = net.ParseCIDR(ipnetStr); err == nil {
			*ipnet = NetIPNet(*cidr)
		}
	}
	return
}

// IndexInfo contains information about a registry
//
// RepositoryInfo Examples:
// {
//   "Index" : {
//     "Name" : "docker.io",
//     "Mirrors" : ["https://registry-2.docker.io/v1/", "https://registry-3.docker.io/v1/"],
//     "Secure" : true,
//     "Official" : true,
//   },
//   "RemoteName" : "library/debian",
//   "LocalName" : "debian",
//   "CanonicalName" : "docker.io/debian"
//   "Official" : true,
// }
//
// {
//   "Index" : {
//     "Name" : "127.0.0.1:5000",
//     "Mirrors" : [],
//     "Secure" : false,
//     "Official" : false,
//   },
//   "RemoteName" : "user/repo",
//   "LocalName" : "127.0.0.1:5000/user/repo",
//   "CanonicalName" : "127.0.0.1:5000/user/repo",
//   "Official" : false,
// }
type IndexInfo struct {
	// Name is the name of the registry, such as "docker.io"
	Name string
	// Mirrors is a list of mirrors, expressed as URIs
	Mirrors []string
	// Secure is set to false if the registry is part of the list of
	// insecure registries. Insecure registries accept HTTP and/or accept
	// HTTPS with certificates from unknown CAs.
	Secure bool
	// Official indicates whether this is an official registry
	Official bool
}

// SearchResult describes a search result returned from a registry
type SearchResult struct {
	// StarCount indicates the number of stars this repository has
	StarCount int `json:"star_count"`
	// IsOfficial is true if the result is from an official repository.
	IsOfficial bool `json:"is_official"`
	// Name is the name of the repository
	Name string `json:"name"`
	// IsAutomated indicates whether the result is automated
	IsAutomated bool `json:"is_automated"`
	// Description is a textual description of the repository
	Description string `json:"description"`
}

// SearchResults lists a collection search results returned from a registry
type SearchResults struct {
	// Query contains the query string that generated the search results
	Query string `json:"query"`
	// NumResults indicates the number of results the query returned
	NumResults int `json:"num_results"`
	// Results is a slice containing the actual results for the search
	Results []SearchResult `json:"results"`
}

// DistributionInspect describes the result obtained from contacting the
// registry to retrieve image metadata
type DistributionInspect struct {
	// Descriptor contains information about the manifest, including
	// the content addressable digest
	Descriptor v1.Descriptor
	// Platforms contains the list of platforms supported by the image,
	// obtained by parsing the manifest
	Platforms []v1.Platform
}
