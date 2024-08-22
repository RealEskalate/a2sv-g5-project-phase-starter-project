"use client";
import React, { useEffect, useState } from "react";
import BarGraph from "@/app/components/Transaction/BarGraph";
import Recent from "@/app/components/Transaction/Recent";
import VisaCard from "@/app/components/Card/VisaCard";
import { TransactionType, ChartData } from "@/types/TransactionValue";
import { useAppSelector } from "@/app/Redux/store/store";
import { Card } from "../../Redux/slices/cardSlice";
import { getExpense } from "@/app/Services/api/fetchTransaction";
import { useSession } from "next-auth/react";
import { ShimmerVisaCard } from "@/app/components/Shimmer/ShimmerVisa";
const Transaction = () => {
  const { data: session } = useSession();
  const accessToken = session?.accessToken as string;
  const [expenseData, setExpenseData] = useState<TransactionType[]>([]);
  const CardData: Card[] = useAppSelector((state) => state.cards.cards);

  const fetchExpense = async () => {
    while (!accessToken) {
      await new Promise((resolve) => setTimeout(resolve, 100)); // Delay to wait for the token
    }
    const res = await getExpense(0, accessToken);
    setExpenseData(res);
  };
  useEffect(() => {
    fetchExpense();
  }, [accessToken]);
  const cardColor = [false, true];

  const convertToChartData = (data: TransactionType[]): ChartData[] => {
    const dayMap: { [key: string]: number } = {
      Mon: 0,
      Tue: 0,
      Wed: 0,
      Thur: 0,
      Fri: 0,
      Sat: 0,
      Sun: 0,
    };

    data.forEach((transaction) => {
      const day = new Date(transaction.date).toLocaleString("en-US", {
        weekday: "short",
      });
      const formattedDay =
        day.charAt(0).toUpperCase() + day.slice(1, 3).toLowerCase();

      if (dayMap[formattedDay] !== undefined) {
        dayMap[formattedDay] += transaction.amount;
      }
    });

    return Object.keys(dayMap).map((day) => ({
      day,
      amount: dayMap[day],
    }));
  };

  const chartData = convertToChartData(expenseData);
  // console.log(chartData, 'chart');
  return (
    <div className="pt-4 px-3 sm:px-6 md:px-8 lg:px-10 w-full">
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-8 w-full">
        <div className="lg:col-span-2 py-4 overflow-x-auto scrollbar-hide">
          <div className="flex justify-between mb-5">
            <p className="font-semibold text-xl sm:text-xl text-[#343C6A] dark:text-gray-300">
              My Cards
            </p>
            <p className="font-semibold text-base sm:text-lg text-[#343C6A] cursor-pointer dark:text-gray-300">
              + Add Card
            </p>
          </div>
          <div className="overflow-x-auto scrollbar-hide">
            <div className="flex gap-8 min-w-[900px] min-h-[170px]">
              <>
                {CardData.length > 0 ? (
                  CardData?.slice(0, 2).map((item, index) => (
                    <VisaCard
                      key={index}
                      data={item}
                      isBlack={cardColor[index] || false}
                      isFade={false}
                      isSimGray={false}
                    />
                  ))
                ) : (
                  <div className="w-full flex gap-6 ">
                    <ShimmerVisaCard />
                    <ShimmerVisaCard />
                  </div>
                )}
              </>
            </div>
          </div>
        </div>
        <div className="w-full flex flex-col gap-4 pt-4">
          <p className="font-semibold text-[20px] text-[#343C6A] mb-1 dark:text-gray-300">
            My Expense
          </p>
          <BarGraph chartData={chartData} />
        </div>
      </div>
      <Recent />
    </div>
  );
};

export default Transaction;
