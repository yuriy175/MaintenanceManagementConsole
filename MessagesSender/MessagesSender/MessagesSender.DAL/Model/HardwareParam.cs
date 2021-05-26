using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class HardwareParam
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        public string HardwareType { get; set; }
        public string ConnectionString { get; set; }
        public string Name { get; set; }
        public bool IsActive { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
