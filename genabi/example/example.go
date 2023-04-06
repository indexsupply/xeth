// View functions and log parsing from example.json
//
// Code generated by "genabi"; DO NOT EDIT.
package example

import (
	"github.com/indexsupply/x/abi"
	"github.com/indexsupply/x/abi/schema"
	"github.com/indexsupply/x/jrpc"
	"math/big"
)

type AccountQueryRequest struct {
	item  *abi.Item
	Addrs [][20]byte
}

func (x AccountQueryRequest) Done() {
	x.item.Done()
}

func DecodeAccountQueryRequest(item *abi.Item) AccountQueryRequest {
	x := AccountQueryRequest{}
	x.item = item
	var (
		addrsItem0 = item.At(0)
		addrs0     = make([][20]byte, addrsItem0.Len())
	)
	for i0 := 0; i0 < addrsItem0.Len(); i0++ {
		addrs0[i0] = addrsItem0.At(i0).Address()
	}
	x.Addrs = addrs0
	return x
}

func (x AccountQueryRequest) Encode() *abi.Item {
	items := make([]*abi.Item, 1)
	var (
		addrs0      = x.Addrs
		addrsItems0 = make([]*abi.Item, len(addrs0))
	)
	for i0 := 0; i0 < len(addrs0); i0++ {
		addrsItems0[i0] = abi.Address(addrs0[i0])
	}
	items[0] = abi.Array(addrsItems0...)
	return abi.Tuple(items...)
}

type AccountQueryResponse struct {
	item    *abi.Item
	Account []AccountQueryResponseAccount
}

func (x AccountQueryResponse) Done() {
	x.item.Done()
}

func DecodeAccountQueryResponse(item *abi.Item) AccountQueryResponse {
	x := AccountQueryResponse{}
	x.item = item
	var (
		accountItem0 = item.At(0)
		account0     = make([]AccountQueryResponseAccount, accountItem0.Len())
	)
	for i0 := 0; i0 < accountItem0.Len(); i0++ {
		account0[i0] = DecodeAccountQueryResponseAccount(accountItem0.At(i0))
	}
	x.Account = account0
	return x
}

func (x AccountQueryResponse) Encode() *abi.Item {
	items := make([]*abi.Item, 1)
	var (
		account0      = x.Account
		accountItems0 = make([]*abi.Item, len(account0))
	)
	for i0 := 0; i0 < len(account0); i0++ {
		accountItems0[i0] = account0[i0].Encode()
	}
	items[0] = abi.Array(accountItems0...)
	return abi.Tuple(items...)
}

type AccountQueryResponseAccount struct {
	item    *abi.Item
	Id      uint16
	Balance *big.Int
}

func (x AccountQueryResponseAccount) Done() {
	x.item.Done()
}

func DecodeAccountQueryResponseAccount(item *abi.Item) AccountQueryResponseAccount {
	x := AccountQueryResponseAccount{}
	x.item = item
	x.Id = item.At(0).Uint16()
	x.Balance = item.At(1).BigInt()
	return x
}

func (x AccountQueryResponseAccount) Encode() *abi.Item {
	items := make([]*abi.Item, 2)
	items[0] = abi.Uint16(x.Id)
	items[1] = abi.BigInt(x.Balance)
	return abi.Tuple(items...)
}

var (
	accountQueryRequestSignature = [32]byte{0x60, 0xcc, 0x7a, 0x74, 0xcf, 0x2a, 0x95, 0x76, 0x90, 0xe6, 0x9a, 0x29, 0x20, 0x16, 0xcd, 0x12, 0x30, 0x80, 0x84, 0x1, 0xfb, 0x30, 0xad, 0x62, 0xb2, 0x87, 0x7e, 0x53, 0xee, 0xcb, 0xed, 0xf2}
	accountQueryResponseSchema   = schema.Parse("((uint16,uint256)[])")
)

func CallAccountQuery(c *jrpc.Client, contract [20]byte, req AccountQueryRequest) (AccountQueryResponse, error) {
	var (
		s4 = accountQueryRequestSignature[:4]
		cd = append(s4, abi.Encode(req.Encode())...)
	)
	respData, err := c.EthCall(contract, cd)
	if err != nil {
		return AccountQueryResponse{}, err
	}
	respItem, _, err := abi.Decode(respData, accountQueryResponseSchema)
	defer respItem.Done()
	if err != nil {
		return AccountQueryResponse{}, err
	}
	return DecodeAccountQueryResponse(respItem), nil
}

