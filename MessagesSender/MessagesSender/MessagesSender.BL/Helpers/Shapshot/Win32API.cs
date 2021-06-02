using System;
using System.Collections.Generic;
using System.Drawing;
using System.Runtime.InteropServices;
using System.Text;

namespace MessagesSender.MessagesSender.BL.Helpers
{
    /// <summary>
    /// Win32API wrapper
    /// </summary>
    internal static class Win32API
    {
        /// <summary>
        /// WM_SYSCOMMAND
        /// </summary>
        internal const int WM_SYSCOMMAND = 0x112;

        /// <summary>
        /// MF_SEPARATOR
        /// </summary>
        internal const int MF_SEPARATOR = 0x800;

        /// <summary>
        /// MF_STRING
        /// </summary>
        internal const int MF_STRING = 0x0;

        /// <summary>
        /// SRCCOPY
        /// </summary>
        internal const int SRCCOPY = 13369376;

		/// <summary>
		/// enumerates threads
		/// </summary>
		/// <param name="hWnd">window handler</param>
		/// <param name="lParam">lParam</param>
		/// <returns>result</returns>
		internal delegate bool EnumThreadDelegate(IntPtr hWnd, IntPtr lParam);

		/// <summary>Enumeration of the different ways of showing a window using 
		/// ShowWindow</summary>
		internal enum WindowShowStyle : uint
		{
			/// <summary>Hides the window and activates another window.</summary>
			/// <remarks>See SW_HIDE</remarks>
			Hide = 0,

			/// <summary>Activates and displays a window. If the window is minimized 
			/// or maximized, the system restores it to its original size and 
			/// position. An application should specify this flag when displaying 
			/// the window for the first time.</summary>
			/// <remarks>See SW_SHOWNORMAL</remarks>
			ShowNormal = 1,

			/// <summary>Activates the window and displays it as a minimized window.</summary>
			/// <remarks>See SW_SHOWMINIMIZED</remarks>
			ShowMinimized = 2,

			/// <summary>Activates the window and displays it as a maximized window.</summary>
			/// <remarks>See SW_SHOWMAXIMIZED</remarks>
			ShowMaximized = 3,

			/// <summary>Maximizes the specified window.</summary>
			/// <remarks>See SW_MAXIMIZE</remarks>
			Maximize = 3,

			/// <summary>Displays a window in its most recent size and position. 
			/// This value is similar to "ShowNormal", except the window is not 
			/// actived.</summary>
			/// <remarks>See SW_SHOWNOACTIVATE</remarks>
			ShowNormalNoActivate = 4,

			/// <summary>Activates the window and displays it in its current size 
			/// and position.</summary>
			/// <remarks>See SW_SHOW</remarks>
			Show = 5,

			/// <summary>Minimizes the specified window and activates the next 
			/// top-level window in the Z order.</summary>
			/// <remarks>See SW_MINIMIZE</remarks>
			Minimize = 6,

			/// <summary>Displays the window as a minimized window. This value is 
			/// similar to "ShowMinimized", except the window is not activated.</summary>
			/// <remarks>See SW_SHOWMINNOACTIVE</remarks>
			ShowMinNoActivate = 7,

			/// <summary>Displays the window in its current size and position. This 
			/// value is similar to "Show", except the window is not activated.</summary>
			/// <remarks>See SW_SHOWNA</remarks>
			ShowNoActivate = 8,

			/// <summary>Activates and displays the window. If the window is 
			/// minimized or maximized, the system restores it to its original size 
			/// and position. An application should specify this flag when restoring 
			/// a minimized window.</summary>
			/// <remarks>See SW_RESTORE</remarks>
			Restore = 9,

			/// <summary>Sets the show state based on the SW_ value specified in the 
			/// STARTUPINFO structure passed to the CreateProcess function by the 
			/// program that started the application.</summary>
			/// <remarks>See SW_SHOWDEFAULT</remarks>
			ShowDefault = 10,

			/// <summary>Windows 2000/XP: Minimizes a window, even if the thread 
			/// that owns the window is hung. This flag should only be used when 
			/// minimizing windows from a different thread.</summary>
			/// <remarks>See SW_FORCEMINIMIZE</remarks>
			ForceMinimized = 11
		}

		/// <summary>
		/// rectangle struct
		/// </summary>
		[StructLayout(LayoutKind.Sequential)]
        internal struct RECT
        {
            /// <summary>
            /// Left
            /// </summary>
            internal int Left;

            /// <summary>
            /// Top
            /// </summary>
            internal int Top;

            /// <summary>
            /// Right
            /// </summary>
            internal int Right;

            /// <summary>
            /// Bottom
            /// </summary>
            internal int Bottom;

