using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class VideoAdapter
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        [Column("card_name")]
        public string CardName { get; set; }
        [Column("memory_gb")]
        public decimal MemoryGb { get; set; }
        [Column("card_status")]
        public string CardStatus { get; set; }
        [Column("drv_date")]
        public string DrvDate { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
