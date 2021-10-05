// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix && ppc
// +build aix,ppc

// Functions to access/create device major and minor numbers matching the
// encoding used by AIX.

package unix

// Major returns the major component of a Linux device number.
func Major(dev uint64) uint32 ***REMOVED***
	return uint32((dev >> 16) & 0xffff)
***REMOVED***

// Minor returns the minor component of a Linux device number.
func Minor(dev uint64) uint32 ***REMOVED***
	return uint32(dev & 0xffff)
***REMOVED***

// Mkdev returns a Linux device number generated from the given major and minor
// components.
func Mkdev(major, minor uint32) uint64 ***REMOVED***
	return uint64(((major) << 16) | (minor))
***REMOVED***
