package funcs

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"errors"
	"fmt"
	"image"
	"log"
	"net"
)

var ImgChan = make(chan *image.RGBA, 10)

// func SendImg(img *image.RGBA, conn net.Conn) {
// 	imgByte, err := EncodeImage(img)
// 	if err != nil {
// 		fmt.Println("encodeImage err =", err)
//
// 	}
// 	// 再将line发送给服务器
// 	_, err = conn.Write(imgByte)
// 	if err != nil {
// 		fmt.Println("conn write err =", err)
// 	}
// }

func ReceiveImageFromClient(conn net.Conn) (*image.RGBA, error) {
	var imgSize int64
	err := binary.Read(conn, binary.BigEndian, &imgSize)
	if err != nil {
		return nil, err
	}

	imgBytes := make([]byte, imgSize)
	_, err = conn.Read(imgBytes)
	if err != nil {
		return nil, err
	}

	imgReader := bytes.NewReader(imgBytes)
	img, _, err := image.Decode(imgReader)
	if err != nil {
		return nil, err
	}

	rgbaImg, ok := img.(*image.RGBA)
	if !ok {
		return nil, errors.New("Received image is not of type *image.RGBA")
	}

	return rgbaImg, nil
}
func EncodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}
func DecodeToRGBA(s []byte) image.RGBA {

	p := image.RGBA{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil {
		log.Fatal(err)
	}
	return p
}
