// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UpdateDbHomeRequest wrapper for the UpdateDbHome operation
type UpdateDbHomeRequest struct {

	// The database home OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	DbHomeId *string `mandatory:"true" contributesTo:"path" name:"dbHomeId"`

	// Request to update the properties of a DB Home.
	UpdateDbHomeDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`
}

func (request UpdateDbHomeRequest) String() string {
	return common.PointerString(request)
}

// UpdateDbHomeResponse wrapper for the UpdateDbHome operation
type UpdateDbHomeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbHome instance
	DbHome `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateDbHomeResponse) String() string {
	return common.PointerString(response)
}
