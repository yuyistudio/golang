// Doc模块演示了如何在Go中编写注释
package doc

import (
	"fmt"
	"sync"
)

// 用于展示的信息
type Info struct {
	Name string // 名称
	age int // 加了注释也不显示,因为没有导出
}

// Deprecated: 建议实验第二个版本ShowV2.
// 显示一段字符串
func Show(info string) {

}

// 显示一段字符串
func ShowV2(info string){

}

