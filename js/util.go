package js

import (
	"github.com/robertkrimen/otto"
	"net/url"
	"strings"
)

func paramsFromObject(o *otto.Object) (params HTTPParams, err error) ***REMOVED***
	if o == nil ***REMOVED***
		return params, nil
	***REMOVED***

	for _, key := range o.Keys() ***REMOVED***
		switch key ***REMOVED***
		case "follow":
			v, err := o.Get(key)
			if err != nil ***REMOVED***
				return params, err
			***REMOVED***
			follow, err := v.ToBoolean()
			if err != nil ***REMOVED***
				return params, err
			***REMOVED***
			params.Follow = follow
		case "quiet":
			v, err := o.Get(key)
			if err != nil ***REMOVED***
				return params, err
			***REMOVED***
			quiet, err := v.ToBoolean()
			if err != nil ***REMOVED***
				return params, err
			***REMOVED***
			params.Quiet = quiet
		case "headers":
			v, err := o.Get(key)
			if err != nil ***REMOVED***
				return params, err
			***REMOVED***
			obj := v.Object()
			if obj == nil ***REMOVED***
				continue
			***REMOVED***

			params.Headers = make(map[string]string)
			for _, name := range obj.Keys() ***REMOVED***
				hv, err := obj.Get(name)
				if err != nil ***REMOVED***
					return params, err
				***REMOVED***
				value, err := hv.ToString()
				if err != nil ***REMOVED***
					return params, err
				***REMOVED***
				params.Headers[name] = value
			***REMOVED***
		***REMOVED***
	***REMOVED***

	return params, nil
***REMOVED***

func bodyFromValue(o otto.Value) (body string, isForm bool, err error) ***REMOVED***
	if o.IsUndefined() || o.IsNull() ***REMOVED***
		return "", false, nil
	***REMOVED***

	if o.IsObject() ***REMOVED***
		query := make(url.Values)
		obj := o.Object()
		for _, key := range obj.Keys() ***REMOVED***
			valObj, _ := obj.Get(key)
			val, err := valObj.ToString()
			if err != nil ***REMOVED***
				return "", false, err
			***REMOVED***
			query.Add(key, val)
		***REMOVED***
		return query.Encode(), true, nil
	***REMOVED***

	body, err = o.ToString()
	if err != nil ***REMOVED***
		return "", false, err
	***REMOVED***

	return body, false, nil
***REMOVED***

func putBodyInURL(url, body string) string ***REMOVED***
	if body == "" ***REMOVED***
		return url
	***REMOVED***

	if !strings.ContainsRune(url, '?') ***REMOVED***
		return url + "?" + body
	***REMOVED*** else ***REMOVED***
		return url + "&" + body
	***REMOVED***
***REMOVED***

func resolveRedirect(from, to string) string ***REMOVED***
	if to == "" ***REMOVED***
		return from
	***REMOVED***

	uFrom, err := url.Parse(from)
	if err != nil ***REMOVED***
		return to
	***REMOVED***

	uTo, err := url.Parse(to)
	if err != nil ***REMOVED***
		return to
	***REMOVED***

	return uFrom.ResolveReference(uTo).String()
***REMOVED***

func Make(vm *otto.Otto, t string) (*otto.Object, error) ***REMOVED***
	val, err := vm.Call("new "+t, nil)
	if err != nil ***REMOVED***
		return nil, err
	***REMOVED***

	return val.Object(), nil
***REMOVED***

func jsCustomError(vm *otto.Otto, t string, err error) otto.Value ***REMOVED***
	return vm.MakeCustomError(t, err.Error())
***REMOVED***

func jsError(vm *otto.Otto, err error) otto.Value ***REMOVED***
	return jsCustomError(vm, "Error", err)
***REMOVED***
