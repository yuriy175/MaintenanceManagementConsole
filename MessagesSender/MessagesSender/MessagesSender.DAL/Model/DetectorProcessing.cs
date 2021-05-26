using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class DetectorProcessing
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        public string ProcessingFilePath { get; set; }
        public decimal? ProcessingType { get; set; }
        public decimal? DetectorId { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
