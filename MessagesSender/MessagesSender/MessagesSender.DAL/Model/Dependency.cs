using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Dependency
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        [Column("dep_name")]
        public string DepName { get; set; }
        [Column("dep_ver")]
        public string DepVer { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
