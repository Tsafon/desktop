package com.github.axet.desktop.os.mac;

import java.io.File;
import java.lang.reflect.Method;

import com.github.axet.desktop.DesktopFolders;
import com.github.axet.desktop.os.mac.cocoa.NSArray;
import com.github.axet.desktop.os.mac.cocoa.NSFileManager;
import com.github.axet.desktop.os.mac.cocoa.NSURL;

public class OSXFolders implements DesktopFolders {

    @Override
    public File getAppData() {
        // From CarbonCore/Folders.h
        return path("asup");
    }

    @Override
    public File getHome() {
        return new File(System.getenv("HOME"));
    }

    @Override
    public File getDocuments() {
        // From CarbonCore/Folders.h
        return path("docs");
    }

    @Override
    public File getDesktop() {
        // From CarbonCore/Folders.h
        return path("desk");
    }

    @Override
    public File getDownloads() {
        NSFileManager f = new NSFileManager();
        NSArray a = f.URLsForDirectoryInDomains(NSFileManager.NSSearchPathDirectory.NSDownloadsDirectory,
                NSFileManager.NSSearchPathDomainMask.NSUserDomainMask);

        long count = a.count();
        if (count != 1)
            throw new RuntimeException("Download folder not found");

        NSURL path = new NSURL(a.objectAtIndex(0));

        return new File(path.path().toString());
    }

    private static Class<?> FileManagerClass;
    private static Method OSTypeToInt;
    private static Short kUserDomain;

    protected static Class<?> getFileManagerClass() {
        if (FileManagerClass == null) {
            try {
                FileManagerClass = Class.forName("com.apple.eio.FileManager");
                OSTypeToInt = FileManagerClass.getMethod("OSTypeToInt", String.class);
                kUserDomain = (Short) FileManagerClass.getField("kUserDomain").get(null);
            } catch (RuntimeException e) {
                throw e;
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        }
        return FileManagerClass;
    }

    File path(String p) {
        try {
            final Method findFolder = getFileManagerClass().getMethod("findFolder", Short.TYPE, Integer.TYPE);
            final String path = (String) findFolder.invoke(null, kUserDomain, OSTypeToInt.invoke(null, p));
            return new File(path);
        } catch (RuntimeException e) {
            throw e;
        } catch (Exception e) {
            throw new RuntimeException(e);
        }
    }
}
