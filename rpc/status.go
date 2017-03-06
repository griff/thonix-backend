package rpc

import (
  "fmt"
)

func (s Status) Valid() error {
  switch s {
  case Status_booting:
    return nil
  case Status_bootFailed:
    return nil
  case Status_installing:
    return nil
  case Status_running:
    return nil
  }
  fmt.Printf("Booting %d\n", Status_booting)
  fmt.Printf("BootFailed %d\n", Status_bootFailed)
  fmt.Printf("Installing %d\n", Status_installing)
  fmt.Printf("Running %d\n", Status_running)
  return fmt.Errorf("Invalid server state %d", s)
}


func ValidStatusFromString(state string) (Status, error) {
  switch state {
  case "booting":
    return Status_booting, nil
  case "bootFailed":
    return Status_bootFailed, nil
  case "installing":
    return Status_installing, nil
  case "running":
    return Status_running, nil
  default:
    return 12, fmt.Errorf("Invalid server state '%s'", state)
  }
}
