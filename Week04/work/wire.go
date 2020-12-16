// +build wireinject
// The build tag makes sure the stub is not built in the final build.
package main

import (
	_"Go-000/Week04/work/models"
	_"github.com/google/wire"
)

// InitializeEvent 声明injector的函数签名
/*func InitializeEvent(msg string) models.Event{
	wire.Build(models.NewEvent, models.NewGreeter, models.NewMessage)
	return models.Event{}
}*/