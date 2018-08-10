package mcuapp

import (
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var MCUAPP *ole.IDispatch
func Init()  {
	ole.CoInitialize(0)
	unknown, _ := oleutil.CreateObject("MCUAPP.KML")
	MCUAPP,_ = unknown.QueryInterface(ole.IID_IDispatch)
	oleutil.MustCallMethod(MCUAPP, "OpenDevice").ToIDispatch()
}

func MoveToR (x int, y int) {
	oleutil.MustCallMethod(MCUAPP, "MoveToR",x,y).ToIDispatch()
}
func LeftClick () {
	oleutil.MustCallMethod(MCUAPP, "LeftClick",1).ToIDispatch()
}
func MiddleClick () {
	oleutil.MustCallMethod(MCUAPP, "MiddleClick",1).ToIDispatch()
}
func KeyPress (key string) {
	oleutil.MustCallMethod(MCUAPP, "KeyPress",key,1).ToIDispatch()
}

