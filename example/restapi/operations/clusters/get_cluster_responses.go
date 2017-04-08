package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mikkeloscar/gin-swagger/example/models"
)

// HTTP code for type GetClusterOK
const GetClusterOKCode int = 200

/*GetClusterOK Cluster information.

swagger:response getClusterOK
*/
type GetClusterOK struct {

	/*
	  In: Body
	*/
	Payload *models.Cluster `json:"body,omitempty"`
}

// NewGetClusterOK creates GetClusterOK with default headers values
func NewGetClusterOK() *GetClusterOK {
	return &GetClusterOK{}
}

// WithPayload adds the payload to the get cluster o k response
func (o *GetClusterOK) WithPayload(payload *models.Cluster) *GetClusterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cluster o k response
func (o *GetClusterOK) SetPayload(payload *models.Cluster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClusterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HTTP code for type GetClusterUnauthorized
const GetClusterUnauthorizedCode int = 401

/*GetClusterUnauthorized Unauthorized

swagger:response getClusterUnauthorized
*/
type GetClusterUnauthorized struct {
}

// NewGetClusterUnauthorized creates GetClusterUnauthorized with default headers values
func NewGetClusterUnauthorized() *GetClusterUnauthorized {
	return &GetClusterUnauthorized{}
}

// WriteResponse to the client
func (o *GetClusterUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
}

// HTTP code for type GetClusterForbidden
const GetClusterForbiddenCode int = 403

/*GetClusterForbidden Forbidden

swagger:response getClusterForbidden
*/
type GetClusterForbidden struct {
}

// NewGetClusterForbidden creates GetClusterForbidden with default headers values
func NewGetClusterForbidden() *GetClusterForbidden {
	return &GetClusterForbidden{}
}

// WriteResponse to the client
func (o *GetClusterForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
}

// HTTP code for type GetClusterNotFound
const GetClusterNotFoundCode int = 404

/*GetClusterNotFound Cluster not found

swagger:response getClusterNotFound
*/
type GetClusterNotFound struct {
}

// NewGetClusterNotFound creates GetClusterNotFound with default headers values
func NewGetClusterNotFound() *GetClusterNotFound {
	return &GetClusterNotFound{}
}

// WriteResponse to the client
func (o *GetClusterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// HTTP code for type GetClusterInternalServerError
const GetClusterInternalServerErrorCode int = 500

/*GetClusterInternalServerError Unexpected error

swagger:response getClusterInternalServerError
*/
type GetClusterInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetClusterInternalServerError creates GetClusterInternalServerError with default headers values
func NewGetClusterInternalServerError() *GetClusterInternalServerError {
	return &GetClusterInternalServerError{}
}

// WithPayload adds the payload to the get cluster internal server error response
func (o *GetClusterInternalServerError) WithPayload(payload *models.Error) *GetClusterInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get cluster internal server error response
func (o *GetClusterInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetClusterInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}