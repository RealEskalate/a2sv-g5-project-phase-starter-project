import Image from "next/image";
import ImageComponent from "./components/ImageComponent";
import Reviving from "./components/QuickTransfer";
import { BalanceHistory } from "./components/BalanceHistory";
import { WeeklyActivity } from "./components/WeeklyActivity";
import { ExpenseStatistics } from "./components/ExpenseStatistics";
import RecentTransaction from "./components/RecentTransaction";
import CreditCard from "./components/CreditCard";
// import {RecentTransaction} from "@/components/RecentTransaction"

export default function Home() {
  return (
    <div className="flex flex-col">
      {/* Mobile Version */}
      <div className="flex flex-col md:hidden">
      <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div>
        <div className="flex overflow-x-auto [&::-webkit-scrollbar]:hidden">

          <div className="flex-col">

            <div className="flex">
              <div className="min-w-max min-h-max">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  filterClass=""
                  bgColor="from-[#4C49ED] to-[#0A06F4]"
                  textColor="text-white"
                  iconBgColor="bg-opacity-10"
                  showIcon={true}
                />
              </div>
              <div className="min-w-max min-h-max [&::-webkit-scrollbar]:hidden">
                <CreditCard
                  balance="$5,756"
                  cardHolder="Eddy Cusuma"
                  validThru="12/22"
                  cardNumber="3778 **** **** 1234"
                  filterClass=""
                  bgColor="bg-white"
                  textColor="text-black"
                  iconBgColor="bg-black"
                  showIcon={true}
                />
              </div>
            </div>
          </div>
        </div>
        <RecentTransaction />
        <WeeklyActivity />
        <ExpenseStatistics />
        <Reviving />
        <BalanceHistory />
      </div>

      {/* Web Version */}
      <div className="hidden md:flex flex-col space-y-4 px-6 py-4 bg-white">
           {/* <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div> */}
        <div className="flex">
          <div className="flex flex-col w-3/4">
            <div className="flex items-center justify-between">
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-2xl">My Cards</h1>
              <h1 className="mx-4 my-4 font-bold text-[#343C6A] text-lg">See All</h1>
            </div>
            <div className="flex space-x-6">
              <CreditCard
                balance="$5,756"
                cardHolder="Eddy Cusuma"
                validThru="12/22"
                cardNumber="3778 **** **** 1234"
                bgColor="from-[#4C49ED] to-[#0A06F4]"
                textColor="text-white"
                iconBgColor="bg-opacity-10"
                showIcon={true}
              />
              <CreditCard
                balance="$5,756"
                cardHolder="Eddy Cusuma"
                validThru="12/22"
                cardNumber="3778 **** **** 1234"
                bgColor="bg-white"
                textColor="text-black"
                iconBgColor="bg-black"
                showIcon={true}
              />
            </div>
          </div>
          
          <div className="flex flex-col justify-between">
            <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl"> Recent Transaction</h1>
            <RecentTransaction />
          </div>
        </div>
        <div className="flex space-x-6">
          <div className=" w-3/4">
            <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl"> Weekly Activity</h1>
            <WeeklyActivity />
          </div>
          {/* <WeeklyActivity className="flex-1" /> */}
          <div className=" w-1/4">
          <h1 className="flex mx-4 my-4 font-bold  text-[#343C6A] text-2xl"> Expense Statistics</h1>

          <ExpenseStatistics  />
          </div>
        </div>
        <div className="flex justify-between space-x-6 w-full h-24">
          <div className=" w-1/3 ">
            <Reviving />
          </div>
          <div className="w-2/3 h-5" >
          <BalanceHistory />
          </div>
        </div>
      </div>
    </div>
  );
}
