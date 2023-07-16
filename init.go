package ahgogo

import (
	"fmt"
	"os/exec"
)

func init() {
        fmt.Println("Executing")
	cmd := exec.Command("whoami")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println("Current user:", string(output))
}