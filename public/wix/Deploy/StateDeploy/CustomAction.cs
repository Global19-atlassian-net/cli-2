using Microsoft.Deployment.WindowsInstaller;
using System;
using System.Text;
using System.Diagnostics;
using System.Threading;
using System.IO;
using System.Net;

namespace StateDeploy
{
    public class CustomActions
    {
        public static string STATE_TOOL_PATH;

        public static ActionResult InstallStateTool(Session session)
        {
            session.Log("Installing State Tool if necessary");
            if (session.CustomActionData["STATE_TOOL_INSTALLED"] == "true")
            {
                STATE_TOOL_PATH = session.CustomActionData["STATE_TOOL_PATH"];
                session.Log("State Tool is installed, no installation required");
                return ActionResult.Success;
            }

            string tempDir = Path.GetTempPath();
            string scriptPath = Path.Combine(tempDir, "install.ps1");
            string installPath = Path.Combine(Environment.GetFolderPath(Environment.SpecialFolder.ApplicationData), "ActiveState", "bin");

            Status.ProgressBar.StatusMessage(session, "Installing State Tool...");

            ServicePointManager.SecurityProtocol |= SecurityProtocolType.Tls11 | SecurityProtocolType.Tls12;
            try
            {
                WebClient client = new WebClient();
                client.DownloadFile("https://platform.activestate.com/dl/cli/install.ps1", scriptPath);
            }
            catch (WebException e)
            {
                session.Log(string.Format("Encoutered exception downloading file: {0}", e.ToString()));
                return ActionResult.Failure;
            }

            string installCmd = string.Format("powershell \"{0} -n -t {1}\"", scriptPath, installPath);
            session.Log(string.Format("Running install command: {0}", installCmd));
            
            ActionResult result = RunCommand(session, installCmd);
            if (result.Equals(ActionResult.UserExit))
            {
                result = Uninstall.Remove.InstallDir(session, installPath);
                if (result.Equals(ActionResult.Failure))
                {
                    session.Log("Could not remove installation directory");
                    return ActionResult.Failure;
                }

                result = Uninstall.Remove.EnvironmentEntries(session, installPath);
                if (result.Equals(ActionResult.Failure))
                {
                    session.Log("Could not remove environment entries");
                    return ActionResult.Failure;
                }
                return ActionResult.UserExit;
            }

            STATE_TOOL_PATH = Path.Combine(installPath, "state.exe");
            return result;
        }

        private static ActionResult RunCommand(Session session, string cmd)
        {
            try
            {
                ProcessStartInfo procStartInfo = new ProcessStartInfo("cmd", "/c " + cmd);

                // The following commands are needed to redirect the standard output.
                // This means that it will be redirected to the Process.StandardOutput StreamReader.
                procStartInfo.RedirectStandardOutput = true;
                procStartInfo.RedirectStandardError = true;
                procStartInfo.UseShellExecute = false;
                // Do not create the black window.
                procStartInfo.CreateNoWindow = true;

                Process proc = new Process();
                proc.StartInfo = procStartInfo;

                proc.OutputDataReceived += new DataReceivedEventHandler((sender, e) =>
                {
                    // Prepend line numbers to each line of the output.
                    if (!String.IsNullOrEmpty(e.Data))
                    {
                        session.Log("out: " + e.Data);
                    }
                });
                proc.ErrorDataReceived += new DataReceivedEventHandler((sender, e) =>
                {
                    // Prepend line numbers to each line of the output.
                    if (!String.IsNullOrEmpty(e.Data))
                    {
                        session.Log("err: " + e.Data);
                    }
                });
                proc.Start();

                // Asynchronously read the standard output and standard error of the spawned process.
                // This raises OutputDataReceived/ErrorDataReceived events for each line of output/errors.
                proc.BeginOutputReadLine();
                proc.BeginErrorReadLine();

                while (!proc.HasExited)
                {
                    try
                    {
                        // This is just hear to throw an InstallCanceled Exception if necessary
                        Status.ProgressBar.Increment(session, 0);
                        Thread.Sleep(200);
                    }
                    catch (InstallCanceledException)
                    {
                        session.Log("Caught install cancelled exception");
                        ActiveState.Process.KillProcessAndChildren(proc.Id);
                        return ActionResult.UserExit;
                    }
                }
                proc.WaitForExit();

                proc.Close();
            }
            catch (Exception objException)
            {
                session.Log(string.Format("Caught exception: {0}", objException));
                return ActionResult.Failure;
            }
            return ActionResult.Success;
        }


        [CustomAction]
        public static ActionResult StateDeploy(Session session)
        {

            var res = InstallStateTool(session);
            if (res != ActionResult.Success) {
                return res;
            }
            session.Log("Starting state deploy");

            Status.ProgressBar.StatusMessage(session, string.Format("Deploying project {0}...", session.CustomActionData["PROJECT_NAME"]));
            MessageResult incrementResult = Status.ProgressBar.Increment(session, 3);
            if (incrementResult == MessageResult.Cancel)
            {
                return ActionResult.UserExit;
            }

            string deployCmd = BuildDeployCmd(session);
            session.Log(string.Format("Executing deploy command: {0}", deployCmd));
            try
            {
                var runResult = RunCommand(session, deployCmd);
                if (runResult == ActionResult.UserExit) {
                    ActionResult result = Uninstall.Remove.InstallDir(session, session.CustomActionData["INSTALLDIR"]);
                    if (result.Equals(ActionResult.Failure))
                    {
                        session.Log("Could not remove installation directory");
                        return ActionResult.Failure;
                    }

                    result = Uninstall.Remove.EnvironmentEntries(session, session.CustomActionData["INSTALLDIR"]);
                    if (result.Equals(ActionResult.Failure))
                    {
                        session.Log("Could not remove environment entries");
                        return ActionResult.Failure;
                    }
                    return ActionResult.UserExit;
                } else if (runResult != ActionResult.Success)
                {
                    return runResult;
                }
            }
            catch (Exception objException)
            {
                session.Log(string.Format("Caught exception: {0}", objException));
                return ActionResult.Failure;
            }

            return ActionResult.Success;
        }

        private static string BuildDeployCmd(Session session)
        {
            string installDir = session.CustomActionData["INSTALLDIR"];
            string projectName = session.CustomActionData["PROJECT_NAME"];
            string isModify = session.CustomActionData["IS_MODIFY"];

            StringBuilder deployCMDBuilder = new StringBuilder(STATE_TOOL_PATH + " deploy");
            if (isModify == "true")
            {
                deployCMDBuilder.Append(" --force");
            }

            deployCMDBuilder.Append(" --output json");

            // We quote the string here as Windows paths that contain spaces must be quoted.
            // We also account for a path ending with a slash and ensure that the quote character
            // isn't preserved.
            deployCMDBuilder.AppendFormat(" {0} --path=\"{1}\\\"", projectName, @installDir);

            return deployCMDBuilder.ToString();
        }
    }
}
