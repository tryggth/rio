package dashboard

import (
	"context"

	"github.com/rancher/rio/modules/dashboard/feature"
	"github.com/rancher/rio/types"
)

func Register(ctx context.Context, rContext *types.Context) error {
	return feature.Register(ctx, rContext)
}
