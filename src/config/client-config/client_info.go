package clientconfig

type Client struct {
	Key         string      `validate:"required"`
	Name        string      `validate:"required"`
	Database    Database    `validate:"required"`
	DataPattern DataPattern `validate:"required"`
}