            /// <summary>
            /// Width
            /// </summary>
            internal int Width
            {
                get { return Math.Abs(Right - Left); }
            }

            /// <summary>
            /// Height
            /// </summary>
            internal int Height
            {
                get { return Math.Abs(Bottom - Top); }
            }

            /// <summary>
            /// string presentation
            /// </summary>
            /// <returns>string</returns>
            public override string ToString()
            {
                return string.Format(
                    "Left = {0}, Top = {1}, Right = {2}, Bottom ={3}",
                    Left, Top, Right, Bottom);
            }
        }

		/// <summary>
		/// Get window class name
		/// </summary>
		/// <param name="hWnd">window handler</param>
		/// <param name="param">param</param>
		/// <param name="length">length</param>
		[DllImport("User32.Dll")]
        internal static extern void GetClassName(IntPtr hWnd, System.Text.StringBuilder param, int length);

		/// <summary>
		/// Get window text length
		/// </summary>
		/// <param name="hWnd">window handler</param>
		/// <returns>length</returns>
		[DllImport("user32.dll", SetLastError = true, CharSet = CharSet.Auto)]
        internal static extern int GetWindowTextLength(IntPtr hWnd);

		/// <summary>
		/// Get window text
		/// </summary>
		/// <param name="hWnd">window handler</param>
		/// <param name="lpString">lpString</param>
		/// <param name="nMaxCount">nMaxCount</param>
		/// <returns>length</returns>
		[DllImport("user32.dll", CharSet = CharSet.Auto, SetLastError = true)]
        internal static extern int GetWindowText(IntPtr hWnd, StringBuilder lpString, int nMaxCount);

        [DllImport("user32.dll")]
        internal static extern bool EnumThreadWindows(uint dwThreadId, Win32API.EnumThreadDelegate lpfn, IntPtr lParam);

        [DllImport("user32.dll")]
        internal static extern IntPtr GetSystemMenu(IntPtr hWnd, bool bRevert);

        [DllImport("user32.dll")]
        internal static extern bool AppendMenu(IntPtr hMenu, int wFlags, int wIDNewItem, string lpNewItem);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool ShowWindow(IntPtr hWnd, WindowShowStyle nCmdShow);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool IsRectEmpty([In] ref RECT lprc);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool ClientToScreen(IntPtr hwnd, ref Point lpPoint);

        [DllImport("user32.dll", SetLastError = true)]
        internal static extern IntPtr FindWindow(string lpClassName, string lpWindowName);

        [DllImport("user32.dll", EntryPoint = "GetDC")]
        internal static extern IntPtr GetDC(IntPtr hWnd);

        [DllImport("user32.dll", EntryPoint = "ReleaseDC")]
        internal static extern IntPtr ReleaseDC(IntPtr hWnd, IntPtr hDc);

        [DllImport("gdi32.dll", EntryPoint = "DeleteDC")]
        internal static extern IntPtr DeleteDC(IntPtr hDc);

        [DllImport("gdi32.dll", EntryPoint = "DeleteObject")]
        internal static extern IntPtr DeleteObject(IntPtr hDc);

        [DllImport("gdi32.dll", EntryPoint = "BitBlt")]
        internal static extern bool BitBlt(IntPtr hdcDest, int xDest, int yDest, int wDest, int hDest, IntPtr hdcSource, int xSrc, int ySrc, int rasterOp);

        [DllImport("gdi32.dll", EntryPoint = "CreateCompatibleBitmap")]
        internal static extern IntPtr CreateCompatibleBitmap(IntPtr hdc, int nWidth, int nHeight);

        [DllImport("gdi32.dll", EntryPoint = "CreateCompatibleDC")]
        internal static extern IntPtr CreateCompatibleDC(IntPtr hdc);

        [DllImport("gdi32.dll", EntryPoint = "SelectObject")]
        internal static extern IntPtr SelectObject(IntPtr hdc, IntPtr bmp);

        [DllImport("user32.dll", SetLastError = false)]
        internal static extern IntPtr GetDesktopWindow();

        [DllImport("user32.dll")]
        internal static extern IntPtr GetWindowDC(IntPtr hWnd);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool SetForegroundWindow(IntPtr hWnd);

        [DllImport("User32.Dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool IsIconic(IntPtr hWnd);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool IsWindow(IntPtr hWnd);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool GetWindowRect(IntPtr hWnd, out RECT lpRect);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool IsWindowVisible(IntPtr hWnd);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool GetClientRect(IntPtr hWnd, out RECT lpRect);

        [DllImport("user32.dll")]
        [return: MarshalAs(UnmanagedType.Bool)]
        internal static extern bool IntersectRect(out RECT lprcDst, [In] ref RECT lprcSrc1,
           [In] ref RECT lprcSrc2);
    }
}
