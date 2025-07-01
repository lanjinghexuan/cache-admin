import axios from 'axios'

export function getCacheList(prefix, page, limit) {
  return axios.get(`/cache/getCacheKeyList/${encodeURIComponent(prefix)}`, { params: { page, limit } })
}

export function cacheDel(prefix, params) {
  return axios.get('/cache/del', { params: { prefix, params } })
}

export function cacheDelByPrefix(prefix) {
  return axios.get('/cache/delByPrefix', { params: { prefix } })
} 