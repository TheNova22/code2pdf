import os
from  pygments import lexers, formatters, styles, highlight
import pdfkit
import sys
from langComment import langComment
from parser import Parser
import requests
from xvfbwrapper import Xvfb
import os

vdisplay = Xvfb()
vdisplay.start()


parser = Parser()
values = parser.get_args()

urls = []

if values.f or values.u:

    if values.f:
        for file in values.f:
            fname = os.path.join(os.getcwd(),file) if str(file).startswith("pdfs/") else os.path.join(os.getcwd(),"pdfs/", file)
            if os.path.isfile(fname) == False:
                raise Exception(f"{fname} couldn't be located. Please check the path and try again")
    if values.u:
        for u in values.u:
            if u[:11] != "https://raw":
                raise Exception(f"{u} is not a raw github url")

            else:
                resp = requests.get(u)
                urls.append(resp.text)

else:
    raise Exception("Missing either files or urls")


authorName = values.a

style = values.s

name = values.n

arr = []

i = 0
for l in urls:
    lex = lexers.guess_lexer_for_filename(values.u[i].split('/')[-1],'int a = 3;')
    i += 1
    if authorName:
        l = langComment[lex.name.split()[0].lower()] + " Authored By: " + authorName + "\n" + l

    x = highlight(l + '  \n' , lex, formatters.HtmlFormatter(linenos = 'inline', noclasses = True, style=style))

    arr.append(x)

    # pdfkit.from_string(x, 'x.pdf')
if values.f:
    for file in values.f:
        fname = os.path.join(os.getcwd(),file) if str(file).startswith("pdfs/") else os.path.join(os.getcwd(),"pdfs/", file)
        lex = lexers.guess_lexer_for_filename(fname,'int a = 3;')
        with open(fname) as f:
            l = f.read()
            
            if authorName:
                l = langComment[lex.name.lower()] + " Authored By: " + authorName + "\n" + l
            
            x = highlight(l + '  \n' , lexers.get_lexer_by_name(lex.name), formatters.HtmlFormatter(linenos = 'inline', noclasses = True,style=style))
            
            arr.append(x)
        
        # pdfkit.from_string(x, 'x.pdf')
html = ("<br></br>" * 2).join(arr)

if not name:
    name = str(html.__hash__() & sys.maxsize)

pdfkit.from_string(html, os.path.join(os.getcwd(), 'pdfs/') + name + '.pdf')
# pdf.output('x.pdf', 'F')
print("Pdf created as " + name + '.pdf')
vdisplay.stop()
