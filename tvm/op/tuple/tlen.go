package tuple

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return TLEN() })
}

func TLEN() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Name:      "TLEN",
		BitPrefix: helpers.BytesPrefix(0x6f, 0x88),
		Action: func(state *vm.State) error {
			tup, err := state.Stack.PopTupleRange(255)
			if err != nil {
				return err
			}
			return state.Stack.PushSmallInt(int64(tup.Len()))
		},
	}
}
