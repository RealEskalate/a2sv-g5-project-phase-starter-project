"use client";
import { useState, useEffect } from "react";
import BalanceCard from "@/components/AccountSmallCard";
import LastTransactionCard from "@/components/LastTransactionCard";
import InvoicesCard from "@/components/InvoicesCard";
import AccountBarChart from "@/components/AccountBarChart";
import Link from "next/link";
import ResponsiveCreditCard from "@/components/CreditCard";
import Cookies from "js-cookie";
import { getAllCards } from "@/services/cardfetch";
import { colors } from "@/constants";
import Image from "next/image";
import MyCardsLoad from "@/components/loadingComponents/MyCardsLoad";

const Accounts = () => {
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
        setCards(data.slice(0, 1));
        console.log("shuluka",cards);
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
    <div className="flex dark:bg-dark text-gray-900 dark:text-white">
      {/* Sidebar */}
      <div className="hidden lg:block w-64 bg-white h-screen fixed top-0 left-0">
        {/* Your Sidebar content goes here */}
      </div>

      {/* Main content */}
      <div className="flex-1 lg:ml-64 p-4 sm:p-8 bg-gray-100 dark:bg-dark text-gray-900 dark:text-white">
        {/* Top Section */}
        <div className="mb-8">
          <h1 className="text-2xl font-semibold mb-6 dark:text-blue-500">
            Accounts
          </h1>
          {/* <div className="grid grid-cols-1 sm:grid-cols-2 gap-4 sm:gap-8"> */}
          <BalanceCard />
          {/* </div> */}
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-10 gap-4 sm:gap-8 mb-8">
          <div className="lg:col-span-7 flex flex-col">
            <h2 className="text-lg font-semibold mb-3 dark:text-blue-500">
              Last Transaction
            </h2>
            <div>
              <LastTransactionCard />
            </div>
          </div>
          <div className="lg:col-span-3 flex flex-col h-full">
            <div className="mb-3 flex justify-between gap-0 md:gap-56 lg:justify-between lg:gap-0  md:justify-start items-center text-lg font-semibold">
              <h2 className="dark:text-blue-500">My Card</h2>
              <Link
                href="/credit-card"
                className="font-normal self-end dark:text-blue-500"
              >
                See All
              </Link>
            </div>
            <div className="flex flex-1 items-stretch">
            {loading ? (
                <MyCardsLoad count={1} />
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
                  <Image
                    src="/icons/null.png"
                    width={80}
                    height={80}
                    alt="null"
                    className="mx-auto pb-2 block"
                  />
                  <span className="mx-auto my-auto md:text-xl text-sm text-[#993d4b] font-bold mb-5">
                    {error ? error : "There are no cards for now!"}
                  </span>
                </div>
              ) : (
                <MyCardsLoad count={1} />
              )}
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
