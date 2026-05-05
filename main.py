from flask import Flask, request, render_template_string


app = Flask(__name__)

def load_page(path):
  filename = f"{path}.txt"
  with open(filename,"r",encoding="utf-8") as f:
    body = f.read()
  return {"path":path,"body":body}

@app.route("/view/<path>")
def view_handler(path):
  p=load_page(path)

  html=f"<h1>{p['path']}</h1><div>{p['body']}</div>"

  return html

if __name__ =="__main__":
  app.run(port=8080)