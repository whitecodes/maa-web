package service

/*
#cgo CFLAGS: -I ../../lib
#cgo LDFLAGS: -L ../../lib -lMaaCore -Wl,-rpath,../../lib
#include <stdlib.h>
#include "AsstCaller.h"
*/
import "C"
import (
	"sync"
)

// AsstHandle 包装 C 的 AsstHandle 指针
type AsstHandle struct {
	handle C.AsstHandle
}

var (
	instance *AsstHandle
	once     sync.Once
)

func GetMaaVersion() string {

	var maaVersion = C.AsstGetVersion()
	var version = C.GoString(maaVersion)
	return version
}

func GetMaaHandle() *AsstHandle {
	once.Do(func() {
		instance = &AsstHandle{handle: C.AsstCreate()}
	})
	return instance
}

func GetMaaConnected(maa *AsstHandle) bool {
	connected := C.AsstConnected(maa.handle)
	return connected != 0
}

func ConnectDevice(maa *AsstHandle, adbPath string, address string, conf string) error {
	asstAsyncCallId := C.AsstAsyncConnect(maa.handle, C.CString(adbPath), C.CString(address), C.CString(conf), C.uint8_t(0))

	if asstAsyncCallId != 0 {
		return nil

	}
	return nil
}
