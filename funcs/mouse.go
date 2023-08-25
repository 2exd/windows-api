package funcs

import "github.com/go-vgo/robotgo"

func MouseMove() {
	x, y := robotgo.GetMousePos()
	// robotgo.MoveMouseSmooth(x+600, y+400)
	robotgo.MoveSmooth(x+10, y+20)
}
