package bot

import (
	"context"
)

type Driver interface {
	Context() context.Context
	Next() bool
}

//go:generate moq -out mock/driver.go -pkg mock . Driver
