package receiver

import (
	"fmt"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

type vmContext struct {
	// Embed the default VM context here,
	// so that we don't need to reimplement all the methods.
	types.DefaultVMContext
}

func SetVMContext() {
	proxywasm.SetVMContext(&vmContext{})
}

// Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &receiverPluginContext{contextID: contextID}
}

type receiverPluginContext struct {
	// Embed the default plugin context here,
	// so that we don't need to reimplement all the methods.
	contextID uint32
	types.DefaultPluginContext
	queueName string
}

// Override types.DefaultPluginContext.
func (ctx *receiverPluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	// Get Plugin configuration.
	config, err := proxywasm.GetPluginConfiguration()
	if err != nil {
		panic(fmt.Sprintf("failed to get plugin config: %v", err))
	}

	// Treat the config as the queue name for receiving.
	ctx.queueName = string(config)

	queueID, err := proxywasm.RegisterSharedQueue(ctx.queueName)
	if err != nil {
		panic("failed register queue")
	}
	proxywasm.LogInfof("queue \"%s\" registered as queueID=%d by contextID=%d", ctx.queueName, queueID, ctx.contextID)
	return types.OnPluginStartStatusOK
}

// Override types.DefaultPluginContext.
func (ctx *receiverPluginContext) OnQueueReady(queueID uint32) {
	data, err := proxywasm.DequeueSharedQueue(queueID)
	switch err {
	case types.ErrorStatusEmpty:
		return
	case nil:
		proxywasm.LogInfof("(contextID=%d) dequeued data from %s(queueID=%d): %s", ctx.contextID, ctx.queueName, queueID, string(data))
	default:
		proxywasm.LogCriticalf("error retrieving data from queue %d: %v", queueID, err)
	}
}
