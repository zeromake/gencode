package testdata

import (
	"unsafe"
)

type LuckScriptIndex struct {
	Id    int
	Stage int
	Type  int
	Play  *StagePlay
	Desks []int
}

func (d *LuckScriptIndex) Size() (s uint64) {

	{
		if d.Play != nil {

			{
				s += (*d.Play).Size()
			}
			s += 0
		}
	}
	{
		l := uint64(len(d.Desks))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		s += 8 * l

	}
	s += 25
	return
}
func (d *LuckScriptIndex) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.Id >> 0)

		buf[1+0] = byte(d.Id >> 8)

		buf[2+0] = byte(d.Id >> 16)

		buf[3+0] = byte(d.Id >> 24)

		buf[4+0] = byte(d.Id >> 32)

		buf[5+0] = byte(d.Id >> 40)

		buf[6+0] = byte(d.Id >> 48)

		buf[7+0] = byte(d.Id >> 56)

	}
	{

		buf[0+8] = byte(d.Stage >> 0)

		buf[1+8] = byte(d.Stage >> 8)

		buf[2+8] = byte(d.Stage >> 16)

		buf[3+8] = byte(d.Stage >> 24)

		buf[4+8] = byte(d.Stage >> 32)

		buf[5+8] = byte(d.Stage >> 40)

		buf[6+8] = byte(d.Stage >> 48)

		buf[7+8] = byte(d.Stage >> 56)

	}
	{

		buf[0+16] = byte(d.Type >> 0)

		buf[1+16] = byte(d.Type >> 8)

		buf[2+16] = byte(d.Type >> 16)

		buf[3+16] = byte(d.Type >> 24)

		buf[4+16] = byte(d.Type >> 32)

		buf[5+16] = byte(d.Type >> 40)

		buf[6+16] = byte(d.Type >> 48)

		buf[7+16] = byte(d.Type >> 56)

	}
	{
		if d.Play == nil {
			buf[i+24] = 0
		} else {
			buf[i+24] = 1

			{
				nbuf, err := (*d.Play).Marshal(buf[25:])
				if err != nil {
					return nil, err
				}
				i += uint64(len(nbuf))
			}
			i += 0
		}
	}
	{
		l := uint64(len(d.Desks))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+25] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+25] = byte(t)
			i++

		}
		for k0 := range d.Desks {

			{

				buf[i+0+25] = byte(d.Desks[k0] >> 0)

				buf[i+1+25] = byte(d.Desks[k0] >> 8)

				buf[i+2+25] = byte(d.Desks[k0] >> 16)

				buf[i+3+25] = byte(d.Desks[k0] >> 24)

				buf[i+4+25] = byte(d.Desks[k0] >> 32)

				buf[i+5+25] = byte(d.Desks[k0] >> 40)

				buf[i+6+25] = byte(d.Desks[k0] >> 48)

				buf[i+7+25] = byte(d.Desks[k0] >> 56)

			}

			i += 8

		}
	}
	return buf[:i+25], nil
}

