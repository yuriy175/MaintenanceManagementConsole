using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class AppParam
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        public string ParamName { get; set; }
        public string ParamValue { get; set; }
        public string Comment { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
