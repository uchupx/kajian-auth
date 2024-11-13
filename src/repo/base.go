package repo

import (
	"github.com/gofrs/uuid"
	"github.com/uchupx/kajian-api/pkg/helper"
	"github.com/uchupx/kajian-api/pkg/logger"
)

type BaseRepo struct{}

func (m BaseRepo) ID() *string {
	val, err := uuid.NewV7()
	if err != nil {
		logger.Logger.Errorf("[BaseRepo - ID] failed to generating uuid v7, %+v", err)
		return nil
	}

	return helper.StringToPointer(val.String())
}
