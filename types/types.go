package types

type Model[T any] interface {
	*T
	TableName() string
}
