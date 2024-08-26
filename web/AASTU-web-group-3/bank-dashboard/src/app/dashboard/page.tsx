"use client";
import React from "react";
import Image from "next/image";
import { useGetAllTransactionsQuery } from "@/lib/redux/api/transactionsApi";
import CreditCard from "../components/CreditCard";
import { BarChartComponent } from "../components/Chart/Barchart";
import {PieChartComponent} from "../components/Chart/PieChart";
import { AreaChartComponent } from "../components/Chart/AreaChartComponent";
import { Transaction } from "@/lib/redux/types/transactions";
import dollar from "../../../public/images/iconfinder_6_4753731 1.png";

const HomePage: React.FC = () => {
  const { data: transactionsData, isLoading } = useGetAllTransactionsQuery({
    size: 3,
    page: 0,
  });

  if (isLoading) return <div>Loading...</div>;

  const getTransactionAmount = (transaction: Transaction) => {
    switch (transaction.type) {
      case "shopping":
      case "transfer":
      case "service":
        return `-$${Math.abs(transaction.amount)}`;
      case "deposit":
        return `+$${transaction.amount}`;
      default:
        return `$${transaction.amount}`; 
    }
  };

  const getAmountStyle = (transaction: Transaction) => {
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

  const imageData = [
    {
      src: "https://via.placeholder.com/48",
      alt: "Placeholder 1",
      name: "Name 1",
      position: "Position 1",
    },
    {
      src: "https://via.placeholder.com/48",
      alt: "Placeholder 2",
      name: "Name 2",
      position: "Position 2",
    },
    {
      src: "https://via.placeholder.com/48",
      alt: "Placeholder 3",
      name: "Name 3",
      position: "Position 3",
    },
  ];

  return (
    <div className="bg-[#F5F7FA] min-h-screen p-5 dark:bg-darkPage">
      <div className="lg:flex lg:space-x-8">
        <div className="lg:w-2/3 lg:flex lg:space-x-8 overflow-x-auto flex bg-[#F5F7Fa] dark:bg-darkPage">
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

        {/* <div className="mt-5 lg:mt-0 lg:w-1/3 lg:bg-white lg:p-5 lg:rounded-lg lg:border lg:shadow-md lg:h-56">
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
                  <p className="text-sm font-medium">
                    {transaction.description}
                  </p>
                  <small className="text-xs text-gray-500">
                    {transaction.date}
                  </small>
                </div>
              </div>
              <p
                className={`font-semibold ml-auto ${getAmountStyle(
                  transaction
                )}`}
              >
                {getTransactionAmount(transaction)}
              </p>
            </div>
          ))}
        </div> */}
      </div>

      <div className="mt-5 space-y-5 lg:space-y-0 lg:flex lg:space-x-8 lg:h-[400px] ">
        <div className="rounded-lg h-full flex-1">
          <BarChartComponent />
        </div>

        <div className="rounded-lg flex-1 h-full">
          <PieChartComponent />
        </div>
      </div>

      <div className="mt-20 space-y-5 lg:space-y-0 lg:flex lg:space-x-8">
        <div className="w-full lg:w-1/2 bg-white p-5 rounded-lg border shadow-md lg:mt-40 dark:bg-darkComponent">
          <h3 className="text-lg font-semibold">Quick Transfer</h3>
          <div className="flex justify-center mt-3 space-x-16 ">
            {imageData.map((image, index) => (
              <div key={index} className="flex flex-col items-center">
                <div className="w-12 h-12 bg-gray-300 rounded-full overflow-hidden">
                  <img
                    src={image.src}
                    alt={image.alt}
                    className="w-full h-full object-cover"
                  />
                </div>
                <p className="text-sm font-medium">{image.name}</p>
                <small className="text-xs text-gray-500">
                  {image.position}
                </small>
              </div>
            ))}
          </div>
          <div className="mt-5 lg:mt-8">
            <h3 className="text-lg font-semibold">Write Amount</h3>
            <div className="relative flex items-center mt-3">
              <input
                type="text"
                placeholder="525.20"
                className="w-full p-3 h-12 rounded-3xl border border-gray-300 bg-white dark:bg-darkComponent"
              />
              <button className="flex items-center justify-center absolute top-0 right-0 h-full px-3 bg-[#1814F3] text-white rounded-3xl">
                <span className="mr-2">Send</span>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                  className="w-5 h-5"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M5 12h14M12 5l7 7-7 7"
                  />
                </svg>
              </button>
            </div>
          </div>
        </div>

        <div className="lg:w-1/2 w-full h-64 lg:h-80">
          <AreaChartComponent />
        </div>
      </div>
    </div>
  );
};

export default HomePage;

