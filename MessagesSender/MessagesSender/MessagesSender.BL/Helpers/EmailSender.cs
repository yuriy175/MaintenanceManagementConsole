using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using Atlas.Common.Core.Interfaces;
using MailKit.Net.Smtp;
using MailKit.Security;
using MessagesSender.BL.BusWrappers.Helpers;
using MessagesSender.Core.Interfaces;
using MessagesSender.Core.Model;
using MessagesSender.MessagesSender.BL.Helpers;
using MimeKit;
using Serilog;

namespace MessagesSender.BL.Helpers
{
    /// <summary>
    /// email service
    /// </summary>
    internal class EmailSender : IEmailSender
    {
        private readonly ILogger _logger;
        private readonly ITopicService _topicService;
        private readonly IConfigurationService _configurationService;

        private (string Smtp, string EmailFrom, string EmailTo, string Login, string Password)?
            _clientInfo =
                (
                      "mail.vko-medprom.ru", // "smtp.mail.ru",
                      "ars@vko-medprom.ru", // "yuriy_nv@mail.ru",
                      "sergey.nikiforov@mskorp.ru",
                      "tech.ars",
                      "lV0uwMU1kFI5lrS5");

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="configurationService">configuration service</param>
        /// <param name="logger">logger</param>
        /// <param name="topicService">topic service</param>
        public EmailSender(
            IConfigurationService configurationService,
            ILogger logger,
            ITopicService topicService)
        {
            _configurationService = configurationService;
            _logger = logger;
            _topicService = topicService;

            CreateConnectionProps();
        }

        /// <summary>
        /// sends teamviewer file
        /// </summary>
        /// <param name="tvPath">file path</param>
        /// <returns>result</returns>
        public async Task<bool> SendTeamViewerAsync(string tvPath)
        {
            if (!File.Exists(tvPath))
            {
                _logger.Error($"SendTeamViewerAsync error : {tvPath} doesn't exists");
            }

            if (!_clientInfo.HasValue)
            {
                return false;
            }

            var clientInfo = _clientInfo.Value;
            var emailMessage = CreateMessage(clientInfo.EmailFrom, clientInfo.EmailTo, $"Team Viewer", $"Team Viewer", tvPath);
            await SendEmailAsync(clientInfo.Smtp, clientInfo.Login, clientInfo.Password, emailMessage);

            return true;
        }

        private MimeMessage CreateMessage(string emailFrom, string emailTo, string subject, string body, string logPath)
        {
            var emailMessage = new MimeMessage();

            emailMessage.From.Add(new MailboxAddress("MMC", emailFrom));
            emailMessage.To.Add(new MailboxAddress(string.Empty, emailTo));
            emailMessage.Subject = subject;

            var builder = new BodyBuilder();

            // Set the plain-text version of the message text
            builder.TextBody = body;

            // We may also want to attach a calendar event for Monica's party...
            builder.Attachments.Add(logPath);

            // Now we just need to set the message body and we're done
            emailMessage.Body = builder.ToMessageBody();

            return emailMessage;
        }

        private async Task SendEmailAsync(string smtpAddress, string fromLogin, string fromPassword, MimeMessage emailMessage)
        {
            using (var client = new SmtpClient())
            {
                // await client.ConnectAsync(smtpAddress, 25, false);
                // await client.ConnectAsync(smtpAddress, 465, true);
                client.Connect(
                    smtpAddress, // int port = 0, 
                    options: SecureSocketOptions.StartTls);
                await client.AuthenticateAsync(fromLogin, fromPassword);
                await client.SendAsync(emailMessage);

                await client.DisconnectAsync(true);
            }
        }

        private void CreateConnectionProps()
        {
            var connectionString = _configurationService.Get<string>(Constants.EmailConnectionStringName, null);
            try
            {
                _clientInfo = ConnectionPropsCreator.CreateEmailProps(connectionString);
            }
            catch (Exception ex)
            {
                _logger.Error(ex, "Ftp client wrong connection string");
            }
        }
    }
}
