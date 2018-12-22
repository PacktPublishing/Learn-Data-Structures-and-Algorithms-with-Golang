//main package has examples shown
// in Hands-On Data Structures and algorithms with Go book
package main

// importing fmt,sort, errors, math/rand and strings packages

import (
	"fmt"
	"math/rand"
	"time"
)

//FeedSource class
type FeedSource struct {
	URL     string
	Channel string
	ID      string
}

//FeedReader class
type FeedReader interface {
	ReadSource() (sources []FeedSource, next time.Time, err error)
}

// Subscriber class
type Subscriber interface {
	GetUpdates() <-chan FeedSource
	EndSubscription() error
}

//Subscription class
type Subscription struct {
	reader  FeedReader
	updates chan FeedSource
	err     chan chan error
}

// GetSubscription method
func GetSubscription(reader FeedReader) *Subscription {
	var s *Subscription
	s = &Subscription{
		reader:  reader,
		updates: make(chan FeedSource),
		err:     make(chan chan error),
	}
	go s.LoopOver()
	return s
}

// GetUpdates method
func (s *Subscription) GetUpdates() <-chan FeedSource {
	return s.updates
}

//EndSubscription method
func (s *Subscription) EndSubscription() error {
	var err chan error
	err = make(chan error)
	s.err <- err
	return <-err
}

// CloseLoop method
func (s *Subscription) CloseLoop() {
	var err error
	for {
		var errc chan error
		select {
		case errc = <-s.err:
			errc <- err
			close(s.updates)
			return
		}
	}
}

//FetchLoop method
func (s *Subscription) FetchLoop() {
	var pending []FeedSource
	var next time.Time
	var err error
	for {
		var fetchDelay time.Duration
		var now time.Time
		if now = time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		startFetch = time.After(fetchDelay)

		select {
		case <-startFetch:
			var fetched []FeedSource
			fetched, next, err = s.reader.ReadSource()
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			pending = append(pending, fetched...)
		}
	}
}

//SendLoop method
func (s *Subscription) SendLoop() {
	var pending []FeedSource
	for {
		var first FeedSource
		var updates chan FeedSource
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}

		select {
		case updates <- first:
			pending = pending[1:]
		}
	}
}

//GetMergedLoop method
func (s *Subscription) GetMergedLoop() {
	var pending []FeedSource
	var next time.Time
	var err error
	for {
		var fetchDelay time.Duration
		var now time.Time
		if now = time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		startFetch = time.After(fetchDelay)
		var first FeedSource
		var updates chan FeedSource
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}
		var errc chan error
		select {
		case errc = <-s.err:
			errc <- err
			close(s.updates)
			return
		case <-startFetch:
			var fetched []FeedSource
			fetched, next, err = s.reader.ReadSource()
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			pending = append(pending, fetched...)
		case updates <- first:
			pending = pending[1:]
		}
	}
}

//dedupeLoop method
func (s *Subscription) DedupeLoop() {
	const maxPending = 10
	var pending []FeedSource
	var next time.Time
	var err error
	var seen = make(map[string]bool)
	for {
		var fetchDelay time.Duration
		var now time.Time
		if now = time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		if len(pending) < maxPending {
			startFetch = time.After(fetchDelay)
		}
		var first FeedSource
		var updates chan FeedSource
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}
		var errc chan error
		select {
		case errc = <-s.err:
			errc <- err
			close(s.updates)
			return
		case <-startFetch:
			var fetched []FeedSource
			fetched, next, err = s.reader.ReadSource()
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			var FeedSource FeedSource
			for _, FeedSource = range fetched {
				if !seen[FeedSource.ID] {
					pending = append(pending, FeedSource)
					seen[FeedSource.ID] = true
				}
			}
		case updates <- first:
			pending = pending[1:]
		}
	}
}

// LoopOver method
func (s *Subscription) LoopOver() {
	const maxPending = 10
	type ReadOutput struct {
		read []FeedSource
		next time.Time
		err  error
	}
	var fetchDone chan ReadOutput
	var pending []FeedSource
	var next time.Time
	var err error
	var seen = make(map[string]bool)
	for {
		var fetchDelay time.Duration
		var now time.Time
		if now = time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending {
			startFetch = time.After(fetchDelay)
		}
		var first FeedSource
		var updates chan FeedSource
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates
		}
		var result ReadOutput
		var fetched []FeedSource
		var errc chan error
		select {
		case <-startFetch:
			fetchDone = make(chan ReadOutput, 1)
			go func() {
				fetched, next, err := s.reader.ReadSource()
				fetchDone <- ReadOutput{fetched, next, err}
			}()
		case result = <-fetchDone:
			fetchDone = nil
			fetched = result.read
			next, err = result.next, result.err
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			var FeedSource FeedSource
			for _, FeedSource = range fetched {
				var id string
				if id = FeedSource.ID; !seen[id] {
					pending = append(pending, FeedSource)
					seen[id] = true
				}
			}
		case errc = <-s.err:
			errc <- err
			close(s.updates)
			return
		case updates <- first:
			pending = pending[1:]
		}
	}
}

//MergedSubscriptions Class
type MergedSubscriptions struct {
	subscriptions []*Subscription
	updates       chan FeedSource
}

