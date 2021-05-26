using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Monitor
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        [Column("monitor")]
        public string MonitorName { get; set; }
        [Column("serial_number")]
        public string SerialNumber { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
