using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class DetectorProcessing
    {
        public int Id { get; set; }
        public string ProcessingFilePath { get; set; }
        public decimal? ProcessingType { get; set; }
        public decimal? DetectorId { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}