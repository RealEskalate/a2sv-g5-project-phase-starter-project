package interfaces

type ProfileUpdateRepository interface {
	SaveProfileImageKey(userId string, imageKey string) error
}
