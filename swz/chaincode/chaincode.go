package main

import (
	"chaincode/api"
	"chaincode/model"
	"chaincode/pkg/utils"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/swz/chaincode/model"
)

type BlockChainTravelService struct {
}

// Init 链码初始化
func (t *BlockChainTravelService) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("链码初始化")
	//初始化商家账户
	var accountIds = [3]string{
		"S55236",
		"S10281",
		"S39120",
	}
	var storeNames = [3]string{"票务1", "票务2", "票务3"}
	var balances = [3]float64{9800000, 4500000, 6200000}
	var addresses = [3]string{"山西省晋中市xxx", "山西省太原市xxx", "山西省介休市xxx"}
	var cerficates = [3]string{"6e885a5c0b", "c8575a62b", "b1269a22c"}
	//初始化账号数据
	for i, val := range accountIds {
		store := &model.Store{
			Name:        storeNames[i],
			StoreID:     val,
			Balance:     balances[i],
			Address:     addresses[i],
			Certificate: cerficates[i],
		}
		// 写入账本
		if err := utils.WriteLedger(store, stub, model.StoreKey, []string{storeNames[i], val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	//初始化顾客账户
	var customerIds = [3]string{
		"C15842",
		"C23364",
		"C22653",
	}

	var customerNames = [3]string{"bob", "alice", "stallone"}
	var balances1 = [3]float64{5000, 2600, 4963}
	for i, val := range customerIds {
		customer := &model.Customer{
			Name:    customerNames[i],
			CTID:    val,
			Balance: balances1[i],
		}
		// 写入账本
		if err := utils.WriteLedger(customer, stub, model.CustomerKey, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	//初始化商品Init
	createTime, _ := stub.GetTxTimestamp()
	var commodityNames = [5]string{"绵山门票", "平遥古城门票", "张壁古堡门票", "五台山门票", "云冈石窟门票"}
	var prices = [5]float64{100, 50, 30, 50, 100}

	var providers = [5]string{"票务3", "票务1", "票务3", "票务2", "票务2"}
	for i, val := range commodityNames {
		commodity := &model.Commodity{
			Name:       val,
			Price:      prices[i],
			CreateTime: time.Unix(int64(createTime.GetSeconds()), int64(createTime.GetNanos())).Local().Format("2006-01-02 15:04:05"),
			Provider:   providers[i],
		}
		// 写入账本
		if err := utils.WriteLedger(commodity, stub, model.CommodityKey, []string{val, providers[i]}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainTravelService) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	funcName, args := stub.GetFunctionAndParameters()
	switch funcName {
	case "hello":
		return api.Hello(stub, args)
	case "queryCustomerList": //
		return api.QueryCustomerList(stub, args)
	case "queryStoreList":
		return api.QueryStoreList(stub, args)
	case "createRemark":
		return api.CreateRemark(stub, args)
	case "createCustomer":
		return api.CreateCustomer(stub, args)
	case "createSold":
		return api.CreateSold(stub, args)
	case "createBuy":
		return api.CreateBuy(stub, args)
	case "queryTxList":
		return api.QueryTxList(stub, args)
	case "createStore":
		return api.CreateStore(stub, args)
	case "queryStoreByNameAndID":
		return api.QueryStoreByNameAndID(stub, args)
	case "queryCommdityList":
		return api.QueryCommdityList(stub, args)

	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainTravelService))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
