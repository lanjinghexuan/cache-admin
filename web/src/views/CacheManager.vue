<template>
  <el-card>
    <el-form :inline="true" @submit.prevent>
      <el-form-item label="前缀">
        <el-input v-model="prefix" placeholder="请输入缓存前缀" />
      </el-form-item>
      <el-form-item label="参数(JSON)">
        <el-input v-model="params" placeholder="可选，参数JSON字符串" />
      </el-form-item>
      <el-form-item>
        <el-button type="danger" :loading="loadingDel" @click="confirmDel">删除缓存</el-button>
        <el-button type="warning" :loading="loadingDelByPrefix" @click="confirmDelByPrefix">模糊删除</el-button>
      </el-form-item>
    </el-form>
    <el-alert v-if="msg" :title="msg" :type="msgType" show-icon style="margin-top: 20px" />
  </el-card>
</template>

<script setup>
import { ref } from 'vue'
import { cacheDel, cacheDelByPrefix } from '../api'
import { ElMessageBox } from 'element-plus'

const prefix = ref('')
const params = ref('')
const msg = ref('')
const msgType = ref('success')
const loadingDel = ref(false)
const loadingDelByPrefix = ref(false)

function confirmDel() {
  if (!prefix.value) {
    msg.value = '前缀不能为空'
    msgType.value = 'error'
    return
  }
  ElMessageBox.confirm('确定要删除该前缀和参数对应的缓存吗？', '提示', { type: 'warning' })
    .then(onDel)
    .catch(() => {})
}

function confirmDelByPrefix() {
  if (!prefix.value) {
    msg.value = '前缀不能为空'
    msgType.value = 'error'
    return
  }
  ElMessageBox.confirm('确定要模糊删除该前缀下所有缓存吗？', '警告', { type: 'warning' })
    .then(onDelByPrefix)
    .catch(() => {})
}

async function onDel() {
  let paramObj = {}
  try {
    paramObj = params.value ? JSON.parse(params.value) : {}
  } catch (e) {
    msg.value = '参数格式错误，应为JSON字符串'
    msgType.value = 'error'
    return
  }
  loadingDel.value = true
  const { data } = await cacheDel(prefix.value, paramObj)
  msg.value = data.msg || (data.code === 200 ? '删除成功' : '删除失败')
  msgType.value = data.code === 200 ? 'success' : 'error'
  loadingDel.value = false
}

async function onDelByPrefix() {
  loadingDelByPrefix.value = true
  const { data } = await cacheDelByPrefix(prefix.value)
  msg.value = data.msg || (data.code === 200 ? '删除成功' : '删除失败')
  msgType.value = data.code === 200 ? 'success' : 'error'
  loadingDelByPrefix.value = false
}
</script> 