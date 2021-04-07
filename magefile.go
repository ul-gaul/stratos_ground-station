// +build mage

package main

import (
    "context"
    "fmt"
    "github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
    "github.com/magefile/mage/sh"
    log "github.com/sirupsen/logrus"
    "os"
    "os/exec"
    "path/filepath"
    "runtime"
    "syscall"
    "time"
)

var (
    packageName     = "github.com/ul-gaul/stratos_ground-station"
    backendDelay    = 10 * time.Second
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

// wait the specified time before returning
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

// Build builds the frontend then builds the application binary
func Build() error {
    log.Info("Building...")
    mg.Deps(Frontend{}.Build)
    return Backend{}.Build()
}

// Run starts the development server for the frontend then starts the backend
func Run(ctx context.Context) error {
    log.Info("Running...")
    
    var cancel context.CancelFunc
    ctx, cancel = context.WithCancel(ctx)
    defer cancel()
    
    mg.CtxDeps(ctx, Frontend{}.Run, func(ctx context.Context) error {
        defer cancel()
        wait(backendDelay)
        return Backend{}.Run(ctx)
    })
    
    return nil
}

// Clean removes compiled files and dependency folders.
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


/***************************************** BACKEND *****************************************/

type Backend mg.Namespace

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
    return sh.RunWithV(env, goexe, "build", "-ldflags", ldFlags, "-o", bin, ".")
}

func (Backend) Run(ctx context.Context) error {
    log.Info("Starting Backend...")
    return sh.RunV(goexe, "run", "-tags", "dev", ".")
}

/***************************************** FRONTEND *****************************************/

type Frontend mg.Namespace

func (Frontend) InstallDeps(ctx context.Context) error {
    log.Info("Installing dependencies...")
    return npm(ctx, "install")
}

func (Frontend) Build(ctx context.Context) error {
    mg.Deps(Frontend{}.InstallDeps)
    log.Info("Building Frontend...")
    return npm(ctx, "run", "build")
}

func (Frontend) Run(ctx context.Context) error {
    mg.CtxDeps(ctx, Frontend{}.InstallDeps)
    log.Info("Starting Frontend...")
    return npm(ctx, "run", "serve")
}

func npm(ctx context.Context, args ...string) error {
    return runFrom(ctx, frontendDir, "npm", args...)
}

// runFrom runs the command from the specified directory
func runFrom(ctx context.Context, dir, command string, args ...string) error {
    cmd := exec.CommandContext(ctx, command, args...)
    cmd.Dir = dir
    cmd.Env = os.Environ()
    cmd.Stderr = os.Stderr
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.SysProcAttr = &syscall.SysProcAttr{
        Pdeathsig: syscall.SIGTERM | syscall.SIGKILL,
    }
    return cmd.Run()
}
