package v1

import (
	bc "application/blockchain"
	"application/pkg/app"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerIdBody struct {
	CustomerId string `json:"customerId"`
}

type CustomerRequestBody struct {
	Args []CustomerIdBody `json:"args"`
}

type StoreIdBody struct {
	StoreId string `json:"storeId"`
}

type StoreRequestBody struct {
	Args []StoreIdBody `json:"args"`
}

type CreateStoreRequestBody struct {
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Certificate string  `json:"certificate"`
	Balance     float64 `json:"balance"`
}

type CreateCustomerRequestBody struct {
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

type QueryAccountListRequestBody struct {
	Name    string `json:"name"`
	StoreID string `json:"storeID"`
}
type QueryAllAccountListRequestBody struct {
}

//id query
func QueryCustomerList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CustomerRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	for _, val := range body.Args {
		bodyBytes = append(bodyBytes, []byte(val.CustomerId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryCustomerList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// name id 查询
func QueryStoreByNameAndID(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(QueryAccountListRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Name == "" && body.StoreID == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(body.StoreID))

	//调用智能合约
	resp, err := bc.ChannelQuery("queryStoreByNameAndID", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

// 列举所有商家
func QueryStoreList(c *gin.Context) {
	appG := app.Gin{C: c}
	/*
		body := new(QueryAllAccountListRequestBody)
		//解析Body参数
		if err := c.ShouldBind(body); err != nil {
			appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
			return
		}
		if body.Name == "" && body.StoreID == "" {
			appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
			return
		}
	*/

	var bodyBytes [][]byte

	//调用智能合约
	resp, err := bc.ChannelQuery("queryStoreList", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	// 反序列化json
	var data []map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

//商家注册 name address certificate balance
func CreateStore(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CreateStoreRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Name == "" || body.Address == "" || body.Certificate == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
		return
	}
	if body.Balance < 0 {
		appG.Response(http.StatusBadRequest, "失败", "Balance账户余额必须大于0")
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(body.Address))
	bodyBytes = append(bodyBytes, []byte(body.Certificate))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.Balance, 'E', -1, 64)))
	//调用智能合约
	resp, err := bc.ChannelExecute("createStore", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}

//顾客注册 name  balance /id由系统随机分配
func CreateCustomer(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CreateCustomerRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Name == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
		return
	}
	if body.Balance < 0 {
		appG.Response(http.StatusBadRequest, "失败", "Balance账户余额必须大于0")
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.Balance, 'E', -1, 64)))
	//调用智能合约
	resp, err := bc.ChannelExecute("createCustomer", bodyBytes)
	if err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	var data map[string]interface{}
	if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
		appG.Response(http.StatusInternalServerError, "失败", err.Error())
		return
	}
	appG.Response(http.StatusOK, "成功", data)
}