func (d *LuckScriptIndex) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Id = 0 | (int(buf[i+0+0]) << 0) | (int(buf[i+1+0]) << 8) | (int(buf[i+2+0]) << 16) | (int(buf[i+3+0]) << 24) | (int(buf[i+4+0]) << 32) | (int(buf[i+5+0]) << 40) | (int(buf[i+6+0]) << 48) | (int(buf[i+7+0]) << 56)

	}
	{

		d.Stage = 0 | (int(buf[i+0+8]) << 0) | (int(buf[i+1+8]) << 8) | (int(buf[i+2+8]) << 16) | (int(buf[i+3+8]) << 24) | (int(buf[i+4+8]) << 32) | (int(buf[i+5+8]) << 40) | (int(buf[i+6+8]) << 48) | (int(buf[i+7+8]) << 56)

	}
	{

		d.Type = 0 | (int(buf[i+0+16]) << 0) | (int(buf[i+1+16]) << 8) | (int(buf[i+2+16]) << 16) | (int(buf[i+3+16]) << 24) | (int(buf[i+4+16]) << 32) | (int(buf[i+5+16]) << 40) | (int(buf[i+6+16]) << 48) | (int(buf[i+7+16]) << 56)

	}
	{
		if buf[i+24] == 1 {
			if d.Play == nil {
				d.Play = new(StagePlay)
			}

			{
				ni, err := (*d.Play).Unmarshal(buf[i+25:])
				if err != nil {
					return 0, err
				}
				i += ni
			}
			i += 0
		} else {
			d.Play = nil
		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+25] & 0x7F)
			for buf[i+25]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+25]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Desks)) >= l {
			d.Desks = d.Desks[:l]
		} else {
			d.Desks = make([]int, l)
		}
		for k0 := range d.Desks {

			{

				d.Desks[k0] = 0 | (int(buf[i+0+25]) << 0) | (int(buf[i+1+25]) << 8) | (int(buf[i+2+25]) << 16) | (int(buf[i+3+25]) << 24) | (int(buf[i+4+25]) << 32) | (int(buf[i+5+25]) << 40) | (int(buf[i+6+25]) << 48) | (int(buf[i+7+25]) << 56)

			}

			i += 8

		}
	}
	return i + 25, nil
}

type StagePlay struct {
	Settings StagePlaySettings
	Cards    []*CardConfig
}

