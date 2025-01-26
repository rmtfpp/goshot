.PHONY: all python-deps build install clean

# Define Python version and environment
PYTHON_VERSION=python3.12  
PYTHON=$(PYTHON_VERSION)
PIP=$(PYTHON_VERSION) -m pip  
VENV_PATH=/home/$(USER)/.venvs/GS
VENV_ACTIVATE=$(VENV_PATH)/bin/activate  

# Ensure the virtual environment exists
all: install

# Install Python dependencies and build the package
python-deps:
	which $(PYTHON_VERSION) || (echo "$(PYTHON_VERSION) not found! Install it first." && exit 1)
	if [ ! -f $(VENV_ACTIVATE) ]; then \
		echo "Virtual environment not found. Creating it..."; \
		$(PYTHON) -m venv $(VENV_PATH); \
	fi
	# Activate the virtual environment and install dependencies
	source $(VENV_ACTIVATE) && \
	$(PIP) install --upgrade pip && \
	$(PIP) install torch torchvision --index-url https://download.pytorch.org/whl/cpu && \
	$(PIP) install -r requirements.txt && \
	$(PIP) install . && \
	$(PYTHON) setup.py build

# Build the Go application
build: python-deps
	source $(VENV_ACTIVATE) && go build -o goshot ./cmd/goshot

# Install the app (copy to .local/bin)
install: build
	sudo cp goshot /usr/local/bin/ && \
	rm goshot && \
	mkdir -p ~/.local/bin && \
	sudo cp run-latex.sh run-text.sh ~/.local/bin/ && \
	sudo mv ~/.local/bin/run-latex.sh ~/.local/bin/run-latex && \
    sudo mv ~/.local/bin/run-text.sh ~/.local/bin/run-text && \
    sudo chmod +x ~/.local/bin/run-latex ~/.local/bin/run-text
	grep -q 'export PATH=$$PATH:$$HOME/.local/bin' ~/.bashrc || echo 'export PATH=$$PATH:$$HOME/.local/bin' >> ~/.bashrc && \
	source ~/.bashrc


# Clean up
clean:
	rm -rf goshot $(PYTHON_ENV)
	rm -rf build dist *.egg-info
	rm -rf $(VENV_PATH)
