// build +windows

package desktop

import (
	"fmt"
	"reflect"
	"strings"
	"syscall"
	"unicode/utf16"
	"unsafe"
)

/*
https://msdn.microsoft.com/en-us/library/ms724832(VS.85).aspx

6.1               Windows 7     / Windows 2008 R2
6.0               Windows Vista / Windows 2008
5.2               Windows 2003
5.1               Windows XP
5.0               Windows 2000
*/
func GetVersion() (int, int, int) {
	v, err := syscall.GetVersion()
	if err != nil {
		panic(err)
	}

	return int(uint8(v)), int(uint8(v >> 8)), int(uint16(v >> 16))
}

func IsWindowsXP() bool {
	v1, v2, _ := GetVersion()

	return v1 == 5 && v2 == 1
}

func String2WString(s string) uintptr {
	return uintptr(unsafe.Pointer(&utf16.Encode([]rune(s + "\x00"))[0]))
}

func WString2String(p uintptr) string {
	var rr []uint16 = make([]uint16, 0, MAX_PATH)
	for p := uintptr(unsafe.Pointer(p)); ; p += 2 {
		u := *(*uint16)(unsafe.Pointer(p))
		if u == 0 {
			return string(utf16.Decode(rr))
		}
		rr = append(rr, u)
	}
	panic("No zero at end of the string")
}

func WArray2String(rr []uint16) string {
	return string(utf16.Decode(rr))
}

func Arg(d interface{}) uintptr {
	v := reflect.ValueOf(d)
	UIntPtr := reflect.TypeOf((uintptr)(0))

	if v.Type().ConvertibleTo(UIntPtr) {
		vv := v.Convert(UIntPtr)
		return vv.Interface().(uintptr)
	} else {
		return v.Pointer()
	}
}

var NULL = Arg(0)

// copy last error from last syscall
var LastError uintptr

type HMENU uintptr

func HMENUPtr(r1, r2 uintptr, err error) HMENU {
	LastError = uintptr(err.(syscall.Errno))
	return HMENU(r1)
}

func (m HMENU) Close() {
	if !BOOLPtr(DestroyMenu.Call(Arg(m))).Bool() {
		panic(GetLastErrorString())
	}
}

type HRESULT uintptr

func HRESULTPtr(r1, r2 uintptr, err error) HRESULT {
	LastError = uintptr(err.(syscall.Errno))
	return HRESULT(r1)
}

func (m HRESULT) String() string {
	msg := [1024]uint16{}
	FormatMessage.Call(Arg(FORMAT_MESSAGE_FROM_SYSTEM), NULL, Arg(m), NULL, Arg(&msg[0]), Arg(len(msg)), NULL)
	return fmt.Sprintf("HRESULT: 0x%08x [%s]", uintptr(m), strings.TrimSpace(WString2String(Arg(&msg[0]))))
}

var Bool2Int = map[bool]int{
	true:  1,
	false: 0,
}

var Int2Bool = map[int]bool{
	1: true,
	0: false,
}

type BOOL uint32

func BOOLPtr(r1, r2 uintptr, err error) BOOL {
	LastError = uintptr(err.(syscall.Errno))
	return BOOL(r1)
}

func (m BOOL) Bool() bool {
	return Int2Bool[int(m)]
}

type DWORD uint32
type UINT uint32
type HWND uintptr

func HWNDPtr(r1, r2 uintptr, err error) HWND {
	LastError = uintptr(err.(syscall.Errno))
	return HWND(r1)
}

type TCHAR uint16

func GetLastErrorString() string {
	return HRESULT(LastError).String()
}

type WNDPROC uintptr

func WNDPROCNew(fn interface{}) WNDPROC {
	return WNDPROC(syscall.NewCallback(fn))
}

type HINSTANCE uintptr

func HINSTANCEPtr(r1, r2 uintptr, err error) HINSTANCE {
	LastError = uintptr(err.(syscall.Errno))
	return HINSTANCE(r1)
}

type BYTE byte
type WORD uint16
type LPCTSTR uintptr
type HCURSOR uintptr
type HBRUSH uintptr
type LONG uint32

type LRESULT uintptr

func LRESULTPtr(r1, r2 uintptr, err error) LRESULT {
	LastError = uintptr(err.(syscall.Errno))
	return LRESULT(r1)
}

type WPARAM uint32
type LPARAM uint32

type ATOM uintptr

func ATOMPtr(r1, r2 uintptr, err error) ATOM {
	LastError = uintptr(err.(syscall.Errno))
	return ATOM(r1)
}

type HDC uintptr

func HDCPtr(r1, r2 uintptr, err error) HDC {
	LastError = uintptr(err.(syscall.Errno))
	return HDC(r1)
}