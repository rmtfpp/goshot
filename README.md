# GoShot

**GoShot** is a golang utility that transcribes the content of an area of the screen. You can select text from any source and GoShot will append its transcription to your system clipboard. GoShot relies on [Tesseract](https://github.com/tesseract-ocr/tesseract) for text transcription and on [LaTeX-OCR](https://github.com/lukas-blecher/LaTeX-OCR). 


# Usage 
GoShot comes in two binaries (user must decide which one to use)
```
run-text
```
```
run-latex
```

# Installation
Installation is currently only possible on Linux systems. Dependencies include gnome-screenshot and tesseract-ocr. Upon installation, a virtual enviroment is created in ~/.venvs/GS in which the python dependencies (pix2tex, torch, torchvision) will be installed. 
```
git clone https://github.com/rmtfpp/goshot
cd goshot
make install
```
The executables will be available in ~/.local/bin and this path will be added to ~/.bashrc. The virtual enviroment will be activated every time ``run-latex`` will be called.

# Disclaimer
This project was done in a hurry for my specific needs, but I thought it might be useful to other people as well. Many parameters are hard-coded and I would recommend checking Makefile before building and installing.  There is torch and torchvision for CPUs in the dependencies, but I think it might be useful to take advantage of CUDA on Nvidia platforms.