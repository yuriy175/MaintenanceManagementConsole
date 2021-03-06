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
    public class AspNetUser
    {
        [Key]
        [Column("rowid")]
        public int RowId { get; set; }
        public string Id { get; set; }
        public string UserName { get; set; }
        public string NormalizedUserName { get; set; }
        public string Email { get; set; }
        public string NormalizedEmail { get; set; }
        public bool? EmailConfirmed { get; set; }
        public string PasswordHash { get; set; }
        public string SecurityStamp { get; set; }
        public string ConcurrencyStamp { get; set; }
        public string PhoneNumber { get; set; }
        public bool? PhoneNumberConfirmed { get; set; }

        public bool? TwoFactorEnabled { get; set; }
        public string LockoutEnd { get; set; }

        public bool? LockoutEnabled { get; set; }
        public int? AccessFailedCount { get; set; }
        public string Language { get; set; }
        public string PersonInfo { get; set; }
        public string PasswordCreation { get; set; }
        public int? PasswordExpirationPeriod { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
