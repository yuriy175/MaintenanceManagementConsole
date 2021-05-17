using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Printer
    {
        public int Id { get; set; }

        [Column("printer_name")]
        public string PrinterName { get; set; }
        [Column("printer_status")]
        public string PrinterStatus { get; set; }
        [Column("printer_location")]
        public string PrinterLocation { get; set; }
        [Column("printer_port")]
        public string PrinterPort { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}