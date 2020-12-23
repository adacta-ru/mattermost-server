// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package jobs

import "github.com/adacta-ru/mattermost-server/v6/model"

type CloudJobInterface interface {
	MakeWorker() model.Worker
	MakeScheduler() model.Scheduler
}
