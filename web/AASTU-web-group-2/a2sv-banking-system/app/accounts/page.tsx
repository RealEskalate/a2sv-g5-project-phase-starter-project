import React from "react";
import {
  MdHome,
  MdSettings,
  MdAttachMoney,
  MdAccountBalance,
} from "react-icons/md";
import ListCard from "./components/ListCard";
import { IconType } from "react-icons";
import BarChartForAccounts from "./components/BarChartForAccounts";
import Card from "../components/Page2/Card";

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
  const ReusableCard: Column = {
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
  const card1: Column = {
    icon: MdAttachMoney, // Updating the icon
    iconStyle: "text-[#396AFF] bg-[#E7EDFF]", // Updating the iconStyle
    data: ReusableCard.data.map((item) => ({
      ...item,
      heading: "Income", // Updating the heading
    })),
  };

  // Example data for the third ListCard
  const card2: Column = {
    icon: MdSettings, // Updating the icon
    iconStyle: "text-[#FF82AC] bg-[#FFE0EB]", // Updating the iconStyle
    data: ReusableCard.data.map((item) => ({
      ...item,
      heading: "Expense", // Updating the heading
    })),
  };

  // Example data for the fourth ListCard
  const card3: Column = {
    icon: MdAccountBalance, // Updating the icon
    iconStyle: "text-[#16DBCC] bg-[#DCFAF8]", // Updating the iconStyle
    data: ReusableCard.data.map((item) => ({
      ...item,
      heading: "Total Savings", // Updating the heading
    })),
  };

  // First column with multiple data items
  const ReusableLastTransaction: Column = {
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

  // First transaction example
  const transaction1: Column = {
    icon: MdAccountBalance, // Different icon
    iconStyle: "text-[#16DBCC] bg-[#DCFAF8]", // Different iconStyle
    data: ReusableLastTransaction.data.map((item, index) => ({
      ...item,
      heading: index === 0 ? "Mobile Services" : item.heading, // Custom heading for the first item
    })),
  };

  const transaction2: Column = {
    icon: MdAttachMoney, // Updating the icon
    iconStyle: "text-[#16DBCC] bg-[#DCFAF8]", // Updating the iconStyle
    data: ReusableLastTransaction.data.map((item, index) => ({
      ...item,
      heading: index === 0 ? "Emilly Wilson " : "+$780",
      headingStyle:
        index === 0 ? item.headingStyle : "text-xs font-bold text-[#16DBAA]",
    })),
  };

  return (
    <>
      <div className="flex flex-col h-full bg-[#F5F7FA] px-3 py-3 gap-5">
        <div>
          <div className="flex flex-wrap gap-2">
            <ListCard column={ReusableCard} width={"w-[48%] md:w-[23%]"} />
            <ListCard column={card1} width={"w-[48%] md:w-[23%]"} />
            <ListCard column={card2} width={"w-[48%] md:w-[23%]"} />
            <ListCard column={card3} width={"w-[48%] md:w-[23%]"} />
          </div>
        </div>

        <div className="flex flex-col md:flex-row gap-5">
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl text-[#333B69] font-semibold">
              Last Transaction
            </span>
            <div className="bg-white flex flex-col justify-between rounded-2xl">
              <ListCard column={ReusableLastTransaction} width={"w-full"} />
              <ListCard column={transaction1} width={"w-full"} />
              <ListCard column={transaction2} width={"w-full"} />
            </div>
          </div>
          <div className="md:w-1/2 gap-1 flex flex-col">
            <div className="flex justify-between mr-2">
              <span className="text-xl text-[#333B69] font-semibold">
                My Card
              </span>
              <span className="text-sm text-[#333B69] font-semibold">
                See All
              </span>
            </div>
            <Card
              balance="$5,756"
              cardHolder="Eddy Cusuma"
              validThru="12/22"
              cardNumber="3778 **** **** 1234"
              filterClass=""
              bgColor="from-[#4C49ED] to-[#0A06F4]"
              textColor="text-white"
              iconBgColor="bg-opacity-10"
              showIcon={true}
            ></Card>
          </div>
        </div>

        <div className="flex flex-col md:flex-row gap-5">
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl text-[#333B69] font-semibold">
              Debit & Credit Overview
            </span>
            <BarChartForAccounts></BarChartForAccounts>
          </div>
          <div className="flex flex-col gap-5 md:w-1/2">
            <span className="text-xl text-[#333B69] font-semibold">
              Invoice Sent
            </span>
            <div className="bg-white flex flex-col justify-between rounded-2xl">
              <ListCard column={ReusableLastTransaction} width={"w-full"} />
              <ListCard column={transaction1} width={"w-full"} />
              <ListCard column={transaction2} width={"w-full"} />
              <ListCard column={transaction2} width={"w-full"} />
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default Page;
