using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class HardDrive
    {
        public string Id { get; set; }
        [Column("device_id")]
        public string DeviceId { get; set; }
        [Column("friendly_name")]
        public string FriendlyName { get; set; }
        [Column("serial_number")]
        public string SerialNumber { get; set; }
        public string mediatype { get; set; }
        [Column("operational_status")]
        public string OperationalStatus { get; set; }
        [Column("health_status")]
        public string HealthStatus { get; set; }
        [Column("size_gb")]
        public string SizeGb { get; set; }
        public int? PowerOnHours { get; set; }
        public int? Temperature { get; set; }
        public int? StartStopCycleCount { get; set; }
        public int? ReadErrorsCorrected { get; set; }
        public int? ReadErrorsUncorrected { get; set; }
        public int? ReadErrorsTotal { get; set; }
        public int? WriteErrorsCorrected { get; set; }
        public int? WriteErrorsUncorrected { get; set; }
        public int? WriteErrorsTotal { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
