package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"github.com/userq11/meetmeup/graph/model"
)

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
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A meetup 2",
		Description: "A description 2",
		UserID:      "2",
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
