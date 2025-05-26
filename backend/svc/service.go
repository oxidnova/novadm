package svc

import (
	"os"
)

// Service it defines a runtime service
type Service interface {
	Name() string
	Load(string) error
	Run()
	Stop(os.Signal)
}
