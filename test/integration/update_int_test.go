// +build !short

package integration

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/testhelpers/integration"
	"github.com/ActiveState/cli/pkg/projectfile"
)

type UpdateIntegrationTestSuite struct {
	integration.Suite
}

func (suite *UpdateIntegrationTestSuite) SetupTest() {
	suite.Suite.SetupTest()
	suite.AppendEnv([]string{"ACTIVESTATE_CLI_AUTO_UPDATE_TIMEOUT=10"})
	suite.AppendEnv([]string{"ACTIVESTATE_CLI_UPDATE_BRANCH=master"})
}

func (suite *UpdateIntegrationTestSuite) getVersion() string {
	suite.Spawn("--version")
	suite.Wait()
	version := strings.Split(suite.Output(), "-")
	return version[len(version)-1]
}

func (suite *UpdateIntegrationTestSuite) TestAutoUpdateDisabled() {
	suite.AppendEnv([]string{"ACTIVESTATE_CLI_DISABLE_UPDATES=true"})
	suite.NotEqual(constants.BuildNumber, suite.getVersion(), "Versions should match as auto-update should not have occurred")
}

func (suite *UpdateIntegrationTestSuite) TestAutoUpdate() {
	suite.AppendEnv([]string{"ACTIVESTATE_CLI_DISABLE_UPDATES=false"})
	suite.NotEqual(constants.BuildNumber, suite.getVersion(), "Versions shouldn't match as auto-update should have occurred")
}

func (suite *UpdateIntegrationTestSuite) TestLocked() {
	// We need a projectfile to be able to version lock
	dir, err := ioutil.TempDir("", "")
	suite.Require().NoError(err)
	os.Chdir(dir)
	projectURL := fmt.Sprintf("https://%s/string/string?commitID=00010001-0001-0001-0001-000100010001", constants.PlatformURL)
	pjfile := projectfile.Project{
		Project: projectURL,
	}
	pjfile.SetPath(filepath.Join(dir, constants.ConfigFileName))
	pjfile.Save()

	suite.AppendEnv([]string{"ACTIVESTATE_CLI_DISABLE_UPDATES=false"})
	suite.Spawn("update", "--lock")
	suite.Expect("Version locked at")
	suite.Wait()

	suite.NotEqual(constants.BuildNumber, suite.getVersion(), "Versions should match because locking is enabled")
}

func (suite *UpdateIntegrationTestSuite) TestUpdate() {
	suite.AppendEnv([]string{"ACTIVESTATE_CLI_DISABLE_UPDATES=true"})
	suite.Spawn("update")
	// on master branch, we might already have the latest version available
	if os.Getenv("GIT_BRANCH") == "master" {
		suite.ExpectRe("(Update completed|You are using the latest version available)", 60*time.Second)
	} else {
		suite.Expect("Update completed", 60*time.Second)
	}
	suite.Wait()

	suite.NotEqual(constants.BuildNumber, suite.getVersion(), "Versions shouldn't match as we ran update")
}

func TestUpdateIntegrationTestSuite(t *testing.T) {
	_ = suite.Run // vscode won't show test helpers unless I use this .. -.-

	if testing.Short() {
		t.Skip("skipping integration test in short mode.")
	}

	suite.Run(t, new(UpdateIntegrationTestSuite))

	// parallel doesn't work with these due to contamination. The RunParallel function does not seem to allow for
	// setting up individual tests
	// integration.RunParallel(t, new(UpdateIntegrationTestSuite))
}
