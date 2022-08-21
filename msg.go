package msg_pb

import "time"

const (
	MsgOnline         = 1
	MsgOffline        = 2
	MsgTaskSecKill    = 3
	MsgTasKMonitor    = 4
	MsgHasStockNotice = 5
	MsgStopWork       = 6
	MsgGetClients     = 7
	MsgSuccessOrder   = 8
	MsgServerTime     = 9  // 服务器时间。当使用阿里云服务器时，与京东服务器时间差较大，使用http请求去获取京东服务器时间时会由于网络环境造成误差较大，改成获取主控服务器时间。
	MsgKillServer     = 10 // 退出进程

	MsgHeart = 100

	AdminClientName = "admin"

	// 下单接口类型
	JdAppApi = 1
	JxAppApi = 2
	JdMApi   = 3
)

// 消息
type Message struct {
	MessageID     int         `json:"message_id"`
	ClientName    string      `json:"client_name"` // MsgOnline使用
	Users         []User      `json:"users"`
	MaxUserNum    int         `json:"max_user_num"` // 用于账号分配时限制每台服务器最多多少用户
	SecKill       SecKillConf `json:"sec_kill"`
	Monitor       MonitorConf `json:"monitor"`
	HasStockSkuID string      `json:"has_stock_sku_id"` // MsgHasStockNotice
	Clients       []string    `json:"clients"`
	SuccessOrders []SuccessOrderInfo
	Now           time.Time
	IsModifyAddr  int `json:"is_modify_addr"`
}

type SecKillConf struct {
	SkuIds       []string      `json:"sku_ids"`
	BuyTime      int64         `json:"buy_time"` // 毫秒时间戳
	StartTime    int           `json:"start_time"`
	EndTime      int           `json:"end_time"`
	Mode         uint32        `json:"mode"`
	RefreshTime  int           `json:"refresh_time"`
	CommitCount  int           `json:"commit_count"`
	BeforeTime   int           `json:"before_time"`
	AddressInfo  AddressInfo   `json:"address_info"`
	ApiType      int           `json:"api_type"`   // 接口类型
	IsYuShou     int           `json:"is_yu_shou"` // 是否是定金模式。预售
	IsMck        int           `json:"is_mck"`     // 是否MCK 0否 1是
	PayType      int           `json:"pay_type"`   // 支付类型 4：普通 5：对公转账
	AddressInfos []AddressInfo `json:"address_infos"`
	Addr2Rate    int           `json:"addr_2_rate"`
}

type MonitorConf struct {
	SkuIds         []string    `json:"sku_ids"`
	IntervalTime   int         `json:"interval_time"` // 提交间隔时间
	SubmitType     int         `json:"submit_type"`   // 1 app 2 qq 3 京喜
	Threshold      int         `json:"threshold"`     // 收藏夹返回多少个以上有货则判断为误判
	AddressInfo    AddressInfo `json:"address_info"`
	MonitorUsers   []User      `json:"monitor_users"`
	SubmitUsers    []User      `json:"submit_users"`
	MonitorUserNum int         `json:"monitor_user_num"`
	SubmitUserNum  int         `json:"submit_user_num"`
}

type User struct {
	NickName string      `json:"nick_name"`
	Cookie   string      `json:"cookie"`
	Address  AddressInfo `json:"address"`
}

type SuccessOrderInfo struct {
	UserName   string    `json:"user_name"`
	Cookie     string    `json:"cookie"`
	SkuID      string    `json:"sku_id"`
	SubmitTime time.Time `json:"submit_time"`
	DiffTime   int       `json:"diff_time"`
	OrderID    string    `json:"order_id"`
	Price      string    `json:"price"`
	MillTime   int       `json:"mill_time"` // 设置毫秒数
}

type AddressInfo struct {
	Province      string `json:"province"`
	ProvinceId    string `json:"province_id"`
	FullAddress   string `json:"full_address"`
	CityID        string `json:"city_id"`
	City          string `json:"city"`
	CountyID      string `json:"county_id"`
	County        string `json:"county"`
	TownID        string `json:"town_id"`
	Town          string `json:"town"`
	AddressDetail string `json:"address_detail"`
	Name          string `json:"name"`
	Mobile        string `json:"mobile"`
	IsRandAddr    int    `json:"is_rand_addr"`
}
