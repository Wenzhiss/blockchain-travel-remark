<template>
  <div class="container">
    <el-table
      :data="storeList"
      height="250"
      border
      style="width: 100%">
      <el-table-column
        prop="name"
        label="店铺名"
        width="120">
      </el-table-column>
      <el-table-column
        prop="storeID"
        label="店铺ID"
        width="120">
      </el-table-column>
      <el-table-column
        prop="certificate"
        label="店铺资质"
        width="120">
      </el-table-column>
      <el-table-column
        prop="address"
        label="店铺地址">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import { queryStoreList } from '@/api/account';

export default {
  data() {
    return {
      storeList: [],
    }
  },
  created() {
    queryStoreList().then((response) => {
      if (response !== null) {
        this.storeList = response;
      }
    })
  },
 /*  methods: {
    listStore() {
      queryStoreList().then((response) => {
        if (response !== null) {
          this.storeList = response;
        }
      })
    }
  },*/

}
</script>


<style>
  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }
</style>
