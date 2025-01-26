package screenshot

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

func Screenshot() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	app := "gnome-screenshot"
	args := []string{"-a", "-f", path.Join(homeDir, ".tmp/tmp.png")}

	cmd := exec.Command(app, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("error:", err)
		log.Println("details:", string(output))
		return
	}
	log.Println("screenshot made")
}
