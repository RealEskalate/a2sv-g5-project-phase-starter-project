import React from 'react';
import { RecentTransactionList } from './components/RecentTransactionList';
import { QuickTransferList } from './components/QuickTransferList';
import { PieChartComponent } from './components/PieChartComponent';
import { BarchartComponent } from './components/BarchartComponent';
import { LineGraphComponent } from './components/LineGraphComponent';
import Link from 'next/link';

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
              <div className='flex-1  p-4 rounded-lg'>
                <div className='text-[#343C6A] border-blue-500'>
                  <div className="flex items-center justify-between">
                    <p className="text-lg font-semibold leading-6 text-left text-[#343C6A]">My cards</p>
                    <Link href="/Transactions" className="text-lg font-semibold leading-6 text-right text-[#343C6A]">See All</Link>
                  </div>
                  <div className='flex gap-2 bg-white rounded-[25px] mt-4'>
                    <div>card1</div>
                    <div>card2</div>
                  </div>
                </div>
              </div>
              <div className='flex-1  p-4 rounded-lg'>
                <div className='text-[#343C6A]'>
                  Recent transactions
                  <RecentTransactionList />
                </div>
              </div>
            </div>

            {/* Second Row: Weekly Activities and Expense Statistics */}
            <div className='flex flex-col md:flex-row gap-6'>
              <div className='flex-1 bg-white rounded-lg'>
                <div className='text-[#343C6A]'>
                  Weekly activities
                  <div className=''>
                    <BarchartComponent />
                  </div>
                </div>
              </div>
              <div className='flex-1 bg-white p-4 rounded-lg'>
                <div className='text-[#343C6A]'>
                  Expense statistics
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
                  Quick transfer
                  <QuickTransferList />
                </div>
              </div>
              <div className='flex-1 bg-white rounded-lg'>
                <div className='text-[#343C6A]'>
                  Balance history
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
