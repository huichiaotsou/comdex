package keeper

import (
	"context"
	bandobi "github.com/bandprotocol/bandchain-packet/obi"
	bandpacket "github.com/bandprotocol/bandchain-packet/packet"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	ibcchanneltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibchost "github.com/cosmos/ibc-go/modules/core/24-host"
	"time"

	"github.com/comdex-official/comdex/x/bandoracle/types"
)

var (
	_ types.MsgServiceServer = (*msgServer)(nil)
)

type msgServer struct {
	Keeper
}

func NewMsgServiceServer(keeper Keeper) types.MsgServiceServer {
	return &msgServer{
		Keeper: keeper,
	}
}

func (k *msgServer) MsgRemoveMarketForAsset(c context.Context, msg *types.MsgRemoveMarketForAssetRequest) (*types.MsgRemoveMarketForAssetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if msg.From != k.assetKeeper.Admin(ctx) {
		return nil, types.ErrorUnauthorized
	}

	if !k.HasMarketForAsset(ctx, msg.Id) {
		return nil, types.ErrorMarketForAssetDoesNotExist
	}

	k.DeleteMarketForAsset(ctx, msg.Id)
	return &types.MsgRemoveMarketForAssetResponse{}, nil
}

func (k *msgServer) MsgAddMarket(c context.Context, msg *types.MsgAddMarketRequest) (*types.MsgAddMarketResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if msg.From != k.assetKeeper.Admin(ctx) {
		return nil, types.ErrorUnauthorized
	}

	if k.HasMarket(ctx, msg.Symbol) {
		return nil, types.ErrorDuplicateMarket
	}

	var (
		market = types.Market{
			Symbol:   msg.Symbol,
			ScriptID: msg.ScriptID,
		}
	)

	k.SetMarket(ctx, market)
	return &types.MsgAddMarketResponse{}, nil
}

func (k *msgServer) MsgUpdateMarket(c context.Context, msg *types.MsgUpdateMarketRequest) (*types.MsgUpdateMarketResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if msg.From != k.assetKeeper.Admin(ctx) {
		return nil, types.ErrorUnauthorized
	}

	market, found := k.GetMarket(ctx, msg.Symbol)
	if !found {
		return nil, types.ErrorMarketDoesNotExist
	}

	if msg.ScriptID != 0 {
		market.ScriptID = msg.ScriptID
	}

	k.SetMarket(ctx, market)
	return &types.MsgUpdateMarketResponse{}, nil
}

