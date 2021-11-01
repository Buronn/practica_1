package functions

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os"
)

func GenerateImage2(res []byte) {
	img, _, _ := image.Decode(bytes.NewReader(res))
	out, err := os.Create("./images/img1.png") //LOCAL
	/* out, err := os.Create("/app/images/img1.png") */    //TEST ENSEÑA
	if err != nil {
		fmt.Println(err)
	}
	err = png.Encode(out, img)
	if err != nil {
		fmt.Println(err)
	}
}
