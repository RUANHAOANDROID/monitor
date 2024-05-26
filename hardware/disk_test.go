package hardware

import "testing"

func TestDisk(t *testing.T) {
	s := GetDiskInfo()
	t.Log(s)
}
