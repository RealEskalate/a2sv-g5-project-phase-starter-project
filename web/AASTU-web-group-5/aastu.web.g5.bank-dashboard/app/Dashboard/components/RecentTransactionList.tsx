import React from 'react';
import { RecentTransactionCard } from './RecentTransactionCard';

export const RecentTransactionList = () => {
  return (
    <div className='transaction-list flex flex-col gap-4 rounded-lg bg-white p-4 w-full'>
      <RecentTransactionCard />
      <RecentTransactionCard />
      <RecentTransactionCard />
      <RecentTransactionCard />
    </div>
  );
};
