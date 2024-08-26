package main

// TickMessage provides the ability to iterate through multiple strings based upon a target tick count for all messages to be displayed by.
type TickMessage struct {
	ticks        int
	currentTicks int
	messages     []string
}

// Update increments current ticks and loops as necessary.
func (t *TickMessage) Update() {
	t.currentTicks++
	if t.currentTicks >= t.ticks {
		t.currentTicks = 0
	}
}

// Message gets the current message.
func (t *TickMessage) Message() string {
	i := float64(t.currentTicks) / float64(t.ticks)
	return t.messages[int(i*float64(len(t.messages)))]
}
