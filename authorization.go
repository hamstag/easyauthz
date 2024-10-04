package easyauthz

import "fmt"

func Authorize(ability string) bool {
	fmt.Println(ability)
	fmt.Println("Check permission in db")
	return true
}
