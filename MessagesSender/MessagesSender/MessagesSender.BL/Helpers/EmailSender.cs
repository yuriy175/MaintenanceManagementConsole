using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;
using MailKit.Net.Smtp;
using MessagesSender.Core.Interfaces;
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

        private readonly (string Smtp, string EmailFrom, string EmailTo, string Password) 
            _clientInfo =
                (
                      "smtp.mail.ru",
                      "yuriy_nv@mail.ru",
                      "sergey.nikiforov@mskorp.ru",
                      "psw");

        /// <summary>
        /// public constructor
        /// </summary>
        /// <param name="logger">logger</param>
        /// <param name="topicService">topic service</param>
        public EmailSender(
            ILogger logger,
            ITopicService topicService)
        {
            _logger = logger;
            _topicService = topicService;
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

            var emailMessage = CreateMessage(_clientInfo.EmailFrom, _clientInfo.EmailTo, $"Team Viewer", $"Team Viewer", tvPath);
            await SendEmailAsync(_clientInfo.Smtp, _clientInfo.EmailFrom, _clientInfo.Password, emailMessage);

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

        private async Task SendEmailAsync(string smtpAddress, string fromAddress, string fromPassword, MimeMessage emailMessage)
        {
            using (var client = new SmtpClient())
            {
                await client.ConnectAsync(smtpAddress, 25, false);
                await client.AuthenticateAsync(fromAddress, fromPassword);
                await client.SendAsync(emailMessage);

                await client.DisconnectAsync(true);
            }
        }
    }
}
