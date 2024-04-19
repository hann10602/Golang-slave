package model

type User struct {
	Username  string `json:"username,omitempty" bson:"username,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Role      string `json:"role" bson:"role"`
	CountryId string `json:"countryId" bson:"countryId"`
}
