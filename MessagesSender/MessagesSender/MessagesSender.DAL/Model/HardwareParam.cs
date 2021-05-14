using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class HardwareParam
    {
        public int Id { get; set; }
        public string HardwareType { get; set; }
        public string ConnectionString { get; set; }
        public string Name { get; set; }
        public bool IsActive { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}