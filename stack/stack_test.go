package stack

import "testing"

func TestPush(t *testing.T) {
	myStack := Stack{}
	myStack.Push(42)

	if myStack.top.data == 42 {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myStack.Push(43)

	if myStack.top.data == 43 {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}
}

func TestPeek(t *testing.T) {
	myStack := Stack{}

	if data, ok := myStack.Peek(); data == 0 && !ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myStack.Push(42)
	myStack.Push(43)
	myStack.Push(44)
	myStack.Push(45)

	if data, ok := myStack.Peek(); data == 45 && ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}
}

func TestPop(t *testing.T) {
	myStack := Stack{}

	if data, ok := myStack.Pop(); data == 0 && !ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myStack.Push(42)
	myStack.Push(43)
	myStack.Push(44)
	myStack.Pop()
	myStack.Push(45)
	myStack.Pop()

	if data, ok := myStack.Pop(); data == 43 && ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	if data, ok := myStack.Pop(); data == 42 && ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	if data, ok := myStack.Pop(); data == 0 && !ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}
}

func TestEmpty(t *testing.T) {
	myStack := Stack{}

	if myStack.IsEmpty() {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myStack.Push(42)
	if !myStack.IsEmpty() {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myStack.Push(43)
	myStack.Push(44)
	myStack.Pop()
	myStack.Pop()
	myStack.Pop()

	if myStack.IsEmpty() {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}
}
