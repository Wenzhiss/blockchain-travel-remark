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

type SellingRequestBody struct {
	CommdityName string `json:"commdityName"`
	StoreName    string `json:"storeName"`
	CustomerID   string `json:"customerID"`
}

type CreateSoldRequestBody struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Provider string  `json:"provider"`
}

type QueryTxListRequestBody struct {
	CustomerId string `json:"customerId"`
}

type RemarkRequestBody struct {
	CustomerID   string `json:"customerID"`
	CommdityName string `json:"commdityName"`
	StoreName    string `json:"storeName"`
	Remark       string `json:"remark"`
}

//创建交易
func CreateBuy(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(SellingRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.CommdityName == "" || body.CustomerID == "" || body.StoreName == "" {
		appG.Response(http.StatusBadRequest, "失败", "商品名、店铺名以及顾客id不能为空")
		return
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.CommdityName))
	bodyBytes = append(bodyBytes, []byte(body.StoreName))
	bodyBytes = append(bodyBytes, []byte(body.CustomerID))
	//调用智能合约
	resp, err := bc.ChannelExecute("createBuy", bodyBytes)
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

//上架商品
func CreateSold(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(CreateSoldRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	if body.Name == "" || body.Provider == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
		return
	}
	if body.Price <= 0 {
		appG.Response(http.StatusBadRequest, "失败", "Price价格必须大于0")
	}

	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.Name))
	bodyBytes = append(bodyBytes, []byte(strconv.FormatFloat(body.Price, 'E', -1, 64)))
	bodyBytes = append(bodyBytes, []byte(body.Provider))
	//调用智能合约
	resp, err := bc.ChannelExecute("createSold", bodyBytes)
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

func QueryTxList(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(QueryTxListRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}
	var bodyBytes [][]byte
	if body.CustomerId != "" {
		bodyBytes = append(bodyBytes, []byte(body.CustomerId))
	}
	//调用智能合约
	resp, err := bc.ChannelQuery("queryTxList", bodyBytes)
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

//customerid commdityname storename remark
func CreateRemark(c *gin.Context) {
	appG := app.Gin{C: c}
	body := new(RemarkRequestBody)
	//解析Body参数
	if err := c.ShouldBind(body); err != nil {
		appG.Response(http.StatusBadRequest, "失败", fmt.Sprintf("参数出错%s", err.Error()))
		return
	}

	if body.CommdityName == "" || body.CustomerID == "" || body.StoreName == "" || body.Remark == "" {
		appG.Response(http.StatusBadRequest, "失败", "参数不能为空")
		return
	}
	var bodyBytes [][]byte
	bodyBytes = append(bodyBytes, []byte(body.CustomerID))
	bodyBytes = append(bodyBytes, []byte(body.CommdityName))
	bodyBytes = append(bodyBytes, []byte(body.StoreName))
	bodyBytes = append(bodyBytes, []byte(body.Remark))
	//调用智能合约
	resp, err := bc.ChannelQuery("createRemark", bodyBytes)
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


//列出所有商品
func QueryCommdityList(c *gin.Context) {
	appG := app.Gin{C: c}

	var bodyBytes [][]byte
	//调用智能合约
	resp, err := bc.ChannelQuery("queryCommdityList", bodyBytes)
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
