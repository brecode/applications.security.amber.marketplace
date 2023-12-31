// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package commercialmarketplace

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// MeteringOperationsClient contains the methods for the MeteringOperations group.
// Don't use this type directly, use NewMeteringOperationsClient() instead.
type MeteringOperationsClient struct {
	con *Connection
}

// NewMeteringOperationsClient creates a new instance of MeteringOperationsClient with the specified values.
func NewMeteringOperationsClient(con *Connection) *MeteringOperationsClient {
	return &MeteringOperationsClient{con: con}
}

// PostBatchUsageEvent - The batch usage event API allows you to emit usage events for more than one purchased entity at once. The batch usage event request
// references the metering services dimension defined by the publisher
// when publishing the offer.
// If the operation fails it returns a generic error.
func (client *MeteringOperationsClient) PostBatchUsageEvent(ctx context.Context, body BatchUsageEvent, options *MeteringOperationsPostBatchUsageEventOptions) (MeteringOperationsPostBatchUsageEventResponse, error) {
	req, err := client.postBatchUsageEventCreateRequest(ctx, body, options)
	if err != nil {
		return MeteringOperationsPostBatchUsageEventResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MeteringOperationsPostBatchUsageEventResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MeteringOperationsPostBatchUsageEventResponse{}, client.postBatchUsageEventHandleError(resp)
	}
	return client.postBatchUsageEventHandleResponse(resp)
}

// postBatchUsageEventCreateRequest creates the PostBatchUsageEvent request.
func (client *MeteringOperationsClient) postBatchUsageEventCreateRequest(ctx context.Context, body BatchUsageEvent, options *MeteringOperationsPostBatchUsageEventOptions) (*azcore.Request, error) {
	urlPath := "/batchUsageEvent"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-08-31")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-requestid", *options.RequestID)
	}
	if options != nil && options.CorrelationID != nil {
		req.Header.Set("x-ms-correlationid", *options.CorrelationID)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(body)
}

// postBatchUsageEventHandleResponse handles the PostBatchUsageEvent response.
func (client *MeteringOperationsClient) postBatchUsageEventHandleResponse(resp *azcore.Response) (MeteringOperationsPostBatchUsageEventResponse, error) {
	result := MeteringOperationsPostBatchUsageEventResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.BatchUsageEventOkResponse); err != nil {
		return MeteringOperationsPostBatchUsageEventResponse{}, err
	}
	return result, nil
}

// postBatchUsageEventHandleError handles the PostBatchUsageEvent error response.
func (client *MeteringOperationsClient) postBatchUsageEventHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

// PostUsageEvent - Posts a single usage event to the marketplace metering service API.
// If the operation fails it returns one of the following error types.
// - *UsageEventBadRequestResponse, *UsageEventConflictResponse
func (client *MeteringOperationsClient) PostUsageEvent(ctx context.Context, body UsageEvent, options *MeteringOperationsPostUsageEventOptions) (MeteringOperationsPostUsageEventResponse, error) {
	req, err := client.postUsageEventCreateRequest(ctx, body, options)
	if err != nil {
		return MeteringOperationsPostUsageEventResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return MeteringOperationsPostUsageEventResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return MeteringOperationsPostUsageEventResponse{}, client.postUsageEventHandleError(resp)
	}
	return client.postUsageEventHandleResponse(resp)
}

// postUsageEventCreateRequest creates the PostUsageEvent request.
func (client *MeteringOperationsClient) postUsageEventCreateRequest(ctx context.Context, body UsageEvent, options *MeteringOperationsPostUsageEventOptions) (*azcore.Request, error) {
	urlPath := "/usageEvent"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2018-08-31")
	req.URL.RawQuery = reqQP.Encode()
	if options != nil && options.RequestID != nil {
		req.Header.Set("x-ms-requestid", *options.RequestID)
	}
	if options != nil && options.CorrelationID != nil {
		req.Header.Set("x-ms-correlationid", *options.CorrelationID)
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(body)
}

// postUsageEventHandleResponse handles the PostUsageEvent response.
func (client *MeteringOperationsClient) postUsageEventHandleResponse(resp *azcore.Response) (MeteringOperationsPostUsageEventResponse, error) {
	result := MeteringOperationsPostUsageEventResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.UsageEventOkResponse); err != nil {
		return MeteringOperationsPostUsageEventResponse{}, err
	}
	return result, nil
}

// postUsageEventHandleError handles the PostUsageEvent error response.
func (client *MeteringOperationsClient) postUsageEventHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
	switch resp.StatusCode {
	case http.StatusBadRequest:
			errType := UsageEventBadRequestResponse{raw: string(body)}
		if err := resp.UnmarshalAsJSON(&errType); err != nil {
			return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
		}
		return azcore.NewResponseError(&errType, resp.Response)
	case http.StatusForbidden:
			if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    	case http.StatusConflict:
			errType := UsageEventConflictResponse{raw: string(body)}
		if err := resp.UnmarshalAsJSON(&errType); err != nil {
			return azcore.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp.Response)
		}
		return azcore.NewResponseError(&errType, resp.Response)
	default:
	if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    	}
}


