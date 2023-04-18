// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.

package console

import (
	"bytes"
	"sync"
	"time"

	"namespacelabs.dev/foundation/std/tasks"
	"namespacelabs.dev/foundation/std/tasks/idtypes"
)

type writesLines interface {
	WriteLines(idtypes.IdAndHash, string, idtypes.CatOutputType, tasks.ActionID, time.Time, [][]byte)
}

type consoleBuffer struct {
	actual   []writesLines
	name     string
	cat      idtypes.CatOutputType
	id       idtypes.IdAndHash
	actionID tasks.ActionID // Optional ActionID in case this buffer is used in a context of an Action.

	mu  sync.Mutex
	buf bytes.Buffer
	ts  *time.Time
}

func (w *consoleBuffer) Write(p []byte) (int, error) {
	w.mu.Lock()

	if w.ts == nil {
		ts := time.Now()
		w.ts = &ts
	}

	w.buf.Write(p)
	var lines [][]byte
	ts := w.ts
	for {
		if i := bytes.IndexByte(w.buf.Bytes(), '\n'); i >= 0 {
			data := make([]byte, i+1)
			_, _ = w.buf.Read(data)
			line := dropCR(data[0 : len(data)-1]) // Drop the \n and the \r.
			lines = append(lines, line)
		} else {
			// Always assume the timestamp of the next write that may end up commiting lines.
			w.ts = nil
			break
		}
	}
	w.mu.Unlock()
	if len(lines) > 0 {
		for _, actual := range w.actual {
			actual.WriteLines(w.id, w.name, w.cat, w.actionID, *ts, lines)
		}
	}
	return len(p), nil
}

func dropCR(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\r' {
		return data[0 : len(data)-1]
	}
	return data
}
