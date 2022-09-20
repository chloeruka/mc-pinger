package mcpinger

import (
	"fmt"
	"testing"
)

func TestReachMcRukaPet(t *testing.T) {
	// Create new Pinger
	pinger := New("mc.ruka.pet", 25565)

	// Get server info
	info, err := pinger.Ping()

	if err != nil {
		t.Error(err)
		return
	}

	// Print server info
	fmt.Printf("Description: \"%s\"\n", info.Description.Text)
	fmt.Printf("Online: %d/%d\n", info.Players.Online, info.Players.Max)
	fmt.Printf("Version: %s\n", info.Version.Name)
}