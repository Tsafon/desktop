package com.github.axet.apple;

import junit.framework.Test;
import junit.framework.TestCase;
import junit.framework.TestSuite;
import fundations.NSFileNanager;

public class NSFileNanagerTest extends TestCase {
    public NSFileNanagerTest(String testName) {
        super(testName);
    }

    public static Test suite() {
        return new TestSuite(NSFileNanagerTest.class);
    }

    public void testApp() {
        CFArrayRef a = NSFileNanager.INSTANCE.NSSearchPathForDirectoriesInDomains(
                NSFileNanager.NSSearchPathDirectory.NSAllApplicationsDirectory,
                NSFileNanager.NSSearchPathDomainMask.NSSystemDomainMask, true);

        for (int i = 0; i < a.getCount(); i++) {
            CFStringRef s = new CFStringRef(a.get(0));
            String ss = s.toString();
            assertEquals(ss, "/Applications");
        }

    }
}
