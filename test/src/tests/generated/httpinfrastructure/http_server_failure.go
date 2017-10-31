package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"tests/pipeline"
)

// HTTPServerFailureClient is the test Infrastructure for AutoRest
type HTTPServerFailureClient struct {
	ManagementClient
}

// NewHTTPServerFailureClient creates an instance of the HTTPServerFailureClient client.
func NewHTTPServerFailureClient(p pipeline.Pipeline) HTTPServerFailureClient {
	return HTTPServerFailureClient{NewManagementClient(p)}
}

// Delete505 return 505 status code - should be represented in the client as an error
//
// booleanValue is simple boolean value true
func (client HTTPServerFailureClient) Delete505(ctx context.Context, booleanValue *bool) (*Error, error) {
	req, err := client.delete505Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.delete505Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Error), err
}

// delete505Preparer prepares the Delete505 request.
func (client HTTPServerFailureClient) delete505Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/failure/server/505"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// delete505Responder handles the response to the Delete505 request.
func (client HTTPServerFailureClient) delete505Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Error{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Get501 return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Get501(ctx context.Context) (*Error, error) {
	req, err := client.get501Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.get501Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Error), err
}

// get501Preparer prepares the Get501 request.
func (client HTTPServerFailureClient) get501Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/failure/server/501"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// get501Responder handles the response to the Get501 request.
func (client HTTPServerFailureClient) get501Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Error{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Head501 return 501 status code - should be represented in the client as an error
func (client HTTPServerFailureClient) Head501(ctx context.Context) (*Error, error) {
	req, err := client.head501Preparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.head501Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Error), err
}

// head501Preparer prepares the Head501 request.
func (client HTTPServerFailureClient) head501Preparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/failure/server/501"
	req, err := pipeline.NewRequest("HEAD", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// head501Responder handles the response to the Head501 request.
func (client HTTPServerFailureClient) head501Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Error{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Post505 return 505 status code - should be represented in the client as an error
//
// booleanValue is simple boolean value true
func (client HTTPServerFailureClient) Post505(ctx context.Context, booleanValue *bool) (*Error, error) {
	req, err := client.post505Preparer(booleanValue)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.post505Responder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*Error), err
}

// post505Preparer prepares the Post505 request.
func (client HTTPServerFailureClient) post505Preparer(booleanValue *bool) (pipeline.Request, error) {
	u := client.url
	u.Path = "/http/failure/server/505"
	req, err := pipeline.NewRequest("POST", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(booleanValue)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// post505Responder handles the response to the Post505 request.
func (client HTTPServerFailureClient) post505Responder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &Error{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}