package graph

import (
	"context"
	"errors"

	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/graph/model"
)

type mutationResolver struct{ *Resolver }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *mutationResolver) CreateMeetup(ctx context.Context, input *model.NewMeetup) (*model.Meetup, error) {
	return r.Domain.CreateMeetup(ctx, input)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	return r.Domain.UpdateMeetup(ctx, id, input)
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteMeetup(ctx, id)
}

func (m *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("Input error")
	}

	return m.Domain.Register(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("Input error")
	}

	return r.Domain.Login(ctx, input)
}
