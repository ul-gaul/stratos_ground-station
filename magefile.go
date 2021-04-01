// +build mage

package main

import (
    "fmt"
    "github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
    "github.com/magefile/mage/sh"
    log "github.com/sirupsen/logrus"
    "os"
    "path/filepath"
    "runtime"
    "time"
)

const (
    packageName     = "github.com/ul-gaul/stratos_ground-station"
    backendDelay    = 15 * time.Second
    frontendDir     = "dashboard"
    appName         = "ground-station"
    binaryOutputDir = "bin/"
    ldFlags         = "-X main.FrontendIndexPath=$FRONTEND_PATH/dist/index.html"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
var Default = Run

// allow user to override go executable by running as GOEXE=xxx make ... on unix-like systems
var goexe = "go"

func init() {
    if exe := os.Getenv("GOEXE"); exe != "" {
        goexe = exe
    }
    
    // We want to use Go 1.11 modules even if the source lives inside GOPATH.
    // The default is "auto".
    _ = os.Setenv("GO111MODULE", "on")
}

type Frontend mg.Namespace

func (Frontend) InstallDeps() error {
    log.Info("Installing dependencies...")
    return npm("install")
}

func (Frontend) Build() error {
    mg.Deps(Frontend{}.InstallDeps)
    log.Info("Building Frontend...")
    return npm("run", "build")
}

func (Frontend) Run() error {
    mg.Deps(Frontend{}.InstallDeps)
    log.Info("Starting Frontend...")
    return npm("run", "serve")
}

type Backend mg.Namespace

func (Backend) Run() error {
    log.Info("Starting Backend...")
    return sh.Run(goexe, "run", "-tags", "dev", ".")
}

func (Backend) Build() error {
    log.Info("Building Application Binary...")
    ext := ""
    if runtime.GOOS == "windows" {
        ext = ".exe"
    }
    
    env := map[string]string{
        "PACKAGE":       packageName,
        "FRONTEND_PATH": frontendDir,
    }
    
    bin := fmt.Sprintf("%s-%s-%s%s", appName, runtime.GOOS, runtime.GOARCH, ext)
    bin = filepath.Join(binaryOutputDir, bin)
    return sh.RunWith(env, goexe, "build", "-ldflags", ldFlags, "-o", bin, ".")
}

func Build() error {
    log.Info("Building...")
    mg.Deps(Frontend{}.Build)
    return Backend{}.Build()
}

func Clean() error {
    log.Info("Cleaning...")
    
    paths := []string{
        filepath.Join(frontendDir, "dist"),
        filepath.Join(frontendDir, "node_modules"),
        binaryOutputDir,
    }
    
    for _, p := range paths {
        if err := sh.Rm(p); err != nil {
            return err
        }
    }
    
    return nil
}

func wait(d time.Duration) {
    ticker := time.NewTicker(time.Second)
    defer ticker.Stop()
    done := time.After(d)
    end := time.Now().Add(d)
    for {
        select {
        case <-ticker.C:
            log.Infof("Backend starts in %s...",
                time.Until(end).Truncate(time.Second))
        case <-done:
            return
        }
    }
}

func Run() error {
    log.Info("Running...")
    
    wd, err := os.Getwd()
    if err != nil {
        return err
    }
    
    mg.Deps(Frontend{}.Run, func() error {
        wait(backendDelay)
        _ = os.Chdir(wd)
        err = Backend{}.Run()
        if err != nil {
            return err
        }
        os.Exit(0)
        return nil
    })
    
    return nil
}

func npm(args ...string) error {
    return runFrom(frontendDir, "npm", args...)
}

func runFrom(dir, cmd string, args ...string) error {
    wd, _ := os.Getwd()
    if err := os.Chdir(dir); err != nil {
        return err
    }
    err := sh.RunV(cmd, args...)
    _ = os.Chdir(wd)
    return err
}
