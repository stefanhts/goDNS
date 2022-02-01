package main

type BytePacketBuffer struct {
	Buff [512]uint8
	Pos  uint16
}

func newBuffer() *BytePacketBuffer {
	b := BytePacketBuffer{
		Buff: [512]uint8{0},
		Pos:  0,
	}
	return &b
}

func (b *BytePacketBuffer) pos() uint16 {
	return b.Pos
}

func (b *BytePacketBuffer) step(steps uint16) {
	if res := steps + b.Pos; res >= 512 {
		panic("buffer exceeded")
	}
	b.Pos += steps
}

func (b *BytePacketBuffer) seek(pos uint16) {
	if pos >= 512 {
		panic("buffer exceeded")
	}
	b.Pos = pos
}

func (b *BytePacketBuffer) read() uint8 {
	if b.Pos >= 512 {
		panic("End of buffer")
	}
	res := b.Buff[b.Pos]
	b.Pos++
	return res
}

func (b *BytePacketBuffer) get() uint8 {
	if b.Pos >= 512 {
		panic("End of buffer")
	}
	return b.Buff[b.Pos]
}

func (b *BytePacketBuffer) getRange(size uint16) []uint8 {
	if b.Pos+size >= 512 {
		panic("End of buffer")
	}
	return b.Buff[b.Pos : b.Pos+size]
}
