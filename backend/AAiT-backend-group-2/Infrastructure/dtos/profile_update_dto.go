package dtos


type ProfileUpdate struct {
	Username		string		`json:"username"`
	Bio				string 		`json:"bio"`
	ProfilePic		string 		`json:"profile_pic"`
	ContactInfo		string		`json:"contact_info"`
}