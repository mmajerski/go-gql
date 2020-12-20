package domain

import (
	"github.com/userq11/meetmeup/graph/model"
	"github.com/userq11/meetmeup/postgres"
)

type Domain struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
}

type Ownable interface {
	IsOwner(user *model.User) bool
}

func NewDomain(usersRepo postgres.UsersRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{UsersRepo: usersRepo, MeetupsRepo: meetupsRepo}
}

func checkOwnership(o Ownable, user *model.User) bool {
	return o.IsOwner(user)
}
