package stack

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return OVER2() })
}

func OVER2() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			if err := requireStackDepth(state, 4); err != nil {
				return err
			}

			val, err := state.Stack.Get(3)
			if err != nil {
				return err
			}
			if err := state.Stack.PushAny(val); err != nil {
				return err
			}
			val, err = state.Stack.Get(3)
			if err != nil {
				return err
			}
			return state.Stack.PushAny(val)
		},
		Name:      "2OVER",
		BitPrefix: helpers.BytesPrefix(0x5D),
	}
}
