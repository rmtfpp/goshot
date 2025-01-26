package transcribe

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func Transcribe_text_image(lang_code string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error getting home directory:", err)
		return
	}

	inputFile := filepath.Join(homeDir, ".tmp/tmp.png")
	outputFile := filepath.Join(homeDir, ".tmp/tmp")

	app := "tesseract"
	args := []string{inputFile, outputFile, "--oem", "1", "-l", lang_code}

	cmd := exec.Command(app, args...)
	output, err := cmd.CombinedOutput()

	if err != nil {
		log.Println("error:", err)
		log.Println("details:", string(output))
		return
	}
	log.Println("screenshot transcribed")
}

func Transcribe_math_image() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Error getting home directory:", err)
		return
	}

	inputFile := filepath.Join(homeDir, ".tmp/tmp.png")
	outputFile := filepath.Join(homeDir, ".tmp/tmp.txt")

	log.Println("Starting OCR process...")

	file, err := os.Create(outputFile)
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	app := "latexocr"
	args := []string{inputFile}

	cmd := exec.Command(app, args...)
	cmd.Stdout = file

	err = cmd.Run()
	if err != nil {
		log.Println("Error running command:", err)
	}
}
