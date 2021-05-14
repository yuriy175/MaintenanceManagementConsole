using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Detector
    {
        public int Id { get; set; }
        public string Name { get; set; }
        public bool IsUse { get; set; }
        public decimal DetectorType { get; set; }
        public string UniqueName { get; set; }
        public string ImagePixelSpacing { get; set; }
        public string CalibratedImageSize { get; set; }
        public string Version { get; set; }
        public string ConnectionString { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}