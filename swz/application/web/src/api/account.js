import request from '@/utils/request'

//id 查询顾客列表
export function queryCustomerList(data) {
	return request({
		url: '/queryCustomerList',
		method: 'post',
		data
	})
}

//QueryStoreByNameAndID 搜索店铺信息
export function queryStoreByNameAndID(data) {
	return request({
		url: '/queryStoreByNameAndID',
		method: 'post',
		data
	})
}

//空args 列举所有店铺
export function queryStoreList(data) {
	return request({
		url: '/queryStoreList',
		method: 'post',
		data
	})
}

//CreateStore; 店铺注册
export function createStore(data) {
	return request({
		url: '/createStore',
		method: 'post',
		data
	})
}

//游客注册
export function createCustomer(data) {
	return request({
		url: '/createCustomer',
		method: 'post',
		data
	})
}

// 登录
export function login(data) {
  return request({
    url: '/queryCustomerList',
    method: 'post',
    data
  })
}
