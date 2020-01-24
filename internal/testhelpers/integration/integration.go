package integration

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/ActiveState/vt10x"
	"github.com/google/uuid"
	"github.com/phayes/permbits"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v2"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/pkg/expect"
	"github.com/ActiveState/cli/pkg/projectfile"
)

var PersistentUsername = "cli-integration-tests"
var PersistentPassword = "test-cli-integration"

var defaultTimeout = 10 * time.Second

// Suite is our integration test suite
type Suite struct {
	suite.Suite
	console    *expect.Console
	state      *vt10x.State
	executable string
	cmd        *exec.Cmd
	env        []string
	wd         *string
}

// SetupTest sets up an integration test suite for testing the state tool executable
func (s *Suite) SetupTest() {
	exe := ""
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}

	root := environment.GetRootPathUnsafe()
	executable := filepath.Join(root, "build/"+constants.CommandName+exe)

	if !fileutils.FileExists(executable) {
		s.FailNow("Integration tests require you to have built a state tool binary. Please run `state run build`.")
	}

	configDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)
	cacheDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)
	binDir, err := ioutil.TempDir("", "")
	s.Require().NoError(err)

	fmt.Println("Configdir: " + configDir)
	fmt.Println("Cachedir: " + cacheDir)
	fmt.Println("Bindir: " + binDir)

	s.executable = filepath.Join(binDir, constants.CommandName+exe)
	fail := fileutils.CopyFile(executable, s.executable)
	s.Require().NoError(fail.ToError())

	permissions, _ := permbits.Stat(s.executable)
	permissions.SetUserExecute(true)
	err = permbits.Chmod(s.executable, permissions)
	s.Require().NoError(err)

	s.ClearEnv()
	s.AppendEnv(os.Environ())
	s.AppendEnv([]string{
		"ACTIVESTATE_CLI_CONFIGDIR=" + configDir,
		"ACTIVESTATE_CLI_CACHEDIR=" + cacheDir,
		"ACTIVESTATE_CLI_DISABLE_UPDATES=true",
		"ACTIVESTATE_CLI_DISABLE_RUNTIME=true",
		"ACTIVESTATE_PROJECT=",
	})
}

// PrepareTemporaryWorkingDirectory prepares a temporary working directory to run the tests in
// It returns the directory name a clean-up function
func (s *Suite) PrepareTemporaryWorkingDirectory(prefix string) (tempDir string, cleanup func()) {

	tempDir, err := ioutil.TempDir("", prefix)
	s.Require().NoError(err)
	err = os.RemoveAll(tempDir)
	s.Require().NoError(err)
	err = os.MkdirAll(tempDir, 0770)
	s.Require().NoError(err)
	s.SetWd(tempDir)

	return tempDir, func() {
		os.RemoveAll(tempDir)
	}
}

// PrepareActiveStateYAML creates a projectfile.Project instance from the
// provided contents and saves the output to an as.y file within the named
// directory.
func (s *Suite) PrepareActiveStateYAML(dir, contents string) {
	msg := "cannot setup activestate.yaml file"

	contents = strings.TrimSpace(contents)
	projectFile := &projectfile.Project{}

	err := yaml.Unmarshal([]byte(contents), projectFile)
	s.Require().NoError(err, msg)

	projectFile.SetPath(filepath.Join(dir, "activestate.yaml"))
	fail := projectFile.Save()
	s.Require().NoError(fail.ToError(), msg)
}

// Executable returns the path to the executable under test (state tool)
func (s *Suite) Executable() string {
	return s.executable
}

// TearDownTest closes the terminal attached to this integration test suite
// Run this to clean-up everything set up with SetupTest()
func (s *Suite) TearDownTest() {
	if s.console != nil {
		s.console.Close()
	}
}

// ClearEnv removes all environment variables
func (s *Suite) ClearEnv() {
	s.env = []string{}
}

// AppendEnv appends new environment variable settings
func (s *Suite) AppendEnv(env []string) {
	s.env = append(s.env, env...)
}

// SetWd specifies a working directory for the spawned processes
// Use this method if you rely on running the test executable in a clean directory
// By default all tests are run in `os.TempDir()`
func (s *Suite) SetWd(dir string) {
	s.wd = &dir
}

