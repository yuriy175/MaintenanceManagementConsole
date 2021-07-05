using System;
using System.IO;
using System.Net;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using MessagesSender.BL.BusWrappers.Helpers;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using Serilog;

namespace MessagesSender.BL.Remoting
{
    /// <summary>
    /// ftp client
    /// </summary>
    internal class FtpClient : IFtpClient
    {
        private readonly IConfigurationService _configurationService;
        private readonly ILogger _logger;
        private readonly ITopicService _topicService;

        private (string HostName, string UserName, string Password)? _connectionProps = null;

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="topicService">topic service</param>
        public FtpClient(
            IConfigurationService configurationService,
            ILogger logger,
            ITopicService topicService)
        {
            _configurationService = configurationService;
            _logger = logger;
            _topicService = topicService;

            _configurationService.AddConfigFile(
                Path.Combine(
                    Path.GetDirectoryName(
                        typeof(IFtpClient).Assembly.Location), "ftpsettings.json"));

            CreateConnectionProps();
        }

        /// <summary>
        /// send file content to ftp
        /// </summary>
        /// <param name="filePath">file path</param>
        /// <param name="destFileName">destination file name</param>
        /// <returns>result</returns>    
        public async Task<bool> SendAsync(string filePath, string destFileName = null)
        {
            if (!_connectionProps.HasValue)
            {
                return false;
            }

            var uri = _connectionProps.Value.HostName +
                Path.GetFileNameWithoutExtension(string.IsNullOrEmpty(destFileName) ? filePath : destFileName) +
                ".zip"; 
            FtpWebRequest request = (FtpWebRequest)WebRequest.Create(uri);
            request.Method = WebRequestMethods.Ftp.UploadFile;

            // This example assumes the FTP site uses anonymous logon.
            // request.Credentials = new NetworkCredential("mqttftp", "medtex");
            request.Credentials = new NetworkCredential(_connectionProps.Value.UserName, _connectionProps.Value.Password);
            request.UsePassive = false;
            request.UseBinary = true;

            try
            {
                using (var sourceStream = File.OpenRead(filePath))
                {
                    using (Stream requestStream = request.GetRequestStream())
                    {
                        var buffer = new byte[4096 * 2];
                        int nRead = 0;
                        while ((nRead = sourceStream.Read(buffer, 0, buffer.Length)) != 0)
                        {
                            requestStream.Write(buffer, 0, nRead);
                        }
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

        private void CreateConnectionProps()
        {
            var connectionString = _configurationService.Get<string>(Constants.FtpClientConnectionStringName, null);
            try
            {
                _connectionProps = ConnectionPropsCreator.CreateFtpProps(connectionString);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "Ftp client wrong connection string");
            }
        }
    }
}
