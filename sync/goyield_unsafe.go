// Copyright 2020 The gVisor Authors.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.14 && !go1.19
// +build go1.14,!go1.19

// Check go:linkname function signatures when updating Go version.

package sync

import (
	_ "unsafe" // for go:linkname
)

//go:linkname goyield runtime.goyield
func goyield()
