package models

type User struct {
	ID        string `json:"id" bson:"_id"`
	Email     string `json:"email,omitempty" bson:"email,omitempty"`
	Password  string `json:"password,omitempty" bson:"password,omitempty"`
	FirstName string `json:"firstName,omitempty" bson:"first_name,omitempty"`
	LastName  string `json:"lastName,omitempty" bson:"last_name,omitempty"`
	Phone     string `json:"phone,omitempty" bson:"phone,omitempty"`
}
