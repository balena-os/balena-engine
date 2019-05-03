package storagemigration

import "testing"

func TestIsWhiteout(t *testing.T) {
	var tcs = map[string]bool{
		".wh.foo.txt":  true,
		"bar.txt":      false,
		".wh..wh.plnk": true,
		".wh..wh..opq": true,
	}
	for file, expect := range tcs {
		if IsWhiteout(file) != expect {
			t.Fatalf("did not detect %v", file)
		}
	}
}
func TestIsWhiteoutMeta(t *testing.T) {
	var tcs = map[string]bool{
		".wh.foo.txt":  false,
		"bar.txt":      false,
		".wh..wh.plnk": true,
		".wh..wh..opq": true,
	}
	for file, expect := range tcs {
		if IsWhiteoutMeta(file) != expect {
			t.Fatalf("did not detect %v", file)
		}
	}
}

func TestIsOpaque(t *testing.T) {
	var tcs = map[string]bool{
		".wh.foo.txt":  false,
		"bar.txt":      false,
		".wh..wh.plnk": false,
		".wh..wh..opq": true,
	}
	for file, expect := range tcs {
		if IsOpaqueParentDir(file) != expect {
			t.Fatalf("did not detect %v", file)
		}
	}
}

func TestStripWhiteoutPrefix(t *testing.T) {
	var tcs = map[string]string{
		".wh.foo.txt":  "foo.txt",
		"bar.txt":      "bar.txt",
		".wh..wh.plnk": ".wh..wh.plnk",
		".wh..wh..opq": ".wh..wh..opq",
	}
	for file, expect := range tcs {
		if result := StripWhiteoutPrefix(file); result != expect {
			t.Fatalf("stripping filename failed, got: %q wanted: %q", result, expect)
		}
	}
}