//GetMergedSubscriptions method
func GetMergedSubscriptions(subs ...*Subscription) *MergedSubscriptions {
	var m *MergedSubscriptions
	m = &MergedSubscriptions{
		subscriptions: subs,
		updates:       make(chan FeedSource),
	}
	var sub *Subscription
	for _, sub = range subs {
		go func(s *Subscription) {
			var it FeedSource
			for it = range s.GetUpdates() {
				m.updates <- it
			}
		}(sub)
	}
	return m
}

//EndSubscription method
func (m *MergedSubscriptions) EndSubscription() (err error) {
	var sub *Subscription
	for _, sub = range m.subscriptions {
		var e error
		if e = sub.EndSubscription(); err == nil && e != nil {
			err = e
		}
	}
	close(m.updates)
	return
}

//GetUpdates method
func (m *MergedSubscriptions) GetUpdates() <-chan FeedSource {
	return m.updates
}

// MergeUpdates class
type MergeUpdates struct {
	subscriptions []*Subscription
	updates       chan FeedSource
	exit          chan struct{}
	errs          chan error
}

//MergedSubscriptions method
func MergeFeedSources(subs ...*Subscription) *MergeUpdates {
	var m *MergeUpdates
	m = &MergeUpdates{
		subscriptions: subs,
		updates:       make(chan FeedSource),
		exit:          make(chan struct{}),
		errs:          make(chan error),
	}
	var sub *Subscription
	for _, sub = range subs {
		go func(s *Subscription) {
			for {
				var it FeedSource
				select {
				case it = <-s.GetUpdates():
				case <-m.exit:
					m.errs <- s.EndSubscription()
					return
				}
				select {
				case m.updates <- it:
				case <-m.exit:
					m.errs <- s.EndSubscription()
					return
				}
			}
		}(sub)
	}
	return m
}

//GetUpdates method
func (m *MergeUpdates) GetUpdates() <-chan FeedSource {
	return m.updates
}

//EndSubscription method
func (m *MergeUpdates) EndSubscription() (err error) {
	close(m.exit)
	for _ = range m.subscriptions {
		var e error
		if e = <-m.errs; e != nil {
			err = e
		}
	}
	close(m.updates)
	return
}

//GetNaiveDedupe method
func GetNaiveDedupe(in <-chan FeedSource) <-chan FeedSource {
	var out chan FeedSource
	out = make(chan FeedSource)
	go func() {
		var seen map[string]bool
		seen = make(map[string]bool)
		var it FeedSource
		for it = range in {
			if !seen[it.ID] {
				out <- it
				seen[it.ID] = true
			}
		}
		close(out)
	}()
	return out
}

//DeDuper class
type DeDuper struct {
	subscriber Subscription
	updates    chan FeedSource
	err        chan chan error
}

//Dedupe method
func Dedupe(s Subscription) *DeDuper {
	var d *DeDuper
	d = &DeDuper{
		subscriber: s,
		updates:    make(chan FeedSource),
		err:        make(chan chan error),
	}
	go d.LoopOver()
	return d
}

// LoopOver method
func (d *DeDuper) LoopOver() {
	var in <-chan FeedSource
	in = d.subscriber.GetUpdates()
	var pending FeedSource
	var out chan FeedSource
	var seen map[string]bool
	seen = make(map[string]bool)
	var errc chan error
	var it FeedSource
	for {
		select {
		case it = <-in:
			if !seen[it.ID] {
				pending = it
				in = nil
				out = d.updates
				seen[it.ID] = true
			}
		case out <- pending:
			in = d.subscriber.GetUpdates()
			out = nil
		case errc = <-d.err:
			var err error
			err = d.subscriber.EndSubscription()
			errc <- err
			close(d.updates)
			return
		}
	}
}

//EndSubscription method
func (d *DeDuper) EndSubscription() error {
	var errc chan error
	errc = make(chan error)
	d.err <- errc
	return <-errc
}

//GetUpdates method
func (d *DeDuper) GetUpdates() <-chan FeedSource {
	return d.updates
}

//ReadFeed method
func ReadFeed(domain string) *Reader {
	return Read(domain)
}

//Reader class
type Reader struct {
	channel string
	sources []FeedSource
}

// Read method
func Read(domain string) *Reader {
	return &Reader{channel: domain}
}

var Duplicates bool

//ReadSource method
func (f *Reader) ReadSource() (FeedSources []FeedSource, next time.Time, err error) {
	var now time.Time
	now = time.Now()
	next = now.Add(time.Duration(rand.Intn(5)) * 500 * time.Millisecond)
	var item FeedSource
	item = FeedSource{
		Channel: f.channel,
		URL:     fmt.Sprintf("FeedSource %d", len(f.sources)),
	}
	item.ID = item.Channel + "/" + item.URL
	f.sources = append(f.sources, item)
	if Duplicates {
		FeedSources = f.sources
	} else {
		FeedSources = []FeedSource{item}
	}
	return
}

// init method
func init() {
	rand.Seed(time.Now().UnixNano())
}

// main method
func main() {
	var merged *MergedSubscriptions
	merged = GetMergedSubscriptions(
		GetSubscription(ReadFeed("wwww.techcrunch.com")),
		GetSubscription(ReadFeed("www.techvibes.com")),
		GetSubscription(ReadFeed("slashdot.com")))

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.EndSubscription())
	})

	var it FeedSource
	for it = range merged.GetUpdates() {
		fmt.Println(it.Channel, it.URL)
	}

	panic("print the stacks")
}
