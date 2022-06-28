package list

import (
	"github.com/analogj/justvanish/pkg/config"
	"github.com/sirupsen/logrus"
)

type ListAction struct {
	logger        *logrus.Entry
	configuration config.Interface
}

func New(logger *logrus.Entry, configuration config.Interface) (ListAction, error) {

	return ListAction{
		logger:        logger,
		configuration: configuration,
	}, nil
}

func (a *ListAction) Start() error {
	return nil
}
