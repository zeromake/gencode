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

		*(*int)(unsafe.Pointer(&buf[0])) = d.Id

	}
	{

		*(*int)(unsafe.Pointer(&buf[8])) = d.Stage

	}
	{

		*(*int)(unsafe.Pointer(&buf[16])) = d.Type

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

				*(*int)(unsafe.Pointer(&buf[i+25])) = d.Desks[k0]

			}

			i += 8

		}
	}
	return buf[:i+25], nil
}

func (d *LuckScriptIndex) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.Id = *(*int)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.Stage = *(*int)(unsafe.Pointer(&buf[i+8]))

	}
	{

		d.Type = *(*int)(unsafe.Pointer(&buf[i+16]))

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

				d.Desks[k0] = *(*int)(unsafe.Pointer(&buf[i+25]))

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

		*(*int)(unsafe.Pointer(&buf[0])) = d.Id

	}
	{

		*(*int)(unsafe.Pointer(&buf[8])) = d.X

	}
	{

		*(*int)(unsafe.Pointer(&buf[16])) = d.Y

	}
	{

		*(*int)(unsafe.Pointer(&buf[24])) = d.W

	}
	{

		*(*int)(unsafe.Pointer(&buf[32])) = d.H

	}
	{

		*(*int)(unsafe.Pointer(&buf[40])) = d.Points

	}
	{

		*(*int)(unsafe.Pointer(&buf[48])) = d.Suit

	}
	{

		*(*int)(unsafe.Pointer(&buf[56])) = d.Rotate

	}
	{

		*(*int)(unsafe.Pointer(&buf[64])) = d.Layer

	}
	{

		*(*int)(unsafe.Pointer(&buf[72])) = d.Tag

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

				*(*int)(unsafe.Pointer(&buf[i+80])) = d.Deps[k0]

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

				*(*int)(unsafe.Pointer(&buf[i+80])) = d.Collect[k0]

			}

			i += 8

		}
	}
	{

		*(*int)(unsafe.Pointer(&buf[i+80])) = d.Count

	}
	{

		*(*int)(unsafe.Pointer(&buf[i+88])) = d.Direction

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

		d.Id = *(*int)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.X = *(*int)(unsafe.Pointer(&buf[i+8]))

	}
	{

		d.Y = *(*int)(unsafe.Pointer(&buf[i+16]))

	}
	{

		d.W = *(*int)(unsafe.Pointer(&buf[i+24]))

	}
	{

		d.H = *(*int)(unsafe.Pointer(&buf[i+32]))

	}
	{

		d.Points = *(*int)(unsafe.Pointer(&buf[i+40]))

	}
	{

		d.Suit = *(*int)(unsafe.Pointer(&buf[i+48]))

	}
	{

		d.Rotate = *(*int)(unsafe.Pointer(&buf[i+56]))

	}
	{

		d.Layer = *(*int)(unsafe.Pointer(&buf[i+64]))

	}
	{

		d.Tag = *(*int)(unsafe.Pointer(&buf[i+72]))

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

				d.Deps[k0] = *(*int)(unsafe.Pointer(&buf[i+80]))

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

				d.Collect[k0] = *(*int)(unsafe.Pointer(&buf[i+80]))

			}

			i += 8

		}
	}
	{

		d.Count = *(*int)(unsafe.Pointer(&buf[i+80]))

	}
	{

		d.Direction = *(*int)(unsafe.Pointer(&buf[i+88]))

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

		*(*int)(unsafe.Pointer(&buf[0])) = d.Type

	}
	{

		*(*int)(unsafe.Pointer(&buf[8])) = d.Count

	}
	{

		*(*int)(unsafe.Pointer(&buf[16])) = d.Rate

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

		d.Type = *(*int)(unsafe.Pointer(&buf[i+0]))

	}
	{

		d.Count = *(*int)(unsafe.Pointer(&buf[i+8]))

	}
	{

		d.Rate = *(*int)(unsafe.Pointer(&buf[i+16]))

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

		*(*int)(unsafe.Pointer(&buf[i+0])) = d.Count

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

		d.Count = *(*int)(unsafe.Pointer(&buf[i+0]))

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

		*(*int)(unsafe.Pointer(&buf[0])) = d.ID

	}
	{

		*(*int)(unsafe.Pointer(&buf[8])) = d.Count

	}
	{

		*(*int)(unsafe.Pointer(&buf[16])) = d.DropId

	}
	{

		*(*float64)(unsafe.Pointer(&buf[24])) = d.Factor

	}
	return buf[:i+32], nil
}

func (d *Reward) Unmarshal(buf []byte) (uint64, error) {
	i := uint64(0)

	{

		d.ID = *(*int)(unsafe.Pointer(&buf[0]))

	}
	{

		d.Count = *(*int)(unsafe.Pointer(&buf[8]))

	}
	{

		d.DropId = *(*int)(unsafe.Pointer(&buf[16]))

	}
	{

		d.Factor = *(*float64)(unsafe.Pointer(&buf[24]))

	}
	return i + 32, nil
}
