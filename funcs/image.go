package funcs

import (
	"bytes"
	"fmt"
	"image"
)

var ImgChan = make(chan *image.RGBA, 10)

// DecodeImageFromBase64 从 Base64 字符串解码为 *image.RGBA
func DecodeImageFromBase64(s []byte) (*image.RGBA, error) {

	img, _, err := image.Decode(bytes.NewReader(s))
	if err != nil {
		return nil, err
	}

	if rgbaImg, ok := img.(*image.RGBA); ok {
		return rgbaImg, nil
	}

	// 如果图像不是 *image.RGBA 类型，可以尝试进行类型转换
	return nil, fmt.Errorf("Decoded image is not of type *image.RGBA")
}
