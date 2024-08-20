package Domain

import "time"

type Account struct {
	RawData           map[string]interface{}
	Provider          string
	Email             string
	Name              string
	FirstName         string
	LastName          string
	NickName          string
	Description       string
	UserID            string
	AvatarURL         string
	Location          string
	AccessToken       string
	AccessTokenSecret string
	RefreshToken      string
	ExpiresAt         time.Time
	IDToken           string
}

// map[
// 	email:abel.bekele@a2sv.org
// 	email_verified:true
// 	family_name:Bekele
// 	given_name:Abel hd:a2sv.org
// 	name:Abel Bekele
// 	picture:https://lh3.googleusercontent.com/a/ACg8ocJCEGdAMzQy7UQ_EcT8tMhEx1b002X_pdT3FJ1MAp2Zgqp_9g=s96-c sub:108246915207483235763

// 	]
