package graph

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/graph/model"
	"github.com/userq11/meetmeup/middleware"
)

type mutationResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) CreateMeetup(ctx context.Context, input *model.NewMeetup) (*model.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errors.New("Unauthenticated")
	}

	if len(input.Name) < 3 {
		return nil, errors.New("Name is not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("Description is not long enough")
	}

	meetup := &model.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}

	return r.MeetupsRepo.CreateMeetup(meetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := r.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return nil, errors.New("Meetup not exist")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("Name is not long enough")
		}

		meetup.Name = *input.Name
		didUpdate = true
	}

	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("Description is not long enough")
		}

		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("No update done")
	}

	meetup, err = r.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("Error while updating meetup: %v", err)
	}

	return meetup, nil
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return false, errors.New("Meetup does not exist")
	}

	err = r.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("Error while deleting meetup: %v", err)
	}

	return true, nil
}

func (m *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := m.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("Email already in use")
	}

	_, err = m.UsersRepo.GetUserByUsername(input.Username)
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

	tx, err := m.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("Error while creating a transaction: %v", err)
		return nil, errors.New("Something went wrong")
	}

	defer tx.Rollback()

	if _, err := m.UsersRepo.CreateUser(tx, user); err != nil {
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

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := r.UsersRepo.GetUserByEmail(input.Email)
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
