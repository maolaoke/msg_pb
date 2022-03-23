package msg_pb

import "time"

const (
	MsgOnline = 1
	MsgOffline = 2
	MsgTaskSecKill = 3
	MsgTasKMonitor = 4
	MsgHasStockNotice = 5
	MsgStopWork = 6
	MsgGetClients = 7
	MsgSuccessOrder = 8

	AdminClientName = "admin"
)

// 消息
type Message struct {
	MessageID 	  int `json:"message_id"`
	ClientName    string `json:"client_name"` // MsgOnline使用
	Users         []User `json:"users"`
	SecKill       SecKillConf   `json:"sec_kill"`
	Monitor       MonitorConf   `json:"monitor"`
	HasStockSkuID string         `json:"has_stock_sku_id"`  // MsgHasStockNotice
	Clients       []string 	`json:"clients"`
	SuccessOrders  []SuccessOrderInfo
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
	NickName string `json:"nick_name"`
	Cookie string `json:"cookie"`
}

type SuccessOrderInfo struct {
	UserName string `json:"user_name"`
	SkuID string	`json:"sku_id"`
	SubmitTime time.Time `json:"submit_time"`
	DiffTime int `json:"diff_time"`
	OrderID string `json:"order_id"`
	Price string `json:"price"`
}