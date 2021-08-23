//go:build !windows
// +build !windows

package main

import "time"

var defaultReconciliationTimeout = 30 * time.Second
