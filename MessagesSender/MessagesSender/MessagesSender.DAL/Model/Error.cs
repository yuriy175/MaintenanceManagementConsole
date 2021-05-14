using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Error
    {
        public int Id { get; set; }
        [Column("error")]
        public string ErrorDescription { get; set; }
        [Column("table_name")]
        public string TableName { get; set; }
    }
}
