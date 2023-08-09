package internal

// TokenQ
// @Description: token 本地内存的存储队列.
type TokenQ struct {
	// token 表
	token []string

	// 维护当前token表的长度
	len int

	// 最大容量
	cap int
}

func NewTQ(cap int) *TokenQ {
	return &TokenQ{token: make([]string, cap), cap: cap}
}

func (t *TokenQ) Push(k string) bool {
	if t.len >= t.cap {
		return false
	}
	t.token = append(t.token, k)
	t.len += 1
	return true
}

func (t *TokenQ) Pop() (string, error) {
	if t.len <= 0 {
		return "", nil
	}
	k := t.token[t.len]
	t.token = t.token[:t.len]
	t.len -= 1
	return k, nil
}

func (t *TokenQ) IsEmpty() bool {
	if t.len == 0 {
		return true
	}
	return false
}
