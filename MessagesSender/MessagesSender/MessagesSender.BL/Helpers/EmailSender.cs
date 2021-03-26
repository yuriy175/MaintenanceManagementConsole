using MailKit.Net.Smtp;
using MimeKit;
using Serilog;
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Threading.Tasks;

namespace MessagesSender.BL.Helpers
{
    internal sealed class EmailSender
    {
        /// <summary>
        /// sends file
        /// </summary>
        /// <param name="logPath">log file path</param>
        /// <returns>result</returns>
        public static async Task<bool> SendAtlasLogsAsync(string logPath)
        {
            (string Smtp, string EmailFrom, string EmailTo, string Password) clientInfo =
                (
                      "smtp.mail.ru",
                      "yuriy_nv@mail.ru",
                      // "yuri.vorobyev@mskorp.ru",
                      "sergey.nikiforov@mskorp.ru",
                      "wrong_psw");

            var emailMessage = CreateMessage(clientInfo.EmailFrom, clientInfo.EmailTo, $"Логи Атлас", $"Логи Атлас", logPath);
            await SendEmailAsync(clientInfo.Smtp, clientInfo.EmailFrom, clientInfo.Password, emailMessage);

            return true;
        }

        private static MimeMessage CreateMessage(string emailFrom, string emailTo, string subject, string body, string logPath)
        {
            var emailMessage = new MimeMessage();

            emailMessage.From.Add(new MailboxAddress("MMC", emailFrom));
            emailMessage.To.Add(new MailboxAddress("", emailTo));
            emailMessage.Subject = subject;
            //emailMessage.Body = new TextPart(MimeKit.Text.TextFormat.Html)
            //{
            //    Text = body,
            //};
            // create our message text, just like before (except don't set it as the message.Body)

            var builder = new BodyBuilder();

            // Set the plain-text version of the message text
            builder.TextBody = body;

            // We may also want to attach a calendar event for Monica's party...
            builder.Attachments.Add(logPath);

            // Now we just need to set the message body and we're done
            emailMessage.Body = builder.ToMessageBody();

            return emailMessage;
        }

        private static async Task SendEmailAsync(string smtpAddress, string fromAddress, string fromPassword, MimeMessage emailMessage)
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
