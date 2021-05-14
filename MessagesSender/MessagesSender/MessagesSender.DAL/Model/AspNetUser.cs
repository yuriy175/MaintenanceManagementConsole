using System;
using System.Collections.Generic;
using System.Text;

namespace MessagesSender.DAL.Model
{
    public class AspNetUser
    {
        public string Id { get; set; }
        public string UserName { get; set; }
        public string NormalizedUserName { get; set; }
        public string Email { get; set; }
        public string NormalizedEmail { get; set; }
        public string EmailConfirmed { get; set; }
        public bool PasswordHash { get; set; }
        public string SecurityStamp { get; set; }
        public string ConcurrencyStamp { get; set; }
        public string PhoneNumber { get; set; }
        public bool PhoneNumberConfirmed { get; set; }

        public bool TwoFactorEnabled { get; set; }
        public DateTime?  LockoutEnd { get; set; }

        public bool LockoutEnabled { get; set; }
        public int AccessFailedCount { get; set; }
        public string Language { get; set; }
        public string PersonInfo { get; set; }
        public DateTime? PasswordCreation { get; set; }
        public int PasswordExpirationPeriod { get; set; }

        public DateTime? Updated { get; set; }
        public bool Active { get; set; }
    }
}
