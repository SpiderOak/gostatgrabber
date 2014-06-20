package gostatgrabber

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"time"
)

const (
	Addr = "127.0.0.1:9119"
)

type statGrabber struct {
	UDPConn *net.UDPConn
}

var (
	fixTagRegex = regexp.MustCompile(`\s+`)
)

// NewStatGrabber returns an object that implements the StatGrabber interface
func NewStatGrabber() (StatGrabber, error) {
	var s statGrabber
	var udpAddr *net.UDPAddr
	var err error

	if udpAddr, err = net.ResolveUDPAddr("udp", Addr); err != nil {
		return s, fmt.Errorf("ResolveUDPAddr %s", err)
	}

	if s.UDPConn, err = net.DialUDP("udp", nil, udpAddr); err != nil {
		return s, fmt.Errorf("DialUDP %s", err)
	}

	return s, nil
}

// Count sends tag to the server to increment a counter
func (s statGrabber) Count(tag string) {
	if _, err := s.UDPConn.Write([]byte(fixTag(tag))); err != nil {
		log.Printf("StatGrabber Count Writer error %s", err)
	}
}

// Average maintains an average of the value
func (s statGrabber) Average(tag string, value int) {
	message := fmt.Sprintf("%s %d", fixTag(tag), value)
	if _, err := s.UDPConn.Write([]byte(message)); err != nil {
		log.Printf("StatGrabber Average Writer error %s", err)
	}
}

// Accumulate accumulates the value
func (s statGrabber) Accumulate(tag string, value int) {
	message := fmt.Sprintf("%s +%d", fixTag(tag), value)
	if _, err := s.UDPConn.Write([]byte(message)); err != nil {
		log.Printf("StatGrabber Average Writer error %s", err)
	}
}

type statTimer struct {
	StartTime time.Time
}

// NewStatTimer returns an object that implements the StatTimer interface
func NewStatTimer() StatTimer {
	return statTimer{StartTime: time.Now()}
}

func (t statTimer) Elapsed() int {
	return int(time.Since(t.StartTime).Seconds())
}

// fixTag replaces spaces with '_'
func fixTag(tag string) string {
	return fixTagRegex.ReplaceAllString(tag, "_")
}
