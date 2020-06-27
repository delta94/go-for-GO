package handler

import (
	"auth/dao"
	"github.com/go-playground/validator/v10"
	"github.com/micro/go-micro/v2/broker"
)

type auth struct {
	mq       broker.Broker
	adc      *dao.AuthDAOCreator
	validate *validator.Validate
}

func NewAuth(mq broker.Broker, adc *dao.AuthDAOCreator, validate *validator.Validate) *auth {
	return &auth{
		mq:       mq,
		adc:      adc,
		validate: validate,
	}
}
