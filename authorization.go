package easyauthz

import "fmt"

func Authorize(ability string) bool {
	fmt.Println(ability)
	return true
}
