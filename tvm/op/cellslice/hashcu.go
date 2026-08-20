package cellslice

import (
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
	"math/big"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return HASHCU() })
}

func HASHCU() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			c, err := state.Stack.PopCell()
			if err != nil {
				return err
			}
			hash := c.HashKey()
			return state.Stack.PushOwnedInt(new(big.Int).SetBytes(hash[:]))
		},
		Name:      "HASHCU",
		BitPrefix: helpers.BytesPrefix(0xF9, 0x00),
	}
}
