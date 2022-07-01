package routers

import (
	v1 "application/server/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.Default()

	/*chaincode{
		QueryCustomerList
		QueryStoreList
		QueryTxList
		CreateRemark
		CreateCustomer
		CreateSold
		CreateBuy
		CreateStore
	}
	*/
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/hello", v1.Hello)

		apiV1.POST("/queryCustomerList", v1.QueryCustomerList)         //查询顾客1
		apiV1.POST("/queryStoreList", v1.QueryStoreList)               //查询全部商家1
		apiV1.POST("/queryStoreByNameAndID", v1.QueryStoreByNameAndID) //按属性查询商家1
		apiV1.POST("/queryTxList", v1.QueryTxList)                     //查询交易列表1
		apiV1.POST("/queryCommdityList", v1.QueryCommdityList)         //查询所有商品

		apiV1.POST("/createRemark", v1.CreateRemark)     //添加评论1
		apiV1.POST("/createCustomer", v1.CreateCustomer) //顾客注册1
		apiV1.POST("/createSold", v1.CreateSold)         //店铺上架商品1
		apiV1.POST("/createBuy", v1.CreateBuy)           //构建交易1
		apiV1.POST("/createStore", v1.CreateStore)       //商家注册1

	}
	// 静态文件路由
	r.StaticFS("/web", http.Dir("./dist/"))
	return r
}
