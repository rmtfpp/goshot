from setuptools import setup, find_packages

setup(
    name='latexocr',
    version='0.1.0',
    packages=find_packages(),
    python_requires='==3.12.*',
    entry_points={
        'console_scripts': [
            'latexocr=latexocr.ocr:main_cli',
        ],
    },
    install_requires=[
        'Pillow',
        'pix2tex',
    ],
)