export default function AuthLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex justify-center items-center overflow-hidden h-screen w-screen p-5">
      {children}
    </div>
  );
}
