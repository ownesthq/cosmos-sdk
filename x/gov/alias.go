// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/ownesthq/cosmos-sdk/x/gov/types
package gov

import (
	"github.com/ownesthq/cosmos-sdk/x/gov/types"
)

const (
	MaxDescriptionLength         = types.MaxDescriptionLength
	MaxTitleLength               = types.MaxTitleLength
	DefaultCodespace             = types.DefaultCodespace
	CodeUnknownProposal          = types.CodeUnknownProposal
	CodeInactiveProposal         = types.CodeInactiveProposal
	CodeAlreadyActiveProposal    = types.CodeAlreadyActiveProposal
	CodeAlreadyFinishedProposal  = types.CodeAlreadyFinishedProposal
	CodeAddressNotStaked         = types.CodeAddressNotStaked
	CodeInvalidContent           = types.CodeInvalidContent
	CodeInvalidProposalType      = types.CodeInvalidProposalType
	CodeInvalidVote              = types.CodeInvalidVote
	CodeInvalidGenesis           = types.CodeInvalidGenesis
	CodeInvalidProposalStatus    = types.CodeInvalidProposalStatus
	CodeProposalHandlerNotExists = types.CodeProposalHandlerNotExists
	ModuleName                   = types.ModuleName
	StoreKey                     = types.StoreKey
	RouterKey                    = types.RouterKey
	QuerierRoute                 = types.QuerierRoute
	DefaultParamspace            = types.DefaultParamspace
	TypeMsgDeposit               = types.TypeMsgDeposit
	TypeMsgVote                  = types.TypeMsgVote
	TypeMsgSubmitProposal        = types.TypeMsgSubmitProposal
	StatusNil                    = types.StatusNil
	StatusDepositPeriod          = types.StatusDepositPeriod
	StatusVotingPeriod           = types.StatusVotingPeriod
	StatusPassed                 = types.StatusPassed
	StatusRejected               = types.StatusRejected
	StatusFailed                 = types.StatusFailed
	ProposalTypeText             = types.ProposalTypeText
	ProposalTypeSoftwareUpgrade  = types.ProposalTypeSoftwareUpgrade
	QueryParams                  = types.QueryParams
	QueryProposals               = types.QueryProposals
	QueryProposal                = types.QueryProposal
	QueryDeposits                = types.QueryDeposits
	QueryDeposit                 = types.QueryDeposit
	QueryVotes                   = types.QueryVotes
	QueryVote                    = types.QueryVote
	QueryTally                   = types.QueryTally
	ParamDeposit                 = types.ParamDeposit
	ParamVoting                  = types.ParamVoting
	ParamTallying                = types.ParamTallying
	OptionEmpty                  = types.OptionEmpty
	OptionYes                    = types.OptionYes
	OptionAbstain                = types.OptionAbstain
	OptionNo                     = types.OptionNo
	OptionNoWithVeto             = types.OptionNoWithVeto
)