func (d *StagePlay) Size() (s uint64) {

	{
		s += d.Settings.Size()
	}
	{
		l := uint64(len(d.Cards))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Cards {

			{
				if d.Cards[k0] != nil {

					{
						s += (*d.Cards[k0]).Size()
					}
					s += 0
				}
			}

			s += 1

		}

	}
	return
}
func (d *StagePlay) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		nbuf, err := d.Settings.Marshal(buf[0:])
		if err != nil {
			return nil, err
		}
		i += uint64(len(nbuf))
	}
	{
		l := uint64(len(d.Cards))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Cards {

			{
				if d.Cards[k0] == nil {
					buf[i+0] = 0
				} else {
					buf[i+0] = 1

					{
						nbuf, err := (*d.Cards[k0]).Marshal(buf[i+1:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
					i += 0
				}
			}

			i += 1

		}
	}
	return buf[:i+0], nil
}

func (d *StagePlay) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		ni, err := d.Settings.Unmarshal(buf[i+0:])
		if err != nil {
			return 0, err
		}
		i += ni
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Cards)) >= l {
			d.Cards = d.Cards[:l]
		} else {
			d.Cards = make([]*CardConfig, l)
		}
		for k0 := range d.Cards {

			{
				if buf[i+0] == 1 {
					if d.Cards[k0] == nil {
						d.Cards[k0] = new(CardConfig)
					}

					{
						ni, err := (*d.Cards[k0]).Unmarshal(buf[i+1:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					i += 0
				} else {
					d.Cards[k0] = nil
				}
			}

			i += 1

		}
	}
	return i + 0, nil
}

type CardConfig struct {
	Id        int
	X         int
	Y         int
	W         int
	H         int
	Points    int
	Suit      int
	Rotate    int
	Layer     int
	Tag       int
	Deps      []int
	Collect   []int
	Count     int
	Direction int
	Cover     bool
}

func (d *CardConfig) Size() (s uint64) {

	{
		l := uint64(len(d.Deps))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		s += 8 * l

	}
	{
		l := uint64(len(d.Collect))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		s += 8 * l

	}
	s += 97
	return
}
func (d *CardConfig) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.Id >> 0)

		buf[1+0] = byte(d.Id >> 8)

		buf[2+0] = byte(d.Id >> 16)

		buf[3+0] = byte(d.Id >> 24)

		buf[4+0] = byte(d.Id >> 32)

		buf[5+0] = byte(d.Id >> 40)

		buf[6+0] = byte(d.Id >> 48)

		buf[7+0] = byte(d.Id >> 56)

	}
	{

		buf[0+8] = byte(d.X >> 0)

		buf[1+8] = byte(d.X >> 8)

		buf[2+8] = byte(d.X >> 16)

		buf[3+8] = byte(d.X >> 24)

		buf[4+8] = byte(d.X >> 32)

		buf[5+8] = byte(d.X >> 40)

		buf[6+8] = byte(d.X >> 48)

		buf[7+8] = byte(d.X >> 56)

	}
	{

		buf[0+16] = byte(d.Y >> 0)

		buf[1+16] = byte(d.Y >> 8)

		buf[2+16] = byte(d.Y >> 16)

		buf[3+16] = byte(d.Y >> 24)

		buf[4+16] = byte(d.Y >> 32)

		buf[5+16] = byte(d.Y >> 40)

		buf[6+16] = byte(d.Y >> 48)

		buf[7+16] = byte(d.Y >> 56)

	}
	{

		buf[0+24] = byte(d.W >> 0)

		buf[1+24] = byte(d.W >> 8)

		buf[2+24] = byte(d.W >> 16)

		buf[3+24] = byte(d.W >> 24)

		buf[4+24] = byte(d.W >> 32)

		buf[5+24] = byte(d.W >> 40)

		buf[6+24] = byte(d.W >> 48)

		buf[7+24] = byte(d.W >> 56)

	}
	{

		buf[0+32] = byte(d.H >> 0)

		buf[1+32] = byte(d.H >> 8)

		buf[2+32] = byte(d.H >> 16)

		buf[3+32] = byte(d.H >> 24)

		buf[4+32] = byte(d.H >> 32)

		buf[5+32] = byte(d.H >> 40)

		buf[6+32] = byte(d.H >> 48)

		buf[7+32] = byte(d.H >> 56)

	}
	{

		buf[0+40] = byte(d.Points >> 0)

		buf[1+40] = byte(d.Points >> 8)

		buf[2+40] = byte(d.Points >> 16)

		buf[3+40] = byte(d.Points >> 24)

		buf[4+40] = byte(d.Points >> 32)

		buf[5+40] = byte(d.Points >> 40)

		buf[6+40] = byte(d.Points >> 48)

		buf[7+40] = byte(d.Points >> 56)

	}
	{

		buf[0+48] = byte(d.Suit >> 0)

		buf[1+48] = byte(d.Suit >> 8)

		buf[2+48] = byte(d.Suit >> 16)

		buf[3+48] = byte(d.Suit >> 24)

		buf[4+48] = byte(d.Suit >> 32)

		buf[5+48] = byte(d.Suit >> 40)

		buf[6+48] = byte(d.Suit >> 48)

		buf[7+48] = byte(d.Suit >> 56)

	}
	{

		buf[0+56] = byte(d.Rotate >> 0)

		buf[1+56] = byte(d.Rotate >> 8)

		buf[2+56] = byte(d.Rotate >> 16)

		buf[3+56] = byte(d.Rotate >> 24)

		buf[4+56] = byte(d.Rotate >> 32)

		buf[5+56] = byte(d.Rotate >> 40)

		buf[6+56] = byte(d.Rotate >> 48)

		buf[7+56] = byte(d.Rotate >> 56)

	}
	{

		buf[0+64] = byte(d.Layer >> 0)

		buf[1+64] = byte(d.Layer >> 8)

		buf[2+64] = byte(d.Layer >> 16)

		buf[3+64] = byte(d.Layer >> 24)

		buf[4+64] = byte(d.Layer >> 32)

		buf[5+64] = byte(d.Layer >> 40)

		buf[6+64] = byte(d.Layer >> 48)

		buf[7+64] = byte(d.Layer >> 56)

	}
	{

		buf[0+72] = byte(d.Tag >> 0)

		buf[1+72] = byte(d.Tag >> 8)

		buf[2+72] = byte(d.Tag >> 16)

		buf[3+72] = byte(d.Tag >> 24)

		buf[4+72] = byte(d.Tag >> 32)

		buf[5+72] = byte(d.Tag >> 40)

		buf[6+72] = byte(d.Tag >> 48)

		buf[7+72] = byte(d.Tag >> 56)

	}
	{
		l := uint64(len(d.Deps))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+80] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+80] = byte(t)
			i++

		}
		for k0 := range d.Deps {

			{

				buf[i+0+80] = byte(d.Deps[k0] >> 0)

				buf[i+1+80] = byte(d.Deps[k0] >> 8)

				buf[i+2+80] = byte(d.Deps[k0] >> 16)

				buf[i+3+80] = byte(d.Deps[k0] >> 24)

				buf[i+4+80] = byte(d.Deps[k0] >> 32)

				buf[i+5+80] = byte(d.Deps[k0] >> 40)

				buf[i+6+80] = byte(d.Deps[k0] >> 48)

				buf[i+7+80] = byte(d.Deps[k0] >> 56)

			}

			i += 8

		}
	}
	{
		l := uint64(len(d.Collect))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+80] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+80] = byte(t)
			i++

		}
		for k0 := range d.Collect {

			{

				buf[i+0+80] = byte(d.Collect[k0] >> 0)

				buf[i+1+80] = byte(d.Collect[k0] >> 8)

				buf[i+2+80] = byte(d.Collect[k0] >> 16)

				buf[i+3+80] = byte(d.Collect[k0] >> 24)

				buf[i+4+80] = byte(d.Collect[k0] >> 32)

				buf[i+5+80] = byte(d.Collect[k0] >> 40)

				buf[i+6+80] = byte(d.Collect[k0] >> 48)

				buf[i+7+80] = byte(d.Collect[k0] >> 56)

			}

			i += 8

		}
	}
	{

		buf[i+0+80] = byte(d.Count >> 0)

		buf[i+1+80] = byte(d.Count >> 8)

		buf[i+2+80] = byte(d.Count >> 16)

		buf[i+3+80] = byte(d.Count >> 24)

		buf[i+4+80] = byte(d.Count >> 32)

		buf[i+5+80] = byte(d.Count >> 40)

		buf[i+6+80] = byte(d.Count >> 48)

		buf[i+7+80] = byte(d.Count >> 56)

	}
	{

		buf[i+0+88] = byte(d.Direction >> 0)

		buf[i+1+88] = byte(d.Direction >> 8)

		buf[i+2+88] = byte(d.Direction >> 16)

		buf[i+3+88] = byte(d.Direction >> 24)

		buf[i+4+88] = byte(d.Direction >> 32)

		buf[i+5+88] = byte(d.Direction >> 40)

		buf[i+6+88] = byte(d.Direction >> 48)

		buf[i+7+88] = byte(d.Direction >> 56)

	}
	{
		if d.Cover {
			buf[i+96] = 1
		} else {
			buf[i+96] = 0
		}
	}
	return buf[:i+97], nil
}

