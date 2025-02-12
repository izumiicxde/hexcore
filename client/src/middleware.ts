import { NextResponse, type NextRequest } from "next/server";

export const middleware = (req: NextRequest) => {
  const token = req.cookies.get("token")?.value;
  if (!token) return NextResponse.redirect(new URL("/login", req.url));
};

export const config = {
  matcher: ["/", "/p/:path*"],
};
