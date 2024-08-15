import React from "react";
import {
  MdHome,
  MdSettings,
  MdAttachMoney,
  MdAccountBalance,
} from "react-icons/md";
import ListCard from "./components/ListCard";
import { IconType } from "react-icons";

type DataItem = {
  heading: string;
  text: string;
  headingStyle: string;
  dataStyle: string;
};

type Column = {
  icon: IconType;
  iconStyle: string;
  data: DataItem[];
};

const Page = () => {
  // Example data for the first ListCard
  const column1: Column = {
    icon: MdHome,
    iconStyle: "text-[#FFBB38] bg-[#FFF5D9]",
    data: [
      {
        heading: "My Balance",
        text: "$12,750",
        headingStyle: "text-sm font-bold text-nowrap text-[#718EBF]",
        dataStyle: "text-xs text-nowrap",
      },
    ],
  };

  // Example data for the second ListCard
  const column2: Column = {
    icon: MdAttachMoney, // Updating the icon
    iconStyle: "text-[#396AFF] bg-[#E7EDFF]", // Updating the iconStyle
    data: column1.data.map((item) => ({
      ...item,
      heading: "Income", // Updating the heading
    })),
  };

  // Example data for the third ListCard
  const column3: Column = {
    icon: MdSettings, // Updating the icon
    iconStyle: "text-[#FF82AC] bg-[#FFE0EB]", // Updating the iconStyle
    data: column1.data.map((item) => ({
      ...item,
      heading: "Expense", // Updating the heading
    })),
  };

  // Example data for the fourth ListCard
  const column4: Column = {
    icon: MdAccountBalance, // Updating the icon
    iconStyle: "text-[#16DBCC] bg-[#DCFAF8]", // Updating the iconStyle
    data: column1.data.map((item) => ({
      ...item,
      heading: "Total Savings", // Updating the heading
    })),
  };

  // First column with multiple data items
  const column5: Column = {
    icon: MdHome,
    iconStyle: "text-[#FFBB38] bg-[#FFF5D9]",
    data: [
      {
        heading: "Spotify Subscription",
        text: "25 Jan 2021",
        headingStyle: "text-sm font-bold text-nowrap",
        dataStyle: "text-xs text-nowrap text-[#718EBF]",
      },
      {
        heading: "-$150",
        text: "",
        headingStyle: "text-xs font-bold text-[#FE5C73]",
        dataStyle: "text-xs text-nowrap",
      },
    ],
  };
  return (
    <>
      <div className="flex flex-col h-svh bg-[#F5F7FA] px-3 py-3 gap-5">
        <div>
          <div className="flex flex-wrap gap-2">
            <ListCard column={column1} width={"w-[48%] md:w-[24.3%]"} />
            <ListCard column={column2} width={"w-[48%] md:w-[24.3%]"} />
            <ListCard column={column3} width={"w-[48%] md:w-[24.3%]"} />
            <ListCard column={column4} width={"w-[48%] md:w-[24.3%]"} />
          </div>
        </div>
        <div className="flex flex-col gap-5">
          <span>Last Transaction</span>
          <div className="bg-white flex flex-col justify-between rounded-2xl">
            <ListCard column={column5} width={"w-full"} />
            <ListCard column={column5} width={"w-full"} />
            <ListCard column={column5} width={"w-full"} />
            <ListCard column={column5} width={"w-full"} />
          </div>
        </div>
      </div>
    </>
  );
};

export default Page;
