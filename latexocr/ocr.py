import sys
from PIL import Image
from pix2tex.cli import LatexOCR


import warnings
warnings.filterwarnings('ignore')

def process_image(image_path):
    img = Image.open(image_path)
    model = LatexOCR()
    return model(img)

def main_cli():

    if len(sys.argv) != 2:
        print("Usage: latexocr <image_path>")
        sys.exit(1)
    
    result = process_image(sys.argv[1])
    print(result)
