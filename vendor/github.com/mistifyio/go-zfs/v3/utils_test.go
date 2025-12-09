package zfs

import (
	"errors"
	"os/exec"
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	for name, test := range map[string]struct {
		prop  string
		value string
		want  Zpool
	}{
		"some fragmentation": {
			prop:  "fragmentation",
			value: "15%",
			want:  Zpool{Fragmentation: 15},
		},
		"no fragmentation": {
			prop:  "fragmentation",
			value: "0",
			want:  Zpool{Fragmentation: 0},
		},
		"untracked fragmentation": {
			prop:  "fragmentation",
			value: "-",
			want:  Zpool{Fragmentation: 0},
		},
	} {
		t.Run(name, func(t *testing.T) {
			got := Zpool{}
			got.parseLine([]string{"", test.prop, test.value})
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("parse failure: wanted: %v, got: %v", test.want, got)
			}
		})
	}
}

func TestCommandError(t *testing.T) {
	cmd := &command{Command: "false"}
	expectedPath, err := exec.LookPath(cmd.Command)
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range []struct {
		name          string
		args          []string
		expectedDebug string
	}{
		{name: "NoArgs", expectedDebug: expectedPath},
		{name: "WithArgs", args: []string{"foo"}, expectedDebug: expectedPath + " foo"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := cmd.Run(tt.args...)
			if err == nil {
				t.Fatal("command.Run: wanted error, got nil")
			}
			var e *Error
			if !errors.As(err, &e) {
				t.Fatalf("command.Run (error): wanted *Error, got %T (%[1]v)", err)
			}
			if e.Debug != tt.expectedDebug {
				t.Fatalf("command.Run (error): wanted Debug %q, got %q", tt.expectedDebug, e.Debug)
			}
		})
	}
}
