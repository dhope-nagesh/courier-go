// This file was auto-generated by Fern from our API Definition.

package auditevents

import (
	context "context"
	fmt "fmt"
	v3 "github.com/trycourier/courier-go/v3"
	core "github.com/trycourier/courier-go/v3/core"
	option "github.com/trycourier/courier-go/v3/option"
	http "net/http"
	url "net/url"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Fetch the list of audit events
func (c *Client) List(
	ctx context.Context,
	request *v3.ListAuditEventsRequest,
	opts ...option.RequestOption,
) (*v3.ListAuditEventsResponse, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := baseURL + "/" + "audit-events"

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *v3.ListAuditEventsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  headers,
			Client:   options.HTTPClient,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Fetch a specific audit event by ID.
func (c *Client) Get(
	ctx context.Context,
	// A unique identifier associated with the audit event you wish to retrieve
	auditEventId string,
	opts ...option.RequestOption,
) (*v3.AuditEvent, error) {
	options := core.NewRequestOptions(opts...)

	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	if options.BaseURL != "" {
		baseURL = options.BaseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"audit-events/%v", auditEventId)

	headers := core.MergeHeaders(c.header.Clone(), options.ToHeader())

	var response *v3.AuditEvent
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodGet,
			Headers:  headers,
			Client:   options.HTTPClient,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