// Spawn executes the state tool executable under test in a pseudo-terminal
func (s *Suite) Spawn(args ...string) {
	s.SpawnCustom(s.executable, args...)
}

// SpawnCustom executes an executable in a pseudo-terminal for integration tests
func (s *Suite) SpawnCustom(executable string, args ...string) {
	var wd string
	if s.wd == nil {
		wd = fileutils.TempDirUnsafe()
	} else {
		wd = *s.wd
	}

	s.cmd = exec.Command(executable, args...)
	s.cmd.Dir = wd
	s.cmd.Env = s.env

	// Create the process in a new process group.
	// This makes the behavior more consistent, as it isolates the signal handling from
	// the parent processes, which are dependent on the test environment.
	s.cmd.SysProcAttr = osutils.SysProcAttrForNewProcessGroup()
	fmt.Printf("Spawning '%s' from %s\n", osutils.CmdString(s.cmd), wd)

	var err error
	s.console, err = expect.NewConsole(
		expect.WithDefaultTimeout(defaultTimeout),
	)
	s.Require().NoError(err)

	err = s.console.Pty.StartProcessInTerminal(s.cmd)
	s.Require().NoError(err)
}

// Output returns the current Terminal snapshot.
func (s *Suite) Output() string {
	return s.console.Pty.State.String()
}

// ExpectRe listens to the terminal output and returns once the expected regular expression is matched or
// a timeout occurs
// Default timeout is 10 seconds
func (s *Suite) ExpectRe(value string, timeout ...time.Duration) {
	opts := []expect.ExpectOpt{expect.RegexpPattern(value)}
	if len(timeout) > 0 {
		opts = append(opts, expect.WithTimeout(timeout[0]))
	}
	_, err := s.console.Expect(opts...)
	if err != nil {
		s.FailNow(
			"Could not meet expectation",
			"Expectation: '%s'\nError: %v\n---\nTerminal snapshot:\n%s\n---\n",
			value, err, s.Output())
	}
}

// TerminalSnapshot returns a snapshot of the terminal output
func (s *Suite) TerminalSnapshot() string {
	return s.console.Pty.State.String()
}

// Expect listens to the terminal output and returns once the expected value is found or
// a timeout occurs
// Default timeout is 10 seconds
func (s *Suite) Expect(value string, timeout ...time.Duration) {
	opts := []expect.ExpectOpt{expect.String(value)}
	if len(timeout) > 0 {
		opts = append(opts, expect.WithTimeout(timeout[0]))
	}
	_, err := s.console.Expect(opts...)
	if err != nil {
		s.FailNow(
			"Could not meet expectation",
			"Expectation: '%s'\nError: %v\n---\nTerminal snapshot:\n%s\n---\n",
			value, err, s.Output())
	}
}

// WaitForInput returns once a shell prompt is active on the terminal
// Default timeout is 10 seconds
func (s *Suite) WaitForInput(timeout ...time.Duration) {
	usr, err := user.Current()
	s.Require().NoError(err)

	msg := "echo wait_ready_$HOME"
	if runtime.GOOS == "windows" {
		msg = "echo wait_ready_%USERPROFILE%"
	}

	s.SendLine(msg)
	s.Expect("wait_ready_"+usr.HomeDir, timeout...)
}

// SendLine sends a new line to the terminal, as if a user typed it
func (s *Suite) SendLine(value string) {
	_, err := s.console.SendLine(value)
	if err != nil {
		s.FailNow("Could not send data to terminal", "error: %v", err)
	}
}

// Send sends a string to the terminal as if a user typed it
func (s *Suite) Send(value string) {
	_, err := s.console.Send(value)
	if err != nil {
		s.FailNow("Could not send data to terminal", "error: %v", err)
	}
}

// Signal sends an arbitrary signal to the running process
func (s *Suite) Signal(sig os.Signal) error {
	return s.cmd.Process.Signal(sig)
}

// SendCtrlC tries to emulate what would happen in an interactive shell, when the user presses Ctrl-C
func (s *Suite) SendCtrlC() {
	s.Send(string([]byte{0x03})) // 0x03 is ASCI character for ^C
}

// Quit sends an interrupt signal to the tested process
func (s *Suite) Quit() error {
	return s.cmd.Process.Signal(os.Interrupt)
}

