package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"ginner/config"
)

// ProcessLock How to use:
//func main() {
//lock, lockFile, err := util.ProcessLock()
//if err != nil {
//logger.Log.Fatal(err)
//}
//defer os.Remove(lockFile)
//defer lock.Close()
//}
func ProcessLock() (*os.File, string, error) {
	abPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return nil, "", err
	}

	logDir := abPath + "/" + config.Conf().Log.Dirname + "/"
	filename := filepath.Base(os.Args[0])
	lockFile := logDir + filename + ".pid"

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

	pid, err := strconv.Atoi(string(filePid))
	if err != nil {
		return lock, lockFile, err
	}

	return lock, lockFile, fmt.Errorf(
		"Found lockfile %s, another copy is running, pid %d",
		lockFile,
		pid,
	)
}
