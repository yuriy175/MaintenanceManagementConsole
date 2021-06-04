using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.Reflection;
using System.Text;

namespace MessagesSender.BL.Helpers
{
    /// <summary>
    /// process helper class
    /// </summary>
    internal static class ProcessHelper
    {
        /// <summary>
        /// Run a process and wait for the exit
        /// </summary>
        /// <param name="exePath">executable path</param>
        /// <param name="args">args</param>
        public static void ProcessRunAndWait(string exePath, string args)
        {
            var path = Path.Combine(Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location), exePath);
            var processStartInfo = new ProcessStartInfo(
                path,
                args
                );
            processStartInfo.WorkingDirectory = Path.GetDirectoryName(path);

            var process = Process.Start(processStartInfo);

            process.WaitForExit();
        }
    }
}
