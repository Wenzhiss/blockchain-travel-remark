package model

/*
	结构体中的必要信息上链，前端信息可以区块链与数据库结合展示。
*/

type Store struct {
	Name        string    `json:name`
	StoreID     string    `json:storeID`
	Address     string    `json:address`
	Commodities Commodity `json:commodities` //店铺商品
	Certificate string    `json:certificate`
	Balance     float64   `json:balance`
}

type Commodity struct {
	Name         string   `json:name`
	Price        float64  `json:price`
	Remarks      []string `json:remarks` //评论
	CreateTime   string   `json:createTime`
	Provider     string   `json:provider`     //供应商
	CustomerList []string `json:customerList` //用来判断用户是否能评论的标志字段
}

type Customer struct {
	Name    string  `json:name`
	CTID    string  `json:cTID`
	Balance float64 `json:balance`
}

//交易记录，供买家查询，  /*可以考虑通过时间限制来评论 */
type TxRecord struct {
	Customers   Customer  `json:customers`   //顾客ID
	CreateTime  string    `json:createTime`  //交易时间
	Commodities Commodity `json:commodities` //商品
	Stores      Store     `json:stores`
}

const (
	CommodityKey = "commodity-key"
	CustomerKey  = "customer-key"
	TxRecordKey  = "txRecord-key"
	StoreKey     = "store-key"
)