// Stop sends an interrupt signal for the tested process and fails if no process has been started yet.
func (s *Suite) Stop() error {
	if s.cmd == nil || s.cmd.Process == nil {
		s.FailNow("stop called without a spawned process")
	}
	return s.Quit()
}

// LoginAsPersistentUser is a common test case after which an integration test user should be logged in to the platform
func (s *Suite) LoginAsPersistentUser() {
	s.Spawn("auth", "--username", PersistentUsername, "--password", PersistentPassword)
	s.Expect("successfully authenticated")
	state, err := s.Wait()
	s.Require().NoError(err)
	s.Require().Equal(0, state.ExitCode())
}

// ExpectExitCode waits for the program under test to terminate, and checks that the returned exit code meets expectations
func (s *Suite) ExpectExitCode(exitCode int, timeout ...time.Duration) {
	ps, err := s.Wait(timeout...)
	if err != nil {
		s.FailNow(
			"Error waiting for process:",
			"\n%v\n---\nTerminal snapshot:\n%s\n---\n",
			err, s.TerminalSnapshot())
	}
	if ps.ExitCode() != exitCode {
		s.FailNow(
			"Process terminated with unexpected exit code\n",
			"Expected: %d, got %d\n---\nTerminal snapshot:\n%s\n---\n",
			exitCode, ps.ExitCode(), s.TerminalSnapshot())
	}
}

// Wait waits for the tested process to finish and returns its state including ExitCode
func (s *Suite) Wait(timeout ...time.Duration) (state *os.ProcessState, err error) {
	if s.cmd == nil || s.cmd.Process == nil {
		return
	}

	t := defaultTimeout
	if len(timeout) > 0 {
		t = timeout[0]
	}

	type processState struct {
		state *os.ProcessState
		err   error
	}
	states := make(chan processState)

	go func() {
		defer close(states)
		s, e := s.cmd.Process.Wait()
		states <- processState{state: s, err: e}
	}()

	select {
	case s := <-states:
		return s.state, s.err
	case <-time.After(t):
		return nil, fmt.Errorf("i/o error")
	}
}

func (s *Suite) TrimSpaceOutput() string {
	// On Windows this means both a carriage return and a new line. Windows
	// also picks up any spaces at the end of the console output, hence all
	// the cleaning we must do here.
	newlineRe := regexp.MustCompile(`\r?\n`)
	return newlineRe.ReplaceAllString(strings.TrimSpace(s.Output()), "")
}

func (s *Suite) CreateNewUser() string {
	uid, err := uuid.NewRandom()
	s.Require().NoError(err)

	username := fmt.Sprintf("user-%s", uid.String()[0:8])
	password := username
	email := fmt.Sprintf("%s@test.tld", username)

	s.Spawn("auth", "signup")
	s.Expect("username:")
	s.SendLine(username)
	s.Expect("password:")
	s.SendLine(password)
	s.Expect("again:")
	s.SendLine(password)
	s.Expect("name:")
	s.SendLine(username)
	s.Expect("email:")
	s.SendLine(email)
	s.Expect("account has been registered", 20*time.Second)
	s.Wait()

	return username
}

func (s *Suite) ExecutablePath() string {
	ext := ""
	if runtime.GOOS == "windows" {
		ext = ".exe"
	}
	name := constants.CommandName + ext
	root := environment.GetRootPathUnsafe()
	subdir := "build"

	return filepath.Join(root, subdir, name)
}

type Dirs struct {
	s      *suite.Suite
	base   string
	Config string
	Cache  string
	Bin    string
	Work   string
}

func (s *Suite) NewDirs() *Dirs {
	noErr := s.Require().NoError

	base, err := ioutil.TempDir("", "")
	noErr(err)

	config := filepath.Join(base, "config")
	cache := filepath.Join(base, "cache")
	bin := filepath.Join(base, "bin")
	work := filepath.Join(base, "work")

	noErr(os.MkdirAll(config, 0700))
	noErr(os.MkdirAll(cache, 0700))
	noErr(os.MkdirAll(bin, 0700))
	noErr(os.MkdirAll(work, 0700))

	return &Dirs{
		s:      &s.Suite,
		base:   base,
		Config: config,
		Cache:  cache,
		Bin:    bin,
		Work:   work,
	}
}

func (d *Dirs) Close() {
	d.s.Require().NoError(os.RemoveAll(d.base))
}

