// Copyright (C) 2019, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"net/http"
	"time"

	"github.com/ava-labs/avalanchego/ids"
	"github.com/ava-labs/avalanchego/utils/formatting"
	"github.com/ava-labs/avalanchego/utils/json"
)

type service struct {
	index *index
}

type FormattedContainer struct {
	ID        ids.ID              `json:"id"`
	Bytes     string              `json:"bytes"`
	Timestamp time.Time           `json:"timestamp"`
	Encoding  formatting.Encoding `json:"encoding"`
	Index     json.Uint64         `json:"index"`
}

func newFormattedContainer(c Container, index uint64, enc formatting.Encoding) (FormattedContainer, error) {
	_ = "STUB: not implemented"
	return *new(FormattedContainer), nil
}

type GetLastAcceptedArgs struct {
	Encoding formatting.Encoding `json:"encoding"`
}

func (s *service) GetLastAccepted(_ *http.Request, args *GetLastAcceptedArgs, reply *FormattedContainer) error {
	_ = "STUB: not implemented"
	return nil
}

type GetContainerByIndexArgs struct {
	Index    json.Uint64         `json:"index"`
	Encoding formatting.Encoding `json:"encoding"`
}

func (s *service) GetContainerByIndex(_ *http.Request, args *GetContainerByIndexArgs, reply *FormattedContainer) error {
	_ = "STUB: not implemented"
	return nil
}

type GetContainerRangeArgs struct {
	StartIndex json.Uint64         `json:"startIndex"`
	NumToFetch json.Uint64         `json:"numToFetch"`
	Encoding   formatting.Encoding `json:"encoding"`
}

type GetContainerRangeResponse struct {
	Containers []FormattedContainer `json:"containers"`
}

// GetContainerRange returns the transactions at index [startIndex], [startIndex+1], ... , [startIndex+n-1]
// If [n] == 0, returns an empty response (i.e. null).
// If [startIndex] > the last accepted index, returns an error (unless the above apply.)
// If [n] > [MaxFetchedByRange], returns an error.
// If we run out of transactions, returns the ones fetched before running out.
func (s *service) GetContainerRange(_ *http.Request, args *GetContainerRangeArgs, reply *GetContainerRangeResponse) error {
	_ = "STUB: not implemented"
	return nil
}

type GetIndexArgs struct {
	ID ids.ID `json:"id"`
}

type GetIndexResponse struct {
	Index json.Uint64 `json:"index"`
}

func (s *service) GetIndex(_ *http.Request, args *GetIndexArgs, reply *GetIndexResponse) error {
	_ = "STUB: not implemented"
	return nil
}

type IsAcceptedArgs struct {
	ID ids.ID `json:"id"`
}

type IsAcceptedResponse struct {
	IsAccepted bool `json:"isAccepted"`
}

func (s *service) IsAccepted(_ *http.Request, args *IsAcceptedArgs, reply *IsAcceptedResponse) error {
	_ = "STUB: not implemented"
	return nil
}

type GetContainerByIDArgs struct {
	ID       ids.ID              `json:"id"`
	Encoding formatting.Encoding `json:"encoding"`
}

func (s *service) GetContainerByID(_ *http.Request, args *GetContainerByIDArgs, reply *FormattedContainer) error {
	_ = "STUB: not implemented"
	return nil
}
