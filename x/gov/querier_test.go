package gov

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/ownesthq/cosmos-sdk/codec"
	sdk "github.com/ownesthq/cosmos-sdk/types"
	"github.com/ownesthq/cosmos-sdk/x/gov/types"
)

const custom = "custom"

func getQueriedParams(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier) (DepositParams, VotingParams, TallyParams) {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryParams, ParamDeposit}, "/"),
		Data: []byte{},
	}

	bz, err := querier(ctx, []string{QueryParams, ParamDeposit}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var depositParams DepositParams
	err2 := cdc.UnmarshalJSON(bz, &depositParams)
	require.Nil(t, err2)

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryParams, ParamVoting}, "/"),
		Data: []byte{},
	}

	bz, err = querier(ctx, []string{QueryParams, ParamVoting}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var votingParams VotingParams
	err2 = cdc.UnmarshalJSON(bz, &votingParams)
	require.Nil(t, err2)

	query = abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryParams, ParamTallying}, "/"),
		Data: []byte{},
	}

	bz, err = querier(ctx, []string{QueryParams, ParamTallying}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var tallyParams TallyParams
	err2 = cdc.UnmarshalJSON(bz, &tallyParams)
	require.Nil(t, err2)

	return depositParams, votingParams, tallyParams
}

func getQueriedProposal(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) Proposal {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryProposal}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{QueryProposal}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var proposal Proposal
	err2 := cdc.UnmarshalJSON(bz, proposal)
	require.Nil(t, err2)
	return proposal
}

func getQueriedProposals(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, depositor, voter sdk.AccAddress, status ProposalStatus, limit uint64) []Proposal {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryProposals}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryProposalsParams(status, limit, voter, depositor)),
	}

	bz, err := querier(ctx, []string{QueryProposals}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var proposals Proposals
	err2 := cdc.UnmarshalJSON(bz, &proposals)
	require.Nil(t, err2)
	return proposals
}

func getQueriedDeposit(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64, depositor sdk.AccAddress) Deposit {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryDeposit}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryDepositParams(proposalID, depositor)),
	}

	bz, err := querier(ctx, []string{QueryDeposit}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var deposit Deposit
	err2 := cdc.UnmarshalJSON(bz, &deposit)
	require.Nil(t, err2)
	return deposit
}

func getQueriedDeposits(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) []Deposit {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryDeposits}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{QueryDeposits}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var deposits []Deposit
	err2 := cdc.UnmarshalJSON(bz, &deposits)
	require.Nil(t, err2)
	return deposits
}

func getQueriedVote(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64, voter sdk.AccAddress) Vote {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryVote}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryVoteParams(proposalID, voter)),
	}

	bz, err := querier(ctx, []string{QueryVote}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var vote Vote
	err2 := cdc.UnmarshalJSON(bz, &vote)
	require.Nil(t, err2)
	return vote
}

func getQueriedVotes(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) []Vote {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryVote}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{QueryVotes}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var votes []Vote
	err2 := cdc.UnmarshalJSON(bz, &votes)
	require.Nil(t, err2)
	return votes
}

