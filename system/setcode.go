package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	eos "github.com/codexnetwork/codexio-go"
)

func NewSetContract(account eos.AccountName, wasmPath, abiPath string) (out []*eos.Action, err error) {
	codeAction, err := NewSetCode(account, wasmPath)
	if err != nil {
		return nil, err
	}

	abiAction, err := NewSetABI(account, abiPath)
	if err != nil {
		return nil, err
	}

	return []*eos.Action{codeAction, abiAction}, nil
}

func NewSetCode(account eos.AccountName, wasmPath string) (out *eos.Action, err error) {
	codeContent, err := ioutil.ReadFile(wasmPath)
	if err != nil {
		return nil, err
	}

	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setcode"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      account,
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(SetCode{
			Account:   account,
			VMType:    0,
			VMVersion: 0,
			Code:      eos.HexBytes(codeContent),
		}),
	}, nil
}

func NewSetABI(account eos.AccountName, abiPath string) (out *eos.Action, err error) {
	abiContent, err := ioutil.ReadFile(abiPath)
	if err != nil {
		return nil, err
	}

	var abiPacked []byte
	if len(abiContent) > 0 {
		var abiDef eos.ABI
		if err := json.Unmarshal(abiContent, &abiDef); err != nil {
			return nil, fmt.Errorf("unmarshal ABI file: %s", err)
		}

		abiPacked, err = eos.MarshalBinary(abiDef)
		if err != nil {
			return nil, fmt.Errorf("packing ABI: %s", err)
		}
	}

	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("setabi"),
		Authorization: []eos.PermissionLevel{
			{
				Actor:      account,
				Permission: eos.PermissionName("active"),
			},
		},
		ActionData: eos.NewActionData(SetABI{
			Account: account,
			ABI:     eos.HexBytes(abiPacked),
		}),
	}, nil
}

//
func NewSetPermission(account eos.AccountName,auth eos.Authority, permission eos.PermissionName,parent eos.PermissionName) (out *eos.Action, err error) {
	return &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("updateauth"),
		Authorization: []eos.PermissionLevel{
			{account,  parent},
		},
		ActionData: eos.NewActionData(SetAuth{
			Account: account,
			Permission:permission,
			Parent:parent,
			Auth:auth,
		}),
	}, nil
}

// NewSetCodeTx is _deprecated_. Use NewSetContract instead, and build
// your transaction yourself.
func NewSetCodeTx(account eos.AccountName, wasmPath, abiPath string) (out *eos.Transaction, err error) {
	actions, err := NewSetContract(account, wasmPath, abiPath)
	if err != nil {
		return nil, err
	}
	return &eos.Transaction{Actions: actions}, nil
}

// SetCode represents the hard-coded `setcode` action.
type SetCode struct {
	Account   eos.AccountName `json:"account"`
	VMType    byte            `json:"vmtype"`
	VMVersion byte            `json:"vmversion"`
	Code      eos.HexBytes    `json:"code"`
}

// SetABI represents the hard-coded `setabi` action.
type SetABI struct {
	Account eos.AccountName `json:"account"`
	ABI     eos.HexBytes    `json:"abi"`
}

type SetAuth struct {
	Account eos.AccountName `json:"account"`
	Permission eos.PermissionName `json:"perssion"`
	Parent eos.PermissionName `json:"parent"`
	Auth eos.Authority `json:"auth"`
}