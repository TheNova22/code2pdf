# code2pdf

Generate pdfs through code files and raw github urls in one go

# Usage
## Through docker
First build it by running
```
docker build --pull --rm -f "Dockerfile" -t code2pdf:latest "." 
```

Then run the docker image by
```
docker run -v  /Users/sumant/code2pdf/run_by_args/pdfs:/pdfs/ code2pdf:latest --u https://raw.githubusercontent.com/TheNova22/tkArt/main/circleDesign.py --f hw.c "sol copy.py" --a "Jayant Sogikar" --n "dracula" --s dracula
```

We set up a volume where code files are stored and pdfs are saved.

ENTRYPOINT is used to accept args and kwargs.

## Python

```
python code2pdf.py --u https://raw.githubusercontent.com/TheNova22/tkArt/main/circleDesign.py --f hw.c "sol copy.py" --a "Jayant Sogikar" --f hw.c "sol copy.py" --n "zenburn" --s zenburn
```
