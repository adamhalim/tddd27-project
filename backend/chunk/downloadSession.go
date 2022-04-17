package chunk

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"
)

// This file handles current file downloads

type DownloadSession struct {
	chunkCount       int
	totalSize        int
	lock             *sync.Mutex
	chunkName        string
	lastModified     time.Time
	directory        string
	chunks           []*chunkFile
	originalFileName string
	fileExtension    string
}

var (
	sessions    map[string]*DownloadSession
	sessionLock *sync.Mutex
)

func init() {
	sessions = make(map[string]*DownloadSession)
	sessionLock = &sync.Mutex{}
}

func newSession(chunkName string, fileName string) error {
	sessionLock.Lock()
	defer sessionLock.Unlock()
	if _, ok := sessions[chunkName]; ok {
		return fmt.Errorf("session for %s already exists", chunkName)
	}

	fileExtension := filepath.Ext(fileName)

	sessions[chunkName] = &DownloadSession{
		chunkCount:       0,
		totalSize:        0,
		lock:             &sync.Mutex{},
		chunkName:        chunkName,
		lastModified:     time.Now(),
		directory:        tmpChunkDir + chunkName,
		chunks:           make([]*chunkFile, 0),
		originalFileName: fileName,
		fileExtension:    fileExtension,
	}

	return nil
}

func (s *DownloadSession) addChunk(chunk *chunkFile) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.chunkCount >= maxChunkCount {
		// TODO: Run cleanup here
		return fmt.Errorf("file too large. maximum # chunks is %d", maxChunkCount)
	}
	s.chunkCount++
	s.lastModified = time.Now()
	s.chunks = append(s.chunks, chunk)
	return nil
}

func GetSession(chunkName string) (*DownloadSession, error) {
	if _, ok := sessions[chunkName]; !ok {
		return nil, fmt.Errorf("no session for %s exists", chunkName)
	}
	session := sessions[chunkName]
	return session, nil
}
