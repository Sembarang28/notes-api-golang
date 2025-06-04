//go:build wireinject
// +build wireinject

package pkg

import "github.com/google/wire"

func InitAuthController() {
	wire.Build()
}
