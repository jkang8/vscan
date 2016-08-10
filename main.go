package main

import (
  "fmt"
  "os"
  "./src/commandHelpers"
)

func backgroundPrompt(scanArgs []string) {
    var option string
    var args = scanArgs
    fmt.Print("Run in background? y/n\n")
    fmt.Scanf("%s", &option)
    if option == "y" {
      scanArgs = append(args, []string{"&"}...)
      commandHelpers.StartCmd(args)
    } else {
      commandHelpers.RunCmd(args)
    }
}

func processOptions(option int) {
    if option == 1 {
      // Full system scan
      scanArgs := []string{"clamscan", "-r", "-i", "--bell", "/"}
      backgroundPrompt(scanArgs)
    } else if option == 2 {
      // Update virus definitions
      scanArgs := []string{"freshclam"}
      commandHelpers.RunCmd(scanArgs)
    } else if option == 3 {
      commandHelpers.KillBackgroundScan()
    } else if option == 4 {
      // Test on virus sample
      scanArgs := []string{"clamscan", "--bell", "-i", "malware_sample"}
      backgroundPrompt(scanArgs)
    }
}

func main() {
  if len(os.Args) < 2 {
    var option int
    fmt.Print("1. Full system scan\n2. Update virus definitions\n3. Stop background scans\n4. Test malware sample\n")
    fmt.Scanf("%d", &option)
    processOptions(option)
  } else if len(os.Args) == 2 {
    // Specified path to scan
    args := os.Args
    if args[1] == "--stop" {
        commandHelpers.KillBackgroundScan()
    } else {
      scanArgs := []string{"clamscan", "-r", "--bell", "-i", "&", args[1]}
      commandHelpers.StartCmd(scanArgs)
    }
  } else {
    fmt.Print("virus_scanner [path]")
  }
}
