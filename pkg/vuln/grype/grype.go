package grype

import (
	"encoding/json"
	"os/exec"

	"github.com/cruise-org/cruise/internal/utils"
	"github.com/cruise-org/cruise/pkg/types"
)

func ScanImage(image string) (*[]types.Vulnerability, error) {
	report := struct {
		Matches []struct {
			Vulnerability struct {
				ID          string   `json:"id"`
				Description string   `json:"description"`
				Severity    string   `json:"severity"`
				URLs        []string `json:"urls"`
			} `json:"vulnerability"`
			Artifact struct {
				Name    string `json:"name"`
				Version string `json:"version"`
			} `json:"artifact"`
		} `json:"matches"`
	}{}

	out, err := exec.Command("grype", image, "-o", "json").Output()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(out, &report)
	if err != nil {
		return nil, err
	}

	arr := make([]types.Vulnerability, 0)
	for _, v := range report.Matches {
		url := "NA"
		if len(v.Vulnerability.URLs) > 0 {
			url = v.Vulnerability.URLs[0]
		}
		vuln := types.Vulnerability{
			VulnID:     v.Vulnerability.ID,
			Title:      v.Vulnerability.Description,
			Severity:   utils.GetSeverity("grype", v.Vulnerability.Severity),
			Pkg:        v.Artifact.Name + " " + v.Artifact.Version,
			Published:  "NA",
			PrimaryURL: url,
		}
		arr = append(arr, vuln)
	}

	return &arr, nil
}
