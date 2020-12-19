package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/graph/model"
)

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	user := new(model.User)

	for _, u := range users {
		if u.ID == obj.User {
			user = u
			break
		}
	}

	if user == nil {
		return nil, errors.New("User with the id does not exist")
	}

	return user, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input *model.NewMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return meetups, nil
}

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	var m []*model.Meetup

	for _, meetup := range meetups {
		if meetup.User == obj.ID {
			m = append(m, meetup)
		}
	}

	return m, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		User:        "1",
	},
	{
		ID:          "2",
		Name:        "A meetup 2",
		Description: "A description 2",
		User:        "2",
	},
}
var users = []*model.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@mail.com",
	},
	{
		ID:       "2",
		Username: "Sam",
		Email:    "sam@mail.com",
	},
}
