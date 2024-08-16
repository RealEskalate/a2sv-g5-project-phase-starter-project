package domain

type Profile struct {
	First_Name      string        `bson:"first_name" json:"first_name"`
	Last_Name       string        `bson:"last_name" json:"last_name" `
	Bio             string        `bson:"bio" json:"bio"`
	Profile_Picture string        `bson:"profile_picture" json:"profile_picture"`
	Contact_Info    []ContactInfo `bson:"contact_info" json:"contact_info"`
}

type ProfileResponse struct {
	First_Name      string        `bson:"first_name" json:"first_name"`
	Last_Name       string        `bson:"last_name" json:"last_name" `
	Bio             string        `bson:"bio" json:"bio"`
	Profile_Picture string        `bson:"profile_picture" json:"profile_picture"`
	Contact_Info    []ContactInfo `bson:"contact_info" json:"contact_info"`
}

type ContactInfo struct {
	Address      string `bson:"address"`
	Phone_number string `bson:"phone_number" json:"phone_number"`
}
