package client

import (
	"github.com/heww/gomock-issue/types"
)

//go:generate go run github.com/golang/mock/mockgen -source=./client.go -destination=mocks/mock.go -package=mocks

type Client[T any, P types.Model[T]] interface {
	Create(*P) error
}
