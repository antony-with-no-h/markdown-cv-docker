# markdown-cv

Dockerfile and ancillary scripts to produce a PDF CV from markdown.

## Usage

In your CV directory create a `Makefile`, calling `pandoc` and `html-to-pdf` passing any stylesheets you want to use in the arguments to `pandoc`.

```
all:
    @pandoc -o cv.html -c css/normalize.css -c css/stylesheet.css -s cv.md
    @html-to-pdf
```

`html-to-pdf` will try to open `file:///cv/cv.html` and write `/cv/cv.pdf`. If you want to change these use `CV_FILE` and `CV_PDF`.

```
all:
    @pandoc -o dog-biscuits.html -c css/normalize.css -c css/stylesheet.css -s cv.md
    @CV_FILE=dog-biscuits.html CV_PDF=dog-biscuits.pdf html-to-pdf
```

Start a container with your CV directory mounted as `/cv`. 

```
docker run --rm --cap-add SYS_ADMIN \
    --mount type=bind,source="$(pwd)/example",target=/cv \
    markdown-cv:latest
```
