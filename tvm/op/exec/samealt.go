package exec

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return SAMEALT() })
}

func SAMEALT() (op *helpers.SimpleOP) {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			state.Reg.C[1] = state.Reg.C[0].Copy()
			return nil
		},
		Name:      "SAMEALT",
		BitPrefix: helpers.BytesPrefix(0xED, 0xFA),
	}
}
