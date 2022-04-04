/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2020 Load Impact
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

package types

import (
	"math/big"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:gochecknoglobals
var max64 = new(big.Int).Exp(big.NewInt(2), big.NewInt(64), nil)

func get128BigInt(hi, lo int64) *big.Int ***REMOVED***
	h := big.NewInt(hi)
	h.Mul(h, max64)
	return h.Add(h, big.NewInt(lo))
***REMOVED***

func TestIpBlock(t *testing.T) ***REMOVED***
	t.Parallel()
	testdata := map[string]struct ***REMOVED***
		count           *big.Int
		firstIP, lastIP net.IP
	***REMOVED******REMOVED***
		"192.168.0.101": ***REMOVED***new(big.Int).SetInt64(1), net.ParseIP("192.168.0.101"), net.ParseIP("192.168.0.101")***REMOVED***,

		"192.168.0.101-192.168.0.200":    ***REMOVED***new(big.Int).SetInt64(100), net.ParseIP("192.168.0.101"), net.ParseIP("192.168.0.200")***REMOVED***,
		"192.168.0.100-192.168.0.200":    ***REMOVED***new(big.Int).SetInt64(101), net.ParseIP("192.168.0.100"), net.ParseIP("192.168.0.200")***REMOVED***,
		"fd00:1:1:0::0-fd00:1:1:ff::3ff": ***REMOVED***get128BigInt(255, 1024), net.ParseIP("fd00:1:1:0::0"), net.ParseIP("fd00:1:1:ff::3ff")***REMOVED***,
		"fd00:1:1:2::1-fd00:1:1:ff::3ff": ***REMOVED***get128BigInt(253, 1023), net.ParseIP("fd00:1:1:2::1"), net.ParseIP("fd00:1:1:ff::3ff")***REMOVED***,

		"192.168.0.0/16":  ***REMOVED***get128BigInt(0, 65534), net.ParseIP("192.168.0.1"), net.ParseIP("192.168.255.254")***REMOVED***,
		"192.168.0.1/16":  ***REMOVED***get128BigInt(0, 65534), net.ParseIP("192.168.0.1"), net.ParseIP("192.168.255.254")***REMOVED***,
		"192.168.0.10/16": ***REMOVED***get128BigInt(0, 65534), net.ParseIP("192.168.0.1"), net.ParseIP("192.168.255.254")***REMOVED***,
		"192.168.0.10/31": ***REMOVED***get128BigInt(0, 2), net.ParseIP("192.168.0.10"), net.ParseIP("192.168.0.11")***REMOVED***,
		"192.168.0.10/32": ***REMOVED***get128BigInt(0, 1), net.ParseIP("192.168.0.10"), net.ParseIP("192.168.0.10")***REMOVED***,
		"fd00::0/120":     ***REMOVED***get128BigInt(0, 256), net.ParseIP("fd00::0"), net.ParseIP("fd00::ff")***REMOVED***,
		"fd00::1/120":     ***REMOVED***get128BigInt(0, 256), net.ParseIP("fd00::0"), net.ParseIP("fd00::ff")***REMOVED***,
		"fd00::3/120":     ***REMOVED***get128BigInt(0, 256), net.ParseIP("fd00::0"), net.ParseIP("fd00::ff")***REMOVED***,
		"fd00::0/112":     ***REMOVED***get128BigInt(0, 65536), net.ParseIP("fd00::0"), net.ParseIP("fd00::ffff")***REMOVED***,
		"fd00::1/112":     ***REMOVED***get128BigInt(0, 65536), net.ParseIP("fd00::0"), net.ParseIP("fd00::ffff")***REMOVED***,
		"fd00::2/112":     ***REMOVED***get128BigInt(0, 65536), net.ParseIP("fd00::0"), net.ParseIP("fd00::ffff")***REMOVED***,
	***REMOVED***
	for name, data := range testdata ***REMOVED***
		name, data := name, data
		t.Run(name, func(t *testing.T) ***REMOVED***
			t.Parallel()
			b, err := getIPBlock(name)
			require.NoError(t, err)
			assert.Equal(t, data.count, b.count)
			pb := ipPoolBlock***REMOVED***firstIP: b.firstIP***REMOVED***
			idx := big.NewInt(0)
			assert.Equal(t, data.firstIP.To16(), pb.getIP(idx).To16())
			idx.Sub(idx.Add(idx, b.count), big.NewInt(1))
			assert.Equal(t, data.lastIP.To16(), pb.getIP(idx).To16())
		***REMOVED***)
	***REMOVED***
***REMOVED***

func TestIPPool(t *testing.T) ***REMOVED***
	t.Parallel()
	testdata := map[string]struct ***REMOVED***
		count   *big.Int
		queries map[uint64]net.IP
	***REMOVED******REMOVED***
		"192.168.0.101": ***REMOVED***
			count:   new(big.Int).SetInt64(1),
			queries: map[uint64]net.IP***REMOVED***0: net.ParseIP("192.168.0.101"), 12: net.ParseIP("192.168.0.101")***REMOVED***,
		***REMOVED***,
		"192.168.0.101,192.168.0.102": ***REMOVED***
			count: new(big.Int).SetInt64(2),
			queries: map[uint64]net.IP***REMOVED***
				0:  net.ParseIP("192.168.0.101"),
				1:  net.ParseIP("192.168.0.102"),
				12: net.ParseIP("192.168.0.101"),
				13: net.ParseIP("192.168.0.102"),
			***REMOVED***,
		***REMOVED***,
		"192.168.0.101-192.168.0.105,fd00::2/112": ***REMOVED***
			count: new(big.Int).SetInt64(65541),
			queries: map[uint64]net.IP***REMOVED***
				0:     net.ParseIP("192.168.0.101"),
				1:     net.ParseIP("192.168.0.102"),
				5:     net.ParseIP("fd00::0"),
				6:     net.ParseIP("fd00::1"),
				65541: net.ParseIP("192.168.0.101"),
			***REMOVED***,
		***REMOVED***,

		"192.168.0.101,192.168.0.102,192.168.0.103,192.168.0.104,192.168.0.105-192.168.0.105,fd00::2/112": ***REMOVED***
			count: new(big.Int).SetInt64(65541),
			queries: map[uint64]net.IP***REMOVED***
				0:     net.ParseIP("192.168.0.101"),
				1:     net.ParseIP("192.168.0.102"),
				2:     net.ParseIP("192.168.0.103"),
				3:     net.ParseIP("192.168.0.104"),
				4:     net.ParseIP("192.168.0.105"),
				5:     net.ParseIP("fd00::0"),
				6:     net.ParseIP("fd00::1"),
				65541: net.ParseIP("192.168.0.101"),
			***REMOVED***,
		***REMOVED***,
	***REMOVED***
	for name, data := range testdata ***REMOVED***
		name, data := name, data
		t.Run(name, func(t *testing.T) ***REMOVED***
			t.Parallel()
			p, err := NewIPPool(name)
			require.NoError(t, err)
			assert.Equal(t, data.count, p.count)
			for q, a := range data.queries ***REMOVED***
				assert.Equal(t, a.To16(), p.GetIP(q).To16(), "index %d", q)
			***REMOVED***
		***REMOVED***)
	***REMOVED***
***REMOVED***

func TestIpBlockError(t *testing.T) ***REMOVED***
	t.Parallel()
	testdata := map[string]string***REMOVED***
		"whatever":                       "not a valid IP",
		"192.168.0.1012":                 "not a valid IP",
		"192.168.0.10/244":               "invalid CIDR",
		"fd00::0/244":                    "invalid CIDR",
		"192.168.0.101-192.168.0.102/32": "wrong IP range format",
		"192.168.0.101-fd00::1":          "mixed IP range format",
		"fd00::1-192.168.0.101":          "mixed IP range format",
		"192.168.0.100-192.168.0.2":      "negative IP range",
		"fd00:1:1:0::0-fd00:1:0:ff::3ff": "negative IP range",
	***REMOVED***
	for name, data := range testdata ***REMOVED***
		name, data := name, data
		t.Run(name, func(t *testing.T) ***REMOVED***
			t.Parallel()
			_, err := getIPBlock(name)
			require.Error(t, err)
			require.Contains(t, err.Error(), data)
		***REMOVED***)
	***REMOVED***
***REMOVED***
