package pkg

import (
	"cache-admin/config"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"
)

type CacheData struct {
	Prefix string
	Params interface{}
	Exprie time.Duration
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
	res, err := config.RedisDB.Get(config.Ctx, prefixParams).Result()
	if err == nil && res != "" {
		err = json.Unmarshal([]byte(res), result)
		if err == nil {
			return nil
		}
	}
	resp, err := query()
	if err != nil {
		return err
	}
	resD, err := json.Marshal(resp)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = config.RedisDB.Set(config.Ctx, prefixParams, resD, opts.Exprie).Err()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	err = json.Unmarshal(resD, result)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

// 模糊删除缓存
func CacheDelByPrefix(prefix string) error {
	iter := config.RedisDB.Scan(config.Ctx, 0, prefix+"*", 0).Iterator()
	for iter.Next(config.Ctx) {
		config.RedisDB.Del(config.Ctx, iter.Val())
	}
	return iter.Err()
}
