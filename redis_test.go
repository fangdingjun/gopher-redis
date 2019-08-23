package redis

import (
	"testing"

	lua "github.com/yuin/gopher-lua"
)

func TestRedis(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("redis", Loader)
	err := L.DoString(`
	 local redis = require "redis"
	 local conn = redis.new({host="127.0.0.1",password="",index=0})
	 print(conn:Set("a","b",0):Val())
	 print(conn:Get("a"):Val())
	 --[[
	 local keys = conn:Do("keys","wallet*"):Val()
	 for k,v in keys() do
		 print(k,v)
	 end
	 ]]
	 print(conn:Do("expire", "b", "100"):Val())
	 print(conn:Do("ttl", "a"):Val())
	 local b = conn:Do("incr", "b"):Val()
	 print(b)
	 conn:Close()
	`)
	if err != nil {
		t.Error(err)
	}
}
