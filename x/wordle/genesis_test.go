package wordle_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "wordle/testutil/keeper"
	"wordle/testutil/nullify"
	"wordle/x/wordle"
	"wordle/x/wordle/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		WordleList: []types.Wordle{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		GuessList: []types.Guess{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WordleKeeper(t)
	wordle.InitGenesis(ctx, *k, genesisState)
	got := wordle.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.WordleList, got.WordleList)
	require.ElementsMatch(t, genesisState.GuessList, got.GuessList)
	// this line is used by starport scaffolding # genesis/test/assert
}
