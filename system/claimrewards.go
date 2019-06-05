package system

import (
	eos "github.com/codexnetwork/codexio-go"
)

// NewClaimRewards will buy at current market price a given number of
// bytes of RAM, and grant them to the `receiver` account.
func NewClaimRewards(voterName eos.AccountName,bpName eos.AccountName) *eos.Action {
	a := &eos.Action{
		Account: AN("eosio"),
		Name:    ActN("claim"),
		Authorization: []eos.PermissionLevel{
			{Actor: voterName, Permission: eos.PermissionName("active")},
		},
		ActionData: eos.NewActionData(ClaimRewards{
			VoterName: voterName,
			BpName:		bpName,
		}),
	}
	return a
}

// ClaimRewards represents the `eosio.system::claimrewards` action.
type ClaimRewards struct {
	VoterName eos.AccountName `json:"voter"`
	BpName		eos.AccountName `json:"bp"`
}
