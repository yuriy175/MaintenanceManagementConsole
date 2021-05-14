using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class AppParam
    {
        public int Id { get; set; }
        public string ParamName { get; set; }
        public string ParamValue { get; set; }
        public string Comment { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
