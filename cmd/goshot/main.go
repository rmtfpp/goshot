package main

import (
	"fmt"
	"os"
	"path"

	"github.com/rmtfpp/goshot/internal/cleanup"
	"github.com/rmtfpp/goshot/internal/clipboard"
	"github.com/rmtfpp/goshot/internal/screenshot"
	"github.com/rmtfpp/goshot/internal/transcribe"
)

func main() {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return
	}

	os.Mkdir(path.Join(homeDir, ".tmp"), os.ModePerm)
	screenshot.Screenshot()
	if os.Args[1] == "-t" {
		transcribe.Transcribe_text_image("ita")
	} else if os.Args[1] == "-l" {
		transcribe.Transcribe_math_image()
	}
	clipboard.File_content_to_clipboard()
	cleanup.Delete_tmp_files()
}
