import type { Metadata } from "next";
import "./globals.css";
import Navigation from "./components/Navigation"
export const metadata: Metadata = {
  title: "A2SV Wallet",
  description: "Built for a2sv",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className="flex">
        <Navigation>{children}</Navigation>
      </body>
    </html>
  );
}
