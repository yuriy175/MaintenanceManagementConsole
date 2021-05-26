using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class RasterParam
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        public int? Code { get; set; }
        public string Name { get; set; }
        public int? FocalDistance { get; set; }
        public int? FocalDistanceMin { get; set; }
        public int? FocalDistanceMax { get; set; }
        public bool IsEnabled { get; set; }
        public string LocalWSName { get; set; }
        public string ConnectionString { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
