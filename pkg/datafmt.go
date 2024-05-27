package pkg

import (
	"fmt"
	"strings"
)

func ProtoDataFmt(s1, s2, s3, s4 interface{}) string {
	var parts []string

	// 添加非空参数到 parts 切片中
	if s1 != nil {
		parts = append(parts, fmt.Sprintf("%v", s1))
	}
	if s2 != nil {
		parts = append(parts, fmt.Sprintf("%v", s2))
	}
	if s3 != nil {
		parts = append(parts, fmt.Sprintf("%v", s3))
	}
	if s4 != nil {
		parts = append(parts, fmt.Sprintf("%v", s4))
	}

	// 使用 strings.Join() 连接所有非空参数，并在参数之间添加逗号
	return strings.Join(parts, ",")
}
