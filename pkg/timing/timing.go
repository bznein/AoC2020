package timing

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"
)

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	packageName, funcName := retrieveCallInfo()
	log.Println(fmt.Sprintf("%s.%s took %s", packageName, funcName, elapsed))
}

func retrieveCallInfo() (string, string) {
	pc, _, _, _ := runtime.Caller(2)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}
	packageArray := strings.Split(packageName, "/")
	packageName = packageArray[len(packageArray)-1]
	return packageName, funcName

}
