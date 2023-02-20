package logic

import (
	"demo/backend/types"
	"demo/public"
)

func SomeJson(*types.SomeJsonQuery) (*types.SomeHandlerResp, *public.ProtocolError) {

	return &types.SomeHandlerResp{Name: "Dynamic", Age: 1, Married: false, RingRatio: float64(-2) / float64(15)}, nil

}
