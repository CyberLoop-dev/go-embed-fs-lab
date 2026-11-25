package main

import (
 "embed"
 "fmt"
 "io/fs"
 "log"
 "strings"
)

//go:embed all:assets
var embedded embed.FS

func main() {
 entries, err := fs.ReadDir(embedded, "assets")
 if err != nil { log.Fatal(err) }
 fmt.Println("Embedded top-level entries:")
 for _, entry := range entries { fmt.Printf("  %-12s dir=%v\n", entry.Name(), entry.IsDir()) }

 matches, err := fs.Glob(embedded, "assets/docs/*.md")
 if err != nil { log.Fatal(err) }
 fmt.Println("\nMarkdown matches:")
 for _, name := range matches { fmt.Println(" ", name) }

 docs, err := fs.Sub(embedded, "assets/docs")
 if err != nil { log.Fatal(err) }
 fmt.Println("\nFiles inside the docs sub-filesystem:")
 err = fs.WalkDir(docs, ".", func(path string, d fs.DirEntry, walkErr error) error {
  if walkErr != nil { return walkErr }
  if !d.IsDir() { fmt.Println(" ", path) }
  return nil
 })
 if err != nil { log.Fatal(err) }

 content, err := fs.ReadFile(docs, "welcome.md")
 if err != nil { log.Fatal(err) }
 fmt.Println("\nRead through fs.Sub:")
 fmt.Println(strings.TrimSpace(string(content)))
}
