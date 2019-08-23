gopher-redis
===========

a redis library for gopher-lua


usage
======

    import (
        lua "github.com/yuin/gopher-lua"
        redis "github.com/fangdingjun/gopher-redis"
    )

    func main(){
        L := lua.NewState()
        defer L.Close()

        L.PreloadModule("redis", redis.Loader)

        err := L.DoString(`
            local redis = require "redis"

            local conn = redis.new({host="127.0.0.1", port="6379", password="", index=0})

            print(conn:Set("a", "b", 0):Val())
            print(conn:Get("a"):Val())

            local keys = conn:Do("keys", "*"):Val()
            for k, v in keys() do
                print(k, v)
            end

            print(conn:Do("expire", "a", "100"):Val())
            print(conn:Do("ttl", "a"):Val())

            local b = conn:Do("incr", "b"):Val()
            print(b)

            conn:Close()
        `)
        if err != nil {
            log.Fatal(err)
        }
    }

API
===

this is a wrapper for go-redis/redis

refer to https://github.com/go-redis/redis
