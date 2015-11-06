package main

import "path/filepath"
import "text/template"
import "os"
import "strings"

func main() {
  names := make([]string, 0)
  t, _ := template.ParseFiles("t/panorama.html")
  files, _ := filepath.Glob("i/*.jpg")

  for _, name := range files {
    name = strings.TrimPrefix(name, "i/")
    name = strings.TrimPrefix(name, "i\\")
    name = strings.TrimSuffix(name, filepath.Ext(name))
    names = append(names, name)
	f, _ := os.Create(name + ".html")
	t.Execute(f, name)
	f.Close()
  }

  t, _ = template.ParseFiles("t/index.html")
  f, _ := os.Create("index.html")
  t.Execute(f, names)
  f.Close()
}
