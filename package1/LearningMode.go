package package1

import (
	"fmt"
)


func Looping(){
	for i := 0 ; i < 5 ; i++{
		fmt.Println("Looping Standart ",i);
	}
}

func LoopingWithArgument()  {
	var i =0;

	for i < 5 {
		fmt.Println("Looping With Argument ",i);
		i++
	}
}

func LoopingWithoutArgument()  {
	var i =0;

	for {
		fmt.Println("Looping Without Argument ",i);
		i++
		if i == 4 {
			break
		}
	}
}