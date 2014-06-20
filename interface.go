package gostatgrabber

// StatGrabber defines the API for communicating with the stat server
type StatGrabber interface {

	// Count sends tag to the server to increment a counter
	Count(tag string)

	// Average maintains an average of the value
	Average(tag string, value int)

	// Accumulate accumulates the value
	Accumulate(tag string, value int)
}

type StatTimer interface {

	// Elapsed time in seconds
	Elapsed() int
}
