package delete

import (
	"github.com/analogj/justvanish/pkg/actions"
	"github.com/analogj/justvanish/pkg/config"
	"github.com/sirupsen/logrus"
)

type DeleteAction struct {
	*actions.CommonAction
}

func New(logger *logrus.Entry, configuration config.Interface) (DeleteAction, error) {
	return DeleteAction{
		CommonAction: &actions.CommonAction{
			Logger:        logger,
			Configuration: configuration,
			ActionType:    "delete",
		},
	}, nil
}
