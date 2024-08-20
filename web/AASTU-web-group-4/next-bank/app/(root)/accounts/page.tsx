"use client";
import { useState, useEffect } from "react";
import BalanceCard from "@/components/AccountSmallCard";
import LastTransactionCard from "@/components/LastTransactionCard";
import DesktopCreditCard from "@/components/DesktopCreditCard";
import InvoicesCard from "@/components/InvoicesCard";
import AccountBarChart from "@/components/AccountBarChart";
import { getAllBankServices, getBankServiceById } from "@/services/bankseervice";

const Accounts = () => {
  const [balanceData, setBalanceData] = useState([]);
  const [lastTransactionCard, setLastTransactionData] = useState([]);
  const [cardData, setCardData] = useState(null);
  const [barChartData, setBarChartData] = useState([]);
  const [invoiceData, setInvoiceData] = useState([]);

  useEffect(() => {
    // Fetch all bank services for the balance data
    const fetchBankServices = async () => {
      const services = await getAllBankServices();
      setBalanceData(services);
    };

    // Fetch last transaction 
    const fetchLastTransaction = async () => {
      const transaction = await getBankServiceById(1);
      setLastTransactionData(transaction);
    };

    // // Fetch card data
    // const fetchCardData = async () => {
    //   const card = await getBankServiceById(2);
    //   setCardData(card);
    // };

    // // Fetch bar chart
    // const fetchBarChartData = async () => {
    //   const chart = await getBankServiceById(3);
    //   setBarChartData(chart);
    // };

    // // Fetch invoice data
    // const fetchInvoiceData = async () => {
    //   const invoice = await getBankServiceById(4);
    //   setInvoiceData(invoice);
    // };

    fetchBankServices();
    fetchLastTransaction();
    // fetchCardData();
    // fetchBarChartData();
    // fetchInvoiceData();
  })

  return (
    <div className="flex">
      {/* Sidebar */}
      <div className="hidden lg:block w-64 bg-white h-screen fixed top-0 left-0">
        {/* Your Sidebar content goes here */}
      </div>

      {/* Main content */}
      <div className="flex-1 lg:ml-64 p-4 sm:p-8 bg-gray-100">
        {/* Top Section */}
        <div className="mb-8">
          <h1 className="text-2xl font-semibold mb-6">Accounts</h1>
          {/* <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-8"> */}
            <BalanceCard />
          {/* </div> */}
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-10 gap-4 sm:gap-8 mb-8">
          <div className="lg:col-span-7 flex flex-col">
            <h2 className="text-lg font-semibold mb-3">Last Transaction</h2>
            <div className="flex-1 flex items-stretch">
              <LastTransactionCard />
            </div>
          </div>
          <div className="lg:col-span-3 flex flex-col h-full">
            <div className="mb-3 flex justify-between items-center text-lg font-semibold">
              <h2>My Card</h2>
              <a href="/credit-card" className="font-normal self-end">
                See All
              </a>
            </div>
            <div className="flex flex-1 items-stretch">
              <DesktopCreditCard bgColor="bg-blue-700" textColor="text-white" />
            </div>
          </div>
        </div>

        {/* Bottom Section */}
        <div className="grid grid-cols-1 lg:grid-cols-10 gap-4 sm:gap-8 mt-8">
          <div className="lg:col-span-7 flex flex-col">
            <h2 className="text-lg font-semibold mb-4">
              Debit & Credit Overview
            </h2>
            <div className="flex-1 flex items-stretch">
              <AccountBarChart />
            </div>
          </div>
          <div className="lg:col-span-3 flex flex-col">
            <h2 className="text-lg font-semibold mb-4">Invoices Sent</h2>
            <div className="flex-1 flex items-stretch">
              <InvoicesCard />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Accounts;