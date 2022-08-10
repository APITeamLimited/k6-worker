package client

import (
	"context"
	"net/http"
	"net/url"

	v1 "go.k6.io/k6/api/v1"
)

// Status returns the current k6 status.
func (c *Client) Status(ctx context.Context) (ret v1.Status, err error) ***REMOVED***
	var resp v1.StatusJSONAPI

	if err = c.CallAPI(ctx, http.MethodGet, &url.URL***REMOVED***Path: "/v1/status"***REMOVED***, nil, &resp); err != nil ***REMOVED***
		return ret, err
	***REMOVED***

	return resp.Status(), nil
***REMOVED***

// SetStatus tries to change the current status and returns the new one if it
// was successful.
func (c *Client) SetStatus(ctx context.Context, patch v1.Status) (ret v1.Status, err error) ***REMOVED***
	var resp v1.StatusJSONAPI

	apiURL := &url.URL***REMOVED***Path: "/v1/status"***REMOVED***
	if err = c.CallAPI(ctx, http.MethodPatch, apiURL, v1.NewStatusJSONAPI(patch), &resp); err != nil ***REMOVED***
		return ret, err
	***REMOVED***

	return resp.Status(), nil
***REMOVED***