type NestedSlicesEvent struct {
	item    *abi.Item
	Strings []string
}

func (x NestedSlicesEvent) Done() {
	x.item.Done()
}

func DecodeNestedSlicesEvent(item *abi.Item) NestedSlicesEvent {
	x := NestedSlicesEvent{}
	x.item = item
	var (
		stringsItem0 = item.At(0)
		strings0     = make([]string, stringsItem0.Len())
	)
	for i0 := 0; i0 < stringsItem0.Len(); i0++ {
		strings0[i0] = stringsItem0.At(i0).String()
	}
	x.Strings = strings0
	return x
}

func (x NestedSlicesEvent) Encode() *abi.Item {
	items := make([]*abi.Item, 1)
	var (
		strings0      = x.Strings
		stringsItems0 = make([]*abi.Item, len(strings0))
	)
	for i0 := 0; i0 < len(strings0); i0++ {
		stringsItems0[i0] = abi.String(strings0[i0])
	}
	items[0] = abi.Array(stringsItems0...)
	return abi.Tuple(items...)
}

var (
	nestedSlicesSignature  = [32]byte{0xee, 0x41, 0x3f, 0x81, 0xe1, 0x39, 0xc0, 0xfa, 0xea, 0xfa, 0xeb, 0xcd, 0x24, 0x55, 0x1f, 0x44, 0x4, 0x1a, 0x81, 0x69, 0xd6, 0x6c, 0x58, 0x8c, 0xe7, 0x71, 0x24, 0x99, 0xda, 0xf4, 0x96, 0x13}
	nestedSlicesSchema     = schema.Parse("(string[])")
	nestedSlicesNumIndexed = int(0)
)

