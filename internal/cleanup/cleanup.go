package cleanup

import (
	"log"
	"os"
	"path/filepath"
)

func Delete_tmp_files() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error getting home directory:", err)
		return
	}

	tmpTxtFile := filepath.Join(homeDir, ".tmp/tmp.txt")
	tmpPngFile := filepath.Join(homeDir, ".tmp/tmp.png")

	err = os.Remove(tmpTxtFile)
	if err != nil {
		log.Println("Error removing tmp.txt:", err)
	}

	err = os.Remove(tmpPngFile)
	if err != nil {
		log.Println("Error removing tmp.png:", err)
	}
}
