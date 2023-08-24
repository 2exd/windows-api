package funcs

import "github.com/go-vgo/robotgo"

func MouseMove() {
	x, y := robotgo.GetMousePos()
	// robotgo.MoveMouseSmooth(x+600, y+400)
	robotgo.MoveMouseSmooth(x+600, y+400, 20.0, 200.0) // 后面两个参数文档上看lowspeed, highspeed，与速度相关，是能改变移动速度，但是没搞清究竟是怎么个改变的
}
