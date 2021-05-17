using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class SqlDatabase
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public string Access { get; set; }
        public string Status { get; set; }
        public string Backup { get; set; }
        public string Data { get; set; }
        public string Log { get; set; }
        public string Compability { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}