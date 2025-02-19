import { NextResponse, NextRequest } from "next/server";

const protectedRoutes = ["/", "/class/", "/verify"];

const isProtectedRoute = (pathname: string): boolean =>
  protectedRoutes.some((route) => pathname.startsWith(route));

export async function authMiddleware(req: NextRequest) {
  try {
    const token = req.cookies.get("token");
    const isAuthenticated = Boolean(token?.value);
    const { pathname } = req.nextUrl;

    if (!isAuthenticated && isProtectedRoute(pathname)) {
      return NextResponse.redirect(new URL("/login", req.url));
    }

    if (
      isAuthenticated &&
      ["/login", "/signup"].some((route) => pathname.startsWith(route))
    ) {
      return NextResponse.redirect(new URL("/", req.url));
    }

    return NextResponse.next();
  } catch (error) {
    console.error("Middleware error:", error);
    return NextResponse.redirect(new URL("/login", req.url));
  }
}

export const config = {
  matcher: ["/", "/class/:path*", "/login", "/signup"],
};
