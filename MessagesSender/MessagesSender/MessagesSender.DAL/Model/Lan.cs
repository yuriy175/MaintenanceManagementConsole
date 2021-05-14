using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Lan
    {
        public string Id { get; set; }
        public string Adapter { get; set; }
        public string Ip { get; set; }
        [Column("default_gateway")]
        public string DefaultGateway { get; set; }
        public string Dns { get; set; }
        public string Status { get; set; }
        [Column("link_speed")]
        public string LinkSpeed { get; set; }
        [Column("interface_description")]
        public string InterfaceDescription { get; set; }
        public string PrefixOrigin { get; set; }
        public string SuffixOrigin { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
