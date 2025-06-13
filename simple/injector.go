//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitializeSimpleService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)

	return NewSimpleService(NewSimpleRepository(isError))
}
