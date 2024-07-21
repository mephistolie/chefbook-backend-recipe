package entity

import "github.com/google/uuid"

const (
	RoleOwner    = "owner"
	RoleCoauthor = "coauthor"
)

type Contributor struct {
	Id   uuid.UUID
	Role string
}
