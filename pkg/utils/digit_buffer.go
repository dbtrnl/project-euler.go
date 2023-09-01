package utils

type DigitBuffer struct {
	Data   int
	count  int
	size   int
	IsFull bool
}

// Creates a digit buffer of size N
func NewDigitBuffer(n int) *DigitBuffer {
	return &DigitBuffer{
		size:   n,
		IsFull: false,
	}
}

// Pushes a rune to the digit buffer. If not a digit, it is ignored.
func (dBf *DigitBuffer) Push(r rune) {
	isRuneDigit := r >= 48 && r <= 57 // Unicode 48 to 57 (digits 0 to 9)
	if !isRuneDigit {
		return
	}
	digit := int(r - '0')
	if dBf.IsFull {
		return
	}
	dBf.Data = dBf.Data*10 + digit
	dBf.count++
	if dBf.count == dBf.size {
		dBf.IsFull = true
	}
}

func (dBf *DigitBuffer) Empty() {
	dBf.Data, dBf.count, dBf.IsFull = 0, 0, false
}
