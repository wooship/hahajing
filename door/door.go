package door

import (
	"hahajing/com"
)

const (
	keywordManagerUpdateCheckTimer = 10 // minute
	keywordUpdateTimer             = 1  // second

	keywordManagerUpdateHour = 1 // 几点开始更新

	keywordCheckReqChSize = 1000
)

// KeywordCheckReq x
type KeywordCheckReq struct {
	MyKeyword *com.MyKeyword

	ResCh chan *KeywordCheckRes
}

// KeywordCheckRes x
type KeywordCheckRes struct {
	Items    []*com.Item
	ErrorStr string
}

// Door x
type Door struct {
	keywordManager    *com.KeywordManager
	KeywordCheckReqCh chan *KeywordCheckReq
}

// Start x
func (d *Door) Start(keywordManager *com.KeywordManager) {
	d.keywordManager = keywordManager
	d.KeywordCheckReqCh = make(chan *KeywordCheckReq, keywordCheckReqChSize)

}
