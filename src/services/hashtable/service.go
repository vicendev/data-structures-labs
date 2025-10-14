package hashtable

import (
	"errors"
	"fmt"
	"hash/fnv"
	"hash/maphash"
	"math/rand"
	"strconv"
	"time"
)

var (
	ErrEmpty          = errors.New("hashtable is empty")
	ErrNotFound       = errors.New("hastable key not found")
	ErrHashNotSupport = errors.New("hashtable hash fn is not supported")
)

type HashFnType string

const (
	HashFnFNV1a32 HashFnType = "fnv1a32"
	HashFnMaphash HashFnType = "maphash"
	HashFnBasic   HashFnType = "basic"
)

const MaxLoad float64 = 0.75

type HashTableService interface {
	// Insertion Methods
	Upsert(key string, value string, hashFnType HashFnType) (oldValue string, replaced bool, err error)
	Seed() int

	// Accesibility Methods
	Get(key string) (valueFound string, err error)

	// Deletion Methods
	Delete(key string) (deletedValue string, err error)

	// Utility Methods
	Clear()
	Reset()
}

type Pair struct {
	Key   string
	Value string
}

type Bucket struct {
	Pairs []Pair
}

type hashtable struct {
	buckets    []Bucket
	size       int
	capacity   int
	hashFn     func(string) uint64
	hashFnType HashFnType
	mhSeed     maphash.Seed
}

func NewHashTable() HashTableService {
	hashtable := &hashtable{
		size:     0,
		capacity: 8,
	}

	hashtable.buckets = make([]Bucket, hashtable.capacity)

	return hashtable
}

// Delete implements HashTableService.
func (h *hashtable) Delete(key string) (deletedValue string, err error) {
	if err := h.validateEmpty(); err != nil {
		return "", err
	}

	hash := h.hashFn(key)
	idx := int(hash % uint64(len(h.buckets)))
	pairs := &h.buckets[idx].Pairs

	for i := range *pairs {
		if (*pairs)[i].Key == key {
			deletedValue = (*pairs)[i].Value

			lastPair := len(*pairs) - 1
			(*pairs)[i] = (*pairs)[lastPair]
			(*pairs)[lastPair] = Pair{}

			*pairs = (*pairs)[:lastPair]

			h.size--

			return deletedValue, nil
		}
	}

	return "", ErrNotFound
}

// Get implements HashTableService.
func (h *hashtable) Get(key string) (valueFound string, err error) {
	if err := h.validateEmpty(); err != nil {
		return "", err
	}

	hash := h.hashFn(key)
	idx := int(hash % uint64(len(h.buckets)))
	pairs := &h.buckets[idx].Pairs

	for i := range *pairs {
		if (*pairs)[i].Key == key {
			return (*pairs)[i].Value, nil
		}
	}

	return "", ErrNotFound
}

// Insert implements HashTableService.
func (h *hashtable) Upsert(key string, value string, hashFnType HashFnType) (oldValue string, replaced bool, err error) {
	if h.hashFnType == "" {
		if err := h.setHashFn(hashFnType); err != nil {
			return "", false, err
		}
	}

	loadFactor := h.loadFactor()

	if loadFactor > MaxLoad {
		h.rehash()
	}

	hash := h.hashFn(key)

	idx := int(hash % uint64(len(h.buckets)))
	pairs := &h.buckets[idx].Pairs

	for i := range *pairs {
		if (*pairs)[i].Key == key {
			oldValue = (*pairs)[i].Value
			(*pairs)[i].Value = value

			return oldValue, true, nil
		}
	}

	*pairs = append(*pairs, Pair{
		Key:   key,
		Value: value,
	})

	h.size++
	h.Print()
	return "", false, nil
}

func (h *hashtable) Seed() (size int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	n := 500
	max := 2000

	nums := make([]int, max)

	for i := 0; i < max; i++ {
		nums[i] = i + 1
	}

	r.Shuffle(len(nums), func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	})

	result := nums[:n]

	for i := 0; i < len(result); i++ {
		resultString := strconv.Itoa(result[i])
		h.Upsert("key_"+resultString, "Data "+resultString, HashFnBasic)
	}

	return h.size
}

func (h *hashtable) Reset() {
	h.buckets = nil
	h.size = 0
	h.capacity = 8
	h.hashFn = nil
	h.hashFnType = ""

	h.buckets = make([]Bucket, h.capacity)
}

func (h *hashtable) Clear() {
	for i := range h.buckets {
		h.buckets[i].Pairs = nil
	}
	h.size = 0
}

/* Private Methods */

func (h *hashtable) loadFactor() float64 {
	if len(h.buckets) == 0 {
		return 0
	}

	return float64(h.size) / float64(len(h.buckets))
}

func (h *hashtable) rehash() {
	oldBuckets := h.buckets
	oldCapaticy := len(h.buckets)

	h.capacity = oldCapaticy * 2
	h.buckets = make([]Bucket, h.capacity)
	h.size = 0

	for _, bucket := range oldBuckets {
		for _, pair := range bucket.Pairs {
			h.insertNoResize(pair.Key, pair.Value)
		}
	}
}

func (h *hashtable) insertNoResize(key string, value string) {
	hash := h.hashFn(key)

	idx := int(hash % uint64(len(h.buckets)))
	h.buckets[idx].Pairs = append(h.buckets[idx].Pairs, Pair{Key: key, Value: value})
	h.size++
}

func (h *hashtable) basicHashFn(key string) uint64 {
	hash := 0

	for i := 0; i < len(key); i++ {
		hash = (31 * hash) + int(key[i])
	}

	if hash < 0 {
		hash = -hash
	}

	return uint64(hash)
}

//lint:ignore U1000 used indirectly via function pointer
func (h *hashtable) fnv1a32HashFn(key string) uint64 {
	hash := fnv.New32a()
	hash.Write([]byte(key))

	return uint64(hash.Sum32())
}

//lint:ignore U1000 used indirectly via function pointer
func (h *hashtable) mapHashFn(key string) uint64 {
	var hash maphash.Hash

	hash.SetSeed(h.mhSeed)
	hash.WriteString(key)

	return hash.Sum64()
}

func (h *hashtable) setHashFn(hashFnType HashFnType) error {
	switch hashFnType {
	case HashFnFNV1a32:
		h.hashFn = h.fnv1a32HashFn
	case HashFnMaphash:
		h.mhSeed = maphash.MakeSeed()
		h.hashFn = h.mapHashFn
	case HashFnBasic:
		h.hashFn = h.basicHashFn
	default:
		return ErrHashNotSupport
	}

	h.hashFnType = hashFnType
	return nil
}

/* Utility Methods */

func (h *hashtable) Print() {
	fmt.Println("========== HashTable ==========")
	for i, bucket := range h.buckets {
		if len(bucket.Pairs) == 0 {
			fmt.Printf("[%d] -> (empty)\n", i)
			continue
		}

		fmt.Printf("[%d]: ", i)
		for j, pair := range bucket.Pairs {
			fmt.Printf("{%s: %s}", pair.Key, pair.Value)
			if j < len(bucket.Pairs)-1 {
				fmt.Print(" -> ")
			}
		}
		fmt.Println()
	}
	//fmt.Printf("Total elements: %d | Capacity: %d | Load factor: %.2f\n",
	//h.size, h.capacity, h.LoadFactor())
	fmt.Printf("Total elements: %d | Capacity: %d \n", h.size, h.capacity)
	fmt.Println("================================")
}

/* Validations */
func (h *hashtable) validateEmpty() error {
	if h.size == 0 {
		return ErrEmpty
	}

	return nil
}
