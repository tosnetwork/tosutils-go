package funcs

import (
	"github.com/tosnetwork/tosutils-go/tlb"
	"github.com/tosnetwork/tosutils-go/tvm/op/helpers"
	"github.com/tosnetwork/tosutils-go/tvm/vm"
	"github.com/tosnetwork/tosutils-go/tvm/vmerr"
)

func init() {
	vm.List = append(vm.List, func() vm.OP { return SENDRAWMSG() })
}

func SENDRAWMSG() *helpers.SimpleOP {
	return &helpers.SimpleOP{
		Action: func(state *vm.State) error {
			if state.Stack.Len() < 2 {
				return vmerr.Error(vmerr.CodeStackUnderflow)
			}

			i0, err := state.Stack.PopIntRangeInt64(0, 255)
			if err != nil {
				return err
			}

			c1, err := state.Stack.PopCell()
			if err != nil {
				return err
			}

			list := tlb.OutList{
				Prev: state.Reg.D[1],
				Out: tlb.ActionSendMsg{
					Mode: uint8(i0),
					Msg:  c1,
				},
			}

			res, err := tlb.ToCell(list)
			if err != nil {
				return vmerr.Error(vmerr.CodeCellOverflow, "cannot serialize raw output message into an output action cell; "+err.Error())
			}
			if err = state.Cells.RegisterCellCreate(); err != nil {
				return err
			}
			state.Reg.D[1] = res
			return nil
		},
		Name:      "SENDRAWMSG",
		BitPrefix: helpers.BytesPrefix(0xFB, 0x00),
	}
}
