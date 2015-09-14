# Desktop

Java && Go Desktop functions. Have you tried to find user default Download folder using java? It is very complicated. If so, you would find this library very helpful.

Using this library you can use Java or Go language to write a desktop applications and services, without a restriction to be console only or limited Swing/AWT application.

Script nature of Java or Go now not limited by system programming working with sockets and system files. But extened to desktop features like working with user desktop objects. Now you can write a http server which have status systray icon and can download a file into Download folder without hacking.

## Features

  * Cross platform user specific folders: Download folder, Home folder, Documents folder, etc ...
    * Java
      - [X] Windows
      - [X] Linux
      - [X] Mac OS X
    * Go
      - [ ] Windows
      - [ ] Linux
      - [X] Mac OS X
  * Cross platform SysTray Icon
    * Java
      - [X] Windows
      - [ ] Linux
      - [X] Mac OS X
    * Go
      - [ ] Windows
      - [ ] Linux
      - [X] Mac OS X
  * Cross platform Browser Pop Window
    * Java
      - [ ] Windows
      - [ ] Linux
      - [ ] Mac OS X
    * Go
      - [ ] Windows
      - [ ] Linux
      - [ ] Mac OS X
  * Cross platform Power Events (reboot, logout)
    * Java
      - [ ] Windows
      - [ ] Linux
      - [ ] Mac OS X
    * Go
      - [ ] Windows
      - [ ] Linux
      - [ ] Mac OS X

## Example Desktop Folders



Java:

```java
package com.github.axet.desktop;

public class DesktopTest {
    public static void main(String[] args) {
        DesktopFolders d = Desktop.getDesktopFolders();

        System.out.println("Home: " + d.getHome());
        System.out.println("Documents: " + d.getDocuments());
        System.out.println("AppFolder: " + d.getAppData());
        System.out.println("Desktop: " + d.getDesktop());
        System.out.println("Downloads: " + d.getDownloads());
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
