package main

import (
	"fmt"

	"github.com/robertobouses/interface_ejercicio6/user"
)

type Inventarirar interface {
	Data() (string, string, user.Artist)
	DataArtist() user.Artist
	Accounting
}

type Accounting interface {
	Amount() int
}

func PrintingAllData(I Inventarirar) {
	fmt.Println(I.Data())
}

func PrintingArtist(I Inventarirar) {
	fmt.Println(I.DataArtist())

}

func PrintingAmount(I Inventarirar) {
	fmt.Println(I.Amount())
}

func main() {
	fmt.Println("holaaa")

	disc1 := user.Disc{"Some Girls", "Rock", user.Artist{"ROLLING STONES", 1978}, 26, 3}
	tape1 := user.Tape{"Help", "Rock", user.Artist{"THE BEATLES STONES", 1965}, 25, 5}

	fmt.Println("----------data of discs and tapes:")
	PrintingAllData(disc1)
	PrintingAllData(tape1)
	fmt.Println("----------data artist of Some Girls:")
	PrintingArtist(disc1)
	fmt.Println("----------amount of discs and tapes:")
	PrintingAmount(disc1)
	PrintingAmount(tape1)

}
