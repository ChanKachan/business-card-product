package internal

import (
	"context"

	postgre "github.com/ChanKachan/business-card-product/pkg/postgreWrapper"
)

type UserInfo struct {
	AccountID int
}

type CtxQuery struct {
	ctx      context.Context
	Tx       postgre.PostgreTx
	UserInfo UserInfo
}
