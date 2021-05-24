using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class News
    {
        public int Id { get; set; }
        public string Tbl { get; set; }
        public int RowId { get; set; }
        public string Type { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}