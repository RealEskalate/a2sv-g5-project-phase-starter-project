<<<<<<< HEAD
import BalanceHistory from "@/components/Charts/BalanceHistory";
import DebiteAndCredit from "@/components/Charts/DebiteAndCredit";
import ExpenseStatistics from "@/components/Charts/ExpenseStatistics";
import WeeklyActivity from "@/components/Charts/WeeklyActivity";
import MyCard from "@/components/MyCard/MyCard";
import QuickTransfer from "@/components/QuickTransfer/QuickTransfer";
import RecentTransaction from "@/components/RecentTransaction/RecentTransaction";
import { Plus } from "lucide-react";
import StoreProvider from "@/providers/StoreProvider";
import TestNaol from "@/components/TestNaol";

export default function Home() {
  return (
    <>
      <div className="w-full lg:flex ">
        <div className="lg:w-2/3 md:pr-3 xl:pr-5 flex-shrink">
          <div className="w-full">
            <div className="flex justify-between">
              <p className="text-[#333B69] pb-3 font-semibold">My Card</p>
              <p className="text-[#333B69] pb-3 font-semibold">See All</p>
            </div>
            <div className="flex  overflow-x-auto space-x-2">
              <StoreProvider>
                <MyCard />
                <MyCard />
                <TestNaol />
              </StoreProvider>
              <div className="w-[295px] h-[175px] bg-gray-200 rounded-3xl justify-center items-center flex flex-shrink-0">
                <Plus size={32} />
              </div>
            </div>
          </div>
        </div>
        <div className="lg:w-1/3 w-full">
          <RecentTransaction />
        </div>
      </div>
      <div className="md:flex my-5">
        <WeeklyActivity />
        {/* <DebiteAndCredit /> */}
        <ExpenseStatistics />
      </div>
      <div className="md:flex justify-between">
        <div className="w-5/12 pe-6">
          <QuickTransfer />
        </div>
        <BalanceHistory />
      </div>
    </>
=======
import Link from 'next/link';

export default function Home() {
  return (
    <div className='w-screen min-h-screen flex justify-center items-center bg-slate-100'>
      <div>
        <div className='my-5'>
          <Link href='/api/auth/signin'>
            <button className='min-w-24 px-4 py-2 bg-indigo-900 text-xl text-gray-50 rounded-2xl'>
              SignIn
            </button>
          </Link>
        </div>
        <div>
          <Link href='/api/auth/signup'>
            <button className=' min-w-24 px-4 py-2 bg-indigo-900 text-xl text-gray-50 rounded-2xl'>
              SignUp
            </button>
          </Link>
        </div>
      </div>
    </div>
>>>>>>> 6df34e5b8006021ff0bbe88e637e8b730943777d
  );
}
