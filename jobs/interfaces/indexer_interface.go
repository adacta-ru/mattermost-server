// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package interfaces

import (
	"github.com/adacta-ru/mattermost-server/v6/model"
)

type IndexerJobInterface interface {
	MakeWorker() model.Worker
}
