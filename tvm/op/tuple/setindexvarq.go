package tuple

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
	"github.com/tosnetwork/tosutils-go/tvm/vmerr"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return SETINDEXVARQ() })
}

func SETINDEXVARQ() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Name:      "SETINDEXVARQ",
		BitPrefix: helpers.BytesPrefix(0x6f, 0x87),
		Action: func(state *vm.State) error {
			if state.Stack.Len() < 3 {
				return vmerr.Error(vmerr.CodeStackUnderflow)
			}
			idx, err := state.Stack.PopIntRangeInt64(0, 254)
			if err != nil {
				return err
			}
			return execSetIndexQuiet(state, int(idx))
		},
	}
}
