'use client';
import React from "react";
import Example from './components/linechart';
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

// Dummy Data
const dataCorner: CardProps[] = [
  {
    title: "My Balance",
    salary: 12000,
    index: 1,
  },
  {
    title: "Income",
    salary: 5600,
    index: 2,
  },
  {
    title: "Expense",
    salary: 3460,
    index: 3,
  },
  {
    title: "Total Saving",
    salary: 7920,
    index: 4,
  },
];

const transactions: Transaction[] = [
  {
    index: 5,
    title: "Spotify Subscription",
    jobtitle: "Shopping",
    creditcard: "1234****",
    status: "Pending",
    value: 150,
    date: "25 Jan 2021", // Added date
  },
  {
    index: 6,
    title: "Mobile Service",
    jobtitle: "Service",
    creditcard: "1234****",
    status: "Pending",
    value: 1200,
    date: "15 Feb 2021", // Added date
  },
  {
    index: 7,
    title: "Grocery Shopping",
    jobtitle: "Supermarket",
    creditcard: "1234****",
    status: "Completed",
    value: 350,
    date: "10 Mar 2021", // Added date
  },
  
];



const Page: React.FC = () => {
  return (
    <div>
      <div className="max-w-screen-lg mx-auto p-4">
        <div className="flex justify-center space-x-4 mb-8">
          {dataCorner.map((card) => (
            <div
              key={card.index}
              className="flex justify-center items-center w-[200px] h-[85px] rounded-lg shadow-lg bg-white"
            >
              <div className="flex items-center">
                <img
                  className="w-8 h-8 text-gray-500 mr-4"
                  src={`/images/${card.index}.png`}
                  alt="Image Icon"
                />
                <div>
                  <h3 className="text-gray-500 text-sm">{card.title}</h3>
                  <p className="text-black text-2xl font-bold">
                    ${card.salary.toLocaleString()}
                  </p>
                </div>
              </div>
            </div>
          ))}
        </div>
       
        <div className="w-1/2">
        <h1 className="text-[22px] font-semibold leading-[26.63px] text-[rgba(51,59,105,1)] text-left bg-white px-4 py-2">
          Last transaction
        </h1>

        <div className="bg-white rounded-3xl shadow-lg p-4 w-[830px] h-[355px] top-80 left-72">
          <div className="flex flex-col">
            {transactions.map((transaction) => (
              <div
                key={transaction.index}
                className="flex items-center p-4 bg-white-50 rounded-lg"
              >
                <div className="p-3 rounded-full">
                  <img
                    src={`/images/${transaction.index}.png`}
                    alt={`${transaction.status} Icon`}
                    className="w-12 h-12"
                  />
                </div>
                <div className="flex-grow flex flex-col">
                  <div className="flex items-center justify-between mb-0">
                    <p className="text-gray-800 font-medium w-1/2">{transaction.title}</p>
                  </div>
                  <div className="flex items-center justify-between">
                    <p className="text-gray-400 text-xs w-1/2 text-left">{transaction.date}</p>
                    <p className="text-gray-400 w-1/4">{transaction.jobtitle}</p>
                    <p className="text-gray-400 w-1/4 text-center">{transaction.creditcard}</p>
                    <p className={` style={{ color: '#718EBF' }} font-medium w-1/4 text-right ${transaction.status  ? "style={{ color: '#718EBF' }}" : "text-green-500"}`}>
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
      </div>
      
      <h1 className="text-[22px] font-semibold leading-[26.63px] text-[rgba(51,59,105,1)] px-4 py-7">
        Debit & Credit Overview
      </h1>
      
      
      <div className="w-auto flex justify-between">

      <div className=" rounded-3xl bg-white gap-12 shadow-xl px-54 w-[60%] h-[370px] top-[487px] left-[256px]">
      <span className="text-left font-inter text-sm font-normal leading-4 px-10 py-2">
        $7,560 Debited & $5,420 Credited in this Week
      </span>
        <Example />
      </div>
      <div className=" w-[33%] shadow-md rounded-3xl p-6">
        <h3 className="text-lg font-semibold">Recent Transactions</h3>

        <div className="flex justify-items-stretch items-center mt-3">
          <div className="flex items-center space-x-2">
          <div className="bg-[#DCFAF8] w-12 h-12 rounded-2xl flex items-center justify-center">

              <img 
                src={`/images/apple.png`}
               
              alt="Deposit Icon" />
            </div>
            <div className="flex flex-col">
              <p className="text-sm font-light text-gray-400">Apple Store</p>
             
              <small className="text-xs" style={{ color: '#718EBF' }}>5h ago</small>

            </div>
          </div>
          
          <p className="text-red-500 font-light ml-auto" style={{ color: '#718EBF' }}>$450</p>

          
        </div>

        <div className="flex justify-between items-center mt-3">
          <div className="flex items-center space-x-2">
          <div className="bg-[#FFF5D9] w-12 h-12 rounded-2xl flex items-center justify-center">

            
              <img 
                src={`/images/u.png`}
                alt="Spotify"
                className="w-6 h- object-contain"
              />
            </div>
            <div className="flex flex-col">
              <p className="text-sm font-light text-gray-400">Michael</p>
              <small className="text-xs" style={{ color: '#718EBF' }}>2 days ago</small>

            </div>
          </div>
          <p className="text-red-500 font-light ml-auto" style={{ color: '#718EBF' }}>$160</p>

          
        </div>
        <div className="flex justify-between items-center mt-3">
          <div className="flex items-center space-x-2">
          <div className="bg-[#E7EDFF] w-12 h-12 rounded-2xl flex items-center justify-center">

           
              <img 
                src={`/images/p.png`}
                alt="Spotify"
                className="w-6 h- object-contain"
              />
            </div>
            <div className="flex flex-col">
              <p className="text-sm font-light text-gray-400">Playstation</p>
              <small className="text-xs" style={{ color: '#718EBF' }}>5 days ago</small>
              
            </div>
          </div>
          <p className="text-red-500 font-light ml-auto" style={{ color: '#718EBF' }}>$1085</p>

          

        </div>
        <div className="flex justify-between items-center mt-3">
          <div className="flex items-center space-x-2">
          <div className="bg-[#FFE0EB] w-12 h-12 rounded-2xl flex items-center justify-center">
<img
                src={`/images/v.png`}
                alt="Spotify"
                className="w-6 h- object-contain"
              />
            </div>
            <div className="flex flex-col">
              <p className="text-sm font-light text-gray-400">William</p>
              <small className="text-xs" style={{ color: '#718EBF' }}>10 days ago</small>


            </div>
          </div>
          <p className="text-red-500 font-light ml-auto" style={{ color: '#718EBF' }}>$90</p>

        </div>
      </div>
    </div>
    </div>
    </div>
  );
};

export default Page;

