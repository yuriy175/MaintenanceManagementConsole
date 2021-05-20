using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class DicomService
    {
        public int Id { get; set; }
        public string LogicalName { get; set; }
        public string AeTitle { get; set; }
        public string IPAddress { get; set; }
        public string Port { get; set; }
        public int PduSize { get; set; }
        public int Timeout { get; set; }
        public Nullable<int> ListenPort { get; set; }
        public int Status { get; set; }
        public int ServiceRole { get; set; }
        public int TransferSyntaxId { get; set; }
        public bool EnableLog { get; set; }
        public string FilePath { get; set; }
        public Nullable<int> AssocNumber { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
