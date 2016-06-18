# Desktop

Java && Go Desktop functions. Have you tried to find user default Download folder using java? It is very complicated. If so, you would find this library very helpful.

Using this library you can use Java or Go language to write a desktop applications and services, without a restriction to be console only or limited Swing/AWT application.

Script nature of Java or Go now not limited by system programming working with sockets and system files. But extened to desktop features like working with user desktop objects. Now you can write a http server which have status systray icon and can download a file into Download folder without hacking.

Go check [Java Docs](java/) or [Go Docs](go/)

## Features

  - [X] Cross platform user specific folders: Download folder, Home folder, Documents folder, etc ...
  - [X] Cross platform SysTray Icon
  - [X] Cross platform Open Browser URL
  - [ ] Cross platform Browser Pop SysTray Window
    - Java/Go Service Can emulate web server, web pages can be opened by browser or mobile device. If we emulate mobile browser client, we can show systray icon dropdown browser mini window. As a result developers need to create only two versions of web site - desktop && mobile and systray icon goin just reuse mobile version.
  - [ ] Cross platform Power Events (control reboot, logout, suspend)
  - [ ] Register URL handlers (open an application from a browser)
  - [ ] Sound and Volume control

## Example Desktop Folders



Java:

```java
package com.github.axet.desktop;

public class DesktopTest {
    public static void main(String[] args) {
        System.out.println("Home: " + Desktop.getHomeFolder());
        System.out.println("Documents: " + Desktop.getDocumentsFolder());
        System.out.println("AppFolder: " + Desktop.getAppDataFolder());
        System.out.println("Desktop: " + Desktop.getDesktopFolder());
        System.out.println("Downloads: " + Desktop.getDownloadsFolder());
    }
}
```

Go:

```go
package main

import (
  "github.com/axet/desktop/go"
)

func main() {
  fmt.Println("Home:", desktop.GetHomeFolder())
  fmt.Println("Documents:" desktop.GetDocumentsFolder())
  fmt.Println("AppFolder:" desktop.GetAppDataFolder())
  fmt.Println("Desktop:" desktop.GetDesktopFolder())
  fmt.Println("Downloads:" desktop.GetDownloadsFolder())
}
```
