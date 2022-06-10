package cars

import "context"

type Repository interface {
	Create(ctx context.Context, car Car) error
	Read(ctx context.Context, registrationNr string) (Car, error)
	Update(ctx context.Context, car Car) error
	Delete(ctx context.Context, car Car) error
}
