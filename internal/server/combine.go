package server

import (
	"github.com/zngue/zng_app/pkg/types"
)

type V1Router []types.Register
type V2Router []types.Register

func NewCombine(v1 V1Router) (routes []types.Register) {
	routes = append(routes, v1...)
	//routes = append(routes, v2...)
	return
}
