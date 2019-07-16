package integration_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	tsuite "github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/test/integration/expectx/suite"
)

type ActivateTestSuite struct {
	Suite
}

func (suite *ActivateTestSuite) TestActivatePython3() {
	root := environment.GetRootPathUnsafe()
	os.Chdir(filepath.Join(root, "test/integration/testdata"))

	suite.LoginAsPersistentUser()
	suite.AppendEnv([]string{"ACTIVESTATE_CLI_DISABLE_RUNTIME=false"})

	tempDir, err := ioutil.TempDir("", "")
	suite.Require().NoError(err)
	os.Remove(tempDir)
	suite.Require().NoError(err)

	suite.Spawn("activate", "ActiveState-CLI/Python3")
	suite.Expect("Where would you like to checkout")
	suite.Send(tempDir)
	suite.Expect("Downloading")
	suite.Expect("Installing", 120*time.Second)
	suite.WaitForInput(120 * time.Second)
	suite.Send("echo ${ACTIVESTATE_ACTIVATED}")
	suite.Send("echo $-")
	//suite.Send("echo \"python3 bin: $(which python3)\"")
	//suite.Send(`echo exit | strace bash -li |& less | grep '^open'`)
	suite.Send("python3 -c \"import sys; print(sys.copyright)\"")
	suite.Expect("ActiveState Software Inc.")
	suite.Send("python3 -c \"import numpy; print(numpy.__doc__)\"")
	suite.Expect("import numpy as np")
	suite.SendQuit()
	suite.Wait()
}

func TestActivateTestSuite(t *testing.T) {
	_ = tsuite.Run // vscode won't show test helpers unless I use this .. -.-

	//suite.Run(t, new(ActivateTestSuite))
	suite.RunParallel(t, new(ActivateTestSuite))
}