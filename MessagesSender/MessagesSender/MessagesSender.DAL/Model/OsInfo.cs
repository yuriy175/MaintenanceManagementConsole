using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class OsInfo
    {
        public int Id { get; set; }

        [Column("os_caption")]
        public string OsCaption { get; set; }
        [Column("os_version")]
        public string OsVersion { get; set; }
        [Column("build_number")]
        public string BuildNumber { get; set; }
        [Column("os_installed")]
        public string OsInstalled { get; set; }
        [Column("current_user")]
        public string CurrentUser { get; set; }
        public string Ip { get; set; }
        public string City { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
