"use client";
import BalanceCard from "@/components/AccountSmallCard";
import LastTransactionCard from "@/components/LastTransactionCard";
import DesktopCreditCard from "@/components/DesktopCreditCard";
import InvoicesCard from "@/components/InvoicesCard";
import AccountBarChart from "@/components/AccountBarChart";
import Link from "next/link";

const Accounts = () => {
  return (
    <div className="flex dark:bg-dark text-gray-900 dark:text-white">
      {/* Sidebar */}
      <div className="hidden lg:block w-64 bg-white h-screen fixed top-0 left-0">
        {/* Your Sidebar content goes here */}
      </div>

      {/* Main content */}
      <div className="flex-1 lg:ml-64 p-4 sm:p-8 bg-gray-100 dark:bg-dark text-gray-900 dark:text-white">
        {/* Top Section */}
        <div className="mb-8">
          <h1 className="text-2xl font-semibold mb-6 dark:text-blue-500">Accounts</h1>
          {/* <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-8"> */}
            <BalanceCard />
          {/* </div> */}
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-10 gap-4 sm:gap-8 mb-8">
          <div className="lg:col-span-7 flex flex-col">
            <h2 className="text-lg font-semibold mb-3 dark:text-blue-500">Last Transaction</h2>
            <div className="flex-1 flex items-stretch">
              <LastTransactionCard />
            </div>
          </div>
          <div className="lg:col-span-3 flex flex-col h-full">
            <div className="mb-3 flex justify-between gap-0 md:gap-56 lg:justify-between lg:gap-0  md:justify-start items-center text-lg font-semibold">
              <h2 className="dark:text-blue-500">My Card</h2>
              <Link href="/credit-card" className="font-normal self-end dark:text-blue-500">
                See All
              </Link>
            </div>
            <div className="flex flex-1 items-stretch">
              <DesktopCreditCard bgColor="bg-blue-700" textColor="text-white" />
            </div>
          </div>
        </div>

        {/* Bottom Section */}
        <div className="grid grid-cols-1 lg:grid-cols-10 gap-4 sm:gap-8 mt-8">
          <div className="lg:col-span-7 flex flex-col">
            <h2 className="text-lg font-semibold mb-4 dark:text-blue-500">
              Debit & Credit Overview
            </h2>
            <div>
              <AccountBarChart />
            </div>
          </div>
          <div className="lg:col-span-3 flex flex-col">
            <h2 className="text-lg font-semibold mb-4">Invoices Sent</h2>
            <div>
              <InvoicesCard />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Accounts;