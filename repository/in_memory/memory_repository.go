package in_memory

import (
	"context"
	"repository/cars"
)

type CarRepository struct {
}

// If you need to initalize your data structure do it in here
func NewCarRepository() *CarRepository {
	return &CarRepository{}
}

func (c CarRepository) Create(ctx context.Context, car cars.Car) error {
	//TODO implement me
	panic("implement me")
}

func (c CarRepository) Read(ctx context.Context, registrationNr string) (cars.Car, error) {
	//TODO implement me
	panic("implement me")
}

func (c CarRepository) Update(ctx context.Context, car cars.Car) error {
	//TODO implement me
	panic("implement me")
}

func (c CarRepository) Delete(ctx context.Context, car cars.Car) error {
	//TODO implement me
	panic("implement me")
}
