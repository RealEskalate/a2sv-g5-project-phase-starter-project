'use client'

// src/pages/transactions/index.tsx

import React from 'react';
import TransactionTable from '@/components/TransactionTable'
import RecentTransactions from '@/components/RecentTransaction';

const Transaction: React.FC = () => {
  return (
    <div className="p-4">
      <h1 className="text-2xl font-bold text-center mb-4">Recent Transactions</h1>
      <RecentTransactions />
    </div>
  );
};

export default Transaction;
