// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package product_notices

import (
	"github.com/adacta-ru/mattermost-server/v6/app"
	tjobs "github.com/adacta-ru/mattermost-server/v6/jobs/interfaces"
)

type ProductNoticesJobInterfaceImpl struct {
	App *app.App
}

func init() {
	app.RegisterProductNoticesJobInterface(func(a *app.App) tjobs.ProductNoticesJobInterface {
		return &ProductNoticesJobInterfaceImpl{a}
	})
}
