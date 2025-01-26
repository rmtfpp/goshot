package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
)

func screenshot() {

	app := "gnome-screenshot"
	args := []string{"-a", "-f", "tmp.png"}

	cmd := exec.Command(app, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("error:", err)
		log.Println("details:", string(output))
		return
	}
	fmt.Println("screenshot made")
}

func transcribe_image(lang_code string) {

	app := "tesseract"
	args := []string{"tmp.png", "tmp", "--oem", "1", "-l", lang_code}

	cmd := exec.Command(app, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("error:", err)
		log.Println("details:", string(output))
		return
	}
	fmt.Println("screenshot transcribed")
}

func file_content_to_clipboard() {
	b, os_err := os.ReadFile("tmp.txt")
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

func delete_tmp_files() {
	os.Remove("tmp.txt")
	os.Remove("tmp.png")
}

func main() {
	screenshot()
	transcribe_image("ita")
	file_content_to_clipboard()
	delete_tmp_files()
}
