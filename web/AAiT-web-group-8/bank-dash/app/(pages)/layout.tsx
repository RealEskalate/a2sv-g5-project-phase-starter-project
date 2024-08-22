import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "../globals.css";

import Header from "@/components/dashboard/header";
import TableUI from "@/components/stats/table";
import Paginate from "@/components/stats/paginate";
import {BarChartt} from "@/components/stats/barchart";
import Sidebar from "@/components/dashboard/sidebar";

const inter = Inter({ subsets: ["latin"] });

const exampleData = [
    { Description: 'Paid', TransactionID: 'Credit Card', Type: 'Paid', Card: 'Credit Card', Date: 'Credit Card', amount: 250.00, Receipt: 'Paid'},
    { Description: 'paid', TransactionID: 'Credit Card', Type: 'paid', Card: 'Credit Card', Date: 'Credit Card', amount: 150.00, Receipt: 'paid'},
    { Description: 'Paid', TransactionID: 'Credit Card', Type: 'Paid', Card: 'Credit Card', Date: 'Credit Card', amount: -250.00, Receipt: 'Paid' },
    { Description: 'paid', TransactionID: 'Credit Card', Type: 'paid', Card: 'Credit Card', Date: 'Credit Card', amount: 150.00, Receipt: 'paid'},
    { Description: 'Paid', TransactionID: 'Credit Card', Type: 'Paid', Card: 'Credit Card', Date: 'Credit Card', amount: -250.00, Receipt: 'Paid' },
];

export const metadata: Metadata = {
  title: "Create Next App",
  description: "Generated by create next app",
};

export default function RootLayout({children}: Readonly<{children: React.ReactNode}>) {
  return (
    <html lang="en">
      <body className={`${inter.className} flex flex-row h-screen`}>
        <Sidebar className="w-64 bg-gray-800 text-white" />
        <div className="flex flex-col w-full">
          <Header className="w-full bg-gray-700 text-white p-4" />
          <main className="flex-1 p-6 bg-slate-200 overflow-auto">
            <TableUI Data={exampleData} />
            <Paginate />
            <BarChartt />
            {children}
          </main>
        </div>
      </body>
    </html>
  )
}

