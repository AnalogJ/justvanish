package request

import (
	"github.com/analogj/justvanish/pkg/actions"
	"github.com/analogj/justvanish/pkg/config"
	"github.com/sirupsen/logrus"
)

type RequestAction struct {
	*actions.CommonAction
}

func New(logger *logrus.Entry, configuration config.Interface) (RequestAction, error) {
	return RequestAction{
		CommonAction: &actions.CommonAction{
			Logger:        logger,
			Configuration: configuration,
			ActionType:    "request",
		},
	}, nil
}
