package main

import "fmt"

func main() {
	gun1, err1 := getGun("ak47")
	gun2, err2 := getGun("shotgun")
	gun3, err3 := getGun("bazooka") //Type Bazooka Doesnt exist

	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(gun1)
	}
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Println(gun2)
	}
	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(gun3)
	}
}
