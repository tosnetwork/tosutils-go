package stack

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return NIP() })
}

func NIP() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			return state.Stack.PopSwapAt(1)
		},
		Name:      "NIP",
		BitPrefix: helpers.BytesPrefix(0x31),
	}
}
