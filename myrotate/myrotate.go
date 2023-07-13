package myrotate

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/jmnote/go-test/myclock"
)

func deleteByAge(age time.Duration) {
	filepath.Walk("data/", func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		point := myclock.Now().Add(-age).UTC().String()[0:10] + ".log"
		if info.Name() < point {
			fmt.Println("delete " + info.Name())
			os.RemoveAll("data/" + info.Name())
		}
		return nil
	})
}
