package unifyPackageLock

import (
	"fmt"
	"os"
	"testing"
)

func TestRunNpmAuditNoPackgeLock(t *testing.T) {
	fmt.Println("Running TestRunNpmAuditNoPackgeLock...")
	output, err := runNpmAudit("/home/aidan/LAB/NPM_IMPROVED/testResources/client", false)
	if err != nil {
		t.Error("error while getting non package lock data", err.Error())
	}
	if len(output) == 0 {
		t.Error("expected output to be non zero")
	}

}

func TestRunNpmAuditWithPackgeLock(t *testing.T) {
	fmt.Println("Running TestRunNpmAuditWithPackgeLock...")
	output, err := runNpmAudit("/home/aidan/LAB/NPM_IMPROVED/testResources/client", true)
	if err != nil {
		t.Error("error while getting non package lock data", err.Error())
		t.Fail()
	}
	if len(output) == 0 {
		t.Error("expected output to be non zero")
		t.Fail()
	}
}

func TestCombinedOutputNpmAudit(t *testing.T) {
	fmt.Println("Running TestCombinedOutputNpmAudit...")
	targetDir := []string{"/home/aidan/LAB/NPM_IMPROVED/testResources/client"}
	CombinedOutputNpmAudit(targetDir, "/home/aidan/LAB/NPM_IMPROVED/testResources/npm.json", true)
	if _, err := os.Stat("/home/aidan/LAB/NPM_IMPROVED/testResources/npm.json"); os.IsNotExist(err) {
		t.Error("expected npm.json file to exist")
		t.Fail()
	} else if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestRunNpmAuditNoPackgeLockFedWebsite(t *testing.T) {
	fmt.Println("Running TestRunNpmAuditNoPackgeLock...")
	output, err := runNpmAudit("/home/aidan/LAB/NPM_IMPROVED/testResources/fed-website.b", false)
	if err != nil {
		t.Error("error while getting non package lock data", err.Error())
	}
	if len(output) == 0 {
		t.Error("expected output to be non zero")
	}

}

func TestRunNpmAuditWithPackgeLockFedWebsite(t *testing.T) {
	fmt.Println("Running TestRunNpmAuditWithPackgeLock...")
	output, err := runNpmAudit("/home/aidan/LAB/NPM_IMPROVED/testResources/fed-website.b", true)
	if err != nil {
		t.Error("error while getting non package lock data", err.Error())
		t.Fail()
	}
	if len(output) == 0 {
		t.Error("expected output to be non zero")
		t.Fail()
	}
}

func TestCombinedOutputNpmAuditFedWebsite(t *testing.T) {
	fmt.Println("Running TestCombinedOutputNpmAudit...")
	targetDir := []string{"/home/aidan/LAB/NPM_IMPROVED/testResources/fed-website.b"}
	CombinedOutputNpmAudit(targetDir, "/home/aidan/LAB/NPM_IMPROVED/testResources/npm.json", true)
	if _, err := os.Stat("/home/aidan/LAB/NPM_IMPROVED/testResources/npm.json"); os.IsNotExist(err) {
		t.Error("expected npm.json file to exist")
		t.Fail()
	} else if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestGetTargetDirs(t *testing.T) {
	fmt.Println("Running TestGetTargetDirs...")
	list, err := GetTargetDirs("/home/aidan/LAB/NPM_IMPROVED/testResources/Noteslify.b", "package.json")
	if err != nil {
		t.Error("error while getting target dirs", err.Error())
	}
	if len(list) != 5 {
		t.Error("expected target dirs to be 4")
	}
}
