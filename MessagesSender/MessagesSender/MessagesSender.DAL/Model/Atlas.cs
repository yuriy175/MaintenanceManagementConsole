using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class AtlasSW
    {
        public int Id { get; set; }
        [Column("atlas_version")]
        public string AtlasVersion { get; set; }
        [Column("complex_type")]
        public string ComplexType { get; set; }
        public string Language { get; set; }
        public string Multimonitor { get; set; }
        [Column("xilibs_version")]
        public string XilibsVersion { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
