package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// ProcessLock How to use:
//func main() {
//abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
//if err != nil {
//panic(err)
//}
//lock, lockFile, err := util.ProcessLock(abPath+"/logs")
//if err != nil {
//logger.Log.Fatal(err)
//}
//defer os.Remove(lockFile)
//defer lock.Close()
//}
func ProcessLock(pidDir string) (*os.File, string, error) {
	pidDir = strings.TrimSuffix(pidDir, "/") + "/"
	filename := filepath.Base(os.Args[0])
	lockFile := pidDir + filename + ".pid"

	lock, err := os.Open(lockFile)
	if err != nil {
		lock, err := os.Create(lockFile)
		if err != nil {
			return lock, lockFile, err
		}
		pid := fmt.Sprint(os.Getpid())
		_, err = lock.WriteString(pid)
		return lock, lockFile, err
	}

	filePid, err := ioutil.ReadAll(lock)
	if err != nil {
		return lock, lockFile, err
	}

	pid, err := strconv.Atoi(string(filePid))
	if err != nil {
		return lock, lockFile, err
	}

	return lock, lockFile, fmt.Errorf(
		"found lockfile %s, another copy is running, pid %d",
		lockFile,
		pid,
	)
}
