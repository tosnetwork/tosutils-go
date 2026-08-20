package exec

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
	"github.com/tosnetwork/tosutils-go/tvm/vmerr"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return JMPXVARARGS() })
}

func JMPXVARARGS() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			if state.Stack.Len() < 2 {
				return vmerr.Error(vmerr.CodeStackUnderflow)
			}

			paramsVal, err := state.Stack.PopIntRangeInt64(-1, 254)
			if err != nil {
				return err
			}

			params := int(paramsVal)

			if params >= 0 && state.Stack.Len() < params+1 {
				return vmerr.Error(vmerr.CodeStackUnderflow)
			}

			cont, err := state.Stack.PopContinuation()
			if err != nil {
				return err
			}

			return state.JumpArgs(cont, params)
		},
		Name:      "JMPXVARARGS",
		BitPrefix: helpers.BytesPrefix(0xDB, 0x3A),
	}
}
