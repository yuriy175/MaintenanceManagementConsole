using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class HospitalInfo
    {
		[Key]
		[Column("rowid")]
		public int RowId { get; set; }		
		public int Id { get; set; }
        [Column("hospital_name")]
        public string HospitalName { get; set; }
        [Column("hospital_code")]
        public string HospitalCode { get; set; }
        [Column("hospital_address")]
        public string HospitalAddress { get; set; }
        [Column("arm_name")]
        public string ArmName { get; set; }
        [Column("complex_model")]
        public string ComplexModel { get; set; }
        [Column("serial_number")]
        public string SerialNumber { get; set; }
        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
