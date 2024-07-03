package service

/*
#cgo CFLAGS: -I ../../lib
#cgo LDFLAGS: -L ../../lib -lMaaCore -Wl,-rpath,../../lib
#include <stdlib.h>
#include "AsstCaller.h"
*/
import "C"

func GetMaaVersion() string {

	var maaVersion = C.AsstGetVersion()
	var version = C.GoString(maaVersion)
	return version

}
