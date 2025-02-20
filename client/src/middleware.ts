import { NextResponse, NextRequest } from "next/server";

const protectedRoutes = ["/", "/dashboard/", "/verify/"];

const isProtectedRoute = (pathname: string): boolean =>
  protectedRoutes.some((route) => pathname.startsWith(route));

export default function middleware(req: NextRequest) {
  try {
    const token = req.cookies.get("token")?.value;
    const isAuthenticated = Boolean(token);
    const { pathname } = req.nextUrl;

    if (!isAuthenticated) {
      if (pathname === "/login" || pathname === "/signup")
        return NextResponse.next();
      if (isProtectedRoute(pathname)) {
        return NextResponse.redirect(new URL("/login", req.url));
      }
    }

    if (isAuthenticated && (pathname === "/login" || pathname === "/signup")) {
      return NextResponse.redirect(new URL("/", req.url));
    }

    return NextResponse.next();
  } catch (error) {
    console.error("Middleware error:", error);
    return NextResponse.redirect(new URL("/login", req.url));
  }
}

export const config = {
  matcher: ["/", "/dashboard/:path*", "/login", "/signup", "/verify"],
};
