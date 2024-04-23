package main

import "fmt"

func getGun(name string) (IGun, error) {

	switch name {
	case "ak47":
		fmt.Println("masuk")
		return newAK47(), nil
	case "shotgun":
		return newShotgun(), nil
	}

	return nil, fmt.Errorf("gagal membuat senjata, jenis senjata %s belum ada dalam factory ini", name)
}
