//------------------------------------------------------------------------------
// <auto-generated>
//     Этот код создан по шаблону.
//
//     Изменения, вносимые в этот файл вручную, могут привести к непредвиденной работе приложения.
//     Изменения, вносимые в этот файл вручную, будут перезаписаны при повторном создании кода.
// </auto-generated>
//------------------------------------------------------------------------------
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.ComponentModel.DataAnnotations.Schema;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class Detector
    {
        [Key]
        [Column("rowid")]
        public int RowId { get; set; }
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
