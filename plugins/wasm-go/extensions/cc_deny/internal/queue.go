package internal

import "github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"

// TokenQ
// @Description: 共享队列.
type TokenQ struct {
	name string
	qid  uint32
	size uint32
}

func Init() {

	qid, err := proxywasm.RegisterSharedQueue("tokenQ")

	// 在 DequeueSharedQueue EnqueueSharedQueue 之前使用，获取 ququeID
	proxywasm.ResolveSharedQueue()

	// 出列,
	proxywasm.DequeueSharedQueue()

	// 进列,
	proxywasm.EnqueueSharedQueue()

	proxywasm.GetSharedData()

	proxywasm.SetSharedData()

}

func refuse() {
	proxywasm.SendHttpResponse(403, nil, []byte("denied by cc"), -1)
}
