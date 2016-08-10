package commandHelpers

import (
  "fmt"
  "os/exec"
  "os"
)

func KillBackgroundScan() {
  killCmd := "kill $(ps aux | grep '[c]lamscan' | awk '{print $2}')"
  cmd:= exec.Command("bash", "-c", killCmd)
  fmt.Print("Background scans ended.\n")
  cmd.Start()
}

func RunCmd(args []string) {
  cmdName := args[0]
  cmd := exec.Command(cmdName, args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}

func StartCmd(args[]string) {
  cmdName := args[0]
  cmd := exec.Command(cmdName, args...)
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  fmt.Print("Starting in scan in background. Scan result will output in shell.\n")
  cmd.Start()
}
