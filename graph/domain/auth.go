package domain

import (
	"context"
	"errors"
	"log"

	"github.com/userq11/meetmeup/graph/model"
)

func (d *Domain) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("Email already in use")
	}

	_, err = d.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("Username already in use")
	}

	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("Error while hashing password: %v", err)
		return nil, errors.New("Something went wrong")
	}

	// TODO: create verification code

	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("Error while creating a transaction: %v", err)
		return nil, errors.New("Something went wrong")
	}

	defer tx.Rollback()

	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("Error creating a user: %v", err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error while commiting: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("Error while generating the token: %v", err)
		return nil, errors.New("Something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (d *Domain) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, errors.New("Incorrect either email or password")
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, errors.New("Incorrect either email or password")
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("Something went wrong")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}
