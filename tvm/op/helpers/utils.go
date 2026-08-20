package helpers

import (
	"github.com/tosnetwork/tosutils-go/tvm/cell"
)

func Builder(b []byte) *cell.Builder {
	return cell.BeginCell().MustStoreSlice(b, uint(len(b)*8))
}
