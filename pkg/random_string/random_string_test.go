package random_string

import "testing"

func TestNew(t *testing.T) {
	random := New(8)
	if len(random) != 8 {
		t.Fatalf("expected 8 characters, got: %d characters", len(random))
	}

	differentRandom := New(8)
	if differentRandom == random {
		t.Fatal("expected different result for differentRandom")
	}
}
