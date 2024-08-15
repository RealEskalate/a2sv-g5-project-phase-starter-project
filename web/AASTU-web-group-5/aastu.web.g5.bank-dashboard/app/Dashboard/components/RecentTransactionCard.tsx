import React from 'react';

export const RecentTransactionCard = () => {
  return (
    <div className="recent-transaction-card bg-white p-4 flex items-center gap-4">
      <div className='transaction-icon w-16 h-16 rounded-full flex items-center justify-center bg-gray-200'>
        <svg className='w-8 h-8 text-black' xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 256 256">
          <path fill="currentColor" d="M124 25.66V128a4 4 0 0 1-8 0V25.66L90.83 50.83a4 4 0 0 1-5.66-5.66l32-32a4 4 0 0 1 5.66 0l32 32a4 4 0 0 1-5.66 5.66Zm64 100.12V96a12 12 0 0 0-12-12h-16a4 4 0 0 0 0 8h16a4 4 0 0 1 4 4v92.9l-3.27-5a24 24 0 0 0-41.51 24.1a1.2 1.2 0 0 0 .12.19l22.26 34a4 4 0 1 0 6.69-4.38l-22.2-33.9A16 16 0 0 1 169.86 188l.12.19l10.67 16.31a4 4 0 0 0 7.35-2.19v-66.84a78.83 78.83 0 0 1 32 63.18V240a4 4 0 0 0 8 0v-41.35a86.84 86.84 0 0 0-40-72.87M80 84H64a12 12 0 0 0-12 12v104a4 4 0 0 0 8 0V96a4 4 0 0 1 4-4h16a4 4 0 0 0 0-8" />
        </svg>
      </div>
      <div className='transaction-name flex flex-col'>
        <p className='text-lg font-medium text-[#232323]'>Deposit from my Card</p>
        <p className='text-sm font-medium text-[#718EBF]'>28 January 2021</p>
      </div>
      <div className='text-red-500 text-lg'>
        -$850
      </div>
    </div>
  );
};
