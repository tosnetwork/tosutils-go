package stack

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return DUP() })
}

func DUP() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			return state.Stack.PushAt(0)
		},
		Name:      "DUP",
		BitPrefix: helpers.BytesPrefix(0x20),
	}
}
