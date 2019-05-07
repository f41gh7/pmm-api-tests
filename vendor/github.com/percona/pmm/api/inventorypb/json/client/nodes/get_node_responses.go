// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// GetNodeReader is a Reader for the GetNode structure.
type GetNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodeOK creates a GetNodeOK with default headers values
func NewGetNodeOK() *GetNodeOK {
	return &GetNodeOK{}
}

/*GetNodeOK handles this case with default header values.

A successful response.
*/
type GetNodeOK struct {
	Payload *GetNodeOKBody
}

func (o *GetNodeOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/Get][%d] getNodeOk  %+v", 200, o.Payload)
}

func (o *GetNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeDefault creates a GetNodeDefault with default headers values
func NewGetNodeDefault(code int) *GetNodeDefault {
	return &GetNodeDefault{
		_statusCode: code,
	}
}

/*GetNodeDefault handles this case with default header values.

An error response.
*/
type GetNodeDefault struct {
	_statusCode int

	Payload *GetNodeDefaultBody
}

// Code gets the status code for the get node default response
func (o *GetNodeDefault) Code() int {
	return o._statusCode
}

func (o *GetNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/Get][%d] GetNode default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNodeBody get node body
swagger:model GetNodeBody
*/
type GetNodeBody struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`
}

// Validate validates this get node body
func (o *GetNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeBody) UnmarshalBinary(b []byte) error {
	var res GetNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model GetNodeDefaultBody
*/
type GetNodeDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this get node default body
func (o *GetNodeDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res GetNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBody get node OK body
swagger:model GetNodeOKBody
*/
type GetNodeOKBody struct {

	// container
	Container *GetNodeOKBodyContainer `json:"container,omitempty"`

	// generic
	Generic *GetNodeOKBodyGeneric `json:"generic,omitempty"`

	// remote
	Remote *GetNodeOKBodyRemote `json:"remote,omitempty"`

	// remote amazon rds
	RemoteAmazonRDS *GetNodeOKBodyRemoteAmazonRDS `json:"remote_amazon_rds,omitempty"`
}

// Validate validates this get node OK body
func (o *GetNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRemoteAmazonRDS(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNodeOKBody) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(o.Container) { // not required
		return nil
	}

	if o.Container != nil {
		if err := o.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "container")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(o.Generic) { // not required
		return nil
	}

	if o.Generic != nil {
		if err := o.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "generic")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	if o.Remote != nil {
		if err := o.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

func (o *GetNodeOKBody) validateRemoteAmazonRDS(formats strfmt.Registry) error {

	if swag.IsZero(o.RemoteAmazonRDS) { // not required
		return nil
	}

	if o.RemoteAmazonRDS != nil {
		if err := o.RemoteAmazonRDS.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNodeOk" + "." + "remote_amazon_rds")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBody) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyContainer ContainerNode represents a Docker container.
swagger:model GetNodeOKBodyContainer
*/
type GetNodeOKBodyContainer struct {

	// Address FIXME https://jira.percona.com/browse/PMM-3786
	Address string `json:"address,omitempty"`

	// Node availability zone. Auto-detected and auto-updated.
	Az string `json:"az,omitempty"`

	// Container identifier. If specified, must be a unique Docker container identifier.
	// Auto-detected and auto-updated.
	ContainerID string `json:"container_id,omitempty"`

	// Container name. Auto-detected and auto-updated.
	ContainerName string `json:"container_name,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs. Auto-detected and auto-updated.
	// If defined, Generic Node with that machine_id must exist.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Node model. Auto-detected and auto-updated.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Node region. Auto-detected and auto-updated.
	Region string `json:"region,omitempty"`
}

// Validate validates this get node OK body container
func (o *GetNodeOKBodyContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyContainer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyContainer) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyGeneric GenericNode represents a bare metal server or virtual machine.
swagger:model GetNodeOKBodyGeneric
*/
type GetNodeOKBodyGeneric struct {

	// Address FIXME https://jira.percona.com/browse/PMM-3786
	Address string `json:"address,omitempty"`

	// Node availability zone. Auto-detected and auto-updated.
	Az string `json:"az,omitempty"`

	// Custom user-assigned labels. Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Linux distribution name and version. Auto-detected and auto-updated.
	Distro string `json:"distro,omitempty"`

	// Linux machine-id. Auto-detected and auto-updated.
	// Must be unique across all Generic Nodes if specified.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Node model. Auto-detected and auto-updated.
	NodeModel string `json:"node_model,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Node region. Auto-detected and auto-updated.
	Region string `json:"region,omitempty"`
}

// Validate validates this get node OK body generic
func (o *GetNodeOKBodyGeneric) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyGeneric) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyGeneric) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyGeneric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyRemote RemoteNode represents generic remote Node. Agents can't run on Remote Nodes.
swagger:model GetNodeOKBodyRemote
*/
type GetNodeOKBodyRemote struct {

	// Custom user-assigned labels. Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this get node OK body remote
func (o *GetNodeOKBodyRemote) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyRemote) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyRemote) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNodeOKBodyRemoteAmazonRDS RemoteAmazonRDSNode represents a Remote Node for Amazon RDS. Agents can't run on Remote Nodes.
swagger:model GetNodeOKBodyRemoteAmazonRDS
*/
type GetNodeOKBodyRemoteAmazonRDS struct {

	// Custom user-assigned labels. Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// DB instance identifier. Unique across all RemoteAmazonRDS Nodes in combination with region. Can be changed.
	Instance string `json:"instance,omitempty"`

	// Unique randomly generated instance identifier. Can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name. Can't be changed.
	NodeName string `json:"node_name,omitempty"`

	// Unique across all RemoteAmazonRDS Nodes in combination with instance. Can't be changed.
	Region string `json:"region,omitempty"`
}

// Validate validates this get node OK body remote amazon RDS
func (o *GetNodeOKBodyRemoteAmazonRDS) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNodeOKBodyRemoteAmazonRDS) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNodeOKBodyRemoteAmazonRDS) UnmarshalBinary(b []byte) error {
	var res GetNodeOKBodyRemoteAmazonRDS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
