package functions

import (
	"bytes"
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"os"
)

func GenerateGif(openFile [70][]byte) {

	anim := gif.GIF{}
	for i := 0; i < 70; i++ {

		xd, _, err := image.Decode(bytes.NewReader(openFile[i]))
		if err != nil {
			log.Println(err)
		}

		paletted := image.NewPaletted(xd.Bounds(), palette.Plan9)

		draw.FloydSteinberg.Draw(paletted, xd.Bounds(), xd, image.ZP)

		anim.Image = append(anim.Image, paletted)
		anim.Delay = append(anim.Delay, 10)

	}
	f, err := ioutil.TempFile("./instagram", "output*.gif") //LOCAL
	/* f, err := ioutil.TempFile("/app/instagram", "output*.gif")  */   //TEST ENSEÑA
	if err != nil {
		log.Fatal(err)
	}
	gif.EncodeAll(f, &anim)
	CreateFile(f.Name())

	defer os.Remove(f.Name())
	GetFile()
}
