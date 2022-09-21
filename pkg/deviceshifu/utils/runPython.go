package utils

import (
	"fmt"
	"k8s.io/klog/v2"
	"os/exec"
)

const (
	SCRIPTDIR = "python_customized_handlers"
	PYTHON    = "python"
	CMDARG    = "-c"
)

func ProcessInstruction(moduleName string, funcName string, rawData string) string {
	cmdString := fmt.Sprintf("import %s; print(%s.%s(%s))", moduleName, moduleName, funcName, rawData)
	cmd := exec.Command(PYTHON, CMDARG, cmdString)
	cmd.Dir = SCRIPTDIR

	processed, err := cmd.CombinedOutput()
	if err != nil {
		klog.Errorf("Failed process command %v\n, error:%v", cmdString, err.Error())
	}
	return string(processed)
}