func (d *CardConfig) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Id = 0 | (int(buf[i+0+0]) << 0) | (int(buf[i+1+0]) << 8) | (int(buf[i+2+0]) << 16) | (int(buf[i+3+0]) << 24) | (int(buf[i+4+0]) << 32) | (int(buf[i+5+0]) << 40) | (int(buf[i+6+0]) << 48) | (int(buf[i+7+0]) << 56)

	}
	{

		d.X = 0 | (int(buf[i+0+8]) << 0) | (int(buf[i+1+8]) << 8) | (int(buf[i+2+8]) << 16) | (int(buf[i+3+8]) << 24) | (int(buf[i+4+8]) << 32) | (int(buf[i+5+8]) << 40) | (int(buf[i+6+8]) << 48) | (int(buf[i+7+8]) << 56)

	}
	{

		d.Y = 0 | (int(buf[i+0+16]) << 0) | (int(buf[i+1+16]) << 8) | (int(buf[i+2+16]) << 16) | (int(buf[i+3+16]) << 24) | (int(buf[i+4+16]) << 32) | (int(buf[i+5+16]) << 40) | (int(buf[i+6+16]) << 48) | (int(buf[i+7+16]) << 56)

	}
	{

		d.W = 0 | (int(buf[i+0+24]) << 0) | (int(buf[i+1+24]) << 8) | (int(buf[i+2+24]) << 16) | (int(buf[i+3+24]) << 24) | (int(buf[i+4+24]) << 32) | (int(buf[i+5+24]) << 40) | (int(buf[i+6+24]) << 48) | (int(buf[i+7+24]) << 56)

	}
	{

		d.H = 0 | (int(buf[i+0+32]) << 0) | (int(buf[i+1+32]) << 8) | (int(buf[i+2+32]) << 16) | (int(buf[i+3+32]) << 24) | (int(buf[i+4+32]) << 32) | (int(buf[i+5+32]) << 40) | (int(buf[i+6+32]) << 48) | (int(buf[i+7+32]) << 56)

	}
	{

		d.Points = 0 | (int(buf[i+0+40]) << 0) | (int(buf[i+1+40]) << 8) | (int(buf[i+2+40]) << 16) | (int(buf[i+3+40]) << 24) | (int(buf[i+4+40]) << 32) | (int(buf[i+5+40]) << 40) | (int(buf[i+6+40]) << 48) | (int(buf[i+7+40]) << 56)

	}
	{

		d.Suit = 0 | (int(buf[i+0+48]) << 0) | (int(buf[i+1+48]) << 8) | (int(buf[i+2+48]) << 16) | (int(buf[i+3+48]) << 24) | (int(buf[i+4+48]) << 32) | (int(buf[i+5+48]) << 40) | (int(buf[i+6+48]) << 48) | (int(buf[i+7+48]) << 56)

	}
	{

		d.Rotate = 0 | (int(buf[i+0+56]) << 0) | (int(buf[i+1+56]) << 8) | (int(buf[i+2+56]) << 16) | (int(buf[i+3+56]) << 24) | (int(buf[i+4+56]) << 32) | (int(buf[i+5+56]) << 40) | (int(buf[i+6+56]) << 48) | (int(buf[i+7+56]) << 56)

	}
	{

		d.Layer = 0 | (int(buf[i+0+64]) << 0) | (int(buf[i+1+64]) << 8) | (int(buf[i+2+64]) << 16) | (int(buf[i+3+64]) << 24) | (int(buf[i+4+64]) << 32) | (int(buf[i+5+64]) << 40) | (int(buf[i+6+64]) << 48) | (int(buf[i+7+64]) << 56)

	}
	{

		d.Tag = 0 | (int(buf[i+0+72]) << 0) | (int(buf[i+1+72]) << 8) | (int(buf[i+2+72]) << 16) | (int(buf[i+3+72]) << 24) | (int(buf[i+4+72]) << 32) | (int(buf[i+5+72]) << 40) | (int(buf[i+6+72]) << 48) | (int(buf[i+7+72]) << 56)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+80] & 0x7F)
			for buf[i+80]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+80]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Deps)) >= l {
			d.Deps = d.Deps[:l]
		} else {
			d.Deps = make([]int, l)
		}
		for k0 := range d.Deps {

			{

				d.Deps[k0] = 0 | (int(buf[i+0+80]) << 0) | (int(buf[i+1+80]) << 8) | (int(buf[i+2+80]) << 16) | (int(buf[i+3+80]) << 24) | (int(buf[i+4+80]) << 32) | (int(buf[i+5+80]) << 40) | (int(buf[i+6+80]) << 48) | (int(buf[i+7+80]) << 56)

			}

			i += 8

		}
	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+80] & 0x7F)
			for buf[i+80]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+80]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Collect)) >= l {
			d.Collect = d.Collect[:l]
		} else {
			d.Collect = make([]int, l)
		}
		for k0 := range d.Collect {

			{

				d.Collect[k0] = 0 | (int(buf[i+0+80]) << 0) | (int(buf[i+1+80]) << 8) | (int(buf[i+2+80]) << 16) | (int(buf[i+3+80]) << 24) | (int(buf[i+4+80]) << 32) | (int(buf[i+5+80]) << 40) | (int(buf[i+6+80]) << 48) | (int(buf[i+7+80]) << 56)

			}

			i += 8

		}
	}
	{

		d.Count = 0 | (int(buf[i+0+80]) << 0) | (int(buf[i+1+80]) << 8) | (int(buf[i+2+80]) << 16) | (int(buf[i+3+80]) << 24) | (int(buf[i+4+80]) << 32) | (int(buf[i+5+80]) << 40) | (int(buf[i+6+80]) << 48) | (int(buf[i+7+80]) << 56)

	}
	{

		d.Direction = 0 | (int(buf[i+0+88]) << 0) | (int(buf[i+1+88]) << 8) | (int(buf[i+2+88]) << 16) | (int(buf[i+3+88]) << 24) | (int(buf[i+4+88]) << 32) | (int(buf[i+5+88]) << 40) | (int(buf[i+6+88]) << 48) | (int(buf[i+7+88]) << 56)

	}
	{
		d.Cover = buf[i+96] == 1
	}
	return i + 97, nil
}

