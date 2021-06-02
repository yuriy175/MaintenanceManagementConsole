using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Drawing;
using System.Runtime.InteropServices;
using System.Text;

namespace MessagesSender.MessagesSender.BL.Helpers
{
    /// <summary>
    /// WindowSnapshot helper
    /// </summary>
    internal class WindowSnapshotHelper
    {
        /// <summary>
        /// Make a window snapshot
        /// </summary>
        /// <param name="appWndHandle">window handler</param>
        /// <param name="isClientWnd">if window client</param>
        /// <param name="nCmdShow">show style</param>
        /// <returns>bitmap</returns>
        public static Bitmap MakeSnapshot(IntPtr appWndHandle, bool isClientWnd, Win32API.WindowShowStyle nCmdShow)
        {
            if (appWndHandle == IntPtr.Zero || !Win32API.IsWindow(appWndHandle) || !Win32API.IsWindowVisible(appWndHandle))
                return null;
            if (Win32API.IsIconic(appWndHandle))
                Win32API.ShowWindow(appWndHandle, nCmdShow);
            if (!Win32API.SetForegroundWindow(appWndHandle))
                return null;
            System.Threading.Thread.Sleep(1000);
            Win32API.RECT appRect;
            bool res = isClientWnd ? Win32API.GetClientRect(appWndHandle, out appRect) : Win32API.GetWindowRect(appWndHandle, out appRect);
            if (!res || appRect.Height == 0 || appRect.Width == 0)
            {
                return null;
            }

            if (isClientWnd)
            {
                Point lt = new Point(appRect.Left, appRect.Top);
                Point rb = new Point(appRect.Right, appRect.Bottom);
                Win32API.ClientToScreen(appWndHandle, ref lt);
                Win32API.ClientToScreen(appWndHandle, ref rb);
                appRect.Left = lt.X;
                appRect.Top = lt.Y;
                appRect.Right = rb.X;
                appRect.Bottom = rb.Y;
            }

            // Intersect with the Desktop rectangle and get what's visible
            IntPtr desktopHandle = Win32API.GetDesktopWindow();
            Win32API.RECT desktopRect;
            Win32API.GetWindowRect(desktopHandle, out desktopRect);
            Win32API.RECT visibleRect;
            if (!Win32API.IntersectRect(out visibleRect, ref desktopRect, ref appRect))
            {
                visibleRect = appRect;
            }

            if (Win32API.IsRectEmpty(ref visibleRect))
                return null;

            int width = visibleRect.Width;
            int height = visibleRect.Height;
            IntPtr hdcTo = IntPtr.Zero;
            IntPtr hdcFrom = IntPtr.Zero;
            IntPtr hBitmap = IntPtr.Zero;
            try
            {
                Bitmap clsRet = null;

                // get device context of the window...
                hdcFrom = isClientWnd ? Win32API.GetDC(appWndHandle) : Win32API.GetWindowDC(appWndHandle);

                // create dc that we can draw to...
                hdcTo = Win32API.CreateCompatibleDC(hdcFrom);
                hBitmap = Win32API.CreateCompatibleBitmap(hdcFrom, width, height);

                // validate...
                if (hBitmap != IntPtr.Zero)
                {
                    // copy...
                    int x = appRect.Left < 0 ? -appRect.Left : 0;
                    int y = appRect.Top < 0 ? -appRect.Top : 0;
                    IntPtr hLocalBitmap = Win32API.SelectObject(hdcTo, hBitmap);
                    Win32API.BitBlt(hdcTo, 0, 0, width, height, hdcFrom, x, y, Win32API.SRCCOPY);
                    Win32API.SelectObject(hdcTo, hLocalBitmap);

                    // create bitmap for window image...
                    clsRet = System.Drawing.Image.FromHbitmap(hBitmap);
                }

                return clsRet;
            }
            finally
            {
                // release ...
                if (hdcFrom != IntPtr.Zero)
                    Win32API.ReleaseDC(appWndHandle, hdcFrom);
                if (hdcTo != IntPtr.Zero)
                    Win32API.DeleteDC(hdcTo);
                if (hBitmap != IntPtr.Zero)
                    Win32API.DeleteObject(hBitmap);
            }
        }

