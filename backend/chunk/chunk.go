package chunk

import (
	"fmt"
	"io/ioutil"
	"os"
)

const (
	tmpChunkDir   = "tmp/"
	oneMB         = 1_048_576
	MaxFileSize   = 200 * oneMB
	ChunkSize     = 5 * oneMB
	maxChunkCount = MaxFileSize / ChunkSize
)

type chunkFile struct {
	Data     []byte
	Id       string
	FileName string
}

func CreateChunk(chunk []byte, id string, filename string, chunkName string) error {
	if len(chunk) > ChunkSize {
		// TOOD: Do cleanup
		return fmt.Errorf("chunk size greater than %db", ChunkSize)
	}
	// New empty directory created at tmp/chunkName/
	if id == "0" {
		if err := createDirectory(chunkName); err != nil {
			return err
		}
		if err := newSession(chunkName, filename); err != nil {
			// Terminate all future entries with this chunkName?
			return err
		}
	}

	session, err := getSession(chunkName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// Chunk created at tmp/chunkName/filename_id.blb
	c := chunkFile{
		Data: chunk,
		Id:   id,
	}
	c.FileName = fmt.Sprintf("%s%s_%s.blb", session.directory, c.Id, session.originalFileName)
	if err := createChunk(c); err != nil {
		return err
	}

	if err := session.addChunk(&c); err != nil {
		return err
	}
	return nil
}

// Combine all chunks for a session into a single result file
func CombineChunks(chunkName string) error {
	session, err := getSession(chunkName)
	if err != nil {
		return err
	}
	defer delete(sessions, chunkName)

	resultFileName := fmt.Sprintf("%s_%s", session.chunkName, session.originalFileName)
	resultFile, err := os.Create(session.directory + "/" + resultFileName)
	defer resultFile.Close()
	if err != nil {
		return err
	}

	for _, chunk := range session.chunks {
		byteFile, err := ioutil.ReadFile(chunk.FileName)
		if err != nil {
			return err
		}
		_, err = resultFile.Write(byteFile)
		if err != nil {
			return err
		}
		err = os.Remove(chunk.FileName)
		if err != nil {
			return err
		}
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
	if err := ioutil.WriteFile(c.FileName, c.Data, 0644); err != nil {
		return err
	}
	// We don't need to keep the data in memory after it's written to file
	c.Data = nil
	return nil
}