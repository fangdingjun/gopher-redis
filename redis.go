package redis

import (
	"fmt"

	"github.com/fangdingjun/go-log"
	"github.com/go-redis/redis"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

// Loader Loader is the module loader
func Loader(L *lua.LState) int {
	t := L.NewTable()
	// L.SetFuncs(t, api)
	L.SetField(t, "new", luar.New(L, newRedis))
	L.Push(t)
	return 1
}

type luaRedis struct {
	*redis.Client
}

func newRedis(L *luar.LState) int {
	opt := L.CheckTable(1)
	host := string(L.RawGet(opt, lua.LString("host")).(lua.LString))

	port := "6379"
	port1 := L.RawGet(opt, lua.LString("port"))
	if port1.Type() == lua.LTString {
		port = string(port1.(lua.LString))
	}

	passwd := ""
	passwd1 := L.RawGet(opt, lua.LString("password"))
	if passwd1.Type() == lua.LTString {
		passwd = string(passwd1.(lua.LString))
	}

	db := 0
	db1 := L.RawGet(opt, lua.LString("index"))
	if db1.Type() == lua.LTNumber {
		db = int(db1.(lua.LNumber))
	}

	log.Debugln(host, port, passwd, db)

	r := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: passwd,
		DB:       db,
	})

	L.Push(luar.New(L.LState, r))

	return 1
}
