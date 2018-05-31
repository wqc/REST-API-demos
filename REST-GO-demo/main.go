package main

import (
	"fmt"
	"github.com/MsloveDl/HuobiProAPI/services"
)

func main() {
	fmt.Println("火币")
	acountData := services.GetAccounts().Data
	if len(acountData) > 0 {
		id := acountData[0].ID
		fmt.Println("%d", id)
	} else {
		fmt.Println("no valid data!!")
	}
}
