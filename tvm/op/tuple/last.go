package tuple

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return LAST() })
}

func LAST() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Name:      "LAST",
		BitPrefix: helpers.BytesPrefix(0x6f, 0x8b),
		Action: func(state *vm.State) error {
			tup, err := state.Stack.PopTupleRange(255, 1)
			if err != nil {
				return err
			}
			val, err := tup.Index(tup.Len() - 1)
			if err != nil {
				return err
			}
			return state.Stack.PushAny(val)
		},
	}
}
