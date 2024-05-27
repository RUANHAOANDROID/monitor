package pkg

import (
	"testing"
)

func Test_Datafmt(t *testing.T) {
	str := ProtoDataFmt("37.8%", 1.35, 2.84, 6)
	t.Log(str)
	str = ProtoDataFmt("BB", nil, 188.8, "Test")
	t.Log(str)
	str = ProtoDataFmt("BB", "测试", nil, nil)
	t.Log(str)
}
