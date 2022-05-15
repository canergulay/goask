package authentication

type EksiAuthEntity struct {
	Token string
}

func (eae *EksiAuthEntity) InitialiseEntity(username string, password string) {
	// TODO: API CALL AND GET THE TOKEN.
	eae.Token = "TOKEN GOT FROM API"
}

func (eae EksiAuthEntity) GetHeader() map[string]string {
	return map[string]string{
		"Authentication": eae.Token,
	}
}
