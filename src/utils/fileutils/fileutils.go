package fileutils

import "os"

func OpenFile(path string) (*os.File, error){
	f, err := os.Open(path)
	
	if err != nil {
		return nil, err
	}

	return f, nil
}


