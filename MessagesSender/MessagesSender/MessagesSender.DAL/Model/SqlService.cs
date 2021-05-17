﻿using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class SqlService
    {
        public int Id { get; set; }
        [Column("sql_name")]
        public string SqlName { get; set; }
        [Column("sql_version")]
        public string SqlVersion { get; set; }
        [Column("sql_status")]
        public string SqlStatus { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
