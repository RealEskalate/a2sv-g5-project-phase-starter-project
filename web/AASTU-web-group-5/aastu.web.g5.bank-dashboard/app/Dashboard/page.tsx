"use client"
import React from 'react';
import { RecentTransactionCard } from "./components/RecentTransactionCard";
import { QuickTransferList } from './components/QuickTransferList';
import { PieChartComponent } from './components/PieChartComponent';
import { BarchartComponent } from './components/BarchartComponent';
import { LineGraphComponent } from './components/LineGraphComponent';
import Link from 'next/link';
import Card from '../components/common/card';
import Chip_card1 from "@/public/assets/image/Chip_Card1.png";
import Chip_card3 from "@/public/assets/image/Chip_Card3.png";

const creditCardColor = {
  cardOne: {
    cardBgColor: "bg-blue-500 rounded-3xl text-white",
    bottomBgColor: "flex justify-between p-4 bg-blue-400 rounded-bl-3xl rounded-br-3xl",
    imageCreditCard: Chip_card1,
    grayCircleColor: false,
  },
  cardThree: {
    cardBgColor: "bg-white rounded-3xl text-black",
    bottomBgColor: "",
    imageCreditCard: Chip_card3,
    grayCircleColor: true,
  },
};

function Dashboard() {
  return (
    <div className='flex flex-col bg-[#f9f9f9]'>
      {/* Main content */}
      <div className='flex-1 flex flex-col'>
        {/* Top Content */}
        <div className='pt-16 md:pt-20 px-6 py-12 flex flex-col gap-6'>
          {/* Main Content Layout */}
          <div className='flex flex-col gap-6'>
            {/* First Row: My Cards and Recent Transactions */}
            <div className='flex flex-col md:flex-row gap-6'>
              <div className='flex-1 p-4 rounded-lg'>
                <div className='text-[#343C6A] border-blue-500'>
                  <div className="flex items-center justify-between">
                    <p className="text-lg font-semibold leading-6 text-left text-[#343C6A]">My Cards</p>
                    <Link href="/Transactions" className="text-lg font-semibold leading-6 text-right text-[#343C6A]">See All</Link>
                  </div>
                  <div className='overflow-x-auto mt-4'>
                    <div className='flex gap-4'>
                      <Card creditCardColor={creditCardColor.cardOne} />
                      <Card creditCardColor={creditCardColor.cardThree} />
                    </div>
                  </div>
                </div>
              </div>
              <div className='flex-1 p-4 rounded-lg'>
                
                  <RecentTransactionCard/>
                
              </div>
            </div>

            {/* Second Row: Weekly Activities and Expense Statistics */}
            <div className='flex flex-col md:flex-row gap-6'>
              <div className='flex-1 bg-white rounded-lg'>
                <div className='text-[#343C6A]'>
                  Weekly Activities
                  <div className=''>
                    <BarchartComponent />
                  </div>
                </div>
              </div>
              <div className='flex-1 bg-white p-4 rounded-lg'>
                <div className='text-[#343C6A]'>
                  Expense Statistics
                  <div className='h-64 flex items-center justify-center'>
                    <PieChartComponent />
                  </div>
                </div>
              </div>
            </div>

            {/* Third Row: Quick Transfer and Balance History */}
            <div className='flex flex-col md:flex-row gap-6'>
              <div className='flex-1 bg-white p-4 rounded-lg'>
                <div className='text-[#343C6A]'>
                  Quick Transfer
                  <QuickTransferList />
                </div>
              </div>
              <div className='flex-1 bg-white rounded-lg'>
                <div className='text-[#343C6A]'>
                  Balance History
                  <div className=''>
                    <LineGraphComponent />
                  </div>
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
