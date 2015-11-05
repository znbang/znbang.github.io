package main

import "path/filepath"
import "text/template"
import "os"
import "strings"

func main() {
  names := make([]string, 0)
  t, _ := template.ParseFiles("panorama.t")
  files, _ := filepath.Glob("images/*.jpg")

  for _, name := range files {
    name = strings.TrimPrefix(name, "images/")
    name = strings.TrimPrefix(name, "images\\")
    name = strings.TrimSuffix(name, filepath.Ext(name))
    names = append(names, name)
	f, _ := os.Create(name + ".html")
	t.Execute(f, name)
	f.Close()
  }

  t, _ = template.ParseFiles("index.t")
  f, _ := os.Create("index.html")
  t.Execute(f, names)
  f.Close()
}
