package com.github.axet.desktop.os.mac.foundation;

import com.sun.jna.Library;
import com.sun.jna.Native;
import com.sun.jna.Pointer;
import com.sun.jna.win32.StdCallLibrary.StdCallCallback;

// thanks to https://gist.github.com/3974488

// https://developer.apple.com/library/mac/#documentation/Cocoa/Reference/ObjCRuntimeRef/Reference/reference.html

public interface Runtime extends Library {

    Runtime INSTANCE = (Runtime) Native.loadLibrary("objc", Runtime.class);

    Pointer objc_lookUpClass(String name);

    String class_getName(Pointer cls);

    Pointer sel_getUid(String str);

    String sel_getName(Pointer aSelector);

    long objc_msgSend(Pointer theReceiver, Pointer theSelector, Object... string);

    Pointer class_createInstance(Pointer cls, int extraBytes);

    boolean class_addMethod(Pointer cls, Pointer selector, StdCallCallback imp, String types);

    Pointer sel_registerName(String name);

    Pointer objc_allocateClassPair(Pointer superClass, String name, long extraBytes);

    void objc_registerClassPair(Pointer cls);

    Pointer class_getInstanceMethod(Pointer cls, Pointer selecter);

    Pointer method_setImplementation(Pointer method, StdCallCallback imp);

    Pointer objc_getProtocol(String protocol);
    
    boolean class_addProtocol(Pointer cls, Pointer protocol);
}
