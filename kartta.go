package main

import (
	"fmt"
	"os"
)

var (
	kartta             Kartta
	pelaajaX, pelaajaY int
)

func init() {
	merkkiRuuduksi := map[rune]Ruutu{
		'#': Seinä,
		'.': Tyhjä,
		'o': Kivi,
		'x': Kuoppa,
		'p': Pelaaja,
		'm': Maali,
	}

	file, err := os.Open("testitaso.mp")
	if err != nil {
		panic("En saanut ladattua tasoa: " + err.Error())
	}

	fmt.Fscan(file, &kartta.Leveys, &kartta.Korkeus, &pelaajaX, &pelaajaY)

	kartta.ruudut = make([]Ruutu, kartta.Leveys*kartta.Korkeus)

	for y := 0; y < kartta.Korkeus; y++ {
		var rivi string
		fmt.Fscan(file, &rivi)
		for x, merkki := range rivi {
			kartta.Laita(x, y, merkkiRuuduksi[merkki])
		}
	}
}

type Kartta struct {
	Korkeus, Leveys int
	ruudut          []Ruutu
}

func (k Kartta) Kohdassa(x, y int) Ruutu {
	return k.ruudut[k.laskeIndeksi(x, y)]
}

func (k *Kartta) Laita(x, y int, r Ruutu) {
	k.ruudut[k.laskeIndeksi(x, y)] = r
}

func (k Kartta) laskeIndeksi(x, y int) int {
	if x < 0 || x >= k.Leveys {
		panic(fmt.Sprintf("x on %d! Kartan leveys on %d.", x, k.Leveys))
	}
	if y < 0 || y >= k.Korkeus {
		panic(fmt.Sprintf("y on %d! Kartan korkeus on %d.", y, k.Korkeus))
	}

	return k.Leveys*y + x
}

type Ruutu uint8

const (
	Tyhjä = iota
	Pelaaja
	Seinä
	Kivi
	Kuoppa
	Maali
	KuviaTextuurissa
)
