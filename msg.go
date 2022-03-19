package msg_pb

const (
	ActionIdSecKill        = 1
	ActionIdMonitor        = 2
	ActionIdHasStockNotice = 3

	MsgOnline = 1
	MsgSendAll = 3
)

// 消息
type Message struct {
	ActionID int `json:"action_id"`
	Users []*User `json:"users"`
	SecKill *SecKillConf `json:"sec_kill"`
	Monitor *MonitorConf `json:"monitor"`
	HasStockSkuID string `json:"has_stock_sku_id"`
}

type SecKillConf struct {
	SkuIds []string
	BuyTime int64 // 毫秒时间戳
	StartTime int
	EndTime int
	Mode uint32
	RefreshTime int
	CommitCount int
}

type MonitorConf struct {
	SkuIds []string
	SingleTime int
	IntervalTime int
	CartInterval int
}

type User struct {
	NickName string
	Cookie string
}
