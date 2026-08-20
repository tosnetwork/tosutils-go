package stack

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
	"github.com/tosnetwork/tosutils-go/tvm/vmerr"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return CHKDEPTH() })
}

func CHKDEPTH() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			idx, err := popSmallIndex(state)
			if err != nil {
				return err
			}
			if idx > state.Stack.Len() {
				return vmerr.Error(vmerr.CodeStackUnderflow)
			}
			return nil
		},
		Name:      "CHKDEPTH",
		BitPrefix: helpers.BytesPrefix(0x69),
	}
}
