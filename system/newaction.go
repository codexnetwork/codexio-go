package system

import (
	eos "github.com/codexnetwork/codexio-go"
)

// NewPropose returns a `propose` action that lives on the
// `eosio.msig` contract.
func NewAction(account_name eos.AccountName, action_name eos.ActionName, requested []eos.PermissionLevel,data string) *eos.Action {
	return &eos.Action{
		Account: account_name,
		Name:    action_name,
		Authorization:requested,
		ActionData: eos.NewActionData(data),
	}
}

