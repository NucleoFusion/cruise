package docker

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type DaemonInfo struct {
	Version string
	OS      string
	Uptime  string
}

func GetDaemonInfo() (*DaemonInfo, error) {
	version, err := cli.ServerVersion(context.Background())
	if err != nil {
		return &DaemonInfo{}, err
	}

	uptimeStart, err := GetDockerDaemonStartTime()
	if err != nil {
		return &DaemonInfo{}, err
	}

	uptime := formatUptime(uptimeStart)

	osname, err := GetOSName()
	if err != nil {
		return &DaemonInfo{}, err
	}

	return &DaemonInfo{
		Version: version.Version,
		OS:      osname,
		Uptime:  uptime,
	}, nil
}

func formatUptime(startedAt time.Time) string {
	uptime := time.Since(startedAt)

	days := int(uptime.Hours()) / 24
	hours := int(uptime.Hours()) % 24
	minutes := int(uptime.Minutes()) % 60

	return fmt.Sprintf("%dd %dh %dm", days, hours, minutes)
}

func GetDockerDaemonStartTime() (time.Time, error) {
	// Run systemctl show to get ExecMainStartTimestamp property
	out, err := exec.Command("systemctl", "show", "-p", "ExecMainStartTimestamp", "docker.service").Output()
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to run systemctl: %w", err)
	}
	line := strings.TrimSpace(string(out))
	// Expected output format: ExecMainStartTimestamp=Mon 2025-07-21 07:03:37 UTC
	parts := strings.SplitN(line, "=", 2)
	if len(parts) != 2 {
		return time.Time{}, fmt.Errorf("unexpected output: %s", line)
	}
	timestampStr := parts[1]

	// Parse the timestamp. Format is "Mon 2006-01-02 15:04:05 MST"
	parsedTime, err := time.Parse("Mon 2006-01-02 15:04:05 MST", timestampStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %w", err)
	}

	return parsedTime, nil
}

func GetOSName() (string, error) {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return "", fmt.Errorf("failed to open /etc/os-release: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments and unrelated lines
		if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		key := parts[0]
		value := parts[1]

		if key == "PRETTY_NAME" {
			// Remove surrounding quotes if any
			value = strings.Trim(value, `"`)
			return value, nil
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error reading /etc/os-release: %w", err)
	}
	return "", fmt.Errorf("PRETTY_NAME not found in /etc/os-release")
}