func (k *msgServer) MsgFetchPrice(c context.Context, msg *types.MsgFetchPriceRequest) (*types.MsgFetchPriceResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	/*var (
		calldata = types.Calldata{
			Multiplier: k.OracleMultiplier(ctx),
		}
	)

	channel, found := k.channel.GetChannel(ctx, msg.SourcePort, msg.SourceChannel)
	if !found {
		return nil, ibcchanneltypes.ErrChannelNotFound
	}

	sequence, found := k.channel.GetNextSequenceSend(ctx, msg.SourcePort, msg.SourceChannel)
	if !found {
		return nil, ibcchanneltypes.ErrSequenceSendNotFound
	}

	capability, found := k.scoped.GetCapability(ctx, ibchost.ChannelCapabilityPath(msg.SourcePort, msg.SourceChannel))
	if !found {
		return nil, ibcchanneltypes.ErrChannelCapabilityNotFound
	}

	var (
		id     = k.GetCalldataID(ctx)
		packet = ibcchanneltypes.Packet{
			Sequence:           sequence,
			SourcePort:         msg.SourcePort,
			SourceChannel:      msg.SourceChannel,
			DestinationPort:    channel.GetCounterparty().GetPortID(),
			DestinationChannel: channel.GetCounterparty().GetChannelID(),
			Data: bandpacket.OracleRequestPacketData{
				ClientID:       fmt.Sprintf("%d", id),
				OracleScriptID: msg.ScriptID,
				Calldata:       bandobi.MustEncode(calldata),
				AskCount:       5,
				MinCount:       3,
				FeeLimit:       msg.FeeLimit,
				PrepareGas:     msg.PrepareGas,
				ExecuteGas:     msg.ExecuteGas,
			}.GetBytes(),
			TimeoutHeight:    msg.TimeoutHeight,
			TimeoutTimestamp: msg.TimeoutTimestamp,
		}
	)

	k.SetCalldataID(ctx, id+1)
	k.SetCalldata(ctx, id, calldata)

	if err := k.channel.SendPacket(ctx, capability, packet); err != nil {
		return nil, err
	}
	--------------------------------

	var (
		calldata = types.Calldata{
			Multiplier: k.OracleMultiplier(ctx),
		}
	)
	sourcePort := types.PortID
	sourceChannelEnd, found := k.channel.GetChannel(ctx, sourcePort, msg.SourceChannel)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port %s",
			msg.SourceChannel,
			sourcePort,
		)
	}
	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channel.GetNextSequenceSend(ctx, sourcePort, msg.SourceChannel)
	if !found {
		return nil, sdkerrors.Wrapf(
			ibcchanneltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, msg.SourceChannel)
	}

	channelCap, ok := k.scoped.GetCapability(ctx, ibchost.ChannelCapabilityPath(sourcePort, msg.SourceChannel))
	if !ok {
		return nil, sdkerrors.Wrap(ibcchanneltypes.ErrChannelCapabilityNotFound,
			"module does not own channel capability")
	}

	encodedCalldata := bandobi.MustEncode(calldata)
	packetData := bandpacket.NewOracleRequestPacketData(
		msg.ClientID,
		msg.ScriptID,
		encodedCalldata,
		3,
		5,
		msg.FeeLimit,
		msg.PrepareGas,
		msg.ExecuteGas,
	)

	err := k.channel.SendPacket(ctx, channelCap, ibcchanneltypes.NewPacket(
		packetData.GetBytes(),
		sequence,
		sourcePort,
		msg.SourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.NewHeight(0, 0),
		uint64(ctx.BlockTime().UnixNano()+int64(10*time.Minute)), // Arbitrary timestamp timeout for now
	))
	if err != nil {
		return nil, err
	}*/

	var (
		calldata = types.Calldata{
			Multiplier: k.OracleMultiplier(ctx),
		}
	)

	sourcePort := types.PortID
	sourceChannelEnd, found := k.channel.GetChannel(ctx, sourcePort, msg.SourceChannel)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port %s",
			msg.SourceChannel,
			sourcePort,
		)
	}
	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channel.GetNextSequenceSend(ctx, sourcePort, msg.SourceChannel)
	if !found {
		return nil, sdkerrors.Wrapf(
			ibcchanneltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, msg.SourceChannel)
	}

	channelCap, ok := k.scoped.GetCapability(ctx, ibchost.ChannelCapabilityPath(sourcePort, msg.SourceChannel))
	if !ok {
		return nil, sdkerrors.Wrap(ibcchanneltypes.ErrChannelCapabilityNotFound,
			"module does not own channel capability")
	}

	encodedCalldata := bandobi.MustEncode(calldata)
	packetData := bandpacket.NewOracleRequestPacketData(
		msg.ClientID,
		msg.OracleScriptID,
		encodedCalldata,
		1,
		1,
		msg.FeeLimit,
		msg.PrepareGas,
		msg.ExecuteGas,
	)

	err := k.channel.SendPacket(ctx, channelCap, ibcchanneltypes.Packet{
		Sequence: sequence,
		SourcePort: sourcePort,
		SourceChannel: msg.SourceChannel,
		DestinationPort: destinationPort,
		DestinationChannel: destinationChannel,
		Data: packetData.GetBytes(),
		TimeoutHeight: clienttypes.NewHeight(0, 0),
		TimeoutTimestamp: uint64(ctx.BlockTime().UnixNano()+int64(10*time.Minute)),
	},// Arbitrary timestamp timeout for now
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgFetchPriceResponse{}, nil
}
