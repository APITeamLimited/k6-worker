#!/bin/bash

# Copyright 2019 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e
shopt -s nullglob

knownfolders="$(printf '%s\n' "/mnt/c/Program Files (x86)/Windows Kits/"/*/Include/*/um/KnownFolders.h | sort -Vr | head -n 1)"
[[ -n $knownfolders ]] || ***REMOVED*** echo "Unable to find KnownFolders.h" >&2; exit 1; ***REMOVED***

***REMOVED***
	echo "// Code generated by 'mkknownfolderids.bash'; DO NOT EDIT."
	echo
	echo "package windows"
	echo "type KNOWNFOLDERID GUID"
	echo "var ("
	while read -r line; do
		[[ $line =~ DEFINE_KNOWN_FOLDER\((FOLDERID_[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+),[\t\ ]*(0x[^,]+)\) ]] || continue
		printf "%s = &KNOWNFOLDERID***REMOVED***0x%08x, 0x%04x, 0x%04x, [8]byte***REMOVED***0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x, 0x%02x***REMOVED******REMOVED***\n" \
			"$***REMOVED***BASH_REMATCH[1]***REMOVED***" $(( "$***REMOVED***BASH_REMATCH[2]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[3]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[4]***REMOVED***" )) \
			$(( "$***REMOVED***BASH_REMATCH[5]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[6]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[7]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[8]***REMOVED***" )) \
			$(( "$***REMOVED***BASH_REMATCH[9]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[10]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[11]***REMOVED***" )) $(( "$***REMOVED***BASH_REMATCH[12]***REMOVED***" ))
	done < "$knownfolders"
	echo ")"
***REMOVED*** | gofmt > "zknownfolderids_windows.go"
