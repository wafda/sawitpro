// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type RegisterRequest struct {
	FullName *string `json:"fullName,omitempty" validate:"required,min=3,max=60"`

	Password *string `json:"password,omitempty" validate:"required,min=6,max=64"`

	PasswordConfirmation *string `json:"passwordConfirmation,omitempty" validate:"required"`

	PhoneNumbers *string `json:"phoneNumbers,omitempty" validate:"required,min=10,max=16,startswith=+62"`
}

type User struct {
	FullName string `json:"fullName" validate:"omitempty,min=3,max=60"`

	Id int64 `json:"id"`

	Password *string `json:"-"`

	PasswordConfirmation *string `json:"passwordConfirmation,omitempty"`

	PhoneNumbers string `json:"phoneNumbers" validate:"omitempty,min=10,max=16,startswith=+62"`

	Token *string `json:"token,omitempty"`
}

type LoginRequest struct {
	Password string `json:"password" validate:"required"`

	PhoneNumbers string `json:"phoneNumbers" validate:"required"`
}

type GetProfileParams struct {
	Token *string `form:"token,omitempty" json:"token,omitempty"`
}

type PutProfileParams struct {
	Token *string `form:"token,omitempty" json:"token,omitempty"`
}
