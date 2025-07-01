<template>
  <el-container style="height: 100vh">
    <el-header>
      <el-menu :default-active="activeMenu" mode="horizontal" @select="onSelect">
        <el-menu-item index="/users">用户列表</el-menu-item>
        <el-menu-item index="/cache">缓存管理</el-menu-item>
      </el-menu>
    </el-header>
    <el-main>
      <router-view />
    </el-main>
  </el-container>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
const router = useRouter()
const route = useRoute()
const activeMenu = ref(route.path)
watch(() => route.path, (val) => activeMenu.value = val)
function onSelect(index) {
  router.push(index)
}
</script>

<style>
body, html, #app {
  height: 100%;
  margin: 0;
}
</style> 