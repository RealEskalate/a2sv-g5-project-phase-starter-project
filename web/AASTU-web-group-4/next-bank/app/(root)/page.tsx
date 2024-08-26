"use client";
import React, { useState, useEffect } from "react";
import { Color } from "chart.js";
import { colors, logo } from "@/constants";
import DesktopCreditCart from "@/components/DesktopCreditCard";
import ResponsiveCreditCard from "@/components/CreditCard";
import RecentTransaction from "@/components/Recent Transaction";
import ExpensesChart from "@/components/ExpensesCart";
import { icons, Import } from "lucide-react";
import { text } from "stream/consumers";
import BarChart from "@/components/BarChart";
import PieChart from "@/components/PieChart";
import QuickTransfer from "@/components/QuickTransfer";
import LineChart from "@/components/LineChart";
import Link from "next/link";
import Cookies from "js-cookie";
import { getAllCards } from "@/services/cardfetch";
import Image from "next/image";
import MyCardsLoad from "@/components/loadingComponents/MyCardsLoad";
import { TbFileSad } from "react-icons/tb";

const Page = () => {
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
    <div className={`${colors.graybg} p-6 md:pl-10 lg:pl-72 md:max-w-full md:px-12  dark:bg-dark text-gray-900 dark:text-white`}>
      <div className="flex flex-col justify-between md:flex-row  gap-10 ">
        <div className=" py-4 md:w-3/5 md:max-w-full">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl dark:text-blue-500">My Cards</h1>
            <Link href="/credit-card" className="py-2 dark:text-blue-500">
              {""}
              See All
            </Link>
          </div>

          <div className="max-w-[345px] md:max-w-full">
            <div className="flex gap-3 overflow-x-auto md:w-auto">
              {loading ? (
                <MyCardsLoad count={2} />
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
                  <span className="mx-auto my-auto md:text-xl text-sm text-red-500 mb-5">
                    {error ? error : "There are no cards for now!"}
                  </span>
                </div>
              ) : (
                <MyCardsLoad count={2} />
              )}
            </div>
          </div>
        </div>
        <div className="  md:w-2/5  flex flex-col ">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl dark:text-blue-500 ">
              Recent Transaction
            </h1>
          </div>
          <div className="flex flex-col rounded-2xl pr-2 w-[100%]">
            <RecentTransaction />
          </div>
        </div>
      </div>
      <div className=" w-[100%] flex flex-col justify-between  md:grid md:grid-cols-5 md:gap-10 ">
        <div className=" md:col-span-3 ">
          <div className={`${colors.navbartext} flex justify-between py-4`}>
            <h1 className="font-bold text-2xl dark:text-blue-500">
              Weekly Activity
            </h1>
          </div>
          <div className="w-[100%]">
            <BarChart />
          </div>
        </div>
        <div className=" w-[100%] py-5 flex flex-col gap-5 md:col-span-2 ">
          <div className={`${colors.navbartext}`}>
            <h1 className="font-bold text-2xl dark:text-blue-500">
              Expense Statstics
            </h1>
          </div>
          <div className="w-[100%] pr-">
            <PieChart />
          </div>
        </div>
      </div>

      <div className="flex flex-col justify-between w-full  md:grid md:grid-cols-5 md:gap-10 ">
        <div className=" md:col-span-2 py-4  ">
          <div className={`${colors.navbartext} flex justify-between `}>
            <h1 className="font-bold text-2xl dark:text-blue-500">
              Quick Transfer
            </h1>
          </div>
          <div className="flex  gap-3 ">
            <div className="flex py-3 ">
              {" "}
              <QuickTransfer />
            </div>
          </div>
        </div>
        <div className=" md:col-span-3 ">
          <div className={`${colors.navbartext} flex justify-between  py-4`}>
            <h1 className="font-bold text-2xl dark:text-blue-500">
              Balance History
            </h1>
          </div>
          <div className="pr-6">
            <LineChart />
          </div>
        </div>
      </div>
    </div>
  );
};

export default Page;
