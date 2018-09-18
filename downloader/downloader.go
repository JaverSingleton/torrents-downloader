package downloader

import (
    "fmt"
    "os/exec"
    "strings"
)

func Download(link string, target string) error {
    linkStr := " -a "  + link
    targetStr := " -w " + "/" + target
    execute("transmission-daemon")
    execute("transmission-remote" + linkStr + targetStr)

    return nil
}

func execute(cmd string) {
  fmt.Println("command is ",cmd)
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head, parts...).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", out)
}
