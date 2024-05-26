package hardware

import (
	"monitor/pkg"
	"testing"
)

func TestCPU(t *testing.T) {
	cpuPercent := GetCPUInfo()
	pkg.Log.Println(cpuPercent)
}
