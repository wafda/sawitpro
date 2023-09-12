package repository

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (r *Repository) GetProfile(ctx context.Context, input GetProfileParams) (dbUser User, err error) {
	return r.getProfile(ctx, *input.Token)
}

func (r *Repository) getProfile(ctx context.Context, token string) (dbUser User, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT id,fullname,password,phone_number FROM user_table WHERE token = $1", token).Scan(
		&dbUser.Id,
		&dbUser.FullName,
		&dbUser.Password,
		&dbUser.PhoneNumbers,
	)
	return
}

func (r *Repository) UpdateProfile(ctx context.Context, input PutProfileParams, newUser User) (dbUser User, err error) {
	if err := r.validate.Struct(input); err != nil {
		return User{}, err
	}
	if err := r.validate.Struct(newUser); err != nil {
		return User{}, err
	}
	dbUser, err = r.getProfile(ctx, *input.Token)
	if err != nil {
		return User{}, err
	}
	if newUser.FullName == "" {
		newUser.FullName = dbUser.FullName
	}
	if newUser.PhoneNumbers == "" {
		newUser.PhoneNumbers = dbUser.PhoneNumbers
	}
	_, err = r.Db.Exec("UPDATE user_table SET fullname = $1, phone_number = $3 WHERE id = $2;", newUser.FullName, dbUser.Id, newUser.PhoneNumbers)
	if err != nil {
		return User{}, err
	}
	dbUser.FullName = newUser.FullName
	dbUser.PhoneNumbers = newUser.PhoneNumbers
	return dbUser, err
}

func (r *Repository) Login(ctx context.Context, input LoginRequest) (User, error) {
	dbUser := User{}
	if err := r.validate.Struct(input); err != nil {
		return dbUser, err
	}

	err := r.Db.QueryRowContext(ctx, "SELECT id,fullname,password,phone_number FROM user_table WHERE phone_number = $1", input.PhoneNumbers).Scan(
		&dbUser.Id,
		&dbUser.FullName,
		&dbUser.Password,
		&dbUser.PhoneNumbers,
	)
	if err != nil {
		return User{}, err
	}

	err = r.ValidatePassword(ctx, dbUser, input.Password)
	if err != nil {
		return User{}, err
	}
	token := GenerateToken(dbUser.Id)
	dbUser.Token = &token
	_, err = r.Db.Exec("UPDATE user_table SET token = $1 WHERE id = $2;", token, dbUser.Id)
	if err != nil {
		return User{}, err
	}

	return dbUser, err
}

func (r *Repository) ValidatePassword(ctx context.Context, user User, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password))
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return errors.New("invalid password or username")
		}
		return err
	}
	return nil
}

func (r *Repository) RegisterUser(ctx context.Context, input RegisterRequest) (User, error) {
	if err := r.validate.Struct(input); err != nil {
		return User{}, err
	}
	if !(*input.Password == *input.PasswordConfirmation) {
		return User{}, errors.New("invalid password confirmation")
	}
	reg := regexp.MustCompile(`([^a-zA-Z\d])+([a-zA-Z\d])+|([a-zA-Z\d])+([^a-zA-Z\d])+`)
	if !reg.Match([]byte(*input.Password)) {
		return User{}, errors.New("password should contain at least 1 numeric, 1 alphabet and 1 non alphanumeric")
	}
	hasehedPassword, err := bcrypt.GenerateFromPassword([]byte(*input.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}

	_, err = r.Db.Exec("INSERT INTO user_table (fullname,phone_number,password) VALUES ($1,$2,$3);", input.FullName, input.PhoneNumbers, hasehedPassword)
	if err != nil {
		return User{}, err
	}

	return User{
		FullName:     *input.FullName,
		PhoneNumbers: *input.PhoneNumbers,
	}, nil
}

func GenerateToken(userID int64) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", userID)))
	h.Write([]byte(time.Now().Format(time.RFC3339Nano)))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
