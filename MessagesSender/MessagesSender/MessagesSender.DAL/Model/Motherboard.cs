using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Motherboard
    {
        public int Id { get; set; }
        public string Mboard { get; set; }
        public string Cpu { get; set; }
        [Column("cpu_status")]
        public string CpuStatus { get; set; }
        [Column("memory_quantity")]
        public decimal MemoryQuantity { get; set; }
        [Column("memory_total")]
        public decimal MemoryTotal { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}