type StagePlaySettings struct {
	Type    int
	Count   int
	Rate    int
	Collect []*StagePlayCollect
}

func (d *StagePlaySettings) Size() (s uint64) {

	{
		l := uint64(len(d.Collect))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Collect {

			{
				if d.Collect[k0] != nil {

					{
						s += (*d.Collect[k0]).Size()
					}
					s += 0
				}
			}

			s += 1

		}

	}
	s += 24
	return
}
func (d *StagePlaySettings) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.Type >> 0)

		buf[1+0] = byte(d.Type >> 8)

		buf[2+0] = byte(d.Type >> 16)

		buf[3+0] = byte(d.Type >> 24)

		buf[4+0] = byte(d.Type >> 32)

		buf[5+0] = byte(d.Type >> 40)

		buf[6+0] = byte(d.Type >> 48)

		buf[7+0] = byte(d.Type >> 56)

	}
	{

		buf[0+8] = byte(d.Count >> 0)

		buf[1+8] = byte(d.Count >> 8)

		buf[2+8] = byte(d.Count >> 16)

		buf[3+8] = byte(d.Count >> 24)

		buf[4+8] = byte(d.Count >> 32)

		buf[5+8] = byte(d.Count >> 40)

		buf[6+8] = byte(d.Count >> 48)

		buf[7+8] = byte(d.Count >> 56)

	}
	{

		buf[0+16] = byte(d.Rate >> 0)

		buf[1+16] = byte(d.Rate >> 8)

		buf[2+16] = byte(d.Rate >> 16)

		buf[3+16] = byte(d.Rate >> 24)

		buf[4+16] = byte(d.Rate >> 32)

		buf[5+16] = byte(d.Rate >> 40)

		buf[6+16] = byte(d.Rate >> 48)

		buf[7+16] = byte(d.Rate >> 56)

	}
	{
		l := uint64(len(d.Collect))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+24] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+24] = byte(t)
			i++

		}
		for k0 := range d.Collect {

			{
				if d.Collect[k0] == nil {
					buf[i+24] = 0
				} else {
					buf[i+24] = 1

					{
						nbuf, err := (*d.Collect[k0]).Marshal(buf[i+25:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
					i += 0
				}
			}

			i += 1

		}
	}
	return buf[:i+24], nil
}

func (d *StagePlaySettings) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Type = 0 | (int(buf[i+0+0]) << 0) | (int(buf[i+1+0]) << 8) | (int(buf[i+2+0]) << 16) | (int(buf[i+3+0]) << 24) | (int(buf[i+4+0]) << 32) | (int(buf[i+5+0]) << 40) | (int(buf[i+6+0]) << 48) | (int(buf[i+7+0]) << 56)

	}
	{

		d.Count = 0 | (int(buf[i+0+8]) << 0) | (int(buf[i+1+8]) << 8) | (int(buf[i+2+8]) << 16) | (int(buf[i+3+8]) << 24) | (int(buf[i+4+8]) << 32) | (int(buf[i+5+8]) << 40) | (int(buf[i+6+8]) << 48) | (int(buf[i+7+8]) << 56)

	}
	{

		d.Rate = 0 | (int(buf[i+0+16]) << 0) | (int(buf[i+1+16]) << 8) | (int(buf[i+2+16]) << 16) | (int(buf[i+3+16]) << 24) | (int(buf[i+4+16]) << 32) | (int(buf[i+5+16]) << 40) | (int(buf[i+6+16]) << 48) | (int(buf[i+7+16]) << 56)

	}
	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+24] & 0x7F)
			for buf[i+24]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+24]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Collect)) >= l {
			d.Collect = d.Collect[:l]
		} else {
			d.Collect = make([]*StagePlayCollect, l)
		}
		for k0 := range d.Collect {

			{
				if buf[i+24] == 1 {
					if d.Collect[k0] == nil {
						d.Collect[k0] = new(StagePlayCollect)
					}

					{
						ni, err := (*d.Collect[k0]).Unmarshal(buf[i+25:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					i += 0
				} else {
					d.Collect[k0] = nil
				}
			}

			i += 1

		}
	}
	return i + 24, nil
}

type StagePlayCollect struct {
	Rewards []*Reward
	Count   int
}

func (d *StagePlayCollect) Size() (s uint64) {

	{
		l := uint64(len(d.Rewards))

		{

			t := l
			for t >= 0x80 {
				t >>= 7
				s++
			}
			s++

		}

		for k0 := range d.Rewards {

			{
				if d.Rewards[k0] != nil {

					{
						s += (*d.Rewards[k0]).Size()
					}
					s += 0
				}
			}

			s += 1

		}

	}
	s += 8
	return
}
func (d *StagePlayCollect) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{
		l := uint64(len(d.Rewards))

		{

			t := uint64(l)

			for t >= 0x80 {
				buf[i+0] = byte(t) | 0x80
				t >>= 7
				i++
			}
			buf[i+0] = byte(t)
			i++

		}
		for k0 := range d.Rewards {

			{
				if d.Rewards[k0] == nil {
					buf[i+0] = 0
				} else {
					buf[i+0] = 1

					{
						nbuf, err := (*d.Rewards[k0]).Marshal(buf[i+1:])
						if err != nil {
							return nil, err
						}
						i += uint64(len(nbuf))
					}
					i += 0
				}
			}

			i += 1

		}
	}
	{

		buf[i+0+0] = byte(d.Count >> 0)

		buf[i+1+0] = byte(d.Count >> 8)

		buf[i+2+0] = byte(d.Count >> 16)

		buf[i+3+0] = byte(d.Count >> 24)

		buf[i+4+0] = byte(d.Count >> 32)

		buf[i+5+0] = byte(d.Count >> 40)

		buf[i+6+0] = byte(d.Count >> 48)

		buf[i+7+0] = byte(d.Count >> 56)

	}
	return buf[:i+8], nil
}

