<template>
  <div class="container">
    <el-form :model="storeForm" ref="storeForm" label-width="100px" class="demo-ruleForm">
      <el-form-item
        label="店铺名"
        prop="storeName"
        :rules="[
          { required: true, message: '店铺名不能为空'},
        ]"
      >
      <el-input type="storeName" v-model.number="storeForm.storeName" autocomplete="off"></el-input>

      </el-form-item>
            <el-form-item
        label="店铺地址"
        prop="address"
        :rules="[
          { required: true, message: '店铺地址不能为空'},
        ]"
      >
        <el-input type="address" v-model.number="storeForm.address" autocomplete="off"></el-input>
      </el-form-item>
      <el-form-item
        label="店铺资质"
        prop="certificate"
        :rules="[
          { required: true, message: '店铺资质不能为空'},
        ]"
      >
      <el-input type="certificate" v-model.number="storeForm.certificate" autocomplete="off"></el-input>
      <el-form-item>
      <el-form-item
        label="店铺余额"
        prop="balance"
        :rules="[
                { required: true, message: '店铺余额不能为空' },
                { type: 'number', message: '余额必须为数字'}
        ]"
      >
      <el-input type="balance" v-model.number="storeForm.balance" autocomplete="off"></el-input>
      <el-form-item>
        <el-button type="primary" @click="submitForm('storeForm')">提交</el-button>
        <el-button @click="resetForm('storeForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import { createStore } from '../../../api/account';
  export default {
    data() {
      return {
        storeForm: {
          storeName: '',
          address: '',
          certificate: '',
          balance:''
        },
        loading:false,
      };
    },
    methods: {
      submitForm(formName) {
        this.$refs[formName].validate((valid) => {
          if (valid) {
            createStore({
              storeName: this.storeName,
              address: this.address,
              certificate: this.certificate,
              balance:this.balance
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
