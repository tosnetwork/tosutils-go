package exec

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return IFNOTRET() })
}

func IFNOTRET() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			b0, err := state.Stack.PopBool()
			if err != nil {
				return err
			}

			if !b0 {
				return state.Return()
			}
			return nil
		},
		Name:      "IFNOTRET",
		BitPrefix: helpers.BytesPrefix(0xDD),
	}
}
