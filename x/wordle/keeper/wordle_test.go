package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "wordle/testutil/keeper"
	"wordle/testutil/nullify"
	"wordle/x/wordle/keeper"
	"wordle/x/wordle/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNWordle(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Wordle {
	items := make([]types.Wordle, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetWordle(ctx, items[i])
	}
	return items
}

func TestWordleGet(t *testing.T) {
	keeper, ctx := keepertest.WordleKeeper(t)
	items := createNWordle(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetWordle(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestWordleRemove(t *testing.T) {
	keeper, ctx := keepertest.WordleKeeper(t)
	items := createNWordle(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveWordle(ctx,
			item.Index,
		)
		_, found := keeper.GetWordle(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestWordleGetAll(t *testing.T) {
	keeper, ctx := keepertest.WordleKeeper(t)
	items := createNWordle(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllWordle(ctx)),
	)
}
