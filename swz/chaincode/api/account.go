package api

import (
	"encoding/json"
	"fmt"
	"go_code/xuperchain-master/core/contractsdk/go/pb"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/swz/chaincode/model"
	"github.com/swz/chaincode/utils"
)

// QueryAccountList 查询顾客账户
func QueryCustomerList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var customerList []model.Customer
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.CustomerKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var customer model.Customer
			err := json.Unmarshal(v, &customer)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryCustomerList-反序列化出错: %s", err))
			}
			customerList = append(customerList, customer)
		}
	}
	customerListByte, err := json.Marshal(customerList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryCustomerList-序列化出错: %s", err))
	}
	return shim.Success(customerListByte)
}

// QueryAccount 查询商家
func QueryStoreByNameAndID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var storeList []model.Store
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.StoreKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var store model.Store
			err := json.Unmarshal(v, &store)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryCustomerList-反序列化出错: %s", err))
			}
			storeList = append(storeList, store)
		}
	}
	storeListByte, err := json.Marshal(storeList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryCustomerList-序列化出错: %s", err))
	}
	return shim.Success(storeListByte)
}

//列举所有的商家 args参数为空
func QueryStoreList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var storeList []model.Store
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.StoreKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var store model.Store
			err := json.Unmarshal(v, &store)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryCustomerList-反序列化出错: %s", err))
			}
			storeList = append(storeList, store)
		}
	}
	storeListByte, err := json.Marshal(storeList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryCustomerList-序列化出错: %s", err))
	}
	return shim.Success(storeListByte)
}

//查询commdity列表
func QueryCommdityList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var commdityList []model.Commodity
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.CommodityKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}

	for _, v := range results {
		if v != nil {
			var commdity model.Commodity
			err := json.Unmarshal(v, &commdity)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryCommdityList-反序列化出错: %s", err))
			}
			commdityList = append(commdityList, commdity)
		}
	}
	commdityListByte, err := json.Marshal(commdityList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryCommdityList-序列化出错: %s", err))
	}
	return shim.Success(commdityListByte)
}
