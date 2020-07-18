package repository

import (
	"github.com/gofrs/uuid"
)

type Reaction struct {
	UserId         uuid.UUID `db:"userId"`
	PresentationId int       `db:"presentationId"`
	Stamp          string    `db:"stamp"`
}

type ReactionStatistics struct {
	PresentationId int
	Counts         []Count
}

type Count struct {
	Stamp string `db:"stamp"`
	Count int    `db:"COUNT(stamp)"`
}

type ReactionRepository interface {
	CreateReaction(reaction *Reaction) error
	GetReactionStatistics(id int) (*ReactionStatistics, error)
}
