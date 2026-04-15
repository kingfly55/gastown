package cmd

import (
	"strings"
	"testing"
)

func TestSpawnPolecatForSling_RequiresHookBead(t *testing.T) {
	t.Parallel()

	_, err := SpawnPolecatForSling("gastown", SlingSpawnOptions{})
	if err == nil {
		t.Fatal("expected error for empty hook bead")
	}
	if !strings.Contains(err.Error(), "hook bead is required for polecat spawn") {
		t.Fatalf("unexpected error: %v", err)
	}
}