type Session struct {
	s       *suite.Suite
	console *expect.Console
	cmd     *exec.Cmd
	dirs    *Dirs
}

func (s *Suite) NewSession(dirs *Dirs, exe string, args ...string) *Session {
	noErr := s.Require().NoError

	if dirs == nil {
		dirs = s.NewDirs()
	}

	execu := filepath.Join(dirs.Bin, filepath.Base(exe))
	fail := fileutils.CopyFile(exe, execu)
	noErr(fail.ToError())

	permissions, _ := permbits.Stat(execu)
	permissions.SetUserExecute(true)
	noErr(permbits.Chmod(execu, permissions))

	var env []string
	env = append(env, os.Environ()...)
	env = append(env, []string{
		"ACTIVESTATE_CLI_CONFIGDIR=" + dirs.Config,
		"ACTIVESTATE_CLI_CACHEDIR=" + dirs.Cache,
		"ACTIVESTATE_CLI_DISABLE_UPDATES=true",
		"ACTIVESTATE_CLI_DISABLE_RUNTIME=true",
		"ACTIVESTATE_PROJECT=",
	}...)

	cmd := exec.Command(execu, args...)
	cmd.Dir = dirs.Work
	cmd.Env = env

	// Create the process in a new process group.
	// This makes the behavior more consistent, as it isolates the signal handling from
	// the parent processes, which are dependent on the test environment.
	cmd.SysProcAttr = osutils.SysProcAttrForNewProcessGroup()
	fmt.Printf("Spawning '%s' from %s\n", osutils.CmdString(cmd), dirs.Work)

	console, err := expect.NewConsole(
		expect.WithDefaultTimeout(defaultTimeout),
	)
	noErr(err)

	noErr(console.Pty.StartProcessInTerminal(cmd))

	return &Session{
		s:       &s.Suite,
		console: console,
		cmd:     cmd,
		dirs:    dirs,
	}
}

func (s *Suite) NewDefaultSession(args ...string) *Session {
	exe := s.ExecutablePath()

	if !fileutils.FileExists(exe) {
		s.FailNow("Integration tests require you to have built a state tool binary. Please run `state run build`.")
	}

	return s.NewSession(nil, exe, args...)
}

func (s *Session) Close() {
	if s.console != nil {
		err := s.console.Close()
		s.s.Require().NoError(err)
	}

	if s.dirs != nil {
		s.dirs.Close()
	}
}

// Output returns the current Terminal snapshot.
func (s *Session) Output() string {
	return s.console.Pty.State.String()
}

// ExpectRe listens to the terminal output and returns once the expected regular expression is matched or
// a timeout occurs
// Default timeout is 10 seconds
func (s *Session) ExpectRe(value string, timeout ...time.Duration) {
	opts := []expect.ExpectOpt{expect.RegexpPattern(value)}
	if len(timeout) > 0 {
		opts = append(opts, expect.WithTimeout(timeout[0]))
	}
	_, err := s.console.Expect(opts...)
	if err != nil {
		s.s.FailNow(
			"Could not meet expectation",
			"Expectation: '%s'\nError: %v\n---\nTerminal snapshot:\n%s\n---\n",
			value, err, s.Output())
	}
}

// Expect listens to the terminal output and returns once the expected value is found or
// a timeout occurs
// Default timeout is 10 seconds
func (s *Session) Expect(value string, timeout ...time.Duration) {
	opts := []expect.ExpectOpt{expect.String(value)}
	if len(timeout) > 0 {
		opts = append(opts, expect.WithTimeout(timeout[0]))
	}
	_, err := s.console.Expect(opts...)
	if err != nil {
		s.s.FailNow(
			"Could not meet expectation",
			"Expectation: '%s'\nError: %v\n---\nTerminal snapshot:\n%s\n---\n",
			value, err, s.Output())
	}
}

// WaitForInput returns once a shell prompt is active on the terminal
// Default timeout is 10 seconds
func (s *Session) WaitForInput(timeout ...time.Duration) {
	usr, err := user.Current()
	s.s.Require().NoError(err)

	msg := "echo wait_ready_$HOME"
	if runtime.GOOS == "windows" {
		msg = "echo wait_ready_%USERPROFILE%"
	}

	s.SendLine(msg)
	s.Expect("wait_ready_"+usr.HomeDir, timeout...)
}

