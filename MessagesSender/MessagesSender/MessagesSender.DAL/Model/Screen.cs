using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Screen
    {
        public int Id { get; set; }
        [Column("screen")]
        public string ScreenName { get; set; }
        public string Width { get; set; }
        public string Height { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}