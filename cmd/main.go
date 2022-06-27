package main

import (
	"context"
	"fmt"
	"log"
	"repository/pkg/cars"
	"repository/pkg/repository/in_memory"
)

func main() {
	ctx := context.Background()
	var carRepository cars.Repository = in_memory.NewCarRepository()
	err := carRepository.Create(ctx, cars.Car{Brand: "volvo", Model: "v70", RegistrationNr: "ABB770"})
	if err != nil {
		log.Fatal(err)
	}
	volvo, err := carRepository.Read(ctx, "ABB770")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(volvo)
}
