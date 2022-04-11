package chunk

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	tmpChunkDir = "tmp/"
)

type chunkFile struct {
	Data             []byte
	Id               string
	Directory        string
	OriginalFileName string
}

func CreateChunk(chunk []byte, id string, filename string, chunkName string) error {
	// New empty directory created at tmp/chunkName/
	if id == "0" {
		if err := createDirectory(chunkName); err != nil {
			return err
		}
	}

	// Chunk created at tmp/chunkName/filename_id.blb
	c := chunkFile{
		Data:             chunk,
		Id:               id,
		Directory:        tmpChunkDir + chunkName + "/",
		OriginalFileName: filename,
	}
	if err := createChunk(c); err != nil {
		return err
	}

	return nil
}

func createDirectory(filename string) error {
	if err := fileExists(tmpChunkDir + filename); err != nil {
		return err
	}

	if err := os.Mkdir(tmpChunkDir+filename, 0755); err != nil {
		return err
	}
	return nil
}

func createChunk(c chunkFile) error {
	if err := ioutil.WriteFile(fmt.Sprintf("%s%s_%s.blb", c.Directory, c.OriginalFileName, c.Id), c.Data, 0644); err != nil {
		return err
	}
	return nil
}
