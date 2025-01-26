package clipboard

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
)

func File_content_to_clipboard() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	b, os_err := os.ReadFile(filepath.Join(homeDir, ".tmp/tmp.txt"))
	if os_err != nil {
		panic(os_err)
	}
	str := string(b)
	fmt.Println(str)

	clip_err := clipboard.WriteAll(str)
	if clip_err != nil {
		panic(clip_err)
	}
}
