<template>
  <div class="container">
    <el-form :model="commodityForm" ref="commodityForm" label-width="100px" class="demo-ruleForm">
      <el-form-item
        label="商品名"
        prop="commodityName"
        :rules="[
          { required: true, message: '商品名不能为空'},
        ]"
      >
      <el-input type="commodityName" v-model.number="commodityForm.commodityName" autocomplete="off"></el-input>

      </el-form-item>
            <el-form-item
        label="价格"
        prop="price"
        :rules="[
          { required: true, message: '价格不能为空'},
          { type: 'number', message: '价格必须为数字'}
        ]"
      >
        <el-input type="price" v-model.number="commodityForm.price" autocomplete="off"></el-input>
      </el-form-item>
            <el-form-item
        label="供应商"
        prop="provider"
        :rules="[
          { required: true, message: '商品名不能为空'},
        ]"
      >
      <el-input type="provider" v-model.number="commodityForm.provider" autocomplete="off"></el-input>
      <el-form-item>
        <el-button type="primary" @click="submitForm('commodityForm')">提交</el-button>
        <el-button @click="resetForm('commodityForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { createSold } from '../../../api/selling';

  export default {
    data() {
      return {
        commodityForm: {
          commodityName: '',
          price: '',
          provider:'',
        },
        loading:false,
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            createSold({
              name:this.commodityForm.commodityName,
              price:this.commodityForm.price,
              procider:this.commodityForm.provider,
            })
								.then((response) => {
									this.loading = false;
									if (response !== null) {
										this.$message({
											type: 'success',
											message: '商品上架成功!',
										});
									} else {
										this.$message({
											type: 'error',
											message: '商品上架失败!',
										});
									}
								})
								.catch((_) => {
									this.loading = false;
								});
          } else {
            return false;
          }
        });
      },
      resetForm(formName) {
        this.$refs[formName].resetFields();
      }
    }
  }
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
