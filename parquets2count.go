package parquets2count

import (
	"bufio"
	"io"
	"iter"
	"os"

	pq "github.com/parquet-go/parquet-go"
)

type ParquetFile struct{ *pq.File }

func (f ParquetFile) NumRows() int64 { return f.File.NumRows() }

func FilenameToRowCount(filename string) (int64, error) {
	file, err := os.Open(filename) //nolint:gosec
	if nil != err {
		return 0, err
	}
	defer file.Close() //nolint:errcheck

	stat, err := file.Stat()
	if nil != err {
		return 0, err
	}

	var sz int64 = stat.Size()
	pfile, err := pq.OpenFile(file, sz)
	if nil != err {
		return 0, err
	}
	return ParquetFile{pfile}.NumRows(), nil
}

func CountAllFiles(filenames iter.Seq[string]) (int64, error) {
	var tot int64 = 0
	for filename := range filenames {
		cnt, e := FilenameToRowCount(filename)
		if nil != e {
			return 0, e
		}
		tot += cnt
	}
	return tot, nil
}

func ReaderToFilenames(rdr io.Reader) iter.Seq[string] {
	return func(yield func(string) bool) {
		var s *bufio.Scanner = bufio.NewScanner(rdr)
		for s.Scan() {
			var filename string = s.Text()
			if !yield(filename) {
				return
			}
		}
	}
}

func StdinToFilenames() iter.Seq[string] { return ReaderToFilenames(os.Stdin) }

func StdinToFilenamesToParquetsToTotalCount() (int64, error) {
	var filenames iter.Seq[string] = StdinToFilenames()
	return CountAllFiles(filenames)
}
