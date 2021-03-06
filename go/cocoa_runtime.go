// +build darwin

package desktop

/*
#cgo LDFLAGS: -lobjc -framework AppKit

#include <stdlib.h>
#include <objc/objc-runtime.h>

BOOL NSApplicationLoad (void);
int NSApplicationMain ( int argc, const char *argv[] );

id objc_msgSend0(id to, SEL sel) {
  return objc_msgSend(to, sel);
}

id objc_msgSend1(id to, SEL sel, void* arg1) {
  return objc_msgSend(to, sel, arg1);
}

id objc_msgSend2(id to, SEL sel, void* arg1, void* arg2) {
  return objc_msgSend(to, sel, arg1, arg2);
}

id objc_msgSend3(id to, SEL sel, void* arg1, void* arg2, void* arg3) {
  return objc_msgSend(to, sel, arg1, arg2, arg3);
}

id objc_msgSend4(id to, SEL sel, void* arg1, void* arg2, void* arg3, void* arg4) {
  return objc_msgSend(to, sel, arg1, arg2, arg3, arg4);
}
*/
import "C"

import (
	"fmt"
	"math"
	"unsafe"
)

func init() {
	b := C.NSApplicationLoad()
	if b == 0 {
		panic("!NSApplicationLoad")
	}
}

var Bool2Int = map[bool]int{
	true:  1,
	false: 0,
}

func Bool2Pointer(b bool) unsafe.Pointer {
	return unsafe.Pointer(uintptr(Bool2Int[b]))
}

var Int2Bool = map[int]bool{
	1: true,
	0: false,
}

func Pointer2Bool(p unsafe.Pointer) bool {
	return Int2Bool[int(uintptr(p))]
}

func Int2Pointer(i int) unsafe.Pointer {
	return unsafe.Pointer(uintptr(i))
}

func Float2Pointer(i float64) unsafe.Pointer {
	return unsafe.Pointer(uintptr(math.Float64bits(i)))
}

func Pointer2Float(p unsafe.Pointer) float64 {
	return math.Float64frombits(uint64(uintptr(p)))
}

func Pointer2Int(p unsafe.Pointer) int {
	return int(uintptr(p))
}

func Pointer2String(p unsafe.Pointer) string {
	return C.GoString((*C.char)(p))
}

func UInt2Pointer(i uint) unsafe.Pointer {
	return unsafe.Pointer(uintptr(i))
}

func Runtime_Loop() {
	app := NSApplicationMainSharedApplication()
	defer app.Release()

	date := NSDateDistantFuture()
	defer date.Release()

	for {
		p := app.NextEventMatchingMaskUntilDateInModeDequeue(NSAnyEventMask, date, NSDefaultRunLoopMode, true)
		fmt.Println(p, p.Type())
		app.SendEvent(p)
		app.UpdateWindows()
		p.Release()
	}
}

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/ObjCRuntimeRef/Reference/reference.html

func Runtime_objc_lookUpClass(s string) unsafe.Pointer {
	c := unsafe.Pointer(C.CString(s))
	defer C.free(c)
	return unsafe.Pointer(C.objc_lookUpClass((*C.char)(c)))
}

func Runtime_class_getName(m unsafe.Pointer) string {
	return C.GoString(C.class_getName((*C.struct_objc_class)(m)))
}

func Runtime_sel_getUid(s string) unsafe.Pointer {
	var c *C.char = C.CString(s)
	defer C.free(unsafe.Pointer(c))
	return unsafe.Pointer(C.sel_getUid(c))
}

func Runtime_sel_getName(sel unsafe.Pointer) string {
	return C.GoString(C.sel_getName((*C.struct_objc_selector)(sel)))
}

func Runtime_objc_msgSend(self unsafe.Pointer, sel unsafe.Pointer, args ...unsafe.Pointer) unsafe.Pointer {
	switch len(args) {
	case 0:
		return unsafe.Pointer(C.objc_msgSend0((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel)))
	case 1:
		return unsafe.Pointer(C.objc_msgSend1((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel), args[0]))
	case 2:
		return unsafe.Pointer(C.objc_msgSend2((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel), args[0], args[1]))
	case 3:
		return unsafe.Pointer(C.objc_msgSend3((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel), args[0], args[1], args[2]))
	case 4:
		return unsafe.Pointer(C.objc_msgSend4((*C.struct_objc_object)(self), (*C.struct_objc_selector)(sel), args[0], args[1], args[2], args[3]))
	default:
		panic(fmt.Sprint("Unsupported number of arguments ", len(args)))
	}
}

func Runtime_class_createInstance(cls unsafe.Pointer, extraBytes int) unsafe.Pointer {
	return unsafe.Pointer(C.class_createInstance((*C.struct_objc_class)(cls), (C.size_t)(extraBytes)))
}

func Runtime_class_addMethod(cls unsafe.Pointer, sel unsafe.Pointer, imp unsafe.Pointer, types string) bool {
	var n *C.char = C.CString(types)
	defer C.free(unsafe.Pointer(n))
	if C.class_addMethod((*C.struct_objc_class)(cls), (*C.struct_objc_selector)(sel), (*[0]byte)(imp), n) == 1 {
		return true
	} else {
		return false
	}
}

func Runtime_sel_registerName(name string) unsafe.Pointer {
	var n *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return unsafe.Pointer(C.sel_registerName(n))
}

func Runtime_objc_allocateClassPair(superClass unsafe.Pointer, name string, extraBytes int) unsafe.Pointer {
	var n *C.char = C.CString(name)
	defer C.free(unsafe.Pointer(n))
	return unsafe.Pointer(C.objc_allocateClassPair((*C.struct_objc_class)(superClass), n, (C.size_t)(extraBytes)))
}

func Runtime_objc_registerClassPair(cls unsafe.Pointer) {
	C.objc_registerClassPair((*C.struct_objc_class)(cls))
}

//Pointer class_getInstanceMethod(Pointer cls, Pointer selecter);

//Pointer method_setImplementation(Pointer method, StdCallCallback imp);

//Pointer objc_getProtocol(String protocol);

//boolean class_addProtocol(Pointer cls, Pointer protocol);
