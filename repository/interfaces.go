// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import "context"

type RepositoryInterface interface {
	GetProfile(ctx context.Context, input GetProfileParams) (dbUser User, err error)
	UpdateProfile(ctx context.Context, input PutProfileParams, newUser User) (dbUser User, err error)
	Login(ctx context.Context, input LoginRequest) (User, error)
	ValidatePassword(ctx context.Context, user User, password string) error
	RegisterUser(ctx context.Context, input RegisterRequest) (User, error)
}
