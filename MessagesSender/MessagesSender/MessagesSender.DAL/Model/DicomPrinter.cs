using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class DicomPrinter
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }
		public int Id { get; set; }
        public Nullable<int> DicomServicesId { get; set; }
        public string Name { get; set; }
        public int Type { get; set; }
        public int Bits { get; set; }

        public string BorderDensity { get; set; }
        public string EmptyImageDensity { get; set; }
        public int MinDensity { get; set; }
        public int MaxDensity { get; set; }

        public string MagnificationType { get; set; }
        public string SmoothingType { get; set; }
        public string PTrim { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
