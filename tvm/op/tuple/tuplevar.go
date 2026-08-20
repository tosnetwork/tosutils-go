package tuple

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return TUPLEVAR() })
}

func TUPLEVAR() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Name:      "TUPLEVAR",
		BitPrefix: helpers.BytesPrefix(0x6f, 0x80),
		Action: func(state *vm.State) error {
			count, err := state.Stack.PopIntRangeInt64(0, 255)
			if err != nil {
				return err
			}
			return execMakeTuple(state, int(count))
		},
	}
}