        /// <summary>
        /// Return window handler
        /// </summary>
        /// <param name="proc">process</param>
        /// <returns>window window</returns>
        public static IntPtr GetWindowHandler(System.Diagnostics.Process proc)
        {
            var realWnd = IntPtr.Zero;
            var windowHandles = new List<IntPtr>();
            GCHandle listHandle = default(GCHandle);
            try
            {
                if (proc.MainWindowHandle == IntPtr.Zero)
                    throw new ApplicationException("Can't add a process with no MainFrame");

                Win32API.RECT maxRect = default(Win32API.RECT);
                if (IsValidUIWnd(proc.MainWindowHandle))
                {
                    realWnd = proc.MainWindowHandle;
                    return realWnd;
                }

                // the mainFrame is size == 0, so we look for the 'real' window
                listHandle = GCHandle.Alloc(windowHandles);
                foreach (ProcessThread pt in proc.Threads)
                {
                    Win32API.EnumThreadWindows((uint)pt.Id, new Win32API.EnumThreadDelegate(EnumThreadCallback), GCHandle.ToIntPtr(listHandle));
                }

                // get the biggest visible window in the current proc
                IntPtr maxHWnd = IntPtr.Zero;
                foreach (IntPtr hWnd in windowHandles)
                {
                    Win32API.RECT crtWndRect;

                    // do we have a valid rect for this window
                    if (Win32API.IsWindowVisible(hWnd) && Win32API.GetWindowRect(hWnd, out crtWndRect) &&
                        crtWndRect.Height > maxRect.Height && crtWndRect.Width > maxRect.Width)
                    {
                        // if the rect is outside the desktop, it's a dummy window
                        Win32API.RECT visibleRect;

                        // if (Win32API.IntersectRect(out visibleRect, ref _DesktopRect, ref CrtWndRect)
                        //    && !Win32API.IsRectEmpty(ref visibleRect))
                        {
                            maxHWnd = hWnd;
                            maxRect = crtWndRect;
                        }
                    }
                }

                if (maxHWnd != IntPtr.Zero && maxRect.Width > 0 && maxRect.Height > 0)
                {
                    realWnd = maxHWnd;
                }
                else
                {
                    realWnd = proc.MainWindowHandle;
                }

                return realWnd;
            }
            finally
            {
                if (listHandle != default(GCHandle) && listHandle.IsAllocated)
                    listHandle.Free();
            }
        }

        /// <summary>
        /// Checks if window handler valid
        /// </summary>
        /// <param name="hWnd">window handler</param>
        /// <returns>result</returns>
        internal static bool IsValidUIWnd(IntPtr hWnd)
        {
            bool res = false;
            if (hWnd == IntPtr.Zero || !Win32API.IsWindow(hWnd) || !Win32API.IsWindowVisible(hWnd))
                return false;
            Win32API.RECT crtWndRect;
            if (!Win32API.GetWindowRect(hWnd, out crtWndRect))
                return false;
            if (crtWndRect.Height > 0 && crtWndRect.Width > 0)
            {// a valid rectangle means the right window is the mainframe and it intersects the desktop
                Win32API.RECT visibleRect;

                // if the rectangle is outside the desktop, it's a dummy window
                // if (Win32API.IntersectRect(out visibleRect, ref _DesktopRect, ref CrtWndRect)
                //    && !Win32API.IsRectEmpty(ref visibleRect))
                res = true;
            }

            return res;
        }

        private static Image GetScreenShot(Point location, Size size)
        {
            IntPtr windowHandle = Win32API.GetDesktopWindow();
            Image myImage = new Bitmap(size.Width, size.Height);
            Graphics g = Graphics.FromImage(myImage);
            IntPtr destDeviceContext = g.GetHdc();
            IntPtr srcDeviceContext = Win32API.GetWindowDC(windowHandle);
            Win32API.BitBlt(destDeviceContext, 0, 0, size.Width, size.Height, srcDeviceContext, location.X, location.Y, Win32API.SRCCOPY);
            Win32API.ReleaseDC(windowHandle, srcDeviceContext);
            g.ReleaseHdc(destDeviceContext);
            return myImage;
        }

        private static bool EnumThreadCallback(IntPtr hWnd, IntPtr lParam)
        {
            GCHandle gch = GCHandle.FromIntPtr(lParam);
            List<IntPtr> list = gch.Target as List<IntPtr>;
            if (list == null)
            {
                throw new InvalidCastException("GCHandle Target could not be cast as List<IntPtr>");
            }

            list.Add(hWnd);
            return true;
        }
    }
}
