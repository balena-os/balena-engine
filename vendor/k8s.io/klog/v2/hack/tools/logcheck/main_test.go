/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	tests := []struct {
		name              string
		allowUnstructured string
		testPackage       string
	}{
		{
			name:              "Allow unstructured logs",
			allowUnstructured: "true",
			testPackage:       "allowUnstructuredLogs",
		},
		{
			name:              "Do not allow unstructured logs",
			allowUnstructured: "false",
			testPackage:       "doNotAllowUnstructuredLogs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			analyzer := analyser()
			analyzer.Flags.Set("allow-unstructured", tt.allowUnstructured)
			analysistest.Run(t, analysistest.TestData(), analyzer, tt.testPackage)
		})
	}
}
