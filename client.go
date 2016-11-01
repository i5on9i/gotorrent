package torrent

import (
	
)

type client struct {
	DataDir string

	noTorrents      chan struct{}
	actorTask       chan func()
}

func NewClient(dataDir string) *client {
	c := &client{
		DataDir: dataDir,

		noTorrents:      make(chan struct{}),
		actorTask:       make(chan func()),
	}
	
	return c
}
