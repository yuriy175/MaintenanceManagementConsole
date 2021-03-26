using System;
using System.Collections.Generic;
using System.IO;
using System.IO.Compression;
using System.Text;

namespace MessagesSender.BL.Helpers
{
    static class ZipHelper
    {
        public static string ZipFolder(string folder)
        {
            var destFolder = folder + DateTime.Now.ToString("_dd_MM_yyyy_HH_mm_ss");
            if (Directory.Exists(destFolder))
            {
                Directory.Delete(destFolder);
            }

            Directory.CreateDirectory(destFolder);
            CopyDirectory(folder, destFolder);

            var zipName = destFolder + ".zip";
            ZipFile.CreateFromDirectory(destFolder, zipName, CompressionLevel.Optimal, true);

            return zipName;
        }

        private static void CopyDirectory(string sourcePath, string destinationPath)
        {
            // Now Create all of the directories
            foreach (string dirPath in Directory.GetDirectories(sourcePath, "*",
                SearchOption.AllDirectories))
                Directory.CreateDirectory(dirPath.Replace(sourcePath, destinationPath));

            // Copy all the files & Replaces any files with the same name
            foreach (string newPath in Directory.GetFiles(sourcePath, "*.*",
                SearchOption.AllDirectories))
                File.Copy(newPath, newPath.Replace(sourcePath, destinationPath), true);
        }
    }
}
