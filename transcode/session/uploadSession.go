package session

import (
	"fmt"
	"sync"
)

type session struct {
	lock               *sync.Mutex
	FileName           string
	ChunkName          string
	OriginalFileName   string
	Dir                string
	Uid                string
	TranscodedFileName string
}

var (
	sessions    map[string]*session
	sessionLock *sync.Mutex
)

func init() {
	sessions = make(map[string]*session)
	sessionLock = &sync.Mutex{}
}

func NewSession(chunkName string, fileName string, dir string, originalFileName string, uid string) (*session, error) {
	sessionLock.Lock()
	defer sessionLock.Unlock()

	if _, ok := sessions[chunkName]; ok {
		return nil, fmt.Errorf("session for %s already exists", chunkName)
	}

	newSession := &session{
		lock:             &sync.Mutex{},
		FileName:         fileName,
		ChunkName:        chunkName,
		OriginalFileName: originalFileName,
		Dir:              dir,
		Uid:              uid,
	}
	sessions[chunkName] = newSession

	return newSession, nil
}

func GetSession(chunkName string) (*session, error) {
	if _, ok := sessions[chunkName]; !ok {
		return nil, fmt.Errorf("no session for %s exists", chunkName)
	}
	session := sessions[chunkName]
	return session, nil
}

func (session *session) RemoveSession() {
	// TODO: Also do cleanup here?
	delete(sessions, session.ChunkName)
}
