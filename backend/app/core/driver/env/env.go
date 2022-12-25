package env

import "os"

type Env string

const (
	prod  Env = "prod"
	prev  Env = "prev"
	dev   Env = "dev"
	local Env = "local"
	empty Env = ""
)

var env Env //nolint:gochecknoglobals

func Init() {
	env = Env(os.Getenv("ENV"))
}

func Get() Env {
	return env
}

func (e Env) String() string {
	return string(e)
}

func (e Env) IsProd() bool {
	return e == prod
}
