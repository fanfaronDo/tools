package rotate

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetFileName(name string) string {
	directory := strings.Split(name, "/")

	return directory[len(directory)-1]
}

func CreateArchive(archiveName string) {
	archive, _ := os.Create(archiveName)
	w := gzip.NewWriter(archive)
	tw := tar.NewWriter(w)
	addToArchive(tw, archiveName)

	archive.Close()
	w.Close()
	tw.Close()
}

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	log.Println(info.Name())
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	header.Name = info.Name()
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}

	return nil
}

func ArchiveNameBuilder(filename string) string {
	timestamp, _ := getUnixTime(filename)

	return regexp.MustCompile(`\.[^.]+$`).ReplaceAllString(filename, "") + "_" + strconv.Itoa(timestamp) + ".tar.gz"
}

func getUnixTime(filename string) (int, error) {
	FILE, err := os.Open(filename)
	defer FILE.Close()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	stat, err := FILE.Stat()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(stat.ModTime().Unix()), nil
}

func MoveTo(oldLoc, newLoc string) error {
	err := os.Rename(oldLoc, newLoc)
	if err != nil {
		return err
	}
	return nil
}
