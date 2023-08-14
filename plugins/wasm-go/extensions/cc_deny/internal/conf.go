package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/alibaba/higress/plugins/wasm-go/pkg/wrapper"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tidwall/gjson"
	"time"
)

var c CCDeny

const (
	RATE  = "rate"
	BURST = "burst"
	QPS   = "qps"
	QPM   = "qpm"
	QPD   = "qpd"

	Header       = "header"
	Cookie       = "cookie"
	BlockSeconds = "block_seconds"

	CCRules = "cc_rules"
)

type (
	// CCDeny cc 识别和防护策略
	CCDeny struct {
		// 令牌桶速率
		Rate int `json:"rate"`

		// 令牌桶大小
		Burst int `json:"burst"`

		// 调用规则集
		RC []rc `json:"cc_rules"`
	}

	rc struct {

		// 从特定 header 识别调用方
		Header string `json:"header,omitempty"`

		// 从特定 cookie 识别调用方
		Cookie string `json:"cookie,omitempty"`

		// 每个调用方每秒最多调用次数
		QPS int `json:"qps,omitempty"`

		// 每个调用方每分钟最多调用次数
		QPM int `json:"qpm,omitempty"`

		// 每个调用方每天最多调用次数
		QPD int `json:"qpd,omitempty"`

		// 超过限制后将该调用方屏蔽，不可访问。(单位秒) 从小黑屋中放出来时会重置请求访问数据
		BlockSeconds int `json:"block_seconds,omitempty"`
	}

	HashK struct {
		headerV    string
		XForwarded string
		hashB      []byte
		hashS      string
	}
)

func NewHashKey(v string, x string) HashK {
	return HashK{
		headerV:    v,
		XForwarded: x,
	}
}

func (h HashK) Generate() {
	hasher := sha256.New()
	hasher.Write([]byte(h.headerV + ":" + h.XForwarded))
	h.hashS = hex.EncodeToString(hasher.Sum(nil))
	h.hashB = []byte(h.hashS)
}

func NewCCDenyDo() CCDenyDo {
	return &c
}

type CCDenyDo interface {
	Rules() []rc
	RateNum() int
	BurstNum() int
}

// ParseConfig 基于 Higress 配置转换传入的 json 初始化解析配置
func ParseConfig(json gjson.Result, config *CCDeny, log wrapper.Log) error {
	log.Info("aa")
	config.Rate = int(json.Get(RATE).Int())
	config.Burst = int(json.Get(BURST).Int())

	rule := json.Get(CCRules).Array()
	config.RC = make([]rc, len(rule))

	for k, v := range rule {
		if v.Get(Header).Exists() {
			config.RC[k].Header = v.Get(Header).String()
		}
		if v.Get(Cookie).Exists() {
			config.RC[k].Cookie = v.Get(Cookie).String()
		}
		if v.Get(QPS).Exists() {
			config.RC[k].QPS = int(v.Get(QPS).Int())
		}
		if v.Get(QPM).Exists() {
			config.RC[k].QPM = int(v.Get(QPM).Int())
		}
		if v.Get(QPD).Exists() {
			config.RC[k].QPD = int(v.Get(QPD).Int())
		}
		if v.Get(BlockSeconds).Exists() {
			config.RC[k].BlockSeconds = int(v.Get(BlockSeconds).Int())
		}
	}

	InitDeny()

	return nil
}

func (c CCDeny) Rules() []rc {
	return c.RC
}

func (c CCDeny) RateNum() int {
	return c.Rate
}

func (c CCDeny) BurstNum() int {
	return c.Burst
}

func OnHttpRequestHeaders(ctx wrapper.HttpContext, config CCDeny, log wrapper.Log) types.Action {
	header, err := proxywasm.GetHttpRequestHeaders()

	// todo get map 从 Higress share data 获取
	//proxywasm.GetSharedData(Header)
	//proxywasm.Share

	hdm := GetHDMap()
	//cdm := GetCDMap()
	accessTime := time.Now()

	for _, v := range header {
		//for _, y := range config.Rules() {
		//
		//}
		if hdm.enable {
			deny, ok := hdm.deny[v[0]]
			if !ok { // 初始化请求存储数据
				continue
			}

			// 当天内的访问量校验
			if deny.day.blocked && accessTime.Before(deny.day.freeTime) {
				_ = proxywasm.SendHttpResponse(200, nil, []byte(err.Error()), -1)
				return types.ActionPause
			}
			if deny.day.blocked && accessTime.After(deny.day.freeTime) {

			}

			if deny.dayUse && deny.day.account+1 <= deny.day.top {
				deny.day.account += 1
			} else {
				deny.day.account += 1
				deny.day.blocked = true
				deny.day.blockTime = accessTime
				deny.day.freeTime = accessTime.Add(time.Duration(deny.BlockSeconds) * time.Second)

				deny.Blocked = true
				deny.BlockReason = "Access limit exceeded today"

				_ = proxywasm.SendHttpResponse(200, nil, []byte(err.Error()), -1)
				return types.ActionPause
			}
		}
	}

	if err != nil {
		log.Errorf("get http request header fail, err: " + err.Error())
		_ = proxywasm.SendHttpResponse(200, nil, []byte(err.Error()), -1)
		return types.ActionPause
	}

	fmt.Println(header)

	//proxywasm.AddHttpRequestHeader("hello", "world")
	//if config.RateNum() > 0 {
	//	proxywasm.SendHttpResponse(200, nil, []byte("hello world"), -1)
	//}
	return types.ActionContinue
}

func hashKey(userAgent string, XForwarded string) string {
	hasher := sha256.New()
	hasher.Write([]byte(userAgent + ":" + XForwarded))
	return hex.EncodeToString(hasher.Sum(nil))
}
