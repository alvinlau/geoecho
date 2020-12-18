// Code generated by go-swagger; DO NOT EDIT.

package geolocate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetGeolocateIPReader is a Reader for the GetGeolocateIP structure.
type GetGeolocateIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGeolocateIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGeolocateIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGeolocateIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGeolocateIPInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGeolocateIPOK creates a GetGeolocateIPOK with default headers values
func NewGetGeolocateIPOK() *GetGeolocateIPOK {
	return &GetGeolocateIPOK{}
}

/*GetGeolocateIPOK handles this case with default header values.

ok
*/
type GetGeolocateIPOK struct {
	Payload string
}

func (o *GetGeolocateIPOK) Error() string {
	return fmt.Sprintf("[GET /geolocate/{ip}][%d] getGeolocateIpOK  %+v", 200, o.Payload)
}

func (o *GetGeolocateIPOK) GetPayload() string {
	return o.Payload
}

func (o *GetGeolocateIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGeolocateIPBadRequest creates a GetGeolocateIPBadRequest with default headers values
func NewGetGeolocateIPBadRequest() *GetGeolocateIPBadRequest {
	return &GetGeolocateIPBadRequest{}
}

/*GetGeolocateIPBadRequest handles this case with default header values.

invalid or missing ip address
*/
type GetGeolocateIPBadRequest struct {
	Payload string
}

func (o *GetGeolocateIPBadRequest) Error() string {
	return fmt.Sprintf("[GET /geolocate/{ip}][%d] getGeolocateIpBadRequest  %+v", 400, o.Payload)
}

func (o *GetGeolocateIPBadRequest) GetPayload() string {
	return o.Payload
}

func (o *GetGeolocateIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGeolocateIPInternalServerError creates a GetGeolocateIPInternalServerError with default headers values
func NewGetGeolocateIPInternalServerError() *GetGeolocateIPInternalServerError {
	return &GetGeolocateIPInternalServerError{}
}

/*GetGeolocateIPInternalServerError handles this case with default header values.

unable to parse response from geolocation API
*/
type GetGeolocateIPInternalServerError struct {
	Payload string
}

func (o *GetGeolocateIPInternalServerError) Error() string {
	return fmt.Sprintf("[GET /geolocate/{ip}][%d] getGeolocateIpInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGeolocateIPInternalServerError) GetPayload() string {
	return o.Payload
}

func (o *GetGeolocateIPInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}