using System;
using System.Collections.Generic;
using System.Linq;
using System.Security.AccessControl;
using System.Security.Principal;
using System.Threading.Tasks;
using Microsoft.Win32;

namespace MessagesSender.BL.Helpers
{
    /// <summary>
    /// Registry manager
    /// </summary>
    internal sealed class RegistryManager
    {
        /// <summary>
        /// Registry policies key
        /// </summary>
        public const string RegistryPoliciesKey = @"Software\Microsoft\Windows\CurrentVersion\Policies";

        /// <summary>
        /// Set policy if disable taskmanager
        /// </summary>
        /// <param name="disableTaskManager">if disable taskmanager</param>
        public static void SetPolicies(bool disableTaskManager)
        {
            try
            {
                var hive = RegistryKey.OpenBaseKey(
                    RegistryHive.CurrentUser, Environment.Is64BitOperatingSystem ? RegistryView.Registry64 : RegistryView.Registry32);
                var registryPoliciesSystemKey = RegistryPoliciesKey + @"\System";
                var baseKey = hive.OpenSubKey(registryPoliciesSystemKey, true);
                if (baseKey == null)
                {
                    var parentBaseKey = hive.OpenSubKey(RegistryPoliciesKey);
                    string user = Environment.UserDomainName + "\\" + Environment.UserName;
                    var registrySecurity = new RegistrySecurity();
                    var parentAccessorsNames = parentBaseKey.GetAccessControl()
                        .GetAccessRules(true, true, typeof(NTAccount))
                        .OfType<AuthorizationRule>()
                        .Where(a => a.IdentityReference.Value != user)
                        .Select(a => a.IdentityReference.Value).ToList();

                    parentAccessorsNames.Add(user);

                    parentAccessorsNames.ForEach(s =>
                    {
                        try
                        {
                            registrySecurity.AddAccessRule(
                                new RegistryAccessRule(
                                    s,
                                    RegistryRights.FullControl,
                                    InheritanceFlags.ContainerInherit,
                                    PropagationFlags.None,
                                    AccessControlType.Allow));
                        }
                        catch
                        {
                            // we don't need an accessor if we can't assign to it the full control
                        }
                    });

                    baseKey = Registry.CurrentUser.CreateSubKey(
                        registryPoliciesSystemKey,
                        RegistryKeyPermissionCheck.ReadWriteSubTree, RegistryOptions.None, registrySecurity);
                    if (baseKey == null)
                        return;
                }

                baseKey.SetValue("DisableTaskMgr", disableTaskManager ? 1 : 0, RegistryValueKind.DWord);
            }
            catch
            {
                return;
            }
        }
    }
}
