//Package ram provides information on the ram of the system.
package ram

import (
	"github.com/jaypipes/ghw"
	"qlova.tech/qty"
)

//Amount is the total usable ram in bytes.
var Amount qty.Data

func init() {
	info, _ := ghw.Memory()
	Amount = qty.Bytes(float64(info.TotalUsableBytes))
}
