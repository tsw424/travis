package governance

import (
	"github.com/CyberMiles/travis/sdk"
	"github.com/ethereum/go-ethereum/common"
)

// Tx
//--------------------------------------------------------------------------------

// register the tx type with its validation logic
// make sure to use the name of the handler as the prefix in the tx type,
// so it gets routed properly
const (
	ByteTxPropose      = 0xA1
	ByteTxVote         = 0xA2
	TypeTransferTxPropose      = governanceModuleName + "/propose/transfer"
	TypeTxVote         = governanceModuleName + "/vote"
)

func init() {
	sdk.TxMapper.RegisterImplementation(TransferTxPropose{}, TypeTransferTxPropose, ByteTxPropose)
	sdk.TxMapper.RegisterImplementation(TxVote{}, TypeTxVote, ByteTxVote)
}

//Verify interface at compile time
var _, _ sdk.TxInner = &TransferTxPropose{}, &TxVote{}

type TransferTxPropose struct {
	Proposer     *common.Address   `json:"proposer"`
	From         *common.Address   `json:"from"`
	To           *common.Address   `json:"to"`
	Amount       string            `json:"amount"`
	Reason       string            `json:"reason"`
	Expire       uint64	           `json:"expire"`
}

func (tx TransferTxPropose) ValidateBasic() error {
	return nil
}

func NewTransferTxPropose(proposer *common.Address, fromAddr *common.Address, toAddr *common.Address, amount string, reason string, expire uint64) sdk.Tx {
	return TransferTxPropose{
		proposer,
		fromAddr,
		toAddr,
		amount,
		reason,
		expire,
	}.Wrap()
}

func (tx TransferTxPropose) Wrap() sdk.Tx { return sdk.Tx{tx} }

type TxVote struct {
	ProposalId       string            `json:"proposal_id"`
	Voter            common.Address    `json:"voter"`
	Answer           string            `json:"answer"`
}

func (tx TxVote) ValidateBasic() error {
	return nil
}

func NewTxVote(pid string, voter common.Address, answer string) sdk.Tx {
	return TxVote{
		pid,
		voter,
		answer,
	}.Wrap()
}

func (tx TxVote) Wrap() sdk.Tx { return sdk.Tx{tx} }
