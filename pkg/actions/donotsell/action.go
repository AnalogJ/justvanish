package donotsell

import (
	"github.com/analogj/justvanish/pkg/actions"
	"github.com/analogj/justvanish/pkg/config"
	"github.com/sirupsen/logrus"
)

type DoNotSellAction struct {
	*actions.CommonAction
}

func New(logger *logrus.Entry, configuration config.Interface) (DoNotSellAction, error) {
	return DoNotSellAction{
		CommonAction: &actions.CommonAction{
			Logger:        logger,
			Configuration: configuration,
			ActionType:    "donotsell",
		},
	}, nil
}
