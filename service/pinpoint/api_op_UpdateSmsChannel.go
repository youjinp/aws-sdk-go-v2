// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateSmsChannelRequest
type UpdateSmsChannelInput struct {
	_ struct{} `type:"structure" payload:"SMSChannelRequest"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	// Specifies the status and settings of the SMS channel for an application.
	//
	// SMSChannelRequest is a required field
	SMSChannelRequest *SMSChannelRequest `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateSmsChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSmsChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSmsChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if s.SMSChannelRequest == nil {
		invalidParams.Add(aws.NewErrParamRequired("SMSChannelRequest"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSmsChannelInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SMSChannelRequest != nil {
		v := s.SMSChannelRequest

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SMSChannelRequest", v, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateSmsChannelResponse
type UpdateSmsChannelOutput struct {
	_ struct{} `type:"structure" payload:"SMSChannelResponse"`

	// Provides information about the status and settings of the SMS channel for
	// an application.
	//
	// SMSChannelResponse is a required field
	SMSChannelResponse *SMSChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateSmsChannelOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateSmsChannelOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.SMSChannelResponse != nil {
		v := s.SMSChannelResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "SMSChannelResponse", v, metadata)
	}
	return nil
}

const opUpdateSmsChannel = "UpdateSmsChannel"

// UpdateSmsChannelRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Updates the status and settings of the SMS channel for an application.
//
//    // Example sending a request using UpdateSmsChannelRequest.
//    req := client.UpdateSmsChannelRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/UpdateSmsChannel
func (c *Client) UpdateSmsChannelRequest(input *UpdateSmsChannelInput) UpdateSmsChannelRequest {
	op := &aws.Operation{
		Name:       opUpdateSmsChannel,
		HTTPMethod: "PUT",
		HTTPPath:   "/v1/apps/{application-id}/channels/sms",
	}

	if input == nil {
		input = &UpdateSmsChannelInput{}
	}

	req := c.newRequest(op, input, &UpdateSmsChannelOutput{})
	return UpdateSmsChannelRequest{Request: req, Input: input, Copy: c.UpdateSmsChannelRequest}
}

// UpdateSmsChannelRequest is the request type for the
// UpdateSmsChannel API operation.
type UpdateSmsChannelRequest struct {
	*aws.Request
	Input *UpdateSmsChannelInput
	Copy  func(*UpdateSmsChannelInput) UpdateSmsChannelRequest
}

// Send marshals and sends the UpdateSmsChannel API request.
func (r UpdateSmsChannelRequest) Send(ctx context.Context) (*UpdateSmsChannelResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSmsChannelResponse{
		UpdateSmsChannelOutput: r.Request.Data.(*UpdateSmsChannelOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSmsChannelResponse is the response type for the
// UpdateSmsChannel API operation.
type UpdateSmsChannelResponse struct {
	*UpdateSmsChannelOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSmsChannel request.
func (r *UpdateSmsChannelResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}