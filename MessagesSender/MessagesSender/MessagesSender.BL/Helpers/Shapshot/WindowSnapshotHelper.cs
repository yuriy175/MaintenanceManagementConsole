using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.Drawing;
using System.Runtime.InteropServices;
using System.Text;

namespace MessagesSender.MessagesSender.BL.Helpers
{
    internal class WindowSnapshotHelper
    {
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

        public static Bitmap MakeSnapshot(IntPtr AppWndHandle, bool IsClientWnd, Win32API.WindowShowStyle nCmdShow)
        {

            if (AppWndHandle == IntPtr.Zero || !Win32API.IsWindow(AppWndHandle) || !Win32API.IsWindowVisible(AppWndHandle))
                return null;
            if (Win32API.IsIconic(AppWndHandle))
                Win32API.ShowWindow(AppWndHandle, nCmdShow);//show it
            if (!Win32API.SetForegroundWindow(AppWndHandle))
                return null;//can't bring it to front
            System.Threading.Thread.Sleep(1000);//give it some time to redraw
            RECT appRect;
            bool res = IsClientWnd ? Win32API.GetClientRect(AppWndHandle, out appRect) : Win32API.GetWindowRect(AppWndHandle, out appRect);
            if (!res || appRect.Height == 0 || appRect.Width == 0)
            {
                return null;//some hidden window
            }
            if (IsClientWnd)
            {
                Point lt = new Point(appRect.Left, appRect.Top);
                Point rb = new Point(appRect.Right, appRect.Bottom);
                Win32API.ClientToScreen(AppWndHandle, ref lt);
                Win32API.ClientToScreen(AppWndHandle, ref rb);
                appRect.Left = lt.X;
                appRect.Top = lt.Y;
                appRect.Right = rb.X;
                appRect.Bottom = rb.Y;
            }
            //Intersect with the Desktop rectangle and get what's visible
            IntPtr DesktopHandle = Win32API.GetDesktopWindow();
            RECT desktopRect;
            Win32API.GetWindowRect(DesktopHandle, out desktopRect);
            RECT visibleRect;
            if (!Win32API.IntersectRect(out visibleRect, ref desktopRect, ref appRect))
            {
                visibleRect = appRect;
            }
            if (Win32API.IsRectEmpty(ref visibleRect))
                return null;

            int Width = visibleRect.Width;
            int Height = visibleRect.Height;
            IntPtr hdcTo = IntPtr.Zero;
            IntPtr hdcFrom = IntPtr.Zero;
            IntPtr hBitmap = IntPtr.Zero;
            try
            {
                Bitmap clsRet = null;

                // get device context of the window...
                hdcFrom = IsClientWnd ? Win32API.GetDC(AppWndHandle) : Win32API.GetWindowDC(AppWndHandle);

                // create dc that we can draw to...
                hdcTo = Win32API.CreateCompatibleDC(hdcFrom);
                hBitmap = Win32API.CreateCompatibleBitmap(hdcFrom, Width, Height);

                //  validate...
                if (hBitmap != IntPtr.Zero)
                {
                    // copy...
                    int x = appRect.Left < 0 ? -appRect.Left : 0;
                    int y = appRect.Top < 0 ? -appRect.Top : 0;
                    IntPtr hLocalBitmap = Win32API.SelectObject(hdcTo, hBitmap);
                    Win32API.BitBlt(hdcTo, 0, 0, Width, Height, hdcFrom, x, y, Win32API.SRCCOPY);
                    Win32API.SelectObject(hdcTo, hLocalBitmap);
                    //  create bitmap for window image...
                    clsRet = System.Drawing.Image.FromHbitmap(hBitmap);
                }
                //MessageBox.Show(string.Format("rect ={0} \n deskrect ={1} \n visiblerect = {2}",rct,drct,VisibleRCT));
                //  return...
                return clsRet;
            }
            finally
            {
                //  release ...
                if (hdcFrom != IntPtr.Zero)
                    Win32API.ReleaseDC(AppWndHandle, hdcFrom);
                if (hdcTo != IntPtr.Zero)
                    Win32API.DeleteDC(hdcTo);
                if (hBitmap != IntPtr.Zero)
                    Win32API.DeleteObject(hBitmap);
            }


        }

        public static IntPtr GetWindowHandler(System.Diagnostics.Process proc)
        {
            var realWnd = IntPtr.Zero;
            var windowHandles = new List<IntPtr>();
            GCHandle listHandle = default(GCHandle);
            try
            {
                if (proc.MainWindowHandle == IntPtr.Zero)
                    throw new ApplicationException("Can't add a process with no MainFrame");

                RECT MaxRect = default(RECT);//init with 0
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


                //get the biggest visible window in the current proc
                IntPtr MaxHWnd = IntPtr.Zero;
                foreach (IntPtr hWnd in windowHandles)
                {
                    RECT CrtWndRect;
                    //do we have a valid rect for this window
                    if (Win32API.IsWindowVisible(hWnd) && Win32API.GetWindowRect(hWnd, out CrtWndRect) &&
                        CrtWndRect.Height > MaxRect.Height && CrtWndRect.Width > MaxRect.Width)
                    {   //if the rect is outside the desktop, it's a dummy window
                        RECT visibleRect;
                        //if (Win32API.IntersectRect(out visibleRect, ref _DesktopRect, ref CrtWndRect)
                        //    && !Win32API.IsRectEmpty(ref visibleRect))
                        {
                            MaxHWnd = hWnd;
                            MaxRect = CrtWndRect;
                        }
                    }
                }
                if (MaxHWnd != IntPtr.Zero && MaxRect.Width > 0 && MaxRect.Height > 0)
                {
                    realWnd = MaxHWnd;
                }
                else
                    realWnd = proc.MainWindowHandle;//just add something even if it's a bad window

                return realWnd;
            }//try ends
            finally
            {
                if (listHandle != default(GCHandle) && listHandle.IsAllocated)
                    listHandle.Free();
            }

        }

        internal static bool IsValidUIWnd(IntPtr hWnd)
        {
            bool res = false;
            if (hWnd == IntPtr.Zero || !Win32API.IsWindow(hWnd) || !Win32API.IsWindowVisible(hWnd))
                return false;
            RECT CrtWndRect;
            if (!Win32API.GetWindowRect(hWnd, out CrtWndRect))
                return false;
            if (CrtWndRect.Height > 0 && CrtWndRect.Width > 0)
            {// a valid rectangle means the right window is the mainframe and it intersects the desktop
                RECT visibleRect;//if the rectangle is outside the desktop, it's a dummy window
                //if (Win32API.IntersectRect(out visibleRect, ref _DesktopRect, ref CrtWndRect)
                //    && !Win32API.IsRectEmpty(ref visibleRect))
                    res = true;
            }
            return res;
        }

        static bool EnumThreadCallback(IntPtr hWnd, IntPtr lParam)
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
