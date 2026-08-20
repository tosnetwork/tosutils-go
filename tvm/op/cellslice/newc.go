package cellslice

import (
	"github.com/tosnetwork/tosutils-go/tvm/cell"
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return NEWC() })
}

func NEWC() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			return state.Stack.PushOwnedBuilder(cell.BeginCell().SetTrace(state.Cells.Trace()))
		},
		Name:      "NEWC",
		BitPrefix: helpers.BytesPrefix(0xC8),
	}
}
