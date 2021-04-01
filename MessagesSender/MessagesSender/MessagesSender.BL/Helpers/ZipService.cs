using MessagesSender.Core.Interfaces;
using Serilog;
using System;
using System.Collections.Generic;
using System.IO;
using System.IO.Compression;
using System.Text;
using System.Threading.Tasks;

namespace MessagesSender.BL.Helpers
{
	/// <summary>
	/// zip service
	/// </summary>
    class ZipService : IZipService
	{
		private readonly ILogger _logger;
		private readonly ITopicService _topicService;

		/// <summary>
		/// public constructor
		/// </summary>
		/// <param name="logger">logger</param>
		/// <param name="topicService">topic service</param>
		public ZipService(
			ILogger logger,
			ITopicService topicService)
		{
			_logger = logger;
			_topicService = topicService;
		}

		/// <summary>
		/// zips folder
		/// </summary>
		/// <param name="folder">folder to zip</param>
		/// <returns>zip file path</returns>		
		public async Task<string> ZipFolderAsync(string folder)
        {
			var destFolder = folder + (await _topicService.GetTopicAsync())?.Replace("/","_") + DateTime.Now.ToString("_dd_MM_yyyy_HH_mm_ss");
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

		/// <summary>
		/// zips file
		/// </summary>
		/// <param name="filePath">file to zip</param>
		/// <returns>zipped file path</returns>
		public async Task<string> ZipFileAsync(string filePath)
		{
			var fileName = Path.GetFileNameWithoutExtension(filePath);
			var destFolder = Path.GetDirectoryName(filePath) + @"\" + (await _topicService.GetTopicAsync())?.Replace("/", "_");
			if (Directory.Exists(destFolder))
			{
				Directory.Delete(destFolder);
			}

			Directory.CreateDirectory(destFolder);
			File.Move(filePath, Path.Combine(destFolder, fileName + Path.GetExtension(filePath)));

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