// Event Signature:
//	nestedSlices(string[])
// Checks the first log topic against the signature hash:
//	ee413f81e139c0faeafaebcd24551f44041a8169d66c588ce7712499daf49613
//
// Copies indexed event inputs from the remaining topics
// into [NestedSlices]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [NestedSlices]:
//	(string[])
func MatchNestedSlices(l abi.Log) (NestedSlicesEvent, error) {
	if len(l.Topics) == 0 {
		return NestedSlicesEvent{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && nestedSlicesSignature != l.Topics[0] {
		return NestedSlicesEvent{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != nestedSlicesNumIndexed {
		return NestedSlicesEvent{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, nestedSlicesSchema)
	if err != nil {
		return NestedSlicesEvent{}, err
	}
	res := DecodeNestedSlicesEvent(item)
	return res, nil
}

type TransferEvent struct {
	item    *abi.Item
	From    [20]byte
	To      [20]byte
	Id      *big.Int
	Extra   [3][2]uint8
	Details [][]TransferEventDetails
}

func (x TransferEvent) Done() {
	x.item.Done()
}

func DecodeTransferEvent(item *abi.Item) TransferEvent {
	x := TransferEvent{}
	x.item = item
	var (
		extraItem0 = item.At(0)
		extra0     = [3][2]uint8{}
	)
	for i0 := 0; i0 < extraItem0.Len(); i0++ {
		var (
			extraItem1 = extraItem0.At(i0)
			extra1     = [2]uint8{}
		)
		for i1 := 0; i1 < extraItem1.Len(); i1++ {
			extra1[i1] = extraItem1.At(i1).Uint8()
		}
		extra0[i0] = extra1
	}
	x.Extra = extra0
	var (
		detailsItem0 = item.At(1)
		details0     = make([][]TransferEventDetails, detailsItem0.Len())
	)
	for i0 := 0; i0 < detailsItem0.Len(); i0++ {
		var (
			detailsItem1 = detailsItem0.At(i0)
			details1     = make([]TransferEventDetails, detailsItem1.Len())
		)
		for i1 := 0; i1 < detailsItem1.Len(); i1++ {
			details1[i1] = DecodeTransferEventDetails(detailsItem1.At(i1))
		}
		details0[i0] = details1
	}
	x.Details = details0
	return x
}

func (x TransferEvent) Encode() *abi.Item {
	items := make([]*abi.Item, 5)
	items[0] = abi.Address(x.From)
	items[1] = abi.Address(x.To)
	items[2] = abi.BigInt(x.Id)
	var (
		extra0      = x.Extra
		extraItems0 = make([]*abi.Item, len(extra0))
	)
	for i0 := 0; i0 < len(extra0); i0++ {
		var (
			extra1      = extra0[i0]
			extraItems1 = make([]*abi.Item, len(extra1))
		)
		for i1 := 0; i1 < len(extra1); i1++ {
			extraItems1[i1] = abi.Uint8(extra1[i1])
		}

		extraItems0[i0] = abi.Array(extraItems0...)
	}
	items[3] = abi.Array(extraItems0...)
	var (
		details0      = x.Details
		detailsItems0 = make([]*abi.Item, len(details0))
	)
	for i0 := 0; i0 < len(details0); i0++ {
		var (
			details1      = details0[i0]
			detailsItems1 = make([]*abi.Item, len(details1))
		)
		for i1 := 0; i1 < len(details1); i1++ {
			detailsItems1[i1] = details1[i1].Encode()
		}

		detailsItems0[i0] = abi.Array(detailsItems0...)
	}
	items[4] = abi.Array(detailsItems0...)
	return abi.Tuple(items...)
}

type TransferEventDetails struct {
	item  *abi.Item
	Other [20]byte
	Key   [32]byte
	Value []byte
	Geo   TransferEventDetailsGeo
}

func (x TransferEventDetails) Done() {
	x.item.Done()
}

func DecodeTransferEventDetails(item *abi.Item) TransferEventDetails {
	x := TransferEventDetails{}
	x.item = item
	x.Other = item.At(0).Address()
	x.Key = item.At(1).Bytes32()
	x.Value = item.At(2).Bytes()
	x.Geo = DecodeTransferEventDetailsGeo(item.At(3))
	return x
}

func (x TransferEventDetails) Encode() *abi.Item {
	items := make([]*abi.Item, 4)
	items[0] = abi.Address(x.Other)
	items[1] = abi.Bytes32(x.Key)
	items[2] = abi.Bytes(x.Value)
	items[3] = x.Geo.Encode()
	return abi.Tuple(items...)
}

type TransferEventDetailsGeo struct {
	item *abi.Item
	X    uint8
	Y    uint8
}

func (x TransferEventDetailsGeo) Done() {
	x.item.Done()
}

func DecodeTransferEventDetailsGeo(item *abi.Item) TransferEventDetailsGeo {
	x := TransferEventDetailsGeo{}
	x.item = item
	x.X = item.At(0).Uint8()
	x.Y = item.At(1).Uint8()
	return x
}

func (x TransferEventDetailsGeo) Encode() *abi.Item {
	items := make([]*abi.Item, 2)
	items[0] = abi.Uint8(x.X)
	items[1] = abi.Uint8(x.Y)
	return abi.Tuple(items...)
}

var (
	transferSignature  = [32]byte{0x70, 0x71, 0x1f, 0x9e, 0xfd, 0x2d, 0x56, 0x86, 0x65, 0x59, 0x2c, 0x1d, 0x62, 0x45, 0xe8, 0x92, 0xea, 0xb7, 0xd9, 0xe5, 0x6c, 0x76, 0x71, 0x46, 0x82, 0x52, 0x60, 0x66, 0xca, 0x69, 0xd6, 0x5e}
	transferSchema     = schema.Parse("(uint8[2][3],(address,bytes32,bytes,(uint8,uint8))[][])")
	transferNumIndexed = int(3)
)

// Event Signature:
//	transfer(address,address,uint256,uint8[2][3],(address,bytes32,bytes,(uint8,uint8))[][])
// Checks the first log topic against the signature hash:
//	70711f9efd2d568665592c1d6245e892eab7d9e56c76714682526066ca69d65e
//
// Copies indexed event inputs from the remaining topics
// into [Transfer]
//
// Uses the the following abi schema to decode the un-indexed
// event inputs from the log's data field into [Transfer]:
//	(uint8[2][3],(address,bytes32,bytes,(uint8,uint8))[][])
func MatchTransfer(l abi.Log) (TransferEvent, error) {
	if len(l.Topics) == 0 {
		return TransferEvent{}, abi.NoTopics
	}
	if len(l.Topics) > 0 && transferSignature != l.Topics[0] {
		return TransferEvent{}, abi.SigMismatch
	}
	if len(l.Topics[1:]) != transferNumIndexed {
		return TransferEvent{}, abi.IndexMismatch
	}
	item, _, err := abi.Decode(l.Data, transferSchema)
	if err != nil {
		return TransferEvent{}, err
	}
	res := DecodeTransferEvent(item)
	res.From = abi.Bytes(l.Topics[1][:]).Address()
	res.To = abi.Bytes(l.Topics[2][:]).Address()
	res.Id = abi.Bytes(l.Topics[3][:]).BigInt()
	return res, nil
}
