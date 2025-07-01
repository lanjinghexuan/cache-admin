package pkg

import (
	"cache-admin/config"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var delCount = 500
var ttlExprie = 60 * time.Second

type CacheData struct {
	Prefix       string
	Params       interface{}
	Exprie       time.Duration
	ForceRefresh bool
}

func PamamsKey(prefix string, param interface{}) string {
	paramsStr, err := json.Marshal(param)
	if err != nil {
		fmt.Println(err)
	}
	paramsHash := md5.Sum(paramsStr)
	paramsHashs := hex.EncodeToString(paramsHash[:])
	return fmt.Sprintf("%s_%s", prefix, paramsHashs)
}

// 删除指定缓存
func CacheDel(opts CacheData) error {
	prefixParams := PamamsKey(opts.Prefix, opts.Params)
	err := config.RedisDB.Del(config.Ctx, prefixParams).Err()
	if err != nil {
		return err
	}
	return nil
}

// 获取添加缓存
func GetCache(result interface{}, opts CacheData, query func() (interface{}, error)) error {
	prefixParams := PamamsKey(opts.Prefix, opts.Params)
	//是否强制刷新缓存
	if !opts.ForceRefresh {
		res, err := config.RedisDB.Get(config.Ctx, prefixParams).Result()
		if err == nil && res != "" {
			err = json.Unmarshal([]byte(res), result)
			if err == nil {
				return nil
			}
		}
		//获取过期时间，如果快过期进行异步刷新增加用户体验
		ttl, err := config.RedisDB.TTL(config.Ctx, prefixParams).Result()
		if err == nil && ttl > 0 && ttl < ttlExprie {
			go func() {
				resp, err := query()
				if err != nil || resp == nil {
					Error("异步缓存失败： %v", err)
					return
				}
				//异步调用无需进行报错处理
				resD, _ := json.Marshal(resp)
				config.RedisDB.Set(config.Ctx, prefixParams, resD, opts.Exprie)
			}()
			return nil
		}

	}
	resp, err := query()
	if err != nil {
		return err
	}
	resD, err := json.Marshal(resp)
	if err != nil {
		Error("获取结果后json编码失败 %v", err)
		return nil
	}
	err = config.RedisDB.Set(config.Ctx, prefixParams, resD, opts.Exprie).Err()
	if err != nil {
		Error("设置缓存失败 %v", err)
		return nil
	}
	err = json.Unmarshal(resD, result)
	if err != nil {
		Error("获取结果后json解码失败 %v", err)
		return err
	}
	return nil
}

// 模糊删除缓存
func CacheDelByPrefix(prefix string) error {
	Info("开始模糊删除缓存,Key:" + prefix)
	iter := config.RedisDB.Scan(config.Ctx, 0, prefix+"*", 0).Iterator()
	var keys []string
	var lens int
	for iter.Next(config.Ctx) {
		keys = append(keys, iter.Val())
		if len(keys) >= delCount {
			lens += delCount
			delAllKey(keys)
			keys = []string{}
		}
	}
	if len(keys) > 0 {
		lens += len(keys)
		delAllKey(keys)
	}
	Info("模糊删除缓存,Key:" + prefix + "结束。删除数量为:" + strconv.Itoa(lens))
	wg.Wait()
	return iter.Err()
}

func delAllKey(b []string) {
	wg.Add(1)
	go func(b []string) {
		defer wg.Done()
		err := config.RedisDB.Del(config.Ctx, b...).Err()
		if err != nil {
			Error("删除缓存出错: %v", err)
		}
	}(b)
}
