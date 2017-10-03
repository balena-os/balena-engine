package progress

import (
	"time"
	"sync"

	"golang.org/x/time/rate"
)

type Sink struct {
	sync.Mutex
	out         Output        // Where to send progress bar to
	Size        int64
	current     int64
	lastUpdate  int64
	id          string
	action      string
	rateLimiter *rate.Limiter
}

// NewProgressSink creates a new ProgressSink.
func NewProgressSink(out Output, size int64, id, action string) *Sink {
	return &Sink{
		out:         out,
		Size:        size,
		id:          id,
		action:      action,
		rateLimiter: rate.NewLimiter(rate.Every(100*time.Millisecond), 1),
	}
}

func (p *Sink) Write(buf []byte) (int, error) {
	p.Lock()
	defer p.Unlock()

	n := len(buf)
	p.current += int64(n)
	updateEvery := int64(1024 * 512) //512kB
	if p.Size > 0 {
		// Update progress for every 1% if 1% < 512kB
		if increment := int64(0.01 * float64(p.Size)); increment < updateEvery {
			updateEvery = increment
		}
	}
	if p.current-p.lastUpdate > updateEvery || p.current == p.Size {
		if p.current == p.Size || p.rateLimiter.Allow() {
			p.out.WriteProgress(Progress{ID: p.id, Action: p.action, Current: p.current, Total: p.Size, LastUpdate: p.current == p.Size})
		}
		p.lastUpdate = p.current
	}

	return n, nil
}
