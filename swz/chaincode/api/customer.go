package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/swz/chaincode/model"
	"github.com/swz/chaincode/utils"
)

//评论函数   customer[id]/remark    commodities--要评论那个商品[name]
func CreateRemark(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 3 {
		return shim.Error(fmt.Sprintf("参数个数不足"))
	}
	customerID := args[0]
	commdityName := args[1]
	storeName := args[2]
	remark := args[3]

	/*
		//根据id获取买家信息
		resultsCustomer, err := utils.GetStateByPartialCompositeKeys(stub, model.CustomerKey, []string{customerID})
		if err != nil || len(resultsCustomer) != 1 {
			return shim.Error(fmt.Sprintf("customer信息获取失败%s", err))
		}
		var customer model.Customer
		if err = json.Unmarshal(resultsCustomer[0], &customer); err != nil {
			return shim.Error(fmt.Sprintf("查询customer信息-反序列化出错: %s", err))
		}
	*/

	//1.先判断该customer是否购买过这商品，是否符合评论的条件
	//1.0 查找与customerid相关的订单 交易记录
	TxListByte := QueryTxList(stub, []string{customerID})
	var TxList []model.TxRecord
	err := json.Unmarshal(TxListByte, &TxList)
	if err != nil {
		return shim.Error(fmt.Sprintf("CreateRemark错误:%s", err))
	}
	//1.1 根据customerID 查询订单内是否存在该顾客 并且加上限制条件--commodity
	var existFlag bool
	var timeTX string
	for _, v := range TxList {
		if v.Commodities.Name == commdityName {
			existFlag = true
			timeTX = v.CreateTime
			break
		}
	}

	//1.2 根据查询到的订单 查看该订单是否在可评论时间内
	if existFlag {
		//将string类型的timeTX转换为time类型的格式
		timeTx, err := time.Parse("2006-01-02 15.04.05", timeTX)
		if err != nil {
			return shim.Error(fmt.Sprintf("txTime格式转换失败:%s", err))
		}
		hour := time.Since(timeTx).Hours()

		//2.添加评论
		//2.1查找商品
		//获取想要购买的商品信息
		resultsCommdity, err := utils.GetStateByPartialCompositeKeys2(stub, model.CommodityKey, []string{commdityName, storeName})
		if err != nil || len(resultsCommdity) != 1 {
			return shim.Error(fmt.Sprintf("根据%s和%s获取商品信息失败: %s", commdityName, storeName, err))
		}
		var commodity model.Commodity
		if err = json.Unmarshal(resultsCommdity[0], &commodity); err != nil {
			return shim.Error(fmt.Sprintf("CreateBuy-反序列化出错：%s", err))
		}

		//2.2添加评论[在24h内,可以评论]
		if hour <= 24 {
			//将评论添加到商品
			commodity.Remarks = append(commodity.Remarks, remark)
		}
		//3.删除商品内顾客名字，过期不能评论  ---评论完也删除
		for i, value := range commodity.CustomerList {
			if value == customerID {
				commodity.CustomerList = append(commodity.CustomerList[:i], commodity.CustomerList[i+1:]...)
				break
			}
		}
		//4.json存数据
		//4.1 save commodity data.
		//写入账本 key：CommodityKey  包含两个索引的key： name和provider；可以根据部分key查询数据 但必须是前一个key查询
		if err := utils.WriteLedger(commodity, stub, model.CommodityKey, []string{commodity.Name, commodity.Provider}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}
	return shim.Success(fmt.Sprintf("评论成功！"))

}

//顾客注册 Name balance [id由系统随机分配]
func CreateCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error(fmt.Sprintf("参数个数不正确！"))
	}
	name := args[0]
	balance := args[1]

	// 参数数据格式转换
	var formattedPrice float64
	if val, err := strconv.ParseFloat(balance, 64); err != nil {
		return shim.Error(fmt.Sprintf("price参数格式转换出错: %s", err))
	} else {
		formattedPrice = val
	}

	//s生成随机id
	rand.Seed(time.Now().UnixNano())
	id := (rand.Intn(10))*10000 + (rand.Intn(10))*1000 + (rand.Intn(10))*100 + (rand.Intn(10))*10 + (rand.Intn(10))

	//id 是否重复？

	//将新注册的顾客存入区块链
	customer := &model.Customer{
		Name:    name,
		CTID:    "C" + string(id),
		Balance: formattedPrice,
	}

	if err := utils.WriteLedger(customer, stub, model.CustomerKey, []string{customer.CTID}); err != nil {
		return shim.Error(fmt.Sprintf("将本次购买交易写入账本失败%s", err))
	}
	customerByte, err := json.Marshal(customer)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	return shim.Success(customerByte)
}
