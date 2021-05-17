using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Modem
    {
        public int Id { get; set; }
        public string Operator { get; set; }
        public string DeviceName { get; set; }
        public string SerialNumber { get; set; }
        public string Imei { get; set; }
        public string workmode { get; set; }
        public string HardwareVersion { get; set; }
        public string SoftwareVersion { get; set; }
        public string WebUIVersion { get; set; }
        public int? CurrentConnectTime { get; set; }
        public int? CurrentUpload { get; set; }
        public int? CurrentDownload { get; set; }
        public int? CurrentDownloadRate { get; set; }
        public int? CurrentUploadRate { get; set; }
        public int? TotalUpload { get; set; }
        public int? TotalDownload { get; set; }
        public int? TotalConnectTime { get; set; }

        public string Rsrq { get; set; }
        public string Rsrp { get; set; }
        public string Rssi { get; set; }
        public string Sinr { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}