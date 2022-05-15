package authentication

type AuthEntity interface {
	InitialiseEntity(string, string)
	GetHeader() map[string]string
}
