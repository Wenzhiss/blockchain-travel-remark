import request from '@/utils/request'


export function createBuy(data) {
	return request({
		url: '/createBuy',
		method: 'post',
		data
	})
}

export function createSold(data) {
	return request({
		url: '/createSold',
		method: 'post',
		data
	})
}

export function queryTxList(data) {
	return request({
		url: '/queryTxList',
		method: 'post',
		data
	})
}

export function createRemark(data) {
	return request({
		url: '/createRemark',
		method: 'post',
		data
	})
}


export function queryCommdityList(data) {
	return request({
		url: '/queryCommdityList',
		method: 'post',
		data
	})
}

