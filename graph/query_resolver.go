package graph

import (
	"context"

	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/graph/model"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupsRepo.GetMeetups()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UsersRepo.GetUserByID(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
