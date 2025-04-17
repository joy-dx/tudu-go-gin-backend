package users

type UserStatus string

const (
	Initial UserStatus = "INITIAL"
	Active  UserStatus = "ACTIVE"
	Frozen  UserStatus = "FROZEN"
	Deleted UserStatus = "DELETED"
)
