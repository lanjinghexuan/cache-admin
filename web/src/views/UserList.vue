<template>
  <el-card>
    <el-form :inline="true" @submit.prevent style="margin-bottom: 10px">
      <el-form-item label="缓存前缀">
        <el-input v-model="prefix" placeholder="请输入缓存前缀" style="width: 240px" />
      </el-form-item>
      <el-form-item label="页码">
        <el-input-number v-model="page" :min="1" />
      </el-form-item>
      <el-form-item label="每页数量">
        <el-input-number v-model="limit" :min="1" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="fetchCacheList">搜索</el-button>
      </el-form-item>
    </el-form>
    <el-table :data="cacheList" v-loading="loading" style="width: 100%; margin-top: 10px" empty-text="暂无数据">
      <el-table-column type="index" label="#" width="50" />
      <el-table-column prop="key" label="缓存Key" />
      <el-table-column prop="ttl" label="TTL(秒)" />
      <el-table-column prop="size" label="内存占用(Byte)" />
    </el-table>
    <el-pagination
      v-if="cacheList.length > 0"
      v-model:current-page="page"
      v-model:page-size="limit"
      :page-sizes="[5, 10, 20, 50]"
      layout="total, sizes, prev, pager, next, jumper"
      :total="total"
      @size-change="fetchCacheList"
      @current-change="fetchCacheList"
      style="margin-top: 16px; text-align: right;"
    />
    <el-alert v-if="errorMsg" :title="errorMsg" type="error" show-icon style="margin-top: 10px" />
  </el-card>
</template>

<script setup>
import { ref } from 'vue'
import { getCacheList } from '../api'

const cacheList = ref([])
const page = ref(1)
const limit = ref(10)
const total = ref(0)
const loading = ref(false)
const errorMsg = ref('')
const prefix = ref('')

async function fetchCacheList() {
  if (!prefix.value) {
    errorMsg.value = '请先输入缓存前缀'
    return
  }
  loading.value = true
  errorMsg.value = ''
  try {
    const { data } = await getCacheList(prefix.value, page.value, limit.value)
    if (data.data) {
      cacheList.value = data.data
      total.value = data.total || 0
    } else {
      cacheList.value = []
      total.value = 0
      errorMsg.value = data.msg || '查询失败'
    }
  } catch (e) {
    cacheList.value = []
    total.value = 0
    errorMsg.value = '网络错误或服务器异常'
  }
  loading.value = false
}
</script> 