// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetSecurityGroupsIDReader is a Reader for the GetSecurityGroupsID structure.
type GetSecurityGroupsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSecurityGroupsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSecurityGroupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetSecurityGroupsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSecurityGroupsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSecurityGroupsIDOK creates a GetSecurityGroupsIDOK with default headers values
func NewGetSecurityGroupsIDOK() *GetSecurityGroupsIDOK {
	return &GetSecurityGroupsIDOK{}
}

/*GetSecurityGroupsIDOK handles this case with default header values.

dummy
*/
type GetSecurityGroupsIDOK struct {
	Payload *models.SecurityGroup
}

func (o *GetSecurityGroupsIDOK) Error() string {
	return fmt.Sprintf("[GET /security_groups/{id}][%d] getSecurityGroupsIdOK  %+v", 200, o.Payload)
}

func (o *GetSecurityGroupsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsIDNotFound creates a GetSecurityGroupsIDNotFound with default headers values
func NewGetSecurityGroupsIDNotFound() *GetSecurityGroupsIDNotFound {
	return &GetSecurityGroupsIDNotFound{}
}

/*GetSecurityGroupsIDNotFound handles this case with default header values.

error
*/
type GetSecurityGroupsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /security_groups/{id}][%d] getSecurityGroupsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetSecurityGroupsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSecurityGroupsIDInternalServerError creates a GetSecurityGroupsIDInternalServerError with default headers values
func NewGetSecurityGroupsIDInternalServerError() *GetSecurityGroupsIDInternalServerError {
	return &GetSecurityGroupsIDInternalServerError{}
}

/*GetSecurityGroupsIDInternalServerError handles this case with default header values.

error
*/
type GetSecurityGroupsIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetSecurityGroupsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /security_groups/{id}][%d] getSecurityGroupsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSecurityGroupsIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
