import React from "react";
import Image from "next/image";
import { PieComp } from "../Charts/PieComp";
import { useAppSelector } from "@/app/Redux/store/store";
import { TransactionType } from "@/app/Redux/slices/TransactionSlice";
import ShimmerRecent from "../Shimmer/SimmerRecent";
const RecentTr = () => {
  const transdata: TransactionType[] = useAppSelector(
    (state) => state.transactions.transactions
  );
  // console.log(transdata, "lates data");

  const dummyTr: any = [
    {
      title: "Deposit from Card",
      date: "28 January 2021",
      price: "$850",
      isNeg: true,
      icon: "/assets/block-card-orange-icon.svg",
      color: "bg-orange-100",
    },
    {
      title: "Deposit Paypal",
      date: "25 January 2021",
      price: "+$2,500",
      isNeg: false,
      icon: "/assets/tr-icon-2.svg", // This should be a valid icon URL or import
      color: "bg-blue-100",
    },
    {
      title: "Jemi Wilson",
      date: "21 January 2021",
      price: "+$5,400",
      isNeg: false,
      icon: "/assets/tr-icon-3.svg", // This should be a valid icon URL or import
      color: "bg-green-100",
    },
  ];

  return (
    <div className="cards-container sm:w-full lg:w-[33%] text-nowrap center-content flex flex-col gap-4">
      <h1 className="flex grow page text-xl font-semibold text-colorBody-1 dark:text-gray-300">
        Recent Transactions
      </h1>
      <>
        {transdata.length > 0 ? (
          <div className="flex flex-col w-full gap-4 bg-white dark:bg-[#232328] rounded-3xl shadow-gray-50 text-colorBody-1 p-6">
            {transdata?.slice(-3).map((data, key: number) => (
              <div
                key={key}
                className="recentTr w-full flex gap-4 items-center justify-center"
              >
                <div
                  className={`{icon flex items-center rounded-full p-4 ${dummyTr[key].color}`}
                >
                  <Image
                    src={dummyTr[key].icon}
                    alt=""
                    width={28}
                    height={28}
                  />
                </div>
                <div className="flex flex-col gap-1 ">
                  <div className="title text-base text-black font-medium dark:text-gray-300">
                    {data.description}
                  </div>
                  <div className="date text-sm text-blue-900 opacity-70 dark:text-gray-400">
                    {data.date}
                  </div>
                </div>
                <div
                  className={`price flex grow justify-end text-base font-medium ${
                    data.amount < 0
                      ? "text-red-500 dark:text-red-400"
                      : "text-green-500 dark:text-green-400"
                  }`}
                >
                  {data.amount} ETB
                </div>
              </div>
            ))}
          </div>
        ) : (
          <div className="flex w-full">
            <ShimmerRecent />
          </div>
        )}
      </>
    </div>
  );
};

export default RecentTr;
