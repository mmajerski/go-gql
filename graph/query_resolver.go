package graph

import (
	"context"

	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/graph/model"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*model.Meetup, error) {
	return r.Domain.MeetupsRepo.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.Domain.UsersRepo.GetUserByID(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
