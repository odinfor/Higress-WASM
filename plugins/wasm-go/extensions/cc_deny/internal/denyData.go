package internal

import (
	"cc_deny/internal/types"
)

var (
	hStore = &types.AccessStore{}
	cStore = &types.AccessStore{}
)

func InitHeaderStore() {
	c := types.NewConfDo().HeaderConf()

	if c.QPD > 0 {

	}
}

func InitCookieStore() {

}

func HeaderStore() *types.AccessStore {
	return hStore
}

func CookieStore() *types.AccessStore {
	return cStore
}
