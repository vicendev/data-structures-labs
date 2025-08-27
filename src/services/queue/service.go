package queue

import "errors"

var (
	ErrUninitialized = errors.New("queue is not initialized")
	ErrEmpty         = errors.New("queue is empty")
	ErrFull          = errors.New("queue is full")
)

type QueueService interface {
	Initialize(capacity int)
	Enqueue(value string) error
	Dequeue() (string, error)
	Tail() (string, error)
	Head() (string, error)
	Size() (int, error)
	IsEmpty() (bool, error)
}

type queue struct {
	elements []string
	head     int
	tail     int
	size     int
	capacity int
}

func NewQueueService() QueueService {
	return &queue{}
}

func (q *queue) Initialize(capacity int) {
	if capacity <= 0 {
		capacity = 10
	}

	q.elements = make([]string, capacity)
	q.head = -1
	q.tail = -1
	q.size = 0
	q.capacity = capacity
}

func (q *queue) Enqueue(value string) error {
	if err := q.validateFullQueue(); err != nil {
		return err
	}

	if q.tail == q.capacity-1 {
		q.tail = -1
	}
	q.size++
	q.tail++

	q.elements[q.tail] = value

	println(q.tail, q.head)
	for i := 0; i < len(q.elements); i++ {
		println(q.elements[i])
	}
	return nil
}

func (q *queue) Dequeue() (string, error) {
	if err := q.validateEmptyQueue(); err != nil {
		return "", err
	}

	if q.head == q.capacity-1 {
		q.head = -1
	}

	q.size--
	q.head++

	value := q.elements[q.head]
	q.elements[q.head] = ""

	println(q.tail, q.head)
	for i := 0; i < len(q.elements); i++ {
		println(q.elements[i])
	}
	return value, nil
}

func (q *queue) Tail() (string, error) {
	if err := q.validateEmptyQueue(); err != nil {
		return "", err
	}

	return q.elements[q.tail], nil
}

func (q *queue) Head() (string, error) {
	if err := q.validateEmptyQueue(); err != nil {
		return "", err
	}

	return q.elements[q.head+1], nil
}

func (q *queue) Size() (int, error) {
	if err := q.validateInitialized(); err != nil {
		return 0, err
	}

	return q.size, nil
}

func (q *queue) IsEmpty() (bool, error) {
	if err := q.validateEmptyQueue(); err != nil {
		return true, err
	}

	return false, nil
}

func (q *queue) validateInitialized() error {
	if q.elements == nil || q.capacity == 0 {
		return ErrUninitialized
	}

	return nil
}

func (q *queue) validateFullQueue() error {
	if err := q.validateInitialized(); err != nil {
		return err
	}

	if q.size >= q.capacity {
		return ErrFull
	}
	return nil
}

func (q *queue) validateEmptyQueue() error {
	if err := q.validateInitialized(); err != nil {
		return err
	}

	if q.size == 0 {
		return ErrEmpty
	}

	return nil
}
