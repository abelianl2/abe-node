package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateWithdraw = "create_withdraw"
	TypeMsgUpdateWithdraw = "update_withdraw"
	TypeMsgDeleteWithdraw = "delete_withdraw"
)

var _ sdk.Msg = &MsgCreateWithdraw{}

func NewMsgCreateWithdraw(
	creator string,
	txHash string,
	from string,
	to string,
	coinType string,
	value uint64,
	data string,
	status string,
	signatures []string,

) *MsgCreateWithdraw {
	return &MsgCreateWithdraw{
		Creator:    creator,
		TxHash:     txHash,
		From:       from,
		To:         to,
		CoinType:   coinType,
		Value:      value,
		Data:       data,
		Status:     status,
		Signatures: signatures,
	}
}

func (msg *MsgCreateWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgCreateWithdraw) Type() string {
	return TypeMsgCreateWithdraw
}

func (msg *MsgCreateWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateWithdraw{}

func NewMsgUpdateWithdraw(
	creator string,
	txHash string,
	from string,
	to string,
	coinType string,
	value uint64,
	data string,
	status string,
	signatures []string,

) *MsgUpdateWithdraw {
	return &MsgUpdateWithdraw{
		Creator:    creator,
		TxHash:     txHash,
		From:       from,
		To:         to,
		CoinType:   coinType,
		Value:      value,
		Data:       data,
		Status:     status,
		Signatures: signatures,
	}
}

func (msg *MsgUpdateWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgUpdateWithdraw) Type() string {
	return TypeMsgUpdateWithdraw
}

func (msg *MsgUpdateWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteWithdraw{}

func NewMsgDeleteWithdraw(
	creator string,
	txHash string,

) *MsgDeleteWithdraw {
	return &MsgDeleteWithdraw{
		Creator: creator,
		TxHash:  txHash,
	}
}
func (msg *MsgDeleteWithdraw) Route() string {
	return RouterKey
}

func (msg *MsgDeleteWithdraw) Type() string {
	return TypeMsgDeleteWithdraw
}

func (msg *MsgDeleteWithdraw) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteWithdraw) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteWithdraw) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
