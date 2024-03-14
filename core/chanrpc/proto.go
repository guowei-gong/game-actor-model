package chanrpc

import (
	"bytes"
	"reflect"
)

// Handler 收到请求消息时, 调用的处理函数
type Handler func(reqCtx *ReqCtx)

// ReqCtx 调用参数
type ReqCtx struct {
	id uint32 // 注册的路由 id
}

// AckCtx 结果信息
type AckCtx struct {
}

func MsgID(m any) uint32 {
	typ := reflect.TypeOf(m)
	if typ.Kind() == reflect.Struct {
		return BKDRHash(typ.Name())
	}
	return 0
}

// BKDRHash Hash 一个字符串
func BKDRHash(s string) uint32 {
	b := bytes.NewBufferString(s).Bytes()
	// Hash 字节序列
	seed := uint32(131)
	hash := uint32(0)

	for _, v := range b {
		hash = hash*seed + uint32(v)
	}
	return hash
}
