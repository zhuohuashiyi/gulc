package gulc

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"io"
)


const (
	N = 1000
)

var hashFuncs = []func([]byte) []byte{_md5, _sha1, _sha256, _sha512}

type Hashable interface {
	Hash() []byte
}

type Filter[T Hashable] interface {
	Add(elem T)                  // 增加一个元素
	Search(elem T) bool          // 查找一个元素是否存在
	Delete(elem T) (bool, error) // 删除一个元素，成功则返回true
	List() ([]T, error)          // 列出当前所有剩余元素
}

type BloomFilter[T Hashable] struct {
	bitMap []uint8
	size   uint64
}

func NewBloomFilter[T Hashable](size uint64) Filter[T] {
	count := size >> 3
	if (size & 7) != 0 {
		count++
	}
	return &BloomFilter[T]{size: count << 3}
}

func (f *BloomFilter[T]) Add(elem T) {
	if f.bitMap == nil {
		f.bitMap = make([]uint8, f.size>>3)
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		f.bitMap[(pos >> 3)] |= (1 << (pos & 7))
	}
}

func (f *BloomFilter[T]) Search(elem T) bool {
	if f.bitMap == nil {
		return false
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		if (f.bitMap[(pos>>3)] & (1 << (pos & 7))) == 0 {
			return false
		}
	}
	return true
}

func (f *BloomFilter[T]) Delete(elem T) (bool, error) {
	return false, ErrFunctionUnSupported
}

func (f *BloomFilter[T]) List() ([]T, error) {
	return nil, ErrFunctionUnSupported
}

type CountBloomFilter[T Hashable] struct {
	bitMap []uint64
	size   uint64
}

func NewCountBloomFilter[T Hashable](size uint64) Filter[T] {
	return &CountBloomFilter[T]{size: size}
}

func (f *CountBloomFilter[T]) Add(elem T) {
	if f.bitMap == nil {
		f.bitMap = make([]uint64, f.size)
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		f.bitMap[pos]++
	}
}

func (f *CountBloomFilter[T]) Search(elem T) bool {
	if f.bitMap == nil {
		return false
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		if (f.bitMap[pos]) == 0 {
			return false
		}
	}
	return true
}

func (f *CountBloomFilter[T]) Delete(elem T) (bool, error) {
	if !f.Search(elem) {
		return false, nil
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		f.bitMap[pos]--
	}
	return true, nil
}

func (f *CountBloomFilter[T]) List() ([]T, error) {
	return nil, ErrFunctionUnSupported
}


type filterCell struct {
	count uint64
	idSum uint64
	hashSum uint64
}

type InvertibleBloomFilter struct {
	tabB []filterCell
	tabC []filterCell
	size uint64
	count uint64
}

func NewInvertibleBloomFilter(size uint64) Filter[Integer] {
	return &InvertibleBloomFilter{size: size}
}


func (f *InvertibleBloomFilter) Add(elem Integer) {
	if f.tabB == nil {
		f.tabB = make([]filterCell, f.size)
		f.tabC = make([]filterCell, f.size)
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		f.tabB[pos].count++
		f.tabB[pos].idSum += uint64(elem.Depacket())
		f.tabB[pos].hashSum += _g(elem.Depacket(), N)
	}
	poses := make([]uint, 2)
	poses[0] = _f1(elem.Depacket(), uint(f.size))
	poses[1] = _f2(elem.Depacket(), uint(f.size))
	for _, pos := range poses {
		f.tabC[pos].count++
		f.tabC[pos].idSum += uint64(elem.Depacket())
		f.tabC[pos].hashSum += _g(elem.Depacket(), N)
	}
	f.count++
}

func (f *InvertibleBloomFilter) Search(elem Integer) bool {
	if f.tabB == nil {
		return false
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		if (f.tabB[pos].count) == 0 {
			return false
		}
	}
	return true
}

func (f *InvertibleBloomFilter) Delete(elem Integer) (bool, error) {
	if !f.Search(elem) {
		return false, nil
	}
	elemHash := elem.Hash()
	for _, hash := range hashFuncs {
		hashBytes := hash(elemHash)
		hashVal := LE2Int64(hashBytes[:8])
		pos := hashVal % f.size
		f.tabB[pos].count--
		f.tabB[pos].idSum -= uint64(elem.Depacket())
		f.tabB[pos].hashSum -= _g(elem.Depacket(), N)
	}
	poses := make([]uint, 2)
	poses[0] = _f1(elem.Depacket(), uint(f.size))
	poses[1] = _f2(elem.Depacket(), uint(f.size))
	for _, pos := range poses {
		f.tabC[pos].count--
		f.tabC[pos].idSum -= uint64(elem.Depacket())
		f.tabC[pos].hashSum -= _g(elem.Depacket(), N)
	}
	f.count--
	return true, nil
}

func (f *InvertibleBloomFilter) List() ([]Integer, error) {
	res := make([]Integer, 0)
	var loopAndFind func(tab []filterCell) = func(tab []filterCell) {
		for {
			flag := false
			for i := uint64(0); i < f.size; i++ {
				if tab[i].count != 0 && _g(int(tab[i].idSum / tab[i].count), N) == tab[i].hashSum / tab[i].count {
					x := IntegerPacket(int(tab[i].idSum / tab[i].count))
					for j := uint64(0); j < tab[i].count; i++ {
						f.Delete(x)
					}
					flag = true
					res = append(res, x)
					break
				}
			}
			if flag {
				break
			}
		}
	}
	loopAndFind(f.tabB)
	if f.count != 0 {
		loopAndFind(f.tabC)
	}
	for _, item := range res {
		f.Add(item)
	}
	return res, nil
}

// hash工具函数
func _sha1(data []byte) []byte {
	hash := sha1.New()
	io.WriteString(hash, string(data))
	return hash.Sum(nil)
}

func _sha256(data []byte) []byte {
	hash := sha256.New()
	io.WriteString(hash, string(data))
	return hash.Sum(nil)
}

func _sha512(data []byte) []byte {
	hash := sha512.New()
	io.WriteString(hash, string(data))
	return hash.Sum(nil)
}

func _md5(data []byte) []byte {
	hash := md5.New()
	io.WriteString(hash, string(data))
	return hash.Sum(nil)
}

func _g(num int, n uint) uint64 {
	data := IntegerPacket(num).Hash()
	data = _sha256(data)
	var m uint64 = uint64(n) * uint64(n)
	k := LE2Int64(data[: 8])
	return k % m
}

func _f1(num int, m uint) uint {
	data := IntegerPacket(num).Hash()
	data = _sha256(data)
	k := LE2Int64(data[: 8])
	return uint(k % uint64(m))
}

func _f2(num int, m uint) uint {
	data := IntegerPacket(num).Hash()
	data = _sha512(data)
	k := LE2Int64(data[: 8])
	return uint(k % uint64(m))
}
