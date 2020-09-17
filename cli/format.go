package cli

import (
	"fmt"
	"strconv"
	"strings"
)

// ColorStr 格式化字符串，用于彩色输出字符串
func ColorStr(format, s string) string {
	colorMap := map[string]int{
		"normal":         0,  // 默认值
		"hightligt":      1,  // 高亮
		"halflight":      2,  // 半亮
		"underline":      4,  // 下划线
		"flash":          5,  // 闪烁
		"closeunderline": 27, // 关闭下划线

		"fgblack":  30, // 前景色为黑色
		"fgred":    31, // 红色
		"fggreen":  32, // 绿色
		"fgbrown":  33, // 棕色
		"fgblue":   34, // 蓝色
		"fgpurple": 35, // 紫色
		"fgcyan":   36, // 青色
		"fgwhite":  37, // 白色

		"bgblack":  40, // 背景色
		"bgred":    41,
		"bggreen":  42,
		"bgbrown":  43,
		"bgblue":   44,
		"bgpurple": 45,
		"bgcyan":   46,
		"bgwhite":  47,
	}
	colors := strings.Split(format, ",")
	var fstring []string
	for i := 0; i < len(colors); i++ {
		str := colors[i]
		v, ok := colorMap[str]
		if ok {
			vstr := strconv.Itoa(v)
			fstring = append(fstring, vstr)

		}
	}
	str := fmt.Sprintf("%c[%sm%s%c[0m", 0x1B, strings.Join(fstring, ";"), s, 0x1B)
	return str
}
