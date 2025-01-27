// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package node_reader

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

type CompletionStatus string

const (
	CompletionStatusUnprocessed                CompletionStatus = "UNPROCESSED"
	CompletionStatusAccepted                   CompletionStatus = "ACCEPTED"
	CompletionStatusRejected                   CompletionStatus = "REJECTED"
	CompletionStatusException                  CompletionStatus = "EXCEPTION"
	CompletionStatusMachineHalted              CompletionStatus = "MACHINE_HALTED"
	CompletionStatusCycleLimitExceeded         CompletionStatus = "CYCLE_LIMIT_EXCEEDED"
	CompletionStatusTimeLimitExceeded          CompletionStatus = "TIME_LIMIT_EXCEEDED"
	CompletionStatusPayloadLengthLimitExceeded CompletionStatus = "PAYLOAD_LENGTH_LIMIT_EXCEEDED"
)

// __getInputStatusInput is used internally by genqlient
type __getInputStatusInput struct {
	Index int `json:"index"`
}

// GetIndex returns __getInputStatusInput.Index, and is useful for accessing the field via an interface.
func (v *__getInputStatusInput) GetIndex() int { return v.Index }

// __getOutputsByInputIndexInput is used internally by genqlient
type __getOutputsByInputIndexInput struct {
	InputIndex int `json:"inputIndex"`
}

// GetInputIndex returns __getOutputsByInputIndexInput.InputIndex, and is useful for accessing the field via an interface.
func (v *__getOutputsByInputIndexInput) GetInputIndex() int { return v.InputIndex }

// getInputStatusInput includes the requested fields of the GraphQL type Input.
// The GraphQL type's documentation follows.
//
// Request submitted to the application to advance its state
type getInputStatusInput struct {
	// Status of the input
	Status CompletionStatus `json:"status"`
}

// GetStatus returns getInputStatusInput.Status, and is useful for accessing the field via an interface.
func (v *getInputStatusInput) GetStatus() CompletionStatus { return v.Status }

// getInputStatusResponse is returned by getInputStatus on success.
type getInputStatusResponse struct {
	// Get input based on its identifier
	Input getInputStatusInput `json:"input"`
}

// GetInput returns getInputStatusResponse.Input, and is useful for accessing the field via an interface.
func (v *getInputStatusResponse) GetInput() getInputStatusInput { return v.Input }

// getOutputsByInputIndexInput includes the requested fields of the GraphQL type Input.
// The GraphQL type's documentation follows.
//
// Request submitted to the application to advance its state
type getOutputsByInputIndexInput struct {
	// Get notices from this particular input with support for pagination
	Notices getOutputsByInputIndexInputNoticesNoticeConnection `json:"notices"`
}

// GetNotices returns getOutputsByInputIndexInput.Notices, and is useful for accessing the field via an interface.
func (v *getOutputsByInputIndexInput) GetNotices() getOutputsByInputIndexInputNoticesNoticeConnection {
	return v.Notices
}

// getOutputsByInputIndexInputNoticesNoticeConnection includes the requested fields of the GraphQL type NoticeConnection.
// The GraphQL type's documentation follows.
//
// Pagination result
type getOutputsByInputIndexInputNoticesNoticeConnection struct {
	// Pagination entries returned for the current page
	Edges []getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdge `json:"edges"`
}

// GetEdges returns getOutputsByInputIndexInputNoticesNoticeConnection.Edges, and is useful for accessing the field via an interface.
func (v *getOutputsByInputIndexInputNoticesNoticeConnection) GetEdges() []getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdge {
	return v.Edges
}

// getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdge includes the requested fields of the GraphQL type NoticeEdge.
// The GraphQL type's documentation follows.
//
// Pagination entry
type getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdge struct {
	// Node instance
	Node getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice `json:"node"`
}

// GetNode returns getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdge.Node, and is useful for accessing the field via an interface.
func (v *getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdge) GetNode() getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice {
	return v.Node
}

// getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice includes the requested fields of the GraphQL type Notice.
// The GraphQL type's documentation follows.
//
// Informational statement that can be validated in the base layer blockchain
type getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice struct {
	// Notice index within the context of the input that produced it
	Index int `json:"index"`
	// Notice data as a payload in Ethereum hex binary format, starting with '0x'
	Payload string `json:"payload"`
}

// GetIndex returns getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice.Index, and is useful for accessing the field via an interface.
func (v *getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice) GetIndex() int {
	return v.Index
}

// GetPayload returns getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice.Payload, and is useful for accessing the field via an interface.
func (v *getOutputsByInputIndexInputNoticesNoticeConnectionEdgesNoticeEdgeNodeNotice) GetPayload() string {
	return v.Payload
}

// getOutputsByInputIndexResponse is returned by getOutputsByInputIndex on success.
type getOutputsByInputIndexResponse struct {
	// Get input based on its identifier
	Input getOutputsByInputIndexInput `json:"input"`
}

// GetInput returns getOutputsByInputIndexResponse.Input, and is useful for accessing the field via an interface.
func (v *getOutputsByInputIndexResponse) GetInput() getOutputsByInputIndexInput { return v.Input }

// The query or mutation executed by getInputStatus.
const getInputStatus_Operation = `
query getInputStatus ($index: Int!) {
	input(index: $index) {
		status
	}
}
`

// Get the input status.
func getInputStatus(
	ctx_ context.Context,
	client_ graphql.Client,
	index int,
) (*getInputStatusResponse, error) {
	req_ := &graphql.Request{
		OpName: "getInputStatus",
		Query:  getInputStatus_Operation,
		Variables: &__getInputStatusInput{
			Index: index,
		},
	}
	var err_ error

	var data_ getInputStatusResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}

// The query or mutation executed by getOutputsByInputIndex.
const getOutputsByInputIndex_Operation = `
query getOutputsByInputIndex ($inputIndex: Int!) {
	input(index: $inputIndex) {
		notices {
			edges {
				node {
					index
					payload
				}
			}
		}
	}
}
`

func getOutputsByInputIndex(
	ctx_ context.Context,
	client_ graphql.Client,
	inputIndex int,
) (*getOutputsByInputIndexResponse, error) {
	req_ := &graphql.Request{
		OpName: "getOutputsByInputIndex",
		Query:  getOutputsByInputIndex_Operation,
		Variables: &__getOutputsByInputIndexInput{
			InputIndex: inputIndex,
		},
	}
	var err_ error

	var data_ getOutputsByInputIndexResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
