package model

// Meetup struct type
type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	User        string `json:"user"`
}
