package dto

type UpdateUserDto struct {
	Email *string `json:"email" bson:"email"`
	Role  *string `json:"role" bson:"role"`
}
