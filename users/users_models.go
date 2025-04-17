package users

import "github.com/uptrace/bun"

// User record
type User struct {
	bun.BaseModel `bun:"table:users"`
	Id            int64      `json:"id" bun:"id,pk,autoincrement"`
	Username      string     `json:"username,omitempty"`
	Password      string     `json:"password"`
	Roles         []string   `json:"roles" bun:"roles,array"`
	Status        UserStatus `json:"status"`
}
