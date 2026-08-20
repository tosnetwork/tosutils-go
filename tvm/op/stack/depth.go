package stack

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return DEPTH() })
}

func DEPTH() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			return state.Stack.PushSmallInt(int64(state.Stack.Len()))
		},
		Name:      "DEPTH",
		BitPrefix: helpers.BytesPrefix(0x68),
	}
}
