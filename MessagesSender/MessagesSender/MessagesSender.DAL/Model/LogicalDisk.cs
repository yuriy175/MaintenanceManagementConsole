using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class LogicalDisk
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        public string Letter { get; set; }
        [Column("total_size")]
        public decimal TotalSize { get; set; }
        [Column("free_size")]
        public decimal FreeSize { get; set; }
        [Column("volume_name")]
        public string VolumeName { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