func (d *StagePlayCollect) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{
		l := uint64(0)

		{

			bs := uint8(7)
			t := uint64(buf[i+0] & 0x7F)
			for buf[i+0]&0x80 == 0x80 {
				i++
				t |= uint64(buf[i+0]&0x7F) << bs
				bs += 7
			}
			i++

			l = t

		}
		if uint64(cap(d.Rewards)) >= l {
			d.Rewards = d.Rewards[:l]
		} else {
			d.Rewards = make([]*Reward, l)
		}
		for k0 := range d.Rewards {

			{
				if buf[i+0] == 1 {
					if d.Rewards[k0] == nil {
						d.Rewards[k0] = new(Reward)
					}

					{
						ni, err := (*d.Rewards[k0]).Unmarshal(buf[i+1:])
						if err != nil {
							return 0, err
						}
						i += ni
					}
					i += 0
				} else {
					d.Rewards[k0] = nil
				}
			}

			i += 1

		}
	}
	{

		d.Count = 0 | (int(buf[i+0+0]) << 0) | (int(buf[i+1+0]) << 8) | (int(buf[i+2+0]) << 16) | (int(buf[i+3+0]) << 24) | (int(buf[i+4+0]) << 32) | (int(buf[i+5+0]) << 40) | (int(buf[i+6+0]) << 48) | (int(buf[i+7+0]) << 56)

	}
	return i + 8, nil
}

