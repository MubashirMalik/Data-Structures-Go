package queue

import "testing"

func TestEnqueue(t *testing.T) {
	myQueue := Queue{}

	myQueue.Enqueue(42)
	if myQueue.front.data == 42 && myQueue.rear.data == 42 {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myQueue.Enqueue(43)
	if myQueue.front.data == 42 && myQueue.rear.data == 43 {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myQueue.Enqueue(44)
	if myQueue.rear.data == 44 && myQueue.front.next.next.data == 44 {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}
}

func TestDequeue(t *testing.T) {
	myQueue := Queue{}

	if data, ok := myQueue.Dequeue(); data == 0 && !ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myQueue.Enqueue(42)
	myQueue.Enqueue(43)
	myQueue.Enqueue(44)
	myQueue.Dequeue()
	myQueue.Enqueue(45)
	myQueue.Dequeue()

	if data, ok := myQueue.Dequeue(); data == 44 && ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	if data, ok := myQueue.Dequeue(); data == 45 && ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	if data, ok := myQueue.Dequeue(); data == 0 && !ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}
}

func TestPeek(t *testing.T) {
	myQueue := Queue{}

	if data, ok := myQueue.Peek(); data == 0 && !ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	myQueue.Enqueue(42)
	myQueue.Enqueue(43)
	myQueue.Enqueue(44)
	myQueue.Dequeue()
	myQueue.Enqueue(45)
	myQueue.Dequeue()

	if data, ok := myQueue.Peek(); data == 44 && ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}

	if data, ok := myQueue.Peek(); data == 44 && ok {
		t.Logf("PASSED.")
	} else {
		t.Errorf("FAILED.")
	}
}