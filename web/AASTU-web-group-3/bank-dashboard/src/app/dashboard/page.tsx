"use client";
import React from "react";
import Image from "next/image";
import img from "./images/profile.png";
import emanuel from "../../../public/images/emanuel-minca-jYv069cQuB8-unsplash 1.png";
import julia from "../../../public/images/pexels-julia-volk-5273755 1.png";
import marcel from "../../../public/images/marcel-strauss-Uc_tOqa_jDY-unsplash 1.png";
import paypal from "../../../public/images/iconfinder_paypal_payment_pay_5340264 1.png";
import deposit from "../../../public/images/iconfinder_business_finance_money-13_2784281 1.png"
import dollar from "../../../public/images/iconfinder_6_4753731 1.png";
import BarChart from "../components/barchart";
import PieChart from "../components/PieChart";
import LineChart from "../components/LineChart";
import CreditCard from "../components/CreditCard";

const imageData = [
  { src: julia.src, alt: "julia", name: "Livia Bator", position: "CEO" },
  { src: marcel.src, alt: "marcel", name: "Randy Press", position: "Director" },
  { src: emanuel.src, alt: "emanuel", name: "Workman", position: "Designer" },
];

const transactionData = [
  {
    src: deposit.src,
    alt: "deposit",
    backgroundColor: "#FFF5D9",
    title: "Deposit from my",
    date: "28 January 2021",
    amount: "-$850",
    amountColor: "text-red-500",
  },
  {
    src: paypal.src,
    alt: "paypal",
    backgroundColor: "#E7EDFF",
    title: "Deposit Paypal",
    date: "25 January 2021",
    amount: "+$2,500",
    amountColor: "text-green-500",
  },
  {
    src: dollar.src,
    alt: "dollar",
    backgroundColor: "#DCFAF8",
    title: "Jemi Wilson",
    date: "21 January 2021",
    amount: "+$5,400",
    amountColor: "text-green-500",
  },
];

const HomePage: React.FC = () => {
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
          {transactionData.map((transaction, index) => (
            <div key={index} className="flex justify-between items-center mt-3">
              <div className="flex items-center space-x-2">
                <div
                  className={`relative ${transaction.backgroundColor} w-12 h-12 rounded-full flex items-center justify-center`}
                >
                  <Image src={transaction.src} alt={transaction.alt} height={100} width ={100} />
                </div>
                <div className="flex flex-col">
                  <p className="text-sm font-medium">{transaction.title}</p>
                  <small className="text-xs text-gray-500">
                    {transaction.date}
                  </small>
                </div>
              </div>
              <p className={`${transaction.amountColor} font-semibold ml-auto`}>
                {transaction.amount}
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
          <div className="flex justify-center mt-3 space-x-16">
            {imageData.map((image, index) => (
              <div key={index} className="flex flex-col items-center">
                <div className="w-12 h-12 bg-gray-300 rounded-full overflow-hidden">
                  <Image
                    src={image.src}
                    alt={image.alt}
                    className="w-full h-full object-cover"
                    height={100}
                    width={100}
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
                className="w-full p-3 pr-20 h-12 rounded-3xl border border-gray-300"
              />
              <button className="absolute top-0 right-0 h-full w-1/2 px-3 bg-[#1814F3] text-white rounded-3xl">
                Send
              </button>
            </div>
          </div>
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
