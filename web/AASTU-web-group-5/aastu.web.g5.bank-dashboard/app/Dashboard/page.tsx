"use client";
import React, { useState, useEffect } from "react";
import RecentTransactionCard from "./components/RecentTransactionCard";
import QuickTransferList from "./components/QuickTransferList";
import PieChartComponent from "./components/PieChartComponent";
import BarchartComponent from "./components/BarchartComponent";
import LineGraphComponent from "./components/LineGraphComponent";
import Link from "next/link";
import Card from "../components/common/card";
import { useSession } from "next-auth/react";
import creditCardColor from "@/app/CreditCards/cardMockData";


interface ExtendedUser {
  name?: string;
  email?: string;
  image?: string;
  accessToken?: string;
}

function Dashboard() {

  const [error, setError] = useState<string | null>(null);
  const [activeLink, setActiveLink] = useState<string>("recent");
  const [cardData, setCardData] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const { data: session } = useSession();
  const user = session?.user as ExtendedUser;
  const accessToken = user?.accessToken;
  const [data, setData] = useState<any[]>([]);
  const [expenseData, setExpenseData] = useState<any[]>([]);


  const fetchCardData = async (page: number) => {
    if (!accessToken) {
      setError("No access token available");
      setLoading(false);
      return;
    }

    try {
      const response = await fetch(
        `https://bank-dashboard-1tst.onrender.com/cards?page=${page}&size=3`,
        {
          headers: {
            Authorization: `Bearer ${accessToken}`,
          },
        }
      );
      if (!response.ok) {
        throw new Error("Failed to fetch cards");
      }

      const data = await response.json();
      setCardData(data.content || []);
    } catch (error) {
      setError((error as Error).message);
    } finally {
      setLoading(false);
    }
  };
  return (
    <div className="flex flex-col bg-[#f9f9f9] min-h-screen">
      {/* Main content */}
      <div className="flex-1 flex flex-col">
        {/* Top Content */}
        <div className="pt-16 md:pt-20 px-6 py-12">
          {/* Main Content Layout */}
          <div className="grid grid-cols-1 gap-6 md:grid-cols-2 lg:grid-cols-2">
            {/* First Row: My Cards and Recent Transactions */}
            <div className="flex flex-col justify-between">
              <div className="p-4 shadow-md rounded-lg h-full flex flex-col">
                <div className="text-[#343C6A] border-blue-500 flex-1">
                  <div className="flex items-center justify-between">
                    <p className="text-lg font-semibold leading-6">My Cards</p>
                    <Link href="/Transactions" className="text-lg font-semibold leading-6">
                      See All
                    </Link>
                  </div>
                  <div className="overflow-x-auto gap-4 mt-4">
                    <div className="flex gap-4 bg-white">
                    {cardData.map((card, index) => (
              <Card
                key={index}
                cardData={card}
                cardColor={creditCardColor[index % creditCardColor.length]}
              />
            ))}
                      </div>
                  </div>
                </div>
              </div>
            </div>
            <div className="p-4 bg-white shadow-md rounded-lg h-full flex flex-col justify-between">
              <RecentTransactionCard />
            </div>

            {/* Second Row: Weekly Activities and Expense Statistics */}
            <div className="p-4 shadow-md rounded-lg h-full flex flex-col justify-between">
              <div className="text-[#343C6A]">
                Weekly Activities
                <div className="bg-white flex-1">
                  <BarchartComponent />
                </div>
              </div>
            </div>
            <div className="p-4 bg-white shadow-md rounded-lg h-full flex flex-col justify-between">
              <div className="text-[#343C6A]">
                Expense Statistics
                <div className="h-64 bg-white flex items-center justify-center flex-1">
                  <PieChartComponent />
                </div>
              </div>
            </div>

            {/* Third Row: Quick Transfer and Balance History */}
            <div className="p-4 bg-white shadow-md rounded-lg h-full flex flex-col justify-between">
              <div className="text-[#343C6A]">
                Quick Transfer
                <QuickTransferList />
              </div>
            </div>
            <div className="p-4 bg-white shadow-md rounded-lg h-full flex flex-col justify-between">
              <div className="text-[#343C6A]">
                Balance History
                <div className="bg-white">
                  <LineGraphComponent />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
