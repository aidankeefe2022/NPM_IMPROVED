package unifyPackageLock

import "maps"

type NpmReport struct {
	Vulnerabilities map[string]any `json:"vulnerabilities"`
	MetaData        map[string]any `json:"metadata"`
}

// note will not include duplicates
func (npmReport *NpmReport) Combine(toAdd *NpmReport) {
	maps.Copy(npmReport.Vulnerabilities, toAdd.Vulnerabilities)
	maps.Copy(npmReport.MetaData, toAdd.MetaData)
}
