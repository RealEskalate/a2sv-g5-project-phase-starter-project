'use client';

import React, { useState, useEffect } from "react";
import RecentTransactionCard from "./components/RecentTransactionCard";
import QuickTransferList from "./components/QuickTransferList";
import PieChartComponent from "./components/PieChartComponent";
import BarchartComponent from "./components/BarchartComponent";
import LineGraphComponent from "./components/LineGraphComponent";
import Link from "next/link";
import Card from "../components/common/card";
import { useSession } from "next-auth/react";
import Chip_card1 from "@/public/assets/image/Chip_Card1.png";
import Chip_card3 from "@/public/assets/image/Chip_Card3.png";

const DashboardCardColor = [
  {
    cardBgColor: "bg-blue-500 rounded-3xl text-white",
    bottomBgColor:
      "flex justify-between p-4 bg-blue-400 rounded-bl-3xl rounded-br-3xl",
    imageCreditCard: Chip_card1,
    grayCircleColor: false,
  },
  {
    cardBgColor:
      "bg-[#fff] rounded-3xl text-[#343C6A] border-2 border-solid border-gray-200 ",
    bottomBgColor:
      "flex justify-between p-4 border-t-2 border-solid border-gray-200 rounded-bl-3xl rounded-br-3xl",
    imageCreditCard: Chip_card3,
    grayCircleColor: true,
  },
];

interface ExtendedUser {
  name?: string;
  email?: string;
  image?: string;
  accessToken?: string;
  refreshToken?: string;
}

function Dashboard() {
  const [error, setError] = useState<string | null>(null);
  const [activeLink, setActiveLink] = useState<string>("recent");
  const [cardData, setCardData] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const { data: session } = useSession();
  const user = session?.user as ExtendedUser;
  const accessToken = user?.accessToken;
  const refreshToken = user?.refreshToken;

  useEffect(() => {
    const fetchCardData = async (page: number) => {
      if (!accessToken) {
        setError("No access token available");
        setLoading(false);
        return;
      }

      try {
        const response = await fetch(
          `https://bank-dashboard-rsf1.onrender.com/cards?page=${page}&size=2`,
          {
            headers: {
              Authorization: `Bearer ${accessToken}`,
            },
          }
        );

        if (!response.ok) {
          // Check if unauthorized, and refresh token if needed
          if (response.status === 401 && refreshToken) {
            try {
              const refreshResponse = await fetch(
                "https://bank-dashboard-rsf1.onrender.com/auth/refresh_token",
                {
                  method: "POST",
                  headers: {
                    "Content-Type": "application/json",
                    Authorization: `Bearer ${refreshToken}`,
                  },
                }
              );

              if (!refreshResponse.ok)
                throw new Error("Failed to refresh token");

              const refreshedTokens = await refreshResponse.json();
              const newAccessToken = refreshedTokens.data.access_token;

              // Retry fetching card data with the new access token
              const retryResponse = await fetch(
                `https://bank-dashboard-rsf1.onrender.com/cards?page=${page}&size=2`,
                {
                  headers: {
                    Authorization: `Bearer ${newAccessToken}`,
                  },
                }
              );

              if (!retryResponse.ok)
                throw new Error("Failed to fetch cards with new token");

              const data = await retryResponse.json();
              setCardData(data.content || []);
            } catch (error) {
              setError("Failed to refresh access token or fetch data");
            }
          } else {
            throw new Error("Failed to fetch cards");
          }
        } else {
          const data = await response.json();
          setCardData(data.content || []);
        }
      } catch (error) {
        setError((error as Error).message);
      } finally {
        setLoading(false);
      }
    };

    fetchCardData(0);
  }, [accessToken, refreshToken]);

  return (
    <div className="bg-[#F5F7FA] space-y-8 w-[95%] pt-3 overflow-hidden mx-auto">

      {/* First Row: My Cards and Recent Transactions */}
      <div className="grid grid-cols-1 md:grid-cols-[67%_33%] gap-4">
        <div className="w-full p-4 rounded-lg flex flex-col justify-between">
          <div className="text-[#343C6A] flex-1">
            <div className="flex items-center justify-between">
              <p className="text-lg font-semibold leading-6">My Cards</p>
              <Link href="/CreditCards" className="text-lg font-semibold leading-6">
                See All
              </Link>
            </div>
            <div className="overflow-x-auto gap-4 mt-4">
              <div className="flex w-full gap-6">
                {cardData.length ? (
                  cardData.map((card, index) => (
                    <Card
                      key={index}
                      cardData={card}
                      cardColor={DashboardCardColor[index % DashboardCardColor.length]}
                    />
                  ))
                ) : (
                  <p>No cards available</p>
                )}
              </div>
            </div>
          </div>
        </div>

        <div className="w-full p-4 rounded-lg flex flex-col justify-between">
          <p className="text-[#343C6A] text-lg font-semibold leading-6">Recent Transactions</p>
          <div className="mt-4">
            <RecentTransactionCard />
          </div>
        </div>
      </div>

      {/* Second Row: Weekly Activities and Expense Statistics */}
      <div className="grid grid-cols-1 md:grid-cols-[67%_33%] gap-4">
        <div className="w-full p-4 rounded-lg flex flex-col justify-between">
          <div className="text-lg font-semibold leading-6 text-[#343C6A]">
            Weekly Activities
            <div className="bg-white mt-4">
              <BarchartComponent />
            </div>
          </div>
        </div>

        <div className="w-full p-4 rounded-lg flex flex-col justify-between">
          <div className="text-lg font-semibold leading-6 text-[#343C6A]">
            Expense Statistics
            <div className="h-64 bg-white rounded-[25px] flex items-center mt-4 justify-center flex-1">
              <PieChartComponent />
            </div>
          </div>
        </div>
      </div>

      {/* Third Row: Quick Transfer and Balance History */}
      <div className="grid grid-cols-1 md:grid-cols-[33%_67%] gap-4">
        <div className="w-full md:w-full p-4 rounded-lg flex flex-col justify-between">
          <div className="text-lg font-semibold leading-6 text-[#343C6A]">
            Quick Transfer
            <div className="mt-4 bg-white rounded-[25px]">
              <QuickTransferList />
            </div>
          </div>
        </div>

        <div className="w-full md:w-full p-4 rounded-lg flex flex-col justify-between">
          <div className="text-lg font-semibold leading-6 text-[#343C6A]">
            Balance History
            <div className="bg-white mt-4">
              <LineGraphComponent />
            </div>
          </div>
        </div>
      </div>

    </div>
  );
}

export default Dashboard;
