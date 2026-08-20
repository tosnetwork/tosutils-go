package stack

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return OVER() })
}

func OVER() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			return state.Stack.PushAt(1)
		},
		Name:      "OVER",
		BitPrefix: helpers.BytesPrefix(0x21),
	}
}
