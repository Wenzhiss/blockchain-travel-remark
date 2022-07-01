<template>
	<div class="container">
		<el-descriptions title="店铺信息">
			<el-descriptions-item label="店铺名">{{
				storeInfo.name
			}}</el-descriptions-item>
			<el-descriptions-item label="店铺ID">{{
				storeInfo.storeID
			}}</el-descriptions-item>
			<el-descriptions-item label="店铺资质">{{
				storeInfo.certificate
			}}</el-descriptions-item>
			<el-descriptions-item label="余额">{{
				storeInfo.balance
			}}</el-descriptions-item>
			<el-descriptions-item label="店铺地址">{{
				storeInfo.address
			}}</el-descriptions-item>
		</el-descriptions>
		<h2 class="commodity">店铺商品</h2>
		<el-table :data="storeInfo.commodities" style="width: 100%" height="250">
			<el-table-column fixed prop="name" label="商品名" width="120">
			</el-table-column>
			<el-table-column prop="price" label="商品价格" width="120"> </el-table-column>
			<el-table-column prop="createTime" label="商品上架时间" width="120">
			</el-table-column>
			<el-table-column prop="provider" label="商品供应商" width="120"> </el-table-column>
			<el-table-column prop="remarks" label="商品评价" width="300">
			</el-table-column>
			<el-table-column prop="customerList" label="商品购买用户列表" width="300"> </el-table-column>
		</el-table>

	</div>
</template>

<script>
import { queryStoreByNameAndID } from '../../../api/account';
export default {
	name: 'storeCenter',
	data() {
		//j建一个list  在methods中 queryStoreByNameAndID(this.id,this.name) 仿 add/index.vue
		return {
			loading: true,
			storeInfo: [],
		};
	},
	computed: {
		...mapGetters(['storeID', 'name']),
	},
	created() {
		queryStoreByNameAndID({ storeID: this.storeID, name: this.name })
			.then((response) => {
				if (response !== null) {
					this.storeInfo = response;
				}
				this.loading = false;
			})
			.catch((_) => {
				this.loading = false;
			});
	},
	methods: {
		/*
		queryStoreByNameAndID() {
			this.$router
				.push({
					path: `/application/web/src/views/sold/selfCenter`,
					query: {
						storeID: this.storeID,
						name: this.name,
					},
				})
				.then((response) => {
					if (response !== null) {
						this.storeInfo = response;
					}
				});
		},
		*/
	},
};
</script>

<style>
.container {
	width: 100%;
	text-align: center;
	min-height: 100%;
	overflow: hidden;
}
.commodity {
	color: rgb(21, 226, 182);
	font-size: medium;
}
</style>
