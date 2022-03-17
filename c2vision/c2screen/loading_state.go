package c2screen

// LoadingState represents the loading state

// chan : channel
// ch <- data // send data to ch
// val <- ch  // block while no data in ch.
type LoadingState struct {
	updates chan loadingUpdate
}
