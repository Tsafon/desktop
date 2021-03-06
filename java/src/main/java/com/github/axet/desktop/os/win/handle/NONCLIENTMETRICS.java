package com.github.axet.desktop.os.win.handle;

import java.util.Arrays;
import java.util.List;

import com.sun.jna.Pointer;
import com.sun.jna.Structure;

// http://msdn.microsoft.com/en-us/library/windows/desktop/ff729175(v=vs.85).aspx
public class NONCLIENTMETRICS extends Structure {

    public static final int ODT_MENU = 1;
    public static final int ODT_LISTBOX = 2;
    public static final int ODT_COMBOBOX = 3;
    public static final int ODT_BUTTON = 4;
    public static final int ODT_STATIC = 5;

    public static class ByValue extends NONCLIENTMETRICS implements Structure.ByValue {
    }

    public static class ByReference extends NONCLIENTMETRICS implements Structure.ByReference {
    }

    public NONCLIENTMETRICS() {
        cbSize = size();
    }

    public NONCLIENTMETRICS(Pointer p) {
        super(p);

        read();
    }

    @Override
    protected List<?> getFieldOrder() {
        return Arrays.asList(new String[] { "cbSize", "iBorderWidth", "iScrollWidth", "iScrollHeight", "iCaptionWidth",
                "iCaptionHeight", "lfCaptionFont", "iSmCaptionWidth", "iSmCaptionHeight", "lfSmCaptionFont",
                "iMenuWidth", "iMenuHeight", "lfMenuFont", "lfStatusFont", "lfMessageFont", 
                //"iPaddedBorderWidth"
                });
    }

    public int cbSize;
    public int iBorderWidth;
    public int iScrollWidth;
    public int iScrollHeight;
    public int iCaptionWidth;
    public int iCaptionHeight;
    public LOGFONT.ByValue lfCaptionFont;
    public int iSmCaptionWidth;
    public int iSmCaptionHeight;
    public LOGFONT.ByValue lfSmCaptionFont;
    public int iMenuWidth;
    public int iMenuHeight;
    public LOGFONT.ByValue lfMenuFont;
    public LOGFONT.ByValue lfStatusFont;
    public LOGFONT.ByValue lfMessageFont;
    //public int iPaddedBorderWidth;
}