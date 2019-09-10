package redis

import (
	"testing"

	log "github.com/fangdingjun/go-log/v5"
	lua "github.com/yuin/gopher-lua"
)

func TestRedis(t *testing.T) {
	log.Default.Level = log.DEBUG
	L := lua.NewState()
	defer L.Close()

	L.PreloadModule("redis", Loader)
	err := L.DoString(`
	 local redis = require "redis"
	 local conn = redis.new({host="127.0.0.1", port=6379, password="", index=0})
	 --[[
	 print(conn:docmd("set", "a", 1))
	 print(conn:docmd("get", "a"))
	 local res, err = conn:docmd("keys", "wallet*")
	 if err ~= nil then
		 error(err)
	 end
	 for k, v in ipairs(res) do
		 print(k, v)
	 end
	 -- ]]
	 local r, err = conn:docmd("hmset","b",{e="e",f="f",g=1})
	 if err ~= nil then
		 error(err)
	 end
	 print(r)
	 -- ]]

	  local function arr2hash(t)
	  	  local t1={}
		  for i=1,#t,2 do
		  	t1[t[i]] = t[i+1]
		  end
		  return t1
	  end
	-- --[[

	 res, err = conn:docmd("hgetall", "b")
	 if err ~= nil then
		 error(err)
	 end

	 for k, v in pairs(arr2hash(res)) do
		 print(k, v)
	 end
	-- ]]
	print(conn:docmd("hexists", "b", "a"))

	 -- print(conn:docmd("eval", "return KEYS[1], KEYS[2]", 2, "aa", "bb"))
	 -- conn:docmd("get", nil)

	 -- print(conn:docmd("incr","cc"))
	 --[[

	 print(conn:docmd("lpush","dd","b"))
	 res, err = conn:docmd("lrange","dd",0,-1)
	 if err ~= nil then
		error(err)
	 end
	 for k, v in pairs(res) do
		 print(k, v)
	 end
	 -- ]]
	 conn:close()
	`)
	if err != nil {
		t.Error(err)
	}
}
