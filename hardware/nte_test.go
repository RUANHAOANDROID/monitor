package hardware

import "testing"

func TestNet(t *testing.T) {
	content, err := Net()
	if err != nil {
	}
	t.Log(content)
}
