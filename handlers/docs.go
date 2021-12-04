// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/denis-zakharov/gowebdev/data"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
//lint:ignore U1000 swagger response model
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
//lint:ignore U1000 swagger response model
type errorValidationWrapper struct {
	// Collection of the errors
	// in: body
	Body ValidationError
}

// A list of products
// swagger:response productsResponse
//lint:ignore U1000 swagger response model
type productsResponseWrapper struct {
	// All current products
	// in: body
	Body []data.Product
}

// Data structure representing a single product
// swagger:response productResponse
//lint:ignore U1000 swagger response model
type productResponseWrapper struct {
	// Newly created product
	// in: body
	Body data.Product
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 swagger response model
type noContentResponseWrapper struct {
}

// swagger:parameters updateProduct createProduct
//lint:ignore U1000 swagger response model
type productParamsWrapper struct {
	// Product data structure to Update or Create.
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters listSingleProduct deleteProduct
//lint:ignore U1000 swagger response model
type productIDParamsWrapper struct {
	// The id of the product for which the operation relates
	// in: path
	// required: true
	ID int `json:"id"`
}
