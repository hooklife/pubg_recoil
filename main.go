package main

import (
	"yaqiang/mcuapp"
	"yaqiang/event"
)


func main() {
	// 初始化键鼠
	mcuapp.Init()
	// 监听鼠标事件
	event.Start()
}