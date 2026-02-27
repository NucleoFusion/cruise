package trivy

import (
	"encoding/json"
	"os/exec"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/types"
)

func ScanImage(image string) (*[]types.Vulnerability, error) {
	report := struct {
		Results []struct {
			Vulns []struct {
				VulnerabilityID  string
				PkgName          string
				InstalledVersion string
				Title            string
				Severity         string
				Description      string
				PrimaryURL       string
				PublishedDate    string
			} `json:"Vulnerabilities"`
		} `json:"Results"`
	}{}

	out, err := exec.Command("trivy", "image", image, "--format", "json").Output()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(out, &report)
	if err != nil {
		return nil, err
	}

	arr := make([]types.Vulnerability, 0)
	for _, res := range report.Results {
		for _, v := range res.Vulns {
			vuln := types.Vulnerability{
				VulnID:     v.VulnerabilityID,
				Title:      v.Title,
				Severity:   utils.GetSeverity("trivy", v.Severity),
				Pkg:        v.PkgName + " " + v.InstalledVersion,
				Published:  v.PublishedDate,
				PrimaryURL: v.PrimaryURL,
			}
			arr = append(arr, vuln)
		}
	}

	return &arr, nil
}
