package stack

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
	"github.com/tosnetwork/tosutils-go/tvm/vmerr"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return ROLL() })
}

func ROLL() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			idx, err := popSmallIndex(state)
			if err != nil {
				return err
			}
			if idx >= state.Stack.Len() {
				return vmerr.Error(vmerr.CodeStackUnderflow)
			}
			if err := consumeLargeStackMoveGas(state, idx); err != nil {
				return err
			}
			for idx > 0 {
				if err := state.Stack.Exchange(idx-1, idx); err != nil {
					return err
				}
				idx--
			}
			return nil
		},
		Name:      "ROLL",
		BitPrefix: helpers.BytesPrefix(0x61),
	}
}
