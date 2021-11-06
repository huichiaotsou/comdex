package bandoracle

import (
	"github.com/bandprotocol/bandchain-packet/obi"
	"github.com/bandprotocol/bandchain-packet/packet"
	"github.com/comdex-official/comdex/x/bandoracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
)

// handleOraclePacket handles the result of the received BandChain oracles
// packet and saves the data into the KV database
func (am AppModule) handleOraclePacket(
	ctx sdk.Context,
	modulePacket channeltypes.Packet,
) (channeltypes.Acknowledgement, error) {
	var ack channeltypes.Acknowledgement
	var modulePacketData packet.OracleResponsePacketData
	if err := types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &modulePacketData); err != nil {
		return ack, nil
	}

	/*id, err := strconv.ParseUint(modulePacketData.ClientID, 10, 64)
	if err != nil {
		return channeltypes.Acknowledgement{}, err
	}

	if modulePacketData.ResolveStatus == packet.RESOLVE_STATUS_SUCCESS {
		calldata, found := am.keeper.GetCalldata(ctx, id)
		if !found {
			return channeltypes.Acknowledgement{}, fmt.Errorf("calldata does not exist for id %d", id)
		}

		var result types.Result
		if err := obi.Decode(modulePacketData.Result, &result); err != nil {
			return channeltypes.Acknowledgement{}, err
		}
		am.keeper.SetPriceForMarket(ctx, calldata.Symbols, result.Rates)
	}

	am.keeper.DeleteCalldata(ctx, id)
	ack = channeltypes.NewResultAcknowledgement(
		types.ModuleCdc.MustMarshalJSON(
			packet.NewOracleRequestPacketAcknowledgement(modulePacketData.RequestID),
		),
	)*/
	var goldPriceResult types.GoldPriceResult
	if err := obi.Decode(modulePacketData.Result, &goldPriceResult); err != nil {
		ack = channeltypes.NewErrorAcknowledgement(err.Error())
		return ack, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest,
			"cannot decode the goldPrice received packet")
	}
	am.keeper.SetGoldPriceResult(ctx, types.OracleRequestID(modulePacketData.RequestID), goldPriceResult)
	return ack, nil
}

// handleOracleAcknowledgment handles the acknowledgment result from the BandChain
// request and saves the request-id into the KV database
func (am AppModule) handleOracleAcknowledgment(
	ctx sdk.Context,
	ack channeltypes.Acknowledgement,
	modulePacket channeltypes.Packet,
) (*sdk.Result, error) {
	switch resp := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Result:
		var oracleAck packet.OracleRequestPacketAcknowledgement
		err := types.ModuleCdc.UnmarshalJSON(resp.Result, &oracleAck)
		if err != nil {
			return nil, nil
		}

		var data packet.OracleRequestPacketData
		if err = types.ModuleCdc.UnmarshalJSON(modulePacket.GetData(), &data); err != nil {
			return nil, nil
		}
		//requestID := types.CalldataKey(oracleAck.RequestID)

		/*var callData types.Calldata
		if err = obi.Decode(data.GetCalldata(), &callData); err != nil {
			return nil, sdkerrors.Wrap(err,
				"cannot decode the coinRates oracle acknowledgment packet")
		}
		am.keeper.SetCalldata(ctx, oracleAck.RequestID, callData)
		return &sdk.Result{}, nil*/

		requestID := types.OracleRequestID(oracleAck.RequestID)

		if requestID == 0 {
			return nil, sdkerrors.Wrap(err,
				"request id 0")
		}

		switch data.GetClientID() {

		case types.GoldPriceClientIDKey:
			var goldPriceData types.GoldPriceCallData
			if err = obi.Decode(data.GetCalldata(), &goldPriceData); err != nil {
				return nil, sdkerrors.Wrap(err,
					"cannot decode the goldPrice oracle acknowledgment packet")
			}
			am.keeper.SetLastGoldPriceID(ctx, requestID)
			return &sdk.Result{}, nil
			// this line is used by starport scaffolding # oracle/module/ack

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
				"oracle acknowledgment packet not found: %s", data.GetClientID())
		}

	case *channeltypes.Acknowledgement_Error:
		return nil, sdkerrors.Wrapf(sdkerrors.Error{},
			"cannot decode the goldPrice oracle acknowledgment packet")
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal,
			"oracle acknowledgment packet not found: %s")

	}
	return nil, nil
}
