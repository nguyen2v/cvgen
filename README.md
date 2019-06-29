# cvgen
cvgen is free and open source CV generator. 
Currently supports only pdf, but in the future it will support other formats.

### Requirements

* [wkhtmltopdf](https://github.com/wkhtmltopdf/wkhtmltopdf)
* go 1.11+

### Installation

You need to build this project by this command:
> go build ./cmd/cvgen

Then download and put wkhtmltopdf compatible with your operating system to cvgen folder or install to your path.

### Usage

At first, create YAML file (you can use example_cv.yaml and fill it with your data).
Then build cvgen and write this command:

> cvgen -f data.yaml -t basic -form pdf -o cv.pdf

This command will use data from data.yaml, fill basic template with provided data and write all to pdf file cv.pdf
You can view, create or edit templates at ./templates folder