var (
	// functions aliases
	RegisterCodec                 = types.RegisterCodec
	RegisterProposalTypeCodec     = types.RegisterProposalTypeCodec
	ValidateAbstract              = types.ValidateAbstract
	NewDeposit                    = types.NewDeposit
	ErrUnknownProposal            = types.ErrUnknownProposal
	ErrInactiveProposal           = types.ErrInactiveProposal
	ErrAlreadyActiveProposal      = types.ErrAlreadyActiveProposal
	ErrAlreadyFinishedProposal    = types.ErrAlreadyFinishedProposal
	ErrAddressNotStaked           = types.ErrAddressNotStaked
	ErrInvalidProposalContent     = types.ErrInvalidProposalContent
	ErrInvalidProposalType        = types.ErrInvalidProposalType
	ErrInvalidVote                = types.ErrInvalidVote
	ErrInvalidGenesis             = types.ErrInvalidGenesis
	ErrNoProposalHandlerExists    = types.ErrNoProposalHandlerExists
	ProposalKey                   = types.ProposalKey
	ActiveProposalByTimeKey       = types.ActiveProposalByTimeKey
	ActiveProposalQueueKey        = types.ActiveProposalQueueKey
	InactiveProposalByTimeKey     = types.InactiveProposalByTimeKey
	InactiveProposalQueueKey      = types.InactiveProposalQueueKey
	DepositsKey                   = types.DepositsKey
	DepositKey                    = types.DepositKey
	VotesKey                      = types.VotesKey
	VoteKey                       = types.VoteKey
	SplitProposalKey              = types.SplitProposalKey
	SplitActiveProposalQueueKey   = types.SplitActiveProposalQueueKey
	SplitInactiveProposalQueueKey = types.SplitInactiveProposalQueueKey
	SplitKeyDeposit               = types.SplitKeyDeposit
	SplitKeyVote                  = types.SplitKeyVote
	NewMsgSubmitProposal          = types.NewMsgSubmitProposal
	NewMsgDeposit                 = types.NewMsgDeposit
	NewMsgVote                    = types.NewMsgVote
	ParamKeyTable                 = types.ParamKeyTable
	NewDepositParams              = types.NewDepositParams
	NewTallyParams                = types.NewTallyParams
	NewVotingParams               = types.NewVotingParams
	NewParams                     = types.NewParams
	NewProposal                   = types.NewProposal
	ProposalStatusFromString      = types.ProposalStatusFromString
	ValidProposalStatus           = types.ValidProposalStatus
	NewTallyResult                = types.NewTallyResult
	NewTallyResultFromMap         = types.NewTallyResultFromMap
	EmptyTallyResult              = types.EmptyTallyResult
	NewTextProposal               = types.NewTextProposal
	NewSoftwareUpgradeProposal    = types.NewSoftwareUpgradeProposal
	RegisterProposalType          = types.RegisterProposalType
	ContentFromProposalType       = types.ContentFromProposalType
	IsValidProposalType           = types.IsValidProposalType
	ProposalHandler               = types.ProposalHandler
	NewQueryProposalParams        = types.NewQueryProposalParams
	NewQueryDepositParams         = types.NewQueryDepositParams
	NewQueryVoteParams            = types.NewQueryVoteParams
	NewQueryProposalsParams       = types.NewQueryProposalsParams
	NewVote                       = types.NewVote
	VoteOptionFromString          = types.VoteOptionFromString
	ValidVoteOption               = types.ValidVoteOption

	// variable aliases
	ModuleCdc                   = types.ModuleCdc
	ProposalsKeyPrefix          = types.ProposalsKeyPrefix
	ActiveProposalQueuePrefix   = types.ActiveProposalQueuePrefix
	InactiveProposalQueuePrefix = types.InactiveProposalQueuePrefix
	ProposalIDKey               = types.ProposalIDKey
	DepositsKeyPrefix           = types.DepositsKeyPrefix
	VotesKeyPrefix              = types.VotesKeyPrefix
	ParamStoreKeyDepositParams  = types.ParamStoreKeyDepositParams
	ParamStoreKeyVotingParams   = types.ParamStoreKeyVotingParams
	ParamStoreKeyTallyParams    = types.ParamStoreKeyTallyParams
)

type (
	Content                 = types.Content
	Handler                 = types.Handler
	Deposit                 = types.Deposit
	Deposits                = types.Deposits
	MsgSubmitProposal       = types.MsgSubmitProposal
	MsgDeposit              = types.MsgDeposit
	MsgVote                 = types.MsgVote
	DepositParams           = types.DepositParams
	TallyParams             = types.TallyParams
	VotingParams            = types.VotingParams
	Params                  = types.Params
	Proposal                = types.Proposal
	Proposals               = types.Proposals
	ProposalQueue           = types.ProposalQueue
	ProposalStatus          = types.ProposalStatus
	TallyResult             = types.TallyResult
	TextProposal            = types.TextProposal
	SoftwareUpgradeProposal = types.SoftwareUpgradeProposal
	QueryProposalParams     = types.QueryProposalParams
	QueryDepositParams      = types.QueryDepositParams
	QueryVoteParams         = types.QueryVoteParams
	QueryProposalsParams    = types.QueryProposalsParams
	Vote                    = types.Vote
	Votes                   = types.Votes
	VoteOption              = types.VoteOption
)
