package api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"math/rand"


	"github.com/swz/chaincode/model"
	"github.com/swz/chaincode/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//创建 添加商品 函数 [name price remarks solder]   key:storeID
func CreateSold(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	name := args[0]
	price := args[1]
	//	remark := args[3]
	provider := args[2]

	if name == "" || price == "" || provider == "" {
		return shim.Error("参数存在空值")
	}
	// 参数数据格式转换
	var formattedPrice float64
	if val, err := strconv.ParseFloat(price, 64); err != nil {
		return shim.Error(fmt.Sprintf("price参数格式转换出错: %s", err))
	} else {
		formattedPrice = val
	}

	//将参数存入commodity实体
	createTime, _ := stub.GetTxTimestamp()
	commodity := &model.Commodity{
		Name:       name,
		Price:      formattedPrice,
		CreateTime: time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05"),
		Provider:   provider,
	}

	//写入账本 key：CommodityKey  包含两个索引的key： name和provider；可以根据部分key查询数据 但必须是前一个key查询
	if err := utils.WriteLedger(commodity, stub, model.CommodityKey, []string{commodity.Name, commodity.Provider}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	//将成功创建的信息返回
	commodityByte, err := json.Marshal(commodity)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(commodityByte)

}

// 用户购买商品函数  commodity/seller/ customerID
func CreateBuy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error("参数个数不满足")
	}
	commdityName := args[0]
	storeName := args[1]
	customerID := args[2]
	if commdityName == "" || storeName == "" || customerID == "" {
		return shim.Error("参数有空值")
	}

	//根据commity和provider获取想要购买的商品信息，确认存在该商品
	resultsCommdity, err := utils.GetStateByPartialCompositeKeys2(stub, model.CommodityKey, []string{commdityName, storeName})
	if err != nil || len(resultsCommdity) != 1 {
		return shim.Error(fmt.Sprintf("根据%s和%s获取商品信息失败: %s", commdityName, storeName, err))
	}
	var commodity model.Commodity
	if err = json.Unmarshal(resultsCommdity[0], &commodity); err != nil {
		return shim.Error(fmt.Sprintf("CreateBuy-反序列化出错：%s", err))
	}

	//获取买家信息
	resultsCustomer, err := utils.GetStateByPartialCompositeKeys(stub, model.CustomerKey, []string{customerID})
	if err != nil || len(resultsCustomer) != 1 {
		return shim.Error(fmt.Sprintf("customer信息获取失败%s", err))
	}
	var buyerAccount model.Customer
	if err = json.Unmarshal(resultsCustomer[0], &buyerAccount); err != nil {
		return shim.Error(fmt.Sprintf("查询customer信息-反序列化出错: %s", err))
	}
	//判断余额是否充足
	if buyerAccount.Balance < commodity.Price {
		return shim.Error(fmt.Sprintf("商品售价%f,您的当前余额为%f,购买失败", commodity.Price, buyerAccount.Balance))
	}

	//获取店铺信息
	resultsStores, err := utils.GetStateByPartialCompositeKeys(stub, model.StoreKey, []string{storeName})
	if err != nil || len(resultsStores) != 1 {
		return shim.Error(fmt.Sprintf("customer信息获取失败%s", err))
	}
	var store model.Store
	if err = json.Unmarshal(resultsStores[0], &store); err != nil {
		return shim.Error(fmt.Sprintf("查询store信息-反序列化出错: %s", err))
	}

	//将buyerAccount加入Commodity的CustomerList列表中
	commodity.CustomerList = append(commodity.CustomerList, buyerAccount.CTID)

	//更新 commodity 的信息
	//写入账本 key：CommodityKey  包含两个索引的key： name和provider；可以根据部分key查询数据 但必须是前一个key查询
	if err := utils.WriteLedger(commodity, stub, model.CommodityKey, []string{commodity.Name, commodity.Provider}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	createTime, _ := stub.GetTxTimestamp()
	//将本次购买交易写入账本,可供买家查询
	sellingBuy := &model.TxRecord{
		Customers:   buyerAccount,
		CreateTime:  time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05"),
		Commodities: commodity,
		Stores:      store,
	}
	if err := utils.WriteLedger(sellingBuy, stub, model.TxRecordKey, []string{sellingBuy.Customers.CTID, sellingBuy.CreateTime}); err != nil {
		return shim.Error(fmt.Sprintf("将本次购买交易写入账本失败%s", err))
	}
	sellingBuyByte, err := json.Marshal(sellingBuy)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}

	//////////////////////////////////注意 店铺余额更新
	//购买成功，扣取余额，更新账本余额，注意，此时需要卖家确认收款，款项才会转入卖家账户，此处先扣除买家的余额
	buyerAccount.Balance -= commodity.Price
	if err := utils.WriteLedger(buyerAccount, stub, model.CustomerKey, []string{buyerAccount.CTID}); err != nil {
		return shim.Error(fmt.Sprintf("扣取买家余额失败%s", err))
	}
	// 成功返回
	return shim.Success(sellingBuyByte)

}

//	查询订单
func QueryTxList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var TxList []model.TxRecord
	resultTxRecord, err := utils.GetStateByPartialCompositeKeys2(stub, model.TxRecordKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range resultTxRecord {
		if v != nil {
			var txRecord model.TxRecord
			err := json.Unmarshal(v, &txRecord)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryTxList-反序列化出错: %s", err))
			}
			TxList = append(TxList, txRecord)
		}
	}
	TxListByte, err := json.Marshal(TxList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryTxList-序列化出错: %s", err))
	}
	return shim.Success(TxListByte)
}


//商家注册 name/ storeid / address / commodity / cerficates / balance
func CreateStore(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("参数个数不正确")
	}
	name := args[0]
	address := args[1]
	certificate := args[2]
	balance := args[3]

	//s生成随机id
	rand.Seed(time.Now().UnixNano())
	id := (rand.Intn(10))*10000 + (rand.Intn(10))*1000 + (rand.Intn(10))*100 + (rand.Intn(10))*10 + (rand.Intn(10))

	storeid := "S"+string(id)

	//id是否重复？？？？

	// 参数数据格式转换
	var formattedPrice float64
	if val, err := strconv.ParseFloat(balance, 64); err != nil {
		return shim.Error(fmt.Sprintf("price参数格式转换出错: %s", err))
	} else {
		formattedPrice = val
	}

	store := &model.Store{
		Name: name,
		StoreID: storeid,
		Address:address,
		Certificate:certificate,
		Balance:formattedPrice,
	}

	//  - storename----storeid
	if err := utils.WriteLedger(store, stub, model.StoreKey, []string{store.Name, store.StoreID}); err != nil {
		return shim.Error(fmt.Sprintf("将本次购买交易写入账本失败%s", err))
	}
	storeByte, err := json.Marshal(store)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}

	return shim.Success(storeByte)

}