func getQueriedTally(t *testing.T, ctx sdk.Context, cdc *codec.Codec, querier sdk.Querier, proposalID uint64) TallyResult {
	query := abci.RequestQuery{
		Path: strings.Join([]string{custom, QuerierRoute, QueryTally}, "/"),
		Data: cdc.MustMarshalJSON(NewQueryProposalParams(proposalID)),
	}

	bz, err := querier(ctx, []string{QueryTally}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	var tally TallyResult
	err2 := cdc.UnmarshalJSON(bz, &tally)
	require.Nil(t, err2)
	return tally
}

func TestQueryParams(t *testing.T) {
	cdc := codec.New()
	input := getMockApp(t, 1000, GenesisState{}, nil)
	querier := NewQuerier(input.keeper)

	header := abci.Header{Height: input.mApp.LastBlockHeight() + 1}
	input.mApp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := input.mApp.NewContext(false, abci.Header{})

	getQueriedParams(t, ctx, cdc, querier)
}

func TestQueries(t *testing.T) {
	cdc := codec.New()
	input := getMockApp(t, 1000, GenesisState{}, nil)
	querier := NewQuerier(input.keeper)
	handler := NewHandler(input.keeper)

	types.RegisterCodec(cdc)

	header := abci.Header{Height: input.mApp.LastBlockHeight() + 1}
	input.mApp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := input.mApp.NewContext(false, abci.Header{})

	depositParams, _, _ := getQueriedParams(t, ctx, cdc, querier)

	// input.addrs[0] proposes (and deposits) proposals #1 and #2
	res := handler(ctx, NewMsgSubmitProposal(testProposal(), sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)}, input.addrs[0]))
	var proposalID1 uint64
	require.True(t, res.IsOK())
	cdc.MustUnmarshalBinaryLengthPrefixed(res.Data, &proposalID1)

	res = handler(ctx, NewMsgSubmitProposal(testProposal(), sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 10000000)}, input.addrs[0]))
	var proposalID2 uint64
	require.True(t, res.IsOK())
	cdc.MustUnmarshalBinaryLengthPrefixed(res.Data, &proposalID2)

	// input.addrs[1] proposes (and deposits) proposals #3
	res = handler(ctx, NewMsgSubmitProposal(testProposal(), sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 1)}, input.addrs[1]))
	var proposalID3 uint64
	require.True(t, res.IsOK())
	cdc.MustUnmarshalBinaryLengthPrefixed(res.Data, &proposalID3)

	// input.addrs[1] deposits on proposals #2 & #3
	res = handler(ctx, NewMsgDeposit(input.addrs[1], proposalID2, depositParams.MinDeposit))
	res = handler(ctx, NewMsgDeposit(input.addrs[1], proposalID3, depositParams.MinDeposit))

	// check deposits on proposal1 match individual deposits
	deposits := getQueriedDeposits(t, ctx, cdc, querier, proposalID1)
	require.Len(t, deposits, 1)
	deposit := getQueriedDeposit(t, ctx, cdc, querier, proposalID1, input.addrs[0])
	require.Equal(t, deposit, deposits[0])

	// check deposits on proposal2 match individual deposits
	deposits = getQueriedDeposits(t, ctx, cdc, querier, proposalID2)
	require.Len(t, deposits, 2)
	deposit = getQueriedDeposit(t, ctx, cdc, querier, proposalID2, input.addrs[0])
	require.True(t, deposit.Equals(deposits[0]))
	deposit = getQueriedDeposit(t, ctx, cdc, querier, proposalID2, input.addrs[1])
	require.True(t, deposit.Equals(deposits[1]))

	// check deposits on proposal3 match individual deposits
	deposits = getQueriedDeposits(t, ctx, cdc, querier, proposalID3)
	require.Len(t, deposits, 1)
	deposit = getQueriedDeposit(t, ctx, cdc, querier, proposalID3, input.addrs[1])
	require.Equal(t, deposit, deposits[0])

	// Only proposal #1 should be in Deposit Period
	proposals := getQueriedProposals(t, ctx, cdc, querier, nil, nil, StatusDepositPeriod, 0)
	require.Len(t, proposals, 1)
	require.Equal(t, proposalID1, proposals[0].ProposalID)

	// Only proposals #2 and #3 should be in Voting Period
	proposals = getQueriedProposals(t, ctx, cdc, querier, nil, nil, StatusVotingPeriod, 0)
	require.Len(t, proposals, 2)
	require.Equal(t, proposalID2, proposals[0].ProposalID)
	require.Equal(t, proposalID3, proposals[1].ProposalID)

	// Addrs[0] votes on proposals #2 & #3
	require.True(t, handler(ctx, NewMsgVote(input.addrs[0], proposalID2, OptionYes)).IsOK())
	require.True(t, handler(ctx, NewMsgVote(input.addrs[0], proposalID3, OptionYes)).IsOK())

	// Addrs[1] votes on proposal #3
	handler(ctx, NewMsgVote(input.addrs[1], proposalID3, OptionYes))

	// Test query voted by input.addrs[0]
	proposals = getQueriedProposals(t, ctx, cdc, querier, nil, input.addrs[0], StatusNil, 0)
	require.Equal(t, proposalID2, (proposals[0]).ProposalID)
	require.Equal(t, proposalID3, (proposals[1]).ProposalID)

	// Test query votes on Proposal 2
	votes := getQueriedVotes(t, ctx, cdc, querier, proposalID2)
	require.Len(t, votes, 1)
	require.Equal(t, input.addrs[0], votes[0].Voter)

	vote := getQueriedVote(t, ctx, cdc, querier, proposalID2, input.addrs[0])
	require.Equal(t, vote, votes[0])

	// Test query votes on Proposal 3
	votes = getQueriedVotes(t, ctx, cdc, querier, proposalID3)
	require.Len(t, votes, 2)
	require.True(t, input.addrs[0].String() == votes[0].Voter.String())
	require.True(t, input.addrs[1].String() == votes[1].Voter.String())

	// Test proposals queries with filters

	// Test query all proposals
	proposals = getQueriedProposals(t, ctx, cdc, querier, nil, nil, StatusNil, 0)
	require.Equal(t, proposalID1, (proposals[0]).ProposalID)
	require.Equal(t, proposalID2, (proposals[1]).ProposalID)
	require.Equal(t, proposalID3, (proposals[2]).ProposalID)

	// Test query voted by input.addrs[1]
	proposals = getQueriedProposals(t, ctx, cdc, querier, nil, input.addrs[1], StatusNil, 0)
	require.Equal(t, proposalID3, (proposals[0]).ProposalID)

	// Test query deposited by input.addrs[0]
	proposals = getQueriedProposals(t, ctx, cdc, querier, input.addrs[0], nil, StatusNil, 0)
	require.Equal(t, proposalID1, (proposals[0]).ProposalID)

	// Test query deposited by addr2
	proposals = getQueriedProposals(t, ctx, cdc, querier, input.addrs[1], nil, StatusNil, 0)
	require.Equal(t, proposalID2, (proposals[0]).ProposalID)
	require.Equal(t, proposalID3, (proposals[1]).ProposalID)

	// Test query voted AND deposited by addr1
	proposals = getQueriedProposals(t, ctx, cdc, querier, input.addrs[0], input.addrs[0], StatusNil, 0)
	require.Equal(t, proposalID2, (proposals[0]).ProposalID)
}
