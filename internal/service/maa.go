package service

/*
#cgo CFLAGS: -I ../../lib
#cgo LDFLAGS: -L ../../lib -lMaaCore -Wl,-rpath,../../lib
#include <stdlib.h>
#include "AsstCaller.h"
*/
import "C"
import (
	"errors"
	"log"
	"strconv"
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

func DestroyMaaHandle(maa *AsstHandle) {
	C.AsstDestroy(maa.handle)
}

func LoadMaaResource(resoucePath string) error {
	loaded := C.AsstLoadResource(C.CString(resoucePath))
	if loaded == 0 {
		return errors.New("load Resource failed")
	}
	return nil
}

func GetMaaConnected(maa *AsstHandle) bool {
	connected := C.AsstConnected(maa.handle)
	// C 中的布尔值是 int 类型，0 为 false，非 0 为 true
	return connected != 0
}

func ConnectDevice(maa *AsstHandle, adbPath string, address string, conf string) error {
	asstAsyncCallId := C.AsstAsyncConnect(maa.handle, C.CString(adbPath), C.CString(address), C.CString(conf), C.uint8_t(1))

	if asstAsyncCallId == 0 {
		log.Println("connect device " + address + " fail: " + strconv.Itoa(int(asstAsyncCallId)))
		return errors.New("connect device failed")
	}
	return nil
}

func BackToHome(maa *AsstHandle) bool {
	back := C.AsstBackToHome(maa.handle)
	return back != 0
}
