package main

import (
	"fmt"

	dog "github.com/listonaj/GoToDog"
	"github.com/listonaj/puppy"
)

func main() {
	fmt.Println("HelloWorld")
	s1 := puppy.Bark()
	fmt.Println(s1)

	s2 := puppy.BigBark()

	fmt.Println(s2)

	s3 := puppy.Barks()

	fmt.Println(s3)

	s4 := dog.WhenGrownUp(s2)
	s5 := dog.WhenGrownUp(s3)

	fmt.Println(s4)
	fmt.Println(s5)

}
