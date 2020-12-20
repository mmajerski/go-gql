package graph

import (
	"context"

	"github.com/userq11/meetmeup/graph/generated"
	"github.com/userq11/meetmeup/graph/model"
)

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) UserID(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }
