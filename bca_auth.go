package bca

import (
	"context"

	"github.com/juju/errors"
	bcaCtx "github.com/purwaren/bca-api/context"
)

func (b *BCA) DoAuthentication(ctx context.Context) (*AuthToken, error) {
	ctx = bcaCtx.With(ctx, bcaCtx.BCASessID(b.bcaSessID))

	b.log(ctx).Info("=== START DO_AUTH ===")

	dtoResp, err := b.api.postGetToken(ctx)
	if err != nil {
		b.log(ctx).Error(errors.Details(err))
		return nil, errors.Trace(err)
	}

	b.setAccessToken(dtoResp.AccessToken)

	b.log(ctx).Info("=== END DO_AUTH ===")

	return dtoResp, nil
}
