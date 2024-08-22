"use client";
import React from "react";
import Image from "next/image";
import { useGetAllTransactionsQuery } from "@/lib/redux/api/transactionsApi";
import CreditCard from "../components/CreditCard";
import BarChart from "../components/barchart";
import PieChart from "../components/PieChart";
import LineChart from "../components/LineChart";
import {Transaction} from "@/lib/redux/types/transactions"
import dollar from "../../../public/images/iconfinder_6_4753731 1.png"

const HomePage: React.FC = () => {
  const { data: transactionsData, isLoading } = useGetAllTransactionsQuery({
    size: 3, 
    page: 0,
  });

  if (isLoading) return <div>Loading...</div>;

  const getTransactionAmount = (transaction:Transaction) => {
    switch (transaction.type) {
      case "shopping":
      case "transfer":
      case "service":
        return `-$${Math.abs(transaction.amount)}`;
      case "deposit":
        return `+$${transaction.amount}`;
      default:
        return `$${transaction.amount}`; // Default case if type is unrecognized
    }
  };

  const getAmountStyle = (transaction:Transaction) => {
    switch (transaction.type) {
      case "shopping":
      case "transfer":
      case "service":
        return "text-red-500";
      case "deposit":
        return "text-green-500";
      default:
        return "";
    }
  };

  return (
    <div className="bg-[#F5F7FA] min-h-screen p-5">
      <div className="lg:flex lg:space-x-8">
        <div className="lg:w-2/3 lg:flex lg:space-x-8 overflow-x-auto flex bg-[#F5F7Fa]">
          <div className="flex-shrink rounded-lg w-full h-60 lg:h-56 lg:w-1/2 lg:rounded-r-none">
            <CreditCard
              name="Karthik P"
              balance="$5,756"
              cardNumber="3778 **** **** 1234"
              validDate="11/15"
              backgroundImg="bg-[#2E2BF0]"
              textColor="text-white"
            />
          </div>

          <div className="flex-shrink rounded-lg lg:w-1/2 lg:rounded-l-none h-60 lg:h-56 w-full">
            <CreditCard
              name="Karthik P"
              balance="$5,756"
              cardNumber="3778 **** **** 1234"
              validDate="11/15"
              backgroundImg="bg-white"
              textColor="text-black"
            />
          </div>
        </div>

        <div className="mt-5 lg:mt-0 lg:w-1/3 lg:bg-white lg:p-5 lg:rounded-lg lg:border lg:shadow-md lg:h-56">
          <h3 className="text-lg font-semibold">Recent Transactions</h3>
          {transactionsData!.data.content.map((transaction, index) => (
            <div key={index} className="flex justify-between items-center mt-2">
              <div className="flex items-center space-x-2">
                <div className="relative w-12 h-12 rounded-full flex items-center justify-center bg-gray-200">
                  <Image
                    src={dollar} // Replace with an appropriate image path or dynamic data
                    alt="transaction"
                    height={100}
                    width={100}
                  />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-medium">{transaction.description}</p>
                  <small className="text-xs text-gray-500">{transaction.date}</small>
                </div>
              </div>
              <p className={`font-semibold ml-auto ${getAmountStyle(transaction)}`}>
                {getTransactionAmount(transaction)}
              </p>
            </div>
          ))}
        </div>
      </div>

      <div className="mt-5 lg:flex lg:space-x-8">
        <div className="lg:w-2/3 lg:bg-white lg:p-5 lg:rounded-lg lg:border lg:shadow-md">
          <h3 className="text-lg font-semibold">Weekly Activity</h3>
          <div className="h-82 bg-white rounded-lg">
            <BarChart />
          </div>
        </div>

        <div className="mt-5 lg:mt-0 lg:w-1/3 lg:bg-white lg:p-5 lg:rounded-lg lg:border lg:shadow-md">
          <h3 className="text-lg font-semibold">Expenses Statistics</h3>
          <div className="h-52 bg-white rounded-lg">
            <PieChart />
          </div>
        </div>
      </div>

      <div className="mt-5 lg:flex lg:space-x-8">
        <div className="lg:w-1/2 lg:bg-white lg:p-5 lg:rounded-lg lg:border lg:shadow-md">
          <h3 className="text-lg font-semibold">Quick Transfer</h3>
          {/* Quick Transfer Section */}
        </div>

        <div className="mt-5 lg:mt-0 lg:w-1/2 lg:bg-white lg:p-5 lg:rounded-lg lg:border lg:shadow-md">
          <h3 className="text-lg font-semibold">Balance History</h3>
          <div className="h-52 bg-white rounded-lg">
            <LineChart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default HomePage;

