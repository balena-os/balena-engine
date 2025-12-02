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

// This fake package is created as golang.org/x/tools/go/analysis/analysistest
// expects it to be here for loading. This package is used to test allow-unstructured
// flag which suppresses errors when unstructured logging is used.
// This is a test file for unstructured logging static check tool unit tests.

package allowUnstructuredLogs

import (
	klog "k8s.io/klog/v2"
)

func allowUnstructuredLogs() {
	// Structured logs
	// Error is expected if structured logging pattern is not used correctly
	klog.InfoS("test log")
	klog.ErrorS(nil, "test log")
	klog.InfoS("Starting container in a pod", "containerID", "containerID", "pod")                // want `Additional arguments to InfoS should always be Key Value pairs. Please check if there is any key or value missing.`
	klog.ErrorS(nil, "Starting container in a pod", "containerID", "containerID", "pod")          // want `Additional arguments to ErrorS should always be Key Value pairs. Please check if there is any key or value missing.`
	klog.InfoS("Starting container in a pod", "测试", "containerID")                                // want `Key positional arguments "测试" are expected to be lowerCamelCase alphanumeric strings. Please remove any non-Latin characters.`
	klog.ErrorS(nil, "Starting container in a pod", "测试", "containerID")                          // want `Key positional arguments "测试" are expected to be lowerCamelCase alphanumeric strings. Please remove any non-Latin characters.`
	klog.InfoS("Starting container in a pod", 7, "containerID")                                   // want `Key positional arguments are expected to be inlined constant strings. Please replace 7 provided with string value`
	klog.ErrorS(nil, "Starting container in a pod", 7, "containerID")                             // want `Key positional arguments are expected to be inlined constant strings. Please replace 7 provided with string value`
	klog.InfoS("Starting container in a pod", map[string]string{"test1": "value"}, "containerID") // want `Key positional arguments are expected to be inlined constant strings. `
	testKey := "a"
	klog.ErrorS(nil, "Starting container in a pod", testKey, "containerID") // want `Key positional arguments are expected to be inlined constant strings. `
	klog.InfoS("test: %s", "testname")                                      // want `structured logging function "InfoS" should not use format specifier "%s"`
	klog.ErrorS(nil, "test no.: %d", 1)                                     // want `structured logging function "ErrorS" should not use format specifier "%d"`

	// Unstructured logs
	// Error is not expected as this package allows unstructured logging
	klog.V(1).Infof("test log")
	klog.Infof("test log")
	klog.Info("test log")
	klog.Infoln("test log")
	klog.InfoDepth(1, "test log")
	klog.Warning("test log")
	klog.Warningf("test log")
	klog.WarningDepth(1, "test log")
	klog.Error("test log")
	klog.Errorf("test log")
	klog.Errorln("test log")
	klog.ErrorDepth(1, "test log")
	klog.Fatal("test log")
	klog.Fatalf("test log")
	klog.Fatalln("test log")
	klog.FatalDepth(1, "test log")
	klog.Exit("test log")
	klog.ExitDepth(1, "test log")
	klog.Exitln("test log")
	klog.Exitf("test log")
}
