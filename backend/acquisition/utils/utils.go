package utils

import (
    log "github.com/sirupsen/logrus"
    "os"
    "path/filepath"
)

// CheckErr logs the error and panics if err is not nil.
func CheckErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}



// GetExecutableDir returns the path to the directory containing the executable.
func GetExecutableDir() (string, error) {
    exe, err := os.Executable()
    if err != nil {
        exe, err = filepath.EvalSymlinks(exe)
        if err != nil {
            return exe, err
        }
    }
    return filepath.Dir(exe), nil
}