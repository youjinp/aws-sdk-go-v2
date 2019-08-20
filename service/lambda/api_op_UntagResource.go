// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UntagResourceRequest
type UntagResourceInput struct {
	_ struct{} `type:"structure"`

	// The function's Amazon Resource Name (ARN).
	//
	// Resource is a required field
	Resource *string `location:"uri" locationName:"ARN" type:"string" required:"true"`

	// A list of tag keys to remove from the function.
	//
	// TagKeys is a required field
	TagKeys []string `location:"querystring" locationName:"tagKeys" type:"list" required:"true"`
}

// String returns the string representation
func (s UntagResourceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagResourceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UntagResourceInput"}

	if s.Resource == nil {
		invalidParams.Add(aws.NewErrParamRequired("Resource"))
	}

	if s.TagKeys == nil {
		invalidParams.Add(aws.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagResourceInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Resource != nil {
		v := *s.Resource

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ARN", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TagKeys != nil {
		v := s.TagKeys

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.QueryTarget, "tagKeys", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UntagResourceOutput
type UntagResourceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagResourceOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UntagResourceOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opUntagResource = "UntagResource"

// UntagResourceRequest returns a request value for making API operation for
// AWS Lambda.
//
// Removes tags (https://docs.aws.amazon.com/lambda/latest/dg/tagging.html)
// from a function.
//
//    // Example sending a request using UntagResourceRequest.
//    req := client.UntagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lambda-2015-03-31/UntagResource
func (c *Client) UntagResourceRequest(input *UntagResourceInput) UntagResourceRequest {
	op := &aws.Operation{
		Name:       opUntagResource,
		HTTPMethod: "DELETE",
		HTTPPath:   "/2017-03-31/tags/{ARN}",
	}

	if input == nil {
		input = &UntagResourceInput{}
	}

	req := c.newRequest(op, input, &UntagResourceOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UntagResourceRequest{Request: req, Input: input, Copy: c.UntagResourceRequest}
}

// UntagResourceRequest is the request type for the
// UntagResource API operation.
type UntagResourceRequest struct {
	*aws.Request
	Input *UntagResourceInput
	Copy  func(*UntagResourceInput) UntagResourceRequest
}

// Send marshals and sends the UntagResource API request.
func (r UntagResourceRequest) Send(ctx context.Context) (*UntagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UntagResourceResponse{
		UntagResourceOutput: r.Request.Data.(*UntagResourceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UntagResourceResponse is the response type for the
// UntagResource API operation.
type UntagResourceResponse struct {
	*UntagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UntagResource request.
func (r *UntagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}