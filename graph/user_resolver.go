package graph

import (
	"context"

	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/graph/model"
)

type userResolver struct{ *Resolver }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	return r.Domain.MeetupsRepo.GetMeetupsForUser(obj)
}
