using MessagesSender.Core.Interfaces;
using Serilog;
using System;
using System.Collections.Generic;
using System.IO;
using System.IO.Compression;
using System.Net;
using System.Text;
using System.Threading.Tasks;

namespace MessagesSender.BL.Remoting
{
	/// <summary>
	/// ftp client
	/// </summary>
    class FtpClient : IFtpClient
	{
		private readonly ILogger _logger;
		private readonly ITopicService _topicService;

		/// <summary>
		/// public constructor
		/// </summary>
		/// <param name="logger">logger</param>
		/// <param name="topicService">topic service</param>
		public FtpClient(
			ILogger logger,
			ITopicService topicService)
		{
			_logger = logger;
			_topicService = topicService;
		}

		/// <summary>
		/// send file content to ftp
		/// </summary>
		/// <param name="filePath">file path</param>
		/// <returns>result</returns>	
		public async Task<bool> SendAsync(string filePath)
		{
			// FtpWebRequest request = (FtpWebRequest)WebRequest.Create("ftp://193.123.58.227:21");
			var uri = "ftp://193.123.58.227:21/files/" + Path.GetFileNameWithoutExtension(filePath) + ".zip"; //@"/" + "logs.zip";
			FtpWebRequest request = (FtpWebRequest)WebRequest.Create(uri);
			request.Method = WebRequestMethods.Ftp.UploadFile;

			// This example assumes the FTP site uses anonymous logon.
			request.Credentials = new NetworkCredential("mqttftp", "medtex");

			// Copy the contents of the file to the request stream.
			//byte[] fileContents;
			//using (StreamReader sourceStream = new StreamReader(filePath)) //  "testfile.txt"))
			//{
			//	fileContents = sourceStream.ReadBlock(fileContents, 0, )// Encoding.UTF8.GetBytes(sourceStream.ReadToEnd());
			//}

			//request.ContentLength = fileContents.Length;
			request.UsePassive = false;
			request.UseBinary = true;

			try
			{
				var sourceStream = File.OpenRead(filePath);
				using (Stream requestStream = request.GetRequestStream())
					{
						//requestStream.Write(fileContents, 0, fileContents.Length);
						var buffer = new byte[4096 * 2];
						int nRead = 0;
						while ((nRead = sourceStream.Read(buffer, 0, buffer.Length)) != 0)
						{
							requestStream.Write(buffer, 0, nRead);
						}
					}

				using (FtpWebResponse response = (FtpWebResponse)request.GetResponse())
				{
					Console.WriteLine($"Upload File Complete, status {response.StatusDescription}");
				}
			}
			catch (Exception ex)
			{
				_logger.Error(ex, "Ftp SendAsync error: ");
				return false;
			}

			return true;
		}
    }
}
