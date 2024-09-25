package filter


import (
	"strings"
	"os"
)

type Options struct {
	D bool
	F bool
	Sl bool
	Ext string
}

type IFiltraeble interface {
	Filtrate(path string, opt Options) bool
  }
  
  type FlagF struct {}
  
  func (f *FlagF) Filtrate(path string, opt Options) bool {
	_, err := os.ReadFile(path)
	return err == nil
  }
  type FlagFWithExt struct {}
  
  func (f *FlagFWithExt) Filtrate(path string, opt Options) bool {
	_, err := os.ReadFile(path)
	return err == nil && strings.HasSuffix(path, opt.Ext)
  }
  
  type FlagD struct {}
  
  func (f *FlagD) Filtrate(path string, opt Options) bool {
	file, _ := os.Open(path)
	f_stat, _ := file.Stat()
	defer file.Close()
  
	return opt.D && f_stat.IsDir()
  }
  
  type FlagLinks struct {}
  
  func (f * FlagLinks) Filtrate(path string, opt Options) bool {
	f_stat, _ := os.Lstat(path)
	return f_stat.Mode()&os.ModeSymlink != 0
  }
  
 func GetOptions() Options{
	return Options{}
 }