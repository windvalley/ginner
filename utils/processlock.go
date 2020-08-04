package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

//ProcessLock How to use:
//func main() {
//lock, lockFile, err := utils.ProcessLock()
//if err != nil {
//logger.Log.Fatal(err)
//}
//defer os.Remove(lockFile)
//defer lock.Close()
//}
func ProcessLock() (*os.File, string, error) {
	filename := filepath.Base(os.Args[0])
	lockFile := "/var/run/" + filename + ".pid"

	lock, err := os.Open(lockFile)
	if err != nil {
		lock, err := os.Create(lockFile)
		if err != nil {
			return lock, lockFile, err
		}
		pid := fmt.Sprint(os.Getpid())
		lock.WriteString(pid)
		return lock, lockFile, nil
	}

	filePid, err := ioutil.ReadAll(lock)
	if err != nil {
		return lock, lockFile, err
	}

	pid, _ := strconv.Atoi(string(filePid))
	return lock, lockFile, fmt.Errorf(
		"Found lockfile %s, another copy is running, pid %d",
		lockFile,
		pid,
	)
}
