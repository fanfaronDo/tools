package main


import (
  "github.com/fanfaronDo/tools/pkg/filter"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var opt Options

func main() {
  flarDirectory := flag.Bool("d", false, "Find only directory")
  flagFile := flag.Bool("f", false, "Find only file")
  flagLinks := flag.Bool("sl", false, "Find only links")
  flagExt := flag.String("ext", "", "Find concretly file extention")
  flag.Parse()

  if len(os.Args) < 2 {
    fmt.Printf("Usage: ./myFind /foo")
    return
  }
  opt.D = *flarDirectory
  opt.F = *flagFile
  opt.Sl = *flagLinks
  opt.Ext = *flagExt

  path := os.Args[len(os.Args)-1]
  err := explainDirectory(path)
  
  if err != nil {
    fmt.Printf("Error walking the path %q: %v\n", path, err)
    return
  }
}


func explainDirectory(path string) error {
  filters, er := SetFilter()
  if er != nil {
    return er
  }
  err := filepath.Walk(
    path,
    func(path string, info fs.FileInfo, err error) error {
      if err != nil {
        fmt.Printf("Prevent panic by handling failure accessing a path %q: %v\n", path, err)
        return err
      }
      if len(filters) == 0 { // default filter
        fmt.Println(path)
      }else{
        for _, filter := range filters {
          Printer(path, filter)
        }
      }
      return nil
    })

    return err 
}

func Printer (path string, filter IFiltraeble) {
  if filter.Filtrate(path, opt){
    if _, ok := filter.(*FlagLinks); ok {
      link, err := os.Readlink(path)
      if err != nil {
        fmt.Println(path, "->", "[broken]")
      }else {
        fmt.Println(path, "->", link)
      }
    }else {
      fmt.Println(path)
    }
  }
}

func SetFilter() ([]IFiltraeble, error) {
  filters := []IFiltraeble{}
  if (!opt.F && !opt.D && !opt.Sl && opt.Ext == ""){
    return filters, nil
  }
  if  (!opt.F && !opt.D && !opt.Sl && opt.Ext != ""){
    return filters, errors.New("No filter specified")
  }
  if opt.F && opt.Ext != "" {
    filters = append(filters, &FlagFWithExt{})
  }else if opt.F && opt.Ext == "" {
    filters = append(filters, &FlagF{})
  }
  if opt.D {
    filters = append(filters, &FlagD{})
  }
  if opt.Sl {
    filters = append(filters, &FlagLinks{})
  }

  return filters, nil
}
