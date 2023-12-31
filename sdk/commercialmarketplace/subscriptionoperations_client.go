// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package commercialmarketplace

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionOperationsClient contains the methods for the SubscriptionOperations group.
// Don't use this type directly, use NewSubscriptionOperationsClient() instead.
type SubscriptionOperationsClient struct {
	con *Connection
}

// NewSubscriptionOperationsClient creates a new instance of SubscriptionOperationsClient with the specified values.
func NewSubscriptionOperationsClient(con *Connection) *SubscriptionOperationsClient {
	return &SubscriptionOperationsClient{con: con}
}

// GetOperationStatus - Enables the publisher to track the status of the specified triggered async operation (such as Subscribe, Unsubscribe, ChangePlan,
// or ChangeQuantity).
// If the operation fails it returns a generic error.
func (client *SubscriptionOperationsClient) GetOperationStatus(ctx context.Context, subscriptionID string, operationID string, options *SubscriptionOperationsGetOperationStatusOptions) (SubscriptionOperationsGetOperationStatusResponse, error) {
	req, err := client.getOperationStatusCreateRequest(ctx, subscriptionID, operationID, options)
	if err != nil {
		return SubscriptionOperationsGetOperationStatusResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SubscriptionOperationsGetOperationStatusResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SubscriptionOperationsGetOperationStatusResponse{}, client.getOperationStatusHandleError(resp)
	}
	return client.getOperationStatusHandleResponse(resp)
}

// getOperationStatusCreateRequest creates the GetOperationStatus request.
func (client *SubscriptionOperationsClient) getOperationStatusCreateRequest(ctx context.Context, subscriptionID string, operationID string, options *SubscriptionOperationsGetOperationStatusOptions) (*azcore.Request, error) {
	urlPath := "/saas/subscriptions/{subscriptionId}/operations/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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
	return req, nil
}

// getOperationStatusHandleResponse handles the GetOperationStatus response.
func (client *SubscriptionOperationsClient) getOperationStatusHandleResponse(resp *azcore.Response) (SubscriptionOperationsGetOperationStatusResponse, error) {
	result := SubscriptionOperationsGetOperationStatusResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.Operation); err != nil {
		return SubscriptionOperationsGetOperationStatusResponse{}, err
	}
	return result, nil
}

// getOperationStatusHandleError handles the GetOperationStatus error response.
func (client *SubscriptionOperationsClient) getOperationStatusHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

// ListOperations - Lists the outstanding operations for the current publisher.
// If the operation fails it returns a generic error.
func (client *SubscriptionOperationsClient) ListOperations(ctx context.Context, subscriptionID string, options *SubscriptionOperationsListOperationsOptions) (SubscriptionOperationsListOperationsResponse, error) {
	req, err := client.listOperationsCreateRequest(ctx, subscriptionID, options)
	if err != nil {
		return SubscriptionOperationsListOperationsResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SubscriptionOperationsListOperationsResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SubscriptionOperationsListOperationsResponse{}, client.listOperationsHandleError(resp)
	}
	return client.listOperationsHandleResponse(resp)
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *SubscriptionOperationsClient) listOperationsCreateRequest(ctx context.Context, subscriptionID string, options *SubscriptionOperationsListOperationsOptions) (*azcore.Request, error) {
	urlPath := "/saas/subscriptions/{subscriptionId}/operations"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *SubscriptionOperationsClient) listOperationsHandleResponse(resp *azcore.Response) (SubscriptionOperationsListOperationsResponse, error) {
	result := SubscriptionOperationsListOperationsResponse{RawResponse: resp.Response}
	if err := resp.UnmarshalAsJSON(&result.OperationList); err != nil {
		return SubscriptionOperationsListOperationsResponse{}, err
	}
	return result, nil
}

// listOperationsHandleError handles the ListOperations error response.
func (client *SubscriptionOperationsClient) listOperationsHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }

// UpdateOperationStatus - Update the status of an operation to indicate success or failure with the provided values.
// If the operation fails it returns a generic error.
func (client *SubscriptionOperationsClient) UpdateOperationStatus(ctx context.Context, subscriptionID string, operationID string, body UpdateOperation, options *SubscriptionOperationsUpdateOperationStatusOptions) (SubscriptionOperationsUpdateOperationStatusResponse, error) {
	req, err := client.updateOperationStatusCreateRequest(ctx, subscriptionID, operationID, body, options)
	if err != nil {
		return SubscriptionOperationsUpdateOperationStatusResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return SubscriptionOperationsUpdateOperationStatusResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return SubscriptionOperationsUpdateOperationStatusResponse{}, client.updateOperationStatusHandleError(resp)
	}
	return SubscriptionOperationsUpdateOperationStatusResponse{RawResponse: resp.Response}, nil
}

// updateOperationStatusCreateRequest creates the UpdateOperationStatus request.
func (client *SubscriptionOperationsClient) updateOperationStatusCreateRequest(ctx context.Context, subscriptionID string, operationID string, body UpdateOperation, options *SubscriptionOperationsUpdateOperationStatusOptions) (*azcore.Request, error) {
	urlPath := "/saas/subscriptions/{subscriptionId}/operations/{operationId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
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
	return req, req.MarshalAsJSON(body)
}

// updateOperationStatusHandleError handles the UpdateOperationStatus error response.
func (client *SubscriptionOperationsClient) updateOperationStatusHandleError(resp *azcore.Response) error {
	body, err := resp.Payload()
	if err != nil {
		return azcore.NewResponseError(err, resp.Response)
	}
		if len(body) == 0 {
      return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
    }
    return azcore.NewResponseError(errors.New(string(body)), resp.Response)
    }


