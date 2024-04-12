package main

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		// Если нет каналов, возвращаем nil
		return nil
	case 1:
		// Если только один канал, возвращаем его
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		select {
		case <-channels[0]:
		case <-channels[1]:
		case <-or(append(channels[2:], orDone)...):
		}
	}()

	return orDone
}
