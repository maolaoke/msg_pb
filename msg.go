package msg_pb

const (
	MsgOnline = 1
	MsgOffline = 2
	MsgTaskSecKill = 3
	MsgTasKMonitor = 4
	MsgHasStockNotice = 5
	MsgStopWork = 6

)

// 消息
type Message struct {
	MessageID 	  int `json:"message_id"`
	Users         []User `json:"users"`
	SecKill       SecKillConf   `json:"sec_kill"`
	Monitor       MonitorConf   `json:"monitor"`
	HasStockSkuID string         `json:"has_stock_sku_id"`
}

type SecKillConf struct {
	SkuIds      []string `json:"sku_ids"`
	BuyTime     int64    `json:"buy_time"` // 毫秒时间戳
	StartTime   int      `json:"start_time"`
	EndTime     int      `json:"end_time"`
	Mode        uint32   `json:"mode"`
	RefreshTime int      `json:"refresh_time"`
	CommitCount int      `json:"commit_count"`
}

type MonitorConf struct {
	SkuIds       []string `json:"sku_ids"`
	SingleTime   int      `json:"single_time"` // 单账号监控时间
	IntervalTime int      `json:"interval_time"` // 提交间隔时间
	CartInterval int      `json:"cart_interval"` // 重新加车时间
}

type User struct {
	NickName string
	Cookie string
}
