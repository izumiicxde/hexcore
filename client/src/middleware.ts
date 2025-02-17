import { NextResponse, NextRequest } from "next/server";
import jwt from "jsonwebtoken";

const secret = process.env.DB_JWT_SECRET!;

export async function middleware(req: NextRequest) {
  try {
    const token = req.cookies.get("token");
    if (!token) {
      return NextResponse.redirect(new URL("/login", req.url));
    }

    const valid = !!token;

    const { pathname } = req.nextUrl;
    if (pathname.startsWith("/login") && !valid) {
      return NextResponse.next();
    }
    if (
      (pathname.startsWith("/") || pathname.startsWith("/class/")) &&
      !valid
    ) {
      return NextResponse.redirect(new URL("/login", req.url));
    }
    return NextResponse.next(); // Allow the request to proceed
  } catch (err) {
    return NextResponse.redirect(new URL("/login", req.url));
  }
}

export const config = {
  matcher: ["/", "/class/:path*"], // Protected routes
};
