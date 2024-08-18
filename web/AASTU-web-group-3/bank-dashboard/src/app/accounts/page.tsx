'use client';
import React from "react";
import Example from '../components/barchart';
import CreditCard from "../components/CreditCard";

interface CardProps {
  title: string;
  salary: number;
  index: number;
}

interface Transaction {
  index: number;
  title: string;
  jobtitle: string;
  creditcard: string;
  status: string;
  value: number;
  date: string;
}

const dataCorner: CardProps[] = [
  { title: "My Balance", salary: 12000, index: 1 },
  { title: "Income", salary: 5600, index: 2 },
  { title: "Expense", salary: 3460, index: 3 },
  { title: "Total Saving", salary: 7920, index: 4 },
];

const transactions: Transaction[] = [
  { index: 5, title: "Spotify Subscription", jobtitle: "Shopping", creditcard: "1234****", status: "Pending", value: 150, date: "25 Jan 2021" },
  { index: 6, title: "Mobile Service", jobtitle: "Service", creditcard: "1234****", status: "Pending", value: 1200, date: "15 Feb 2021" },
  { index: 7, title: "Grocery Shopping", jobtitle: "Supermarket", creditcard: "1234****", status: "Completed", value: 350, date: "10 Mar 2021" },
];

const Page: React.FC = () => {
  return (
    <div className="max-w-screen max-h-screen mx-auto px-4 py-4 sm:px-6 sm:py-6 lg:px-10 lg:py-8">
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4 mb-4">
        {dataCorner.map((card) => (
          <div key={card.index} className="flex justify-center items-center h-[85px] rounded-2xl shadow-xl bg-white p-4">
            <div className="flex items-center">
              <img className="w-8 h-8 text-gray-500 mr-4" src={`/images/${card.index}.png`} alt="Image Icon" />
              <div>
                <h3 className="text-gray-500 text-sm">{card.title}</h3>
                <p className="text-black text-2xl font-bold">${card.salary.toLocaleString()}</p>
              </div>
            </div>
          </div>
        ))}
      </div>

      <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] text-left px-4 py-2">
        Last transaction
      </h1>

      <div className=" grid grid-cols-1 lg:grid-cols-3 gap-4">
        <div className="rounded-3xl bg-white shadow-xl p- w-full lg:col-span-2">
          <div className="flex flex-col space gap-y-2">
            {transactions.map((transaction) => (
              <div key={transaction.index} className="flex items-center p-4 bg-white rounded-lg shadow-sm">
                <div className="p-3 rounded-full">
                  <img src={`/images/${transaction.index}.png`} alt={`${transaction.status} Icon`} className="w-12 h-12" />
                </div>
                <div className="flex-grow flex flex-col ml-4">
                  <div className="flex items-center justify-between mb-1">
                    <p className="text-gray-800 font-medium">{transaction.title}</p>
                  </div>
                  <div className="flex items-center justify-between overflow-hidden">
                    <p className="text-gray-400 text-xs w-1/2 text-left">{transaction.date}</p>
                    <p className="text-gray-400 w-1/4 hidden sm:block">{transaction.jobtitle}</p>
                    <p className="text-gray-400 w-1/4 text-center hidden sm:block">{transaction.creditcard}</p>
                    <p className={` style={{ color: '#718EBF' }} hidden sm:block font-medium w-1/4 text-right ${transaction.status  ? "style={{ color: '#718EBF' }}" : "text-green-500"}`}>
                      {transaction.status}
                    </p>
                    <p className={`text-400 font-medium w-1/4 text-right ${transaction.value < 0 ? "text-red-500" : "text-green-500"}`}>
                      ${Math.abs(transaction.value).toLocaleString()}
                    </p>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </div>

        <div className="w-full h-56 md:h-[90%] rounded-3xl overflow-hidden col-span-1">
        <CreditCard
                name="Karthik P"
                balance="$5,756"
                cardNumber="3778 **** **** 1234"
                validDate="11/15"
                backgroundImg="bg-blue-600"
                textColor ="text-white"
            />

        </div>
      </div>

      <h1 className="text-[22px] font-bold leading-[26.63px] text-[rgba(51,59,105,1)] px-4 py-7">
        Debit & Credit Overview
      </h1>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4 ">
        <div className="rounded-3xl bg-white shadow-xl p-6 lg:col-span-2 space gap-y-3 h-80">
          <span className="text-left font-inter text-sm font-normal leading-4 px-6 py-2">
            $7,560 Debited & $5,420 Credited in this Week
          </span>
          <Example />
        </div>

        <div className="rounded-3xl bg-white shadow-md p-6">
          <h3 className="text-lg font-semibold">Invoices Sent</h3>

          <div className="flex flex-col space-y-3 mt-3">
            {/* Invoice 1 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#DCFAF8] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <img src={`/images/apple.png`} alt="Deposit Icon" />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">Apple Store</p>
                  <small className="text-xs text-blue-500">5h ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$450</p>
            </div>

            {/* Invoice 2 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#FFF5D9] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <img src={`/images/u.png`} alt="Spotify" className="w-6 h-6 object-contain" />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">Michael</p>
                  <small className="text-xs text-blue-500">2 days ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$160</p>
            </div>

            {/* Invoice 3 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#E7EDFF] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <img src={`/images/p.png`} alt="Playstation" className="w-6 h-6 object-contain" />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">Playstation</p>
                  <small className="text-xs text-blue-500">5 days ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$1085</p>
            </div>

            {/* Invoice 4 */}
            <div className="flex justify-between items-center">
              <div className="flex items-center space-x-2">
                <div className="bg-[#FFE0EB] w-12 h-12 rounded-2xl flex items-center justify-center">
                  <img src={`/images/v.png`} alt="William" className="w-6 h-6 object-contain" />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-light text-gray-400">William</p>
                  <small className="text-xs text-blue-500">10 days ago</small>
                </div>
              </div>
              <p className="text-blue-500 font-light ml-auto">$90</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Page;
