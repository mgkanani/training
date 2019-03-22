// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// GetTransactionReceivedOKCode is the HTTP code returned for type GetTransactionReceivedOK
const GetTransactionReceivedOKCode int = 200

/*GetTransactionReceivedOK Success!

swagger:response getTransactionReceivedOK
*/
type GetTransactionReceivedOK struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`

	/*
	  In: Body
	*/
	Payload *model.Transaction `json:"body,omitempty"`
}

// NewGetTransactionReceivedOK creates GetTransactionReceivedOK with default headers values
func NewGetTransactionReceivedOK() *GetTransactionReceivedOK {

	return &GetTransactionReceivedOK{}
}

// WithVersion adds the version to the get transaction received o k response
func (o *GetTransactionReceivedOK) WithVersion(version string) *GetTransactionReceivedOK {
	o.Version = version
	return o
}

// SetVersion sets the version to the get transaction received o k response
func (o *GetTransactionReceivedOK) SetVersion(version string) {
	o.Version = version
}

// WithPayload adds the payload to the get transaction received o k response
func (o *GetTransactionReceivedOK) WithPayload(payload *model.Transaction) *GetTransactionReceivedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get transaction received o k response
func (o *GetTransactionReceivedOK) SetPayload(payload *model.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTransactionReceivedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTransactionReceivedNotFoundCode is the HTTP code returned for type GetTransactionReceivedNotFound
const GetTransactionReceivedNotFoundCode int = 404

/*GetTransactionReceivedNotFound Transaction not found

swagger:response getTransactionReceivedNotFound
*/
type GetTransactionReceivedNotFound struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewGetTransactionReceivedNotFound creates GetTransactionReceivedNotFound with default headers values
func NewGetTransactionReceivedNotFound() *GetTransactionReceivedNotFound {

	return &GetTransactionReceivedNotFound{}
}

// WithVersion adds the version to the get transaction received not found response
func (o *GetTransactionReceivedNotFound) WithVersion(version string) *GetTransactionReceivedNotFound {
	o.Version = version
	return o
}

// SetVersion sets the version to the get transaction received not found response
func (o *GetTransactionReceivedNotFound) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *GetTransactionReceivedNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// GetTransactionReceivedInternalServerErrorCode is the HTTP code returned for type GetTransactionReceivedInternalServerError
const GetTransactionReceivedInternalServerErrorCode int = 500

/*GetTransactionReceivedInternalServerError Internal server error

swagger:response getTransactionReceivedInternalServerError
*/
type GetTransactionReceivedInternalServerError struct {
	/*Version of the microservice that responded

	 */
	Version string `json:"version"`
}

// NewGetTransactionReceivedInternalServerError creates GetTransactionReceivedInternalServerError with default headers values
func NewGetTransactionReceivedInternalServerError() *GetTransactionReceivedInternalServerError {

	return &GetTransactionReceivedInternalServerError{}
}

// WithVersion adds the version to the get transaction received internal server error response
func (o *GetTransactionReceivedInternalServerError) WithVersion(version string) *GetTransactionReceivedInternalServerError {
	o.Version = version
	return o
}

// SetVersion sets the version to the get transaction received internal server error response
func (o *GetTransactionReceivedInternalServerError) SetVersion(version string) {
	o.Version = version
}

// WriteResponse to the client
func (o *GetTransactionReceivedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header version

	version := o.Version
	if version != "" {
		rw.Header().Set("version", version)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}