// SendLine sends a new line to the terminal, as if a user typed it
func (s *Session) SendLine(value string) {
	_, err := s.console.SendLine(value)
	if err != nil {
		s.s.FailNow("Could not send data to terminal", "error: %v", err)
	}
}

// Send sends a string to the terminal as if a user typed it
func (s *Session) Send(value string) {
	_, err := s.console.Send(value)
	if err != nil {
		s.s.FailNow("Could not send data to terminal", "error: %v", err)
	}
}

// Signal sends an arbitrary signal to the running process
func (s *Session) Signal(sig os.Signal) {
	err := s.cmd.Process.Signal(sig)
	s.s.Require().NoError(err)
}

// SendCtrlC tries to emulate what would happen in an interactive shell, when the user presses Ctrl-C
func (s *Session) SendCtrlC() {
	s.Send(string([]byte{0x03})) // 0x03 is ASCI character for ^C
}

// Quit sends an interrupt signal to the tested process
func (s *Session) Quit() {
	s.Signal(os.Interrupt)
}

// Stop sends an interrupt signal for the tested process and fails if no process has been started yet.
func (s *Session) Stop() {
	if s.cmd == nil || s.cmd.Process == nil {
		s.s.FailNow("stop called without a spawned process")
	}
	s.Quit()
}

// ExpectExitCode waits for the program under test to terminate, and checks that the returned exit code meets expectations
func (s *Session) ExpectExitCode(exitCode int, timeout time.Duration) {
	ps, err := s.WaitFor(timeout)
	if err != nil {
		s.s.FailNow(
			"Error waiting for process:",
			"\n%v\n---\nTerminal snapshot:\n%s\n---\n",
			err, s.Output())
	}
	if ps.ExitCode() != exitCode {
		s.s.FailNow(
			"Process terminated with unexpected exit code\n",
			"Expected: %d, got %d\n---\nTerminal snapshot:\n%s\n---\n",
			exitCode, ps.ExitCode(), s.Output())
	}
}

// Wait waits for the tested process to finish and returns its state including ExitCode
func (s *Session) Wait() (*os.ProcessState, error) {
	return s.wait(0)
}

func (s *Session) WaitFor(d time.Duration) (*os.ProcessState, error) {
	return s.wait(d)
}

func (s *Session) wait(timeout time.Duration) (state *os.ProcessState, err error) {
	if s.cmd == nil || s.cmd.Process == nil {
		return
	}

	d := defaultTimeout
	if timeout > 0 {
		d = timeout
	}

	t := time.AfterFunc(d, func() { s.s.FailNow("wait took too long") })
	defer t.Stop()

	return s.cmd.Process.Wait()
}

func (s *Session) TrimSpaceOutput() string {
	// When the PTY reaches 80 characters it continues output on a new line.
	// On Windows this means both a carriage return and a new line. Windows
	// also picks up any spaces at the end of the console output, hence all
	// the cleaning we must do here.
	newlineRe := regexp.MustCompile(`\r?\n`)
	return newlineRe.ReplaceAllString(strings.TrimSpace(s.Output()), "")
}

// LoginAsPersistentUser is a common test case after which an integration test user should be logged in to the platform
func (s *Suite) XLoginAsPersistentUser() {
	ts := s.NewDefaultSession("auth", "--username", PersistentUsername, "--password", PersistentPassword)
	defer ts.Close()

	ts.Expect("successfully authenticated")
	state, err := ts.Wait()
	s.Require().NoError(err)
	s.Require().Equal(0, state.ExitCode())
}

func (s *Suite) XCreateNewUser() string {
	uid, err := uuid.NewRandom()
	s.Require().NoError(err)

	username := fmt.Sprintf("user-%s", uid.String()[0:8])
	password := username
	email := fmt.Sprintf("%s@test.tld", username)

	ts := s.NewDefaultSession("auth", "signup")
	defer ts.Close()

	ts.Expect("username:")
	ts.SendLine(username)
	ts.Expect("password:")
	ts.SendLine(password)
	ts.Expect("again:")
	ts.SendLine(password)
	ts.Expect("name:")
	ts.SendLine(username)
	ts.Expect("email:")
	ts.SendLine(email)
	ts.Expect("account has been registered", 20*time.Second)
	ts.Wait()

	return username
}
