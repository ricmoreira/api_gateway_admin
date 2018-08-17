package models

type Role struct {
	ID    string `json:"id,omitempty"`
	Role  string `json:"role" valid:"required~Field role cannot be empty"`
	Level int    `json:"level" valid:"required~Field level cannot be empty or is missing"`
}
