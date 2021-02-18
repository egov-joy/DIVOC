// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/divoc/portal-api/swagger_gen/models"
)

// GetFacilities4PublicOKCode is the HTTP code returned for type GetFacilities4PublicOK
const GetFacilities4PublicOKCode int = 200

/*GetFacilities4PublicOK OK

swagger:response getFacilities4PublicOK
*/
type GetFacilities4PublicOK struct {

	/*
	  In: Body
	*/
	Payload []*models.PublicFacility `json:"body,omitempty"`
}

// NewGetFacilities4PublicOK creates GetFacilities4PublicOK with default headers values
func NewGetFacilities4PublicOK() *GetFacilities4PublicOK {

	return &GetFacilities4PublicOK{}
}

// WithPayload adds the payload to the get facilities4 public o k response
func (o *GetFacilities4PublicOK) WithPayload(payload []*models.PublicFacility) *GetFacilities4PublicOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get facilities4 public o k response
func (o *GetFacilities4PublicOK) SetPayload(payload []*models.PublicFacility) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFacilities4PublicOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.PublicFacility, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
