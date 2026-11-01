package servers

import "testing"

func TestCleanSudoCmd(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"sudo apt update", "apt update"},
		{"sudo  systemctl restart nginx", "systemctl restart nginx"},
		{"sudo\tsystemctl restart nginx", "systemctl restart nginx"},
		{"sudo sudo cat /etc/shadow", "cat /etc/shadow"},
		{"uptime", "uptime"},
		{"  sudo docker ps  ", "docker ps"},
		{"sudo -s", "-s"},
	}

	for _, tt := range tests {
		got := CleanSudoCmd(tt.input)
		if got != tt.expected {
			t.Errorf("CleanSudoCmd(%q) = %q; expected %q", tt.input, got, tt.expected)
		}
	}
}
