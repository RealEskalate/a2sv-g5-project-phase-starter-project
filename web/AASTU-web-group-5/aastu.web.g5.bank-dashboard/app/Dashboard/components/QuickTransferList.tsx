import React from 'react';
import { QuickTransferCard } from './QuickTransferCard';

export const QuickTransferList = () => {
  return (
    <div className='container rounded-[20px] flex flex-col gap-4 p-4'>
      <div className='flex flex-col md:flex-row justify-between items-center'>
        <div className='flex flex-row gap-4 overflow-x-auto'>
          <QuickTransferCard />
          <QuickTransferCard />
          <QuickTransferCard />
        </div>
        <svg className='w-6 h-6 shadow-md mt-2 md:mt-0' xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24">
          <g fill="none" stroke="black" strokeLinecap="round" strokeLinejoin="round" strokeWidth="1.5">
            <path d="m11 8.5l3.5 3.5l-3.5 3.5" />
            <path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10" />
          </g>
        </svg>
      </div>
      <div className="flex flex-col md:flex-row items-center gap-4">
        <p className="text-sm font-bold w-full md:w-1/2">Write Amount</p>
        <div className="relative w-full md:w-1/2">
          <input
            type="text"
            className="w-full p-2 rounded-[50px] border border-gray-300 pr-12"
            placeholder='525.50'
          />
          <button type="submit" className="absolute right-0 top-0 bottom-0 text-white flex items-center bg-blue-700 rounded-[50px] py-2 px-4">
            Send
            <svg className='w-6 h-6 ml-2' xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256">
              <path fill="currentColor" d="M226.27 29.22a5 5 0 0 0-5.1-.87L18.51 107.66a10.22 10.22 0 0 0 1.75 19.56L76 138.16V200a12 12 0 0 0 7.51 11.13A12.1 12.1 0 0 0 88 212a12 12 0 0 0 8.62-3.68l28-29l43 37.71a12 12 0 0 0 7.89 3a12.5 12.5 0 0 0 3.74-.59a11.87 11.87 0 0 0 8-8.72l40.62-176.6a5 5 0 0 0-1.6-4.9M20 117.38a2.13 2.13 0 0 1 1.42-2.27l174.65-68.35l-117 83.85l-57.26-11.24a2.12 2.12 0 0 1-1.81-1.99m70.87 85.38A4 4 0 0 1 84 200v-56.3l34.58 30.3Zm88.58 6.14a4 4 0 0 1-6.57 2.09l-86.45-75.81l131.7-94.38Z" />
            </svg>
          </button>
        </div>
      </div>
    </div>
  );
}
