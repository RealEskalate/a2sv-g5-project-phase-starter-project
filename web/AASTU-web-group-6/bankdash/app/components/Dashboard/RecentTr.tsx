import React from "react";
import Image from "next/image";
import { PieComp } from "../Charts/PieComp";
const RecentTr = () => {
  const dummyTr: any = [
    {
      title: "Deposit from my Card",
      date: "28 January 2021",
      price: "$850",
      isNeg: true,
      icon: "/assets/tr-icon-1.svg",
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
    <div className="cards-container text-nowrap center-content flex flex-col gap-4 sm:scale-75 md:scale-90 lg:scale-100">
      <h1 className="flex grow page text-xl font-semibold text-colorBody-1">
        Recent Transactions
      </h1>
      <div className="flex flex-col w-full gap-4 bg-white rounded-3xl shadow-gray-50 text-colorBody-1 p-6">
        {dummyTr.map((data: any, key: number) => (
          <div
            key={key}
            className="recentTr w-full flex gap-4 items-center justify-center"
          >
            <div className={`{icon rounded-full p-4 ${data.color}`}>
              <Image src={data.icon || ""} alt="" width={28} height={42} />
            </div>
            <div className="flex flex-col gap-1">
              <div className="title text-base text-black font-medium">
                {data.title}
              </div>
              <div className="date text-sm text-blue-900 opacity-70">
                {data.date}
              </div>
            </div>
            <div
              className={`price flex grow justify-end font-medium ${
                data.isNeg ? "text-red-500" : "text-green-500"
              }`}
            >
              {data.price}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default RecentTr;
