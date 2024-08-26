"use client";

import React, { Suspense, useEffect, useState } from "react";
import RecentTransactions from "@/components/RecentTransaction";
import ExpensesChart from "@/components/ExpensesCart";
import SlidingCards from "@/components/SlidingCards"; // Import the sliding cards component
import Cookies from "js-cookie";
import { getAllCards } from "@/services/cardfetch";
import Image from "next/image";
import MyCardsLoad from "@/components/loadingComponents/MyCardsLoad";
import ResponsiveCreditCard from "@/components/CreditCard";
import { colors } from "@/constants";
import { TbFileSad } from "react-icons/tb";

const Transaction: React.FC = () => {
  const [cards, setCards] = useState<any[]>([]);
  const [token, setToken] = useState<any>(null); // Initialize with null to indicate it's being fetched
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(true); // Add a loading state

  useEffect(() => {
    const fetchCards = async () => {
      try {
        const storedToken = Cookies.get("accessToken");
        setToken(storedToken);

        if (!storedToken) {
          setLoading(false); // Stop loading if no token is found
          throw new Error("Token not found. Please log in again.");
        }

        const data = await getAllCards(storedToken);
        setCards(data.slice(0, 2));
        console.log(data);
        setError(null);
      } catch (err) {
        setError("Failed to fetch cards data!");
      } finally {
        setLoading(false); // Stop loading once the fetch is done
      }
    };

    fetchCards();
  }, []);
  return (
    <div className=" w-[100%]">
      {/* Large Screens Layout */}
      <div className=" hidden lg:grid lg:grid-cols-2 lg:gap-5 lg:space-x-8 lg:pb-8 lg:ml-72 lg:pt-16">
        {/* Cards Section */}
        <div className="flex flex-col">
          <h1 className="text-2xl font-bold mb-4 dark:text-blue-500">
            My Cards
          </h1>
          <div className=" w-[100%] overflow-x-auto flex space-x-4 scrollbar-thin scrollbar-track-[#F5F7FA] dark:scrollbar-track-dark scrollbar-thumb-[#92a7c5] scrollbar-thumb-rounded-full">
          {loading ? (
                <MyCardsLoad count={2}/>
              ) : Array.isArray(cards) && cards.length > 0 ? (
                cards.map((card: any, index: number) => (
                  <div key={index} className="p-1 flex gap-1">
                    <ResponsiveCreditCard
                      backgroundColor={
                        index % 2 === 0 ? colors.blue : colors.white
                      }
                      balance={card.balance}
                      cardHolder={card.cardHolder}
                      expiryDate={card.expiryDate.slice(0, 10)}
                      cardNumber={card.semiCardNumber}
                    />
                  </div>
                ))
              ) : token ? (
                <div className="w-screen bg-white py-16 rounded-xl flex flex-col justify-center dark:bg-dark dark:border-[1px] dark:border-gray-700">
                  <TbFileSad
                    className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
                    strokeWidth={1}
                  />
                  <span className="mx-auto my-auto md:text-xl text-sm text-[#993d4b] mb-5">
                    {error ? error : "There are no cards for now!"}
                  </span>
                </div>
              ) : (
                <MyCardsLoad count={2} />
              )}
          </div>
        </div>

        {/* Chart Section */}
        <div className="flex-1 ml-0">
          <h1 className="text-2xl font-bold dark:text-blue-500">My Expenses</h1>

          <div className="pl-8">
            <ExpensesChart />
          </div>
        </div>
      </div>

      {/* Mobile Layout */}
      <div className="lg:hidden pt-10">
        <h1 className="text-2xl font-bold mb-4 dark:text-blue-500">My Cards</h1>
        {loading ? (
          <MyCardsLoad count={2}/>
        ) : Array.isArray(cards) && cards.length > 0 ? (
          cards.map((card: any, index: number) => (
            <div key={index} className="p-1 flex gap-1">
              <ResponsiveCreditCard
                backgroundColor={
                  index % 2 === 0 ? colors.blue : colors.white
                }
                balance={card.balance}
                cardHolder={card.cardHolder}
                expiryDate={card.expiryDate.slice(0, 10)}
                cardNumber={card.semiCardNumber}
              />
            </div>
          ))
        ) : token ? (
          <div className="w-screen bg-white py-16 rounded-xl flex flex-col justify-center dark:bg-dark dark:border-[1px] dark:border-gray-700">
            <TbFileSad
              className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
              strokeWidth={1}
            />
            <span className="mx-auto my-auto md:text-xl text-sm text-[#993d4b] mb-5">
              {error ? error : "There are no cards for now!"}
            </span>
          </div>
        ) : (
          <MyCardsLoad count={2} />
        )}
        <h1 className="text-2xl font-bold mb-4 dark:text-blue-500">
          My Expenses
        </h1>
        <div className="w-full">
          <ExpensesChart />
        </div>
      </div>

      {/* Recent Transactions Section */}
      <div>
        <h1 className="text-2xl font-bold text-balance lg:text-center mb-4 dark:text-blue-500">
          Recent Transactions
        </h1>

        <RecentTransactions />
      </div>
    </div>
  );
};

export default Transaction;
