package clientconfig

type Database struct {
	Database string `validate:"required"`
	Name     string `validate:"required"`
	Host     string `validate:"required"`
	Port     string `validate:"required"`
	User     string `validate:"required"`
	Password string `validate:"required"`
}
