/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2017 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package encoding

import (
	"context"
	"testing"

	"github.com/dop251/goja"
	"github.com/loadimpact/k6/js/common"
	"github.com/stretchr/testify/assert"
)

func TestEncodingAlgorithms(t *testing.T) ***REMOVED***
	if testing.Short() ***REMOVED***
		return
	***REMOVED***

	rt := goja.New()
	rt.SetFieldNameMapper(common.FieldNameMapper***REMOVED******REMOVED***)
	ctx := context.Background()
	ctx = common.WithRuntime(ctx, rt)
	rt.Set("encoding", common.Bind(rt, New(), &ctx))

	t.Run("Base64", func(t *testing.T) ***REMOVED***
		t.Run("DefaultEnc", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "aGVsbG8gd29ybGQ=";
			let encoded = encoding.b64encode("hello world");
			if (encoded !== correct) ***REMOVED***
				throw new Error("Encoding mismatch: " + encoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("DefaultDec", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "hello world";
			let decoded = encoding.b64decode("aGVsbG8gd29ybGQ=");
			if (decoded !== correct) ***REMOVED***
				throw new Error("Decoding mismatch: " + decoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("DefaultUnicodeEnc", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "44GT44KT44Gr44Gh44Gv5LiW55WM";
			let encoded = encoding.b64encode("こんにちは世界", "std");
			if (encoded !== correct) ***REMOVED***
				throw new Error("Encoding mismatch: " + encoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("DefaultUnicodeDec", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "こんにちは世界";
			let decoded = encoding.b64decode("44GT44KT44Gr44Gh44Gv5LiW55WM");
			if (decoded !== correct) ***REMOVED***
				throw new Error("Decoding mismatch: " + decoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("StdEnc", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "aGVsbG8gd29ybGQ=";
			let encoded = encoding.b64encode("hello world", "std");
			if (encoded !== correct) ***REMOVED***
				throw new Error("Encoding mismatch: " + encoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("StdDec", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "hello world";
			let decoded = encoding.b64decode("aGVsbG8gd29ybGQ=", "std");
			if (decoded !== correct) ***REMOVED***
				throw new Error("Decoding mismatch: " + decoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("RawStdEnc", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "aGVsbG8gd29ybGQ";
			let encoded = encoding.b64encode("hello world", "rawstd");
			if (encoded !== correct) ***REMOVED***
				throw new Error("Encoding mismatch: " + encoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("RawStdDec", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "hello world";
			let decoded = encoding.b64decode("aGVsbG8gd29ybGQ", "rawstd");
			if (decoded !== correct) ***REMOVED***
				throw new Error("Decoding mismatch: " + decoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("URLEnc", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "5bCP6aO85by-Li4=";
			let encoded = encoding.b64encode("小飼弾..", "url");
			if (encoded !== correct) ***REMOVED***
				throw new Error("Encoding mismatch: " + encoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("URLDec", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "小飼弾..";
			let decoded = encoding.b64decode("5bCP6aO85by-Li4=", "url");
			if (decoded !== correct) ***REMOVED***
				throw new Error("Decoding mismatch: " + decoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("RawURLEnc", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "5bCP6aO85by-Li4";
			let encoded = encoding.b64encode("小飼弾..", "rawurl");
			if (encoded !== correct) ***REMOVED***
				throw new Error("Encoding mismatch: " + encoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
		t.Run("RawURLDec", func(t *testing.T) ***REMOVED***
			_, err := common.RunString(rt, `
			const correct = "小飼弾..";
			let decoded = encoding.b64decode("5bCP6aO85by-Li4", "rawurl");
			if (decoded !== correct) ***REMOVED***
				throw new Error("Decoding mismatch: " + decoded);
			***REMOVED***`)
			assert.NoError(t, err)
		***REMOVED***)
	***REMOVED***)
***REMOVED***
