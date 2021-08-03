package goupbit

var Holder = &holder{}

type holder struct {
	accessKey string
	secretKey string
}

func (pHolder *holder) GetAccessKey() string {
	return pHolder.accessKey
}

func (pHolder *holder) GetSecretKey() string {
	return pHolder.secretKey
}