type Reward struct {
	ID     int
	Count  int
	DropId int
	Factor float64
}

func (d *Reward) Size() (s uint64) {

	s += 32
	return
}
func (d *Reward) Marshal(buf []byte) ([]byte, error) {
	size := d.Size()
	{
		if uint64(cap(buf)) >= size {
			buf = buf[:size]
		} else {
			buf = make([]byte, size)
		}
	}
	i := uint64(0)

	{

		buf[0+0] = byte(d.ID >> 0)

		buf[1+0] = byte(d.ID >> 8)

		buf[2+0] = byte(d.ID >> 16)

		buf[3+0] = byte(d.ID >> 24)

		buf[4+0] = byte(d.ID >> 32)

		buf[5+0] = byte(d.ID >> 40)

		buf[6+0] = byte(d.ID >> 48)

		buf[7+0] = byte(d.ID >> 56)

	}
	{

		buf[0+8] = byte(d.Count >> 0)

		buf[1+8] = byte(d.Count >> 8)

		buf[2+8] = byte(d.Count >> 16)

		buf[3+8] = byte(d.Count >> 24)

		buf[4+8] = byte(d.Count >> 32)

		buf[5+8] = byte(d.Count >> 40)

		buf[6+8] = byte(d.Count >> 48)

		buf[7+8] = byte(d.Count >> 56)

	}
	{

		buf[0+16] = byte(d.DropId >> 0)

		buf[1+16] = byte(d.DropId >> 8)

		buf[2+16] = byte(d.DropId >> 16)

		buf[3+16] = byte(d.DropId >> 24)

		buf[4+16] = byte(d.DropId >> 32)

		buf[5+16] = byte(d.DropId >> 40)

		buf[6+16] = byte(d.DropId >> 48)

		buf[7+16] = byte(d.DropId >> 56)

	}
	{

		v := *(*uint64)(unsafe.Pointer(&(d.Factor)))

		buf[0+24] = byte(v >> 0)

		buf[1+24] = byte(v >> 8)

		buf[2+24] = byte(v >> 16)

		buf[3+24] = byte(v >> 24)

		buf[4+24] = byte(v >> 32)

		buf[5+24] = byte(v >> 40)

		buf[6+24] = byte(v >> 48)

		buf[7+24] = byte(v >> 56)

	}
	return buf[:i+32], nil
}

