package vuln

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/config"
	"github.com/cruise-org/cruise/pkg/types"
	"github.com/cruise-org/cruise/pkg/vuln/grype"
	"github.com/cruise-org/cruise/pkg/vuln/trivy"
)

func ScanImage(image, provider string) (*[]types.Vulnerability, error) {
	switch provider {
	case "grype":
		return grype.ScanImage(image)
	case "trivy":
		return trivy.ScanImage(image)
	default:
		return nil, fmt.Errorf("invalid provider name, provider not supported")
	}
}

func ScannerAvailable(scanner string) bool {
	_, err := exec.Command("bash", "-c", scanner+" --version").CombinedOutput() // TODO: Config file based checker
	return err == nil
}

func VulnHeaders(totalWidth int) string {
	w := totalWidth / 15
	return fmt.Sprintf(
		"%-*s %-*s %-*s %-*s %-*s %-*s",
		2*w, "ID",
		2*w, "Pkg",
		totalWidth-15*w, "Severity",
		5*w, "Title",
		2*w, "Date",
		3*w, "URL")
}

func Format(totalWidth int, vuln *types.Vulnerability) string {
	w := totalWidth / 15
	return fmt.Sprintf(
		"%-*s %-*s %-*s %-*s %-*s %-*s",
		2*w, utils.Shorten(vuln.VulnID, 2*w),
		2*w, utils.Shorten(vuln.Pkg, 2*w),
		totalWidth-15*w, utils.SeverityText(vuln.Severity),
		5*w, utils.Shorten(vuln.Title, 5*w),
		2*w, utils.Shorten(vuln.Published, 2*w),
		3*w, utils.Shorten(vuln.PrimaryURL, 3*w))
}

func Export(content []string, page string) error {
	filename := fmt.Sprintf("%d:%d_%d-%d_%s", time.Now().Hour(), time.Now().Minute(), time.Now().Day(), time.Now().Month(), page)

	f, err := os.Create(filepath.Join(config.Cfg.Global.ExportDir, filename))
	if err != nil {
		return err
	}

	f.WriteString(strings.Join(content, "\n"))

	return nil
}
