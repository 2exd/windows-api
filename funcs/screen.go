package funcs

import (
	"fmt"
	"github.com/kbinani/screenshot"
	hook "github.com/robotn/gohook"
	"image"
	"image/png"
	"os"
	"time"
)

func GetScreen() {
	fmt.Println("--- Please press 1+2+3 to send screenshot ---")
	hook.Register(hook.KeyDown, []string{"1", "2", "3"}, func(e hook.Event) {
		fmt.Println("1+2+3")
		ScreenshotFunc()
	})

	fmt.Println("--- Please press ctrl + shift + q to stop screenshot ---")
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("ctrl-shift-q")
		hook.End()
	})
	s := hook.Start()
	<-hook.Process(s)
}

func ScreenshotFunc() {
	n := screenshot.NumActiveDisplays()
	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			panic(err)
		}
		ImgChan <- img
	}
}

func SaveScreen(img *image.RGBA) {
	fileName := fmt.Sprintf("%d.png", time.Now().UnixNano())
	file, _ := os.Create(fileName)
	err := png.Encode(file, img)
	if err != nil {
		file.Close() // 关闭文件
		panic(err)
	}

	// err = file.Sync() // 强制将文件内容写入磁盘
	// if err != nil {
	// 	file.Close()
	// 	panic(err)
	// }

	err = file.Close() // 保存并关闭文件
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", fileName)
}
