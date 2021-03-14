//Package cpu provides information about the cpu this program is running on.
package cpu

import "github.com/jaypipes/ghw"

//Cores is the number of reported cpu cores on the system.
var Cores int = -1

//Threads is the number of reported hardware threads on this system.
var Threads int = -1

func init() {
	info, _ := ghw.CPU()

	Cores = int(info.TotalCores)
	Threads = int(info.TotalThreads)
}
