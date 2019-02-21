package render

import "testing"

func TestInsertsIntoQueueCorrectly(t *testing.T) {
	l := NewLayer()

	l.Queue(DrawRequest{
		Sort: 2,
	})

	l.Queue(DrawRequest{
		Sort: 6,
	})

	l.Queue(DrawRequest{
		Sort: 4,
	})

	if len(l.queue) != 3 {
		t.Fatalf("Expected 3 to be in queue, but have %d", len(l.queue))
	}

	if l.queue[0].Sort != 2 {
		t.Errorf("Expected sort 2 to be first, but got sort %f", l.queue[0].Sort)
	}

	if l.queue[1].Sort != 4 {
		t.Errorf("Expected sort 2 to be first, but got sort %f", l.queue[1].Sort)
	}

	if l.queue[2].Sort != 6 {
		t.Errorf("Expected sort 6 to be last, but got sort %f", l.queue[2].Sort)
	}
}
