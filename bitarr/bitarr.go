package bitarr

import (
	"bytes"
	"fmt"
)

// 一个bit数组通常会用一个无符号数或者称之为“字”的slice来表示，每一个元素的每一位都表示集合里的一个值。
// 当集合的第i位被设置时，我们才说这个集合包含元素i。
// 对于words，每个元素可存储的值有bitNum个，每超过bitNum个则进位，即添加一个元素。（注意，0也占了一位，所以bitNum才要进位，第一个元素可存储0-63）。
// 所以，对于words中的一个元素，要转换为具体的值时：首先取到其位置i，用 bitNum * i 作为已进位数（类似于每10位要进位）， 然后将这个元素转换
// 为二进制数，从右往左数，第多少位为1则表示相应的有这个值，用这个位数 x+bitNum *i 即为我们存入的值。
type bitArr struct {
	words []uint
}

const (
	// 根据平台自动判断决定是32还是bitNum
	bitNum = 32 << (^uint(0) >> 63)
)

func New() *bitArr {
	return &bitArr{}
}

func (b bitArr) Has(x int) bool {
	word, bit := x/bitNum, uint(x%bitNum)
	return word < len(b.words) && b.words[word]&(1<<bit) != 0

}

func (b *bitArr) Add(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	for word >= len(b.words) {
		b.words = append(b.words, 0)
	}
	b.words[word] |= 1 << bit
}

func (b *bitArr) AddAll(args ...int) {
	for _, x := range args {
		b.Add(x)
	}
}

func (b *bitArr) Len() int {
	var len int
	for _, word := range b.words {
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

func (b *bitArr) Remove(x int) {
	word, bit := x/bitNum, uint(x%bitNum)
	if b.Has(x) {
		b.words[word] ^= 1 << bit
	}
}

func (b *bitArr) Clear() {
	b.words = append([]uint{})
}

func (b *bitArr) Copy() *bitArr {
	ba := &bitArr{words: []uint{}}
	for _, value := range b.words {
		ba.words = append(ba.words, value)
	}
	return ba
}

// A与B的并集
func (b *bitArr) UnionWith(t *bitArr) {
	for i, tWord := range t.words {
		if i < len(b.words) {
			b.words[i] |= tWord
		} else {
			b.words = append(b.words, tWord)
		}
	}
}

// A与B的交集
func (b *bitArr) IntersectWith(t *bitArr) {
	for i, tWord := range t.words {
		if i > len(b.words) {
			continue
		}
		b.words[i] &= tWord
	}
}

// A与B的差集
func (b *bitArr) DiffWith(t *bitArr) {
	t1 := t.Copy()
	t1.IntersectWith(b)
	for i, tWord := range t1.words {
		if i < len(b.words) {
			b.words[i] ^= tWord
		}
	}
}

// A与B的并差集，元素出现在A没有出现在B，或出现在B没有出现在A
func (b *bitArr) SymmetricDiff(t *bitArr) {
	for i, tWord := range t.words {
		if i < len(b.words) {
			b.words[i] ^= tWord
		} else {
			b.words = append(b.words, tWord)
		}
	}
}

// 获取比特数组中的所有元素的slice集合
func (b *bitArr) Elems() []int {
	var elems []int
	for i, word := range b.words {
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, bitNum*i+j)
			}
		}
	}
	return elems
}

func (b bitArr) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range b.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitNum; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitNum*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
