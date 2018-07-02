package registry // import "github.com/docker/docker/registry"

import (
	"reflect"
	"sort"
	"strings"
	"testing"

	registrytypes "github.com/docker/docker/api/types/registry"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestLoadValidRegistries(t *testing.T) {
	var (
		secReg   registrytypes.Registry
		insecReg registrytypes.Registry
		config   *serviceConfig
		err      error
	)
	// secure with mirrors
	secReg, err = registrytypes.NewRegistry("https://secure.registry.com")
	secMirrors := []string{"https://secure.mirror1.com", "https://secure.mirror2.com"}
	if err != nil {
		t.Fatal(err)
	}
	if err := secReg.AddMirror(secMirrors[0]); err != nil {
		t.Fatal(err)
	}
	if err := secReg.AddMirror(secMirrors[1]); err != nil {
		t.Fatal(err)
	}

	// insecure without mirrors
	insecReg, err = registrytypes.NewRegistry("http://insecure.registry.com")
	if err != nil {
		t.Fatal(err)
	}

	// docker.io mirrors to test backwards compatibility
	officialMirrors := []string{"https://official.mirror1.com", "https://official.mirror2.com"}

	// create serciveConfig
	config, err = newServiceConfig(
		ServiceOptions{
			Mirrors:    officialMirrors,
			Registries: []registrytypes.Registry{secReg, insecReg},
		})
	if err != nil {
		t.Fatal(err)
	}

	// now test if the config looks as expected
	getMirrors := func(reg registrytypes.Registry) []string {
		mirrors := []string{}
		for _, mir := range reg.Mirrors {
			mirrors = append(mirrors, mir.URL.String())
		}
		return mirrors
	}

	if reg, loaded := config.Registries["secure.registry.com"]; !loaded {
		t.Fatalf("registry not loaded")
	} else {
		assert.Equal(t, true, reg.URL.IsSecure())
		assert.Equal(t, false, reg.URL.IsOfficial())
		mirrors := getMirrors(reg)
		assert.Equal(t, len(secMirrors), len(mirrors))
		sort.Strings(mirrors)
		sort.Strings(secMirrors)
		assert.Equal(t, secMirrors[0], mirrors[0])
		assert.Equal(t, secMirrors[1], mirrors[1])
	}

	if reg, loaded := config.Registries["insecure.registry.com"]; !loaded {
		t.Fatalf("registry not loaded")
	} else {
		assert.Equal(t, false, reg.URL.IsSecure())
		assert.Equal(t, false, reg.URL.IsOfficial())
		mirrors := getMirrors(reg)
		assert.Equal(t, 0, len(mirrors))
	}

	// backwards compatibility: "docker.io" will be loaded due to the config.Mirrors
	if reg, loaded := config.Registries["docker.io"]; !loaded {
		t.Fatalf("registry not loaded")
	} else {
		assert.Equal(t, true, reg.URL.IsSecure())
		assert.Equal(t, true, reg.URL.IsOfficial())
		mirrors := getMirrors(reg)
		assert.Equal(t, len(officialMirrors), len(mirrors))
		sort.Strings(mirrors)
		sort.Strings(officialMirrors)
		// append '/' (see ValidateMirror())
		assert.Equal(t, officialMirrors[0]+"/", mirrors[0])
		assert.Equal(t, officialMirrors[1]+"/", mirrors[1])
	}
}

//func TestLoadInvalidRegistries(t *testing.T) {
// XXX: this has to be tested manually as the v17.09.X doesn't have a proper
//	error handling for service configs (errors are silently ignored), so
//	the backported patch panics() instead.
//}

func TestFindRegistry(t *testing.T) {
	var (
		regA   registrytypes.Registry
		regB   registrytypes.Registry
		config *serviceConfig
		err    error
	)

	regA, err = registrytypes.NewRegistry("https://registry-a.com/my-prefix")
	if err != nil {
		t.Fatal(err)
	}

	regB, err = registrytypes.NewRegistry("http://registry-b.com")
	if err != nil {
		t.Fatal(err)
	}

	// create serciveConfig
	config, err = newServiceConfig(
		ServiceOptions{
			Registries: []registrytypes.Registry{regA, regB},
		})
	if err != nil {
		t.Fatal(err)
	}

	// no match -> nil
	reg := config.FindRegistry("foo")
	assert.Assert(t, is.Nil(reg))

	// prefix match -> registry
	reg = config.FindRegistry("registry-a.com/my-prefix/image:latest")
	assert.Assert(t, reg != nil)
	assert.Equal(t, "registry-a.com", reg.URL.Host())
	// no prefix match -> nil
	reg = config.FindRegistry("registry-a.com/not-my-prefix/image:42")
	assert.Assert(t, is.Nil(reg))

	// prefix match -> registry
	reg = config.FindRegistry("registry-b.com/image:latest")
	assert.Assert(t, reg != nil)
	assert.Equal(t, "registry-b.com", reg.URL.Host())
	// prefix match -> registry
	reg = config.FindRegistry("registry-b.com/also-in-namespaces/image:latest")
	assert.Assert(t, reg != nil)
	assert.Equal(t, "registry-b.com", reg.URL.Host())
}

func TestLoadAllowNondistributableArtifacts(t *testing.T) {
	testCases := []struct {
		registries []string
		cidrStrs   []string
		hostnames  []string
		err        string
	}{
		{
			registries: []string{"1.2.3.0/24"},
			cidrStrs:   []string{"1.2.3.0/24"},
		},
		{
			registries: []string{"2001:db8::/120"},
			cidrStrs:   []string{"2001:db8::/120"},
		},
		{
			registries: []string{"127.0.0.1"},
			hostnames:  []string{"127.0.0.1"},
		},
		{
			registries: []string{"127.0.0.1:8080"},
			hostnames:  []string{"127.0.0.1:8080"},
		},
		{
			registries: []string{"2001:db8::1"},
			hostnames:  []string{"2001:db8::1"},
		},
		{
			registries: []string{"[2001:db8::1]:80"},
			hostnames:  []string{"[2001:db8::1]:80"},
		},
		{
			registries: []string{"[2001:db8::1]:80"},
			hostnames:  []string{"[2001:db8::1]:80"},
		},
		{
			registries: []string{"1.2.3.0/24", "2001:db8::/120", "127.0.0.1", "127.0.0.1:8080"},
			cidrStrs:   []string{"1.2.3.0/24", "2001:db8::/120"},
			hostnames:  []string{"127.0.0.1", "127.0.0.1:8080"},
		},

		{
			registries: []string{"http://mytest.com"},
			err:        "allow-nondistributable-artifacts registry http://mytest.com should not contain '://'",
		},
		{
			registries: []string{"https://mytest.com"},
			err:        "allow-nondistributable-artifacts registry https://mytest.com should not contain '://'",
		},
		{
			registries: []string{"HTTP://mytest.com"},
			err:        "allow-nondistributable-artifacts registry HTTP://mytest.com should not contain '://'",
		},
		{
			registries: []string{"svn://mytest.com"},
			err:        "allow-nondistributable-artifacts registry svn://mytest.com should not contain '://'",
		},
		{
			registries: []string{"-invalid-registry"},
			err:        "Cannot begin or end with a hyphen",
		},
		{
			registries: []string{`mytest-.com`},
			err:        `allow-nondistributable-artifacts registry mytest-.com is not valid: invalid host "mytest-.com"`,
		},
		{
			registries: []string{`1200:0000:AB00:1234:0000:2552:7777:1313:8080`},
			err:        `allow-nondistributable-artifacts registry 1200:0000:AB00:1234:0000:2552:7777:1313:8080 is not valid: invalid host "1200:0000:AB00:1234:0000:2552:7777:1313:8080"`,
		},
		{
			registries: []string{`mytest.com:500000`},
			err:        `allow-nondistributable-artifacts registry mytest.com:500000 is not valid: invalid port "500000"`,
		},
		{
			registries: []string{`"mytest.com"`},
			err:        `allow-nondistributable-artifacts registry "mytest.com" is not valid: invalid host "\"mytest.com\""`,
		},
		{
			registries: []string{`"mytest.com:5000"`},
			err:        `allow-nondistributable-artifacts registry "mytest.com:5000" is not valid: invalid host "\"mytest.com"`,
		},
	}
	for _, testCase := range testCases {
		config := emptyServiceConfig
		err := config.LoadAllowNondistributableArtifacts(testCase.registries)
		if testCase.err == "" {
			if err != nil {
				t.Fatalf("expect no error, got '%s'", err)
			}

			var cidrStrs []string
			for _, c := range config.AllowNondistributableArtifactsCIDRs {
				cidrStrs = append(cidrStrs, c.String())
			}

			sort.Strings(testCase.cidrStrs)
			sort.Strings(cidrStrs)
			if (len(testCase.cidrStrs) > 0 || len(cidrStrs) > 0) && !reflect.DeepEqual(testCase.cidrStrs, cidrStrs) {
				t.Fatalf("expect AllowNondistributableArtifactsCIDRs to be '%+v', got '%+v'", testCase.cidrStrs, cidrStrs)
			}

			sort.Strings(testCase.hostnames)
			sort.Strings(config.AllowNondistributableArtifactsHostnames)
			if (len(testCase.hostnames) > 0 || len(config.AllowNondistributableArtifactsHostnames) > 0) && !reflect.DeepEqual(testCase.hostnames, config.AllowNondistributableArtifactsHostnames) {
				t.Fatalf("expect AllowNondistributableArtifactsHostnames to be '%+v', got '%+v'", testCase.hostnames, config.AllowNondistributableArtifactsHostnames)
			}
		} else {
			if err == nil {
				t.Fatalf("expect error '%s', got no error", testCase.err)
			}
			if !strings.Contains(err.Error(), testCase.err) {
				t.Fatalf("expect error '%s', got '%s'", testCase.err, err)
			}
		}
	}
}

func TestValidateMirror(t *testing.T) {
	valid := []string{
		"http://mirror-1.com",
		"http://mirror-1.com/",
		"https://mirror-1.com",
		"https://mirror-1.com/",
		"http://localhost",
		"https://localhost",
		"http://localhost:5000",
		"https://localhost:5000",
		"http://127.0.0.1",
		"https://127.0.0.1",
		"http://127.0.0.1:5000",
		"https://127.0.0.1:5000",
	}

	invalid := []string{
		"!invalid!://%as%",
		"ftp://mirror-1.com",
		"http://mirror-1.com/?q=foo",
		"http://mirror-1.com/v1/",
		"http://mirror-1.com/v1/?q=foo",
		"http://mirror-1.com/v1/?q=foo#frag",
		"http://mirror-1.com?q=foo",
		"https://mirror-1.com#frag",
		"https://mirror-1.com/#frag",
		"http://foo:bar@mirror-1.com/",
		"https://mirror-1.com/v1/",
		"https://mirror-1.com/v1/#",
		"https://mirror-1.com?q",
	}

	for _, address := range valid {
		if ret, err := ValidateMirror(address); err != nil || ret == "" {
			t.Errorf("ValidateMirror(`"+address+"`) got %s %s", ret, err)
		}
	}

	for _, address := range invalid {
		if ret, err := ValidateMirror(address); err == nil || ret != "" {
			t.Errorf("ValidateMirror(`"+address+"`) got %s %s", ret, err)
		}
	}
}

func TestLoadInsecureRegistries(t *testing.T) {
	testCases := []struct {
		registries []string
		index      string
		err        string
	}{
		{
			registries: []string{"127.0.0.1"},
			index:      "127.0.0.1",
		},
		{
			registries: []string{"127.0.0.1:8080"},
			index:      "127.0.0.1:8080",
		},
		{
			registries: []string{"2001:db8::1"},
			index:      "2001:db8::1",
		},
		{
			registries: []string{"[2001:db8::1]:80"},
			index:      "[2001:db8::1]:80",
		},
		{
			registries: []string{"http://mytest.com"},
			index:      "mytest.com",
		},
		{
			registries: []string{"https://mytest.com"},
			index:      "mytest.com",
		},
		{
			registries: []string{"HTTP://mytest.com"},
			index:      "mytest.com",
		},
		{
			registries: []string{"svn://mytest.com"},
			err:        "insecure registry svn://mytest.com should not contain '://'",
		},
		{
			registries: []string{"-invalid-registry"},
			err:        "Cannot begin or end with a hyphen",
		},
		{
			registries: []string{`mytest-.com`},
			err:        `insecure registry mytest-.com is not valid: invalid host "mytest-.com"`,
		},
		{
			registries: []string{`1200:0000:AB00:1234:0000:2552:7777:1313:8080`},
			err:        `insecure registry 1200:0000:AB00:1234:0000:2552:7777:1313:8080 is not valid: invalid host "1200:0000:AB00:1234:0000:2552:7777:1313:8080"`,
		},
		{
			registries: []string{`mytest.com:500000`},
			err:        `insecure registry mytest.com:500000 is not valid: invalid port "500000"`,
		},
		{
			registries: []string{`"mytest.com"`},
			err:        `insecure registry "mytest.com" is not valid: invalid host "\"mytest.com\""`,
		},
		{
			registries: []string{`"mytest.com:5000"`},
			err:        `insecure registry "mytest.com:5000" is not valid: invalid host "\"mytest.com"`,
		},
	}
	for _, testCase := range testCases {
		config := emptyServiceConfig
		err := config.LoadInsecureRegistries(testCase.registries)
		if testCase.err == "" {
			if err != nil {
				t.Fatalf("expect no error, got '%s'", err)
			}
			match := false
			for index := range config.IndexConfigs {
				if index == testCase.index {
					match = true
				}
			}
			if !match {
				t.Fatalf("expect index configs to contain '%s', got %+v", testCase.index, config.IndexConfigs)
			}
		} else {
			if err == nil {
				t.Fatalf("expect error '%s', got no error", testCase.err)
			}
			if !strings.Contains(err.Error(), testCase.err) {
				t.Fatalf("expect error '%s', got '%s'", testCase.err, err)
			}
		}
	}
}

func TestNewServiceConfig(t *testing.T) {
	testCases := []struct {
		opts   ServiceOptions
		errStr string
	}{
		{
			ServiceOptions{},
			"",
		},
		{
			ServiceOptions{
				Mirrors: []string{"example.com:5000"},
			},
			`invalid mirror: unsupported scheme "example.com" in "example.com:5000"`,
		},
		{
			ServiceOptions{
				Mirrors: []string{"http://example.com:5000"},
			},
			"",
		},
		{
			ServiceOptions{
				InsecureRegistries: []string{"[fe80::]/64"},
			},
			`insecure registry [fe80::]/64 is not valid: invalid host "[fe80::]/64"`,
		},
		{
			ServiceOptions{
				InsecureRegistries: []string{"102.10.8.1/24"},
			},
			"",
		},
		{
			ServiceOptions{
				AllowNondistributableArtifacts: []string{"[fe80::]/64"},
			},
			`allow-nondistributable-artifacts registry [fe80::]/64 is not valid: invalid host "[fe80::]/64"`,
		},
		{
			ServiceOptions{
				AllowNondistributableArtifacts: []string{"102.10.8.1/24"},
			},
			"",
		},
	}

	for _, testCase := range testCases {
		_, err := newServiceConfig(testCase.opts)
		if testCase.errStr != "" {
			assert.Check(t, is.Error(err, testCase.errStr))
		} else {
			assert.Check(t, err)
		}
	}
}

func TestValidateIndexName(t *testing.T) {
	valid := []struct {
		index  string
		expect string
	}{
		{
			index:  "index.docker.io",
			expect: "docker.io",
		},
		{
			index:  "example.com",
			expect: "example.com",
		},
		{
			index:  "127.0.0.1:8080",
			expect: "127.0.0.1:8080",
		},
		{
			index:  "mytest-1.com",
			expect: "mytest-1.com",
		},
		{
			index:  "mirror-1.com/v1/?q=foo",
			expect: "mirror-1.com/v1/?q=foo",
		},
	}

	for _, testCase := range valid {
		result, err := ValidateIndexName(testCase.index)
		if assert.Check(t, err) {
			assert.Check(t, is.Equal(testCase.expect, result))
		}

	}
}

func TestValidateIndexNameWithError(t *testing.T) {
	invalid := []struct {
		index string
		err   string
	}{
		{
			index: "docker.io-",
			err:   "invalid index name (docker.io-). Cannot begin or end with a hyphen",
		},
		{
			index: "-example.com",
			err:   "invalid index name (-example.com). Cannot begin or end with a hyphen",
		},
		{
			index: "mirror-1.com/v1/?q=foo-",
			err:   "invalid index name (mirror-1.com/v1/?q=foo-). Cannot begin or end with a hyphen",
		},
	}
	for _, testCase := range invalid {
		_, err := ValidateIndexName(testCase.index)
		assert.Check(t, is.Error(err, testCase.err))
	}
}
