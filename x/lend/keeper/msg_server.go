package keeper

import (
	"context"
	"github.com/comdex-official/comdex/x/lend/types"
)

type msgServer struct {
	Keeper
}

func (m msgServer) Lend(ctx context.Context, lend *types.MsgLend) (*types.MsgLendResponse, error) {
	panic("implement me")
}

func (m msgServer) Withdraw(ctx context.Context, withdraw *types.MsgWithdraw) (*types.MsgWithdrawResponse, error) {
	panic("implement me")
}

func (m msgServer) Borrow(ctx context.Context, borrow *types.MsgBorrow) (*types.MsgBorrowResponse, error) {
	panic("implement me")
}

func (m msgServer) Repay(ctx context.Context, repay *types.MsgRepay) (*types.MsgRepayResponse, error) {
	panic("implement me")
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
