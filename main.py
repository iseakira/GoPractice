from flask import Flask, request, render_template_string


app = Flask(__name__)

def load_page(title):
  filename = f"{title}.txt"
  with open(filename,"r",encoding="utf-8") as f:
    body = f.read()
  return {"title":title,"body":body}

@app.route("/view/<title>")
def view_handler(title):
  p=load_page(title)

  html=f"<h1>{p['title']}</h1><div>{p['body']}</div>"

  return html

if __name__ =="__main__":
  app.run(port=8080)