func (d *Reward) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = 0 | (int(buf[0+0]) << 0) | (int(buf[1+0]) << 8) | (int(buf[2+0]) << 16) | (int(buf[3+0]) << 24) | (int(buf[4+0]) << 32) | (int(buf[5+0]) << 40) | (int(buf[6+0]) << 48) | (int(buf[7+0]) << 56)

	}
	{

		d.Count = 0 | (int(buf[0+8]) << 0) | (int(buf[1+8]) << 8) | (int(buf[2+8]) << 16) | (int(buf[3+8]) << 24) | (int(buf[4+8]) << 32) | (int(buf[5+8]) << 40) | (int(buf[6+8]) << 48) | (int(buf[7+8]) << 56)

	}
	{

		d.DropId = 0 | (int(buf[0+16]) << 0) | (int(buf[1+16]) << 8) | (int(buf[2+16]) << 16) | (int(buf[3+16]) << 24) | (int(buf[4+16]) << 32) | (int(buf[5+16]) << 40) | (int(buf[6+16]) << 48) | (int(buf[7+16]) << 56)

	}
	{

		v := 0 | (uint64(buf[0+24]) << 0) | (uint64(buf[1+24]) << 8) | (uint64(buf[2+24]) << 16) | (uint64(buf[3+24]) << 24) | (uint64(buf[4+24]) << 32) | (uint64(buf[5+24]) << 40) | (uint64(buf[6+24]) << 48) | (uint64(buf[7+24]) << 56)
		d.Factor = *(*float64)(unsafe.Pointer(&v))

	}
	return i + 32, nil
}
