// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package einterfaces

import (
	"github.com/adacta-ru/mattermost-server/v6/model"
)

type ComplianceInterface interface {
	StartComplianceDailyJob()
	RunComplianceJob(job *model.Compliance) *model.AppError
}
