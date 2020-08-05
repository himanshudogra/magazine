package main

import (
	"fmt"

	"github.com/himanshudogra/magazine"
)

func main() {

	subscriber := magazine.Subscriber{Name: "Himanshu"}
	subscriber.Street = "1-KA-277,278"
	subscriber.City = "Alwar"
	subscriber.State = "RJ"
	subscriber.PostalCode = "301001"
	fmt.Println("Name:", subscriber.Name)
	fmt.Println("Address:", subscriber.Address)
}
