import { NextResponse, type NextRequest } from "next/server";

export const middleware = (req: NextRequest) => {
  const token = req.cookies.get("token")?.value;
  const pathname = req.nextUrl.pathname;

  const isAuthenticated = !!token; // Only checks if the token exists, no validation

  // **1. Redirect unauthenticated users away from protected pages**
  if (!isAuthenticated && ["/", "/p"].includes(pathname)) {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  // **2. Redirect authenticated users away from /login and /signup**
  if (isAuthenticated && ["/login", "/signup"].includes(pathname)) {
    return NextResponse.redirect(new URL("/", req.url));
  }

  return NextResponse.next();
};

export const config = {
  matcher: ["/", "/p/:path*", "/login", "/signup"],
};
