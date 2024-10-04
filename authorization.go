package easyauthz

import "fmt"

func Authorize(ability string) bool {
	fmt.Println(ability)
	fmt.Println("Check permission in db")
	fmt.Println("Test change version to v1.0.1")
	return true
}
