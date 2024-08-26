import type { Metadata } from "next";
import { Inter, Lato } from "next/font/google";
import "./globals.css";
import Navigation from "./components/Navigation";
import { DarkModeProvider } from "./components/Context/DarkModeContext";
import ProgressBar from "./components/loadingprovider/ProgressBar";
const inter = Inter({ subsets: ["latin"] });
const lato = Lato({
  subsets: ["latin"],
  weight: ["400", "700"],
});

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
      <body className={`${inter.className} ${lato.className} flex w-full`}>
        <ProgressBar>
          <DarkModeProvider>
            <Navigation>{children}</Navigation>
          </DarkModeProvider>
        </ProgressBar>
      </body>
    </html>
  );
}
