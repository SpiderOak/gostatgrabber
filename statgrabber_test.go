package gostatgrabber

import (
	"bytes"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	receiveChan, err := statReceiver(t)
	if err != nil {
		t.Fatalf("statReceiver() %s", err)
	}
	<-receiveChan

	s, err := NewStatGrabber()
	if err != nil {
		t.Fatalf("NewStatGrabber() %s", err)
	}
	tag := "pork"
	s.Count(tag)

	result := <-receiveChan
	t.Logf("TestCounter received '%s'", result)
	if !bytes.Equal(result, []byte(tag)) {
		t.Fatalf("unexpected result '%q' expecting '%q' %d",
			result, tag, bytes.Compare(result, []byte(tag)))
	}
}

func TestAverage(t *testing.T) {
	receiveChan, err := statReceiver(t)
	if err != nil {
		t.Fatalf("statReceiver() %s", err)
	}
	<-receiveChan

	s, err := NewStatGrabber()
	if err != nil {
		t.Fatalf("NewStatGrabber() %s", err)
	}
	tag := "pork"
	s.Average(tag, 42)

	result := <-receiveChan
	t.Logf("TestCounter received '%s'", result)
	if !bytes.Equal(result, []byte("pork 42")) {
		t.Fatalf("unexpected result '%q' %d", result,
			bytes.Compare(result, []byte("pork 42")))
	}
}

func TestAccumulate(t *testing.T) {
	receiveChan, err := statReceiver(t)
	if err != nil {
		t.Fatalf("statReceiver() %s", err)
	}
	<-receiveChan

	s, err := NewStatGrabber()
	if err != nil {
		t.Fatalf("NewStatGrabber() %s", err)
	}
	tag := "pork"
	s.Accumulate(tag, 42)

	result := <-receiveChan
	t.Logf("TestCounter received '%s'", result)
	if !bytes.Equal(result, []byte("pork +42")) {
		t.Fatalf("unexpected result '%q' %d", result,
			bytes.Compare(result, []byte("pork +42")))
	}
}

func TestStatTimer(t *testing.T) {
	timer := NewStatTimer()
	time.Sleep(time.Second * 1)
	elapsed := timer.Elapsed()
	if elapsed != 1 {
		t.Fatalf("unexpected elapsed %d", elapsed)
	}
}

// statReceiver is a test utility that receives a single UDP packet and passes
// it through the channel
func statReceiver(t *testing.T) (<-chan []byte, error) {
	receiveChan := make(chan []byte)

	udpAddr, err := net.ResolveUDPAddr("udp", Addr)
	if err != nil {
		return nil, fmt.Errorf("ResolveUDPAddr %s", err)
	}

	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, fmt.Errorf("ListenUDP %s", err)
	}

	go func() {
		buffer := make([]byte, 1024)
		receiveChan <- []byte("ready")
		t.Logf("waiting read")
		n, err := udpConn.Read(buffer)
		if err != nil {
			t.Fatalf("udpConn.Read %s", err)
		}
		receiveChan <- buffer[:n]
		close(receiveChan)
		udpConn.Close()
	}()

	return receiveChan, nil
}
