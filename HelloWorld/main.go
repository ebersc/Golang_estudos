package main

import (
	"fmt"
)

func main(){
	// fmt.Println("Olá mundo!")
	// fmt.Println("TreinaWeb" + " Cursos")
	// fmt.Println("1 + 1 = ", 1+1)
	// fmt.Println("1.1 + 2.2 = ", 1.1 + 2.2)
	// fmt.Println(true)

	fmt.Println("INTEIROS SEM SINAL")

	var v1 byte = 255 //uint8 = byte = 0 -> 255
	fmt.Println(v1)
	var v2 uint16 = 33
	fmt.Println(v2)
	var v3 uint32 = 44
	fmt.Println(v3)
	var v4 uint64 = 55
	fmt.Println(v4)

	fmt.Println("INTEIROS COM SINAL")

	var i1 int8 = 127 //-128 a 127
	fmt.Println(i1)
	var i2 int16 = -1000
	fmt.Println(i2)
	var i3 rune = 15000 //int32
	fmt.Println(i3)
	var i4 int64 = -55000
	fmt.Println(i4)

	var t1 uint = 1000 //segue a arquitetura do processador 32 ou 64 bits
	fmt.Println(t1)

	var t2 int = 2000 // mesma coisa
	fmt.Println(t2)

	var f1 float32 = 1
	fmt.Println(f1)

	var f2 float64 = 2
	fmt.Println(f2)

	var f3 complex64 = 3
	fmt.Println(f3)

	var f4 complex128 = 4
	fmt.Println(f4)

	fmt.Println("STRINGS")

	var s1 string = "treinaweb \nCursos"
	fmt.Println(s1)

	var s2 string = `treinaweb
	Cursos`

	fmt.Println(s2)
}