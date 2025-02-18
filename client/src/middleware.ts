import { NextResponse, NextRequest } from "next/server";

const secret = process.env.DB_JWT_SECRET!;

export async function middleware(req: NextRequest) {
  try {
    const token = req.cookies.get("token");
    const isAuthenticated = !!token?.value;
    const { pathname } = req.nextUrl;

    if (
      !isAuthenticated &&
      (pathname === "/" || pathname.startsWith("/class/"))
    ) {
      return NextResponse.redirect(new URL("/login", req.url));
    }

    if (
      isAuthenticated &&
      (pathname.startsWith("/login") || pathname.startsWith("/signup"))
    ) {
      return NextResponse.redirect(new URL("/", req.url));
    }

    return NextResponse.next();
  } catch {
    return NextResponse.redirect(new URL("/login", req.url));
  }
}

export const config = {
  matcher: ["/", "/class/:path*", "/login", "/signup"], // Protected routes
};
