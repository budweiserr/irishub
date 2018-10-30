package iservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSvcDef{}, "iris-hub/iservice/MsgSvcDef", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}
