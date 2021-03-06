// +build darwin

package desktop

import "C"

import (
	"bufio"
	"bytes"
	"image"
	"image/png"
	"unsafe"
)

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/ApplicationKit/Classes/NSImage_Class

var NSImageClass unsafe.Pointer = Runtime_objc_lookUpClass("NSImage")
var NSImageInitWithData unsafe.Pointer = Runtime_sel_getUid("initWithData:")

type NSImage struct {
	NSObject
}

func Image2Bytes(p image.Image) []byte {
	var b bytes.Buffer
	w := bufio.NewWriter(&b)
	err := png.Encode(w, p)
	// it we have correct image on memory we have to write it down.
	// error would mean out of memroy or similar error. so panic.
	if err != nil {
		panic(err)
	}
	w.Flush()
	return b.Bytes()
}

func NSImageNew() NSImage {
	return NSImage{NSObjectPointer(Runtime_class_createInstance(NSImageClass, 0))}
}

func NSImageData(p NSData) NSImage {
	var m NSImage = NSImageNew()
	Runtime_objc_msgSend(m.Pointer, NSImageInitWithData, p.Pointer)
	return m
}

func NSImageImage(p image.Image) NSImage {
	var d NSData = NSDataNew(Image2Bytes(p))
	defer d.Release()
	return NSImageData(d)
}

func NSImagePointer(p unsafe.Pointer) NSImage {
	return NSImage{NSObjectPointer(p)}
}
