package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/userq11/meetmeup/graph/model"
	"github.com/userq11/meetmeup/middleware"
)

func (d *Domain) CreateMeetup(ctx context.Context, input *model.NewMeetup) (*model.Meetup, error) {
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

	return d.MeetupsRepo.CreateMeetup(meetup)
}

func (d *Domain) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return nil, errors.New("Unauthenticated")
	}

	meetup, err := d.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return nil, errors.New("Meetup not exist")
	}

	if !meetup.IsOwner(currentUser) {
		return nil, errors.New("Not authorized")
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

	meetup, err = d.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("Error while updating meetup: %v", err)
	}

	return meetup, nil
}

func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware.GetCurrentUserFromCtx(ctx)
	if err != nil {
		return false, errors.New("Unauthenticated")
	}

	meetup, err := d.MeetupsRepo.GetById(id)
	if err != nil || meetup == nil {
		return false, errors.New("Meetup does not exist")
	}

	if !checkOwnership(meetup, currentUser) {
		return false, errors.New("Not authorized")
	}

	err = d.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("Error while deleting meetup: %v", err)
	}

	return true, nil
}
