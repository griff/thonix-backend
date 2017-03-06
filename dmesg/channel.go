package dmesg

func Channel(l LogReadCloser) (<-chan LogEntry, <-chan error) {
	ch := make(chan LogEntry)
	errch := make(chan error)
	go func() {
		for {
			entry, err := l.Read()
			if err != nil {
				errch <- err
				l.Close()
				close(ch)
				close(errch)
				break
			} else if entry == nil {
				close(ch)
				close(errch)
				l.Close()
				break
			} else {
				ch <- *entry
			}
		}
	}()
	return ch, errch
}
