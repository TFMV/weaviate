//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2020 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new contextionary api API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contextionary api API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	C11yConcepts(params *C11yConceptsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yConceptsOK, error)

	C11yCorpusGet(params *C11yCorpusGetParams, authInfo runtime.ClientAuthInfoWriter) error

	C11yExtensions(params *C11yExtensionsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yExtensionsOK, error)

	C11yWords(params *C11yWordsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yWordsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  C11yConcepts checks if a concept is part of the contextionary

  Checks if a concept is part of the contextionary. Concepts should be concatenated as described here: https://github.com/semi-technologies/weaviate/blob/master/docs/en/use/ontology-schema.md#camelcase
*/
func (a *Client) C11yConcepts(params *C11yConceptsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yConceptsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yConceptsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.concepts",
		Method:             "GET",
		PathPattern:        "/c11y/concepts/{concept}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yConceptsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*C11yConceptsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for c11y.concepts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  C11yCorpusGet checks if a word or word string is part of the contextionary

  Analyzes a sentence based on the contextionary
*/
func (a *Client) C11yCorpusGet(params *C11yCorpusGetParams, authInfo runtime.ClientAuthInfoWriter) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yCorpusGetParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.corpus.get",
		Method:             "POST",
		PathPattern:        "/c11y/corpus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yCorpusGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil
}

/*
  C11yExtensions extends the contextionary with custom concepts

  Extend the contextionary with your own custom concepts
*/
func (a *Client) C11yExtensions(params *C11yExtensionsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yExtensionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yExtensionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.extensions",
		Method:             "POST",
		PathPattern:        "/c11y/extensions/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yExtensionsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*C11yExtensionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for c11y.extensions: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  C11yWords checks if a word or word string is part of the contextionary

  Checks if a word or wordString is part of the contextionary. Words should be concatenated as described here: https://github.com/semi-technologies/weaviate/blob/master/docs/en/use/ontology-schema.md#camelcase
*/
func (a *Client) C11yWords(params *C11yWordsParams, authInfo runtime.ClientAuthInfoWriter) (*C11yWordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewC11yWordsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "c11y.words",
		Method:             "GET",
		PathPattern:        "/c11y/words/{words}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &C11yWordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*C11yWordsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for c11y.words: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
