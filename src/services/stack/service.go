package stack

import "errors"

var (
	ErrUninitialized = errors.New("stack is not initialized")
	ErrEmpty         = errors.New("stack is empty")
	ErrFull          = errors.New("stack is full")
)

type StackService interface {
	Initialize(capacity int)
	Push(value string) error
	Pop() (string, error)
	Peek() (string, error)
	Size() (int, error)
	IsEmpty() (bool, error)
}

type stack struct {
	elements []string
	top      int
	capacity int
}

func NewStackService() StackService {
	return &stack{}
}

func (s *stack) Initialize(capacity int) {
	if capacity <= 0 {
		capacity = 10
	}

	s.elements = make([]string, capacity)
	s.top = -1
	s.capacity = capacity
}

func (s *stack) Push(value string) error {

	if err := s.validateForPush(); err != nil {
		return err
	}

	s.top++
	s.elements[s.top] = value
	return nil
}

func (s *stack) Pop() (string, error) {

	if err := s.validateForPopOrPeek(); err != nil {
		return "", err
	}

	value := s.elements[s.top]
	s.elements[s.top] = ""
	s.top--

	return value, nil
}

func (s *stack) Peek() (string, error) {
	if err := s.validateForPopOrPeek(); err != nil {
		return "", err
	}

	return s.elements[s.top], nil
}

func (s *stack) Size() (int, error) {
	if err := s.validateInitialized(); err != nil {
		return -1, err
	}
	return s.top + 1, nil
}

func (s *stack) IsEmpty() (bool, error) {
	if err := s.validateInitialized(); err != nil {
		return true, err
	}

	return s.top == -1, nil
}

func (s *stack) validateInitialized() error {
	if s.elements == nil || s.capacity == 0 {
		return ErrUninitialized
	}

	return nil
}

func (s *stack) validateForPush() error {
	if err := s.validateInitialized(); err != nil {
		return err
	}

	if s.top >= s.capacity-1 {
		return ErrFull
	}
	return nil
}

func (s *stack) validateForPopOrPeek() error {
	if err := s.validateInitialized(); err != nil {
		return err
	}

	if s.top == -1 {
		return ErrEmpty
	}
	return nil
}
