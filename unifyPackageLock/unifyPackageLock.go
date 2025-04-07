package unifyPackageLock

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
)

/*
run on the file out put from CreateCombinedJsonLock()
*/
func runNpmAudit(targetDir string, packageLockReq bool) ([]byte, error) {
	var cmd *exec.Cmd
	if packageLockReq {
		cmd = exec.Command("npm", "audit", "--json")
	} else {
		if entries, err := os.ReadDir(targetDir); err != nil {
			fmt.Println(err)
		} else {
			count := 0
			for _, entry := range entries {
				if entry.Name() == "package-lock.json" {
					count++
				}
			}
			if count != 0 {
				cmdFirst := exec.Command("npm", "i", "--package-lock-only", "--force")
				cmdFirst.Dir = targetDir
				_, _ = cmdFirst.Output()
			}
		}
		cmd = exec.Command("npm", "audit", "--json")
	}
	cmd.Dir = targetDir
	output, _ := cmd.Output()
	if len(output) == 0 {
		return output, errors.New(fmt.Sprint("NPM output is empty"))
	}
	return output, nil
}

func CombinedOutputNpmAudit(targetDir []string, outputPath string, packageLockReq bool) error {
	reports := make([]NpmReport, len(targetDir))
	for i, dir := range targetDir {
		reports[i] = NpmReport{}
		output, err := runNpmAudit(dir, packageLockReq)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(output, &reports[i])
		if err != nil {
			fmt.Println("issue unmarshalling json: ", err.Error())
			return err
		}
	}
	for _, report := range reports[:1] {
		reports[0].Combine(&report)
	}
	data, err := json.Marshal(reports)
	if err != nil {
		fmt.Println("Issue while marshaling json data from npm reports: ", err.Error())
		return err
	}
	if err := os.WriteFile(outputPath+"/npm.json", data, 0644); err != nil {
		fmt.Println("error while writing to file", err.Error())
		os.Exit(1)
	}
	return nil
}

func GetTargetDirs(targetDir string, targetFileName string) ([]string, error) {
	var targetFiles []string
	err := findTargetFiles(targetDir, &targetFiles, targetFileName)
	if err != nil {
		return targetFiles, err
	}
	return targetFiles, nil
}

func findTargetFiles(targetDir string, targetFiles *[]string, targetFileName string) error {
	entries, err := os.ReadDir(targetDir)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			_ = findTargetFiles(targetDir+"/"+entry.Name(), targetFiles, targetFileName)
		} else {
			if entry.Name() == targetFileName {
				*targetFiles = append(*targetFiles, targetDir)
			}
		}
	}

	if len(*targetFiles) == 0 {
		return errors.New("no target files found: " + targetDir)
	}
	return nil
}
