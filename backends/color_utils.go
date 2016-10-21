package backends
import (
	color "github.com/fatih/color"
)

// 警告信息采用红色显示
var RedF = color.New(color.FgRed).SprintFunc()

// 新增服务等采用绿色显示
var GreenF = color.New(color.FgGreen).SprintFunc()

var MagentaF = color.New(color.FgMagenta).SprintFunc()
var CyanF = color.New(color.FgCyan).SprintFunc()

var BlueF = color.New(color.FgBlue).SprintFunc()