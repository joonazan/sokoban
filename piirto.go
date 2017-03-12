package main

import (
	"github.com/go-gl/gl/v2.1/gl"
	"image"
	_ "image/png"
	"os"
	"unsafe"
)

func LataaKuvat() {
	file, err := os.Open("tiles.png")
	if err != nil {
		panic("En saa avattua kuvatiedostoa: " + err.Error())
	}

	kuva, _, err := image.Decode(file)
	if err != nil {
		panic("Kuva on rikki: " + err.Error())
	}

	var texture uint32
	gl.GenTextures(1, &texture)

	gl.BindTexture(gl.TEXTURE_2D, texture)
	gl.TexImage2D(gl.TEXTURE_2D, 0, 4,
		int32(kuva.Bounds().Dx()), int32(kuva.Bounds().Dy()), 0,
		gl.RGBA, gl.UNSIGNED_BYTE,
		unsafe.Pointer(&kuva.(*image.RGBA).Pix[0]),
	)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	gl.Enable(gl.TEXTURE_2D)
}

func Piirrä() {
	gl.Begin(gl.QUADS)
	for i, r := range kartta.ruudut {
		x := i % kartta.Leveys
		y := i / kartta.Leveys
		piirräRuutu(x, y, r)
	}
	piirräRuutu(pelaajaX, pelaajaY, Pelaaja)
	gl.End()
}

func piirräRuutu(x, y int, r Ruutu) {
	const sivu = 0.1
	const tekstuuriaskel = 1.0 / KuviaTextuurissa

	left := float32(x)*sivu - 1
	top := -float32(y)*sivu + 1

	texAlku := tekstuuriaskel * float32(r)
	texLoppu := texAlku + tekstuuriaskel

	gl.TexCoord2f(texAlku, 0)
	gl.Vertex2f(left, top)

	gl.TexCoord2f(texLoppu, 0)
	gl.Vertex2f(left+sivu, top)

	gl.TexCoord2f(texLoppu, 1)
	gl.Vertex2f(left+sivu, top-sivu)

	gl.TexCoord2f(texAlku, 1)
	gl.Vertex2f(left, top-sivu)
}
