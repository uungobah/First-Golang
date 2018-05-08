package main

import (
	"fmt"
	"HelloWorld/package1"
)

func main()  {
	/*test()
	variable()
	variableWithoutDataType()
	multiVariable()*/
	package1.Looping()
	package1.LoopingWithArgument()
	package1.LoopingWithoutArgument()
}

func test(){
	fmt.Println("Test")
}

func variable(){
	fmt.Printf("Variable Dengan Tipe Data \n")
	var firstName string
	firstName = "Jhon"

	var lastName string = "Kennedy"
	fmt.Println("hello","world!!", firstName, lastName)
}

func variableWithoutDataType(){
	fmt.Printf("Variable Tanpa Tipe Data \n")

	var firstName string = "john"
	lastName := "wick"

	fmt.Printf("halo john wick!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName + "!")
}

func multiVariable(){
	fmt.Printf("Multi Variable \n")
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	fmt.Println("halo", first, second, third + "!")

	// Langsung diisi
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	fmt.Println("halo", fourth, fifth, sixth + "!")

	// tanpa var
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Println("halo", seventh, eight, ninth + "!")

	/*Dengan menggunakan teknik type inference, deklarasi multi variabel bisa
	 dilakukan untuk variabel-variabel yang tipe data satu sama lainnya berbeda.*/
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println("halo", one, isFriday, twoPointTwo, say + "!")

	/*Underscore (_) adalah predefined variabel yang bisa dimanfaatkan untuk menampung nilai yang tidak dipakai.
	 Bisa dibilang variabel ini merupakan keranjang sampah.*/
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	name, _ := "john", "wick"
	fmt.Printf(name)
}

