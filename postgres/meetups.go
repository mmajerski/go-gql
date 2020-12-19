package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/userq11/meetmeup/graph/model"
)

type MeetupsRepo struct {
	DB *pg.DB
}

func (m *MeetupsRepo) GetMeetups() ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}

	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	return meetup, err
}
