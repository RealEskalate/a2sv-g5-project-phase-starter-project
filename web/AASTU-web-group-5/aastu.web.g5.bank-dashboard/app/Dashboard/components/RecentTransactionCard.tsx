"use client"

import React from 'react';
import Image from 'next/image'; // Assuming you're using Next.js
import { RecentTransaction } from './mockData'; // Adjust the path as needed
import { Slice } from 'lucide-react';

export const RecentTransactionCard = () => {
  const slicedTransactions = RecentTransaction.slice(0, 3);

  return (
    <div>
      <p>Recent Transactions</p>
      <div>
        {slicedTransactions.map((transaction, index) => (
          <div key={index} className="flex items-center justify-between space-x-8 mb-4">
            <div className="flex items-center justify-between space-x-8 mb-4">
              <Image height={44} width={44} src={transaction.image} alt="invoice" className="rounded-full object-cover" />
              <div>
                <p className="font-semibold text-sm md:text-base">{transaction.name}</p>
                <p className="text-xs md:text-sm text-gray-500">{transaction.date}</p>
              </div>
            </div>
            {transaction.amount[0]=="+" ? 
            <p className="font-semibold text-green-600 text-sm md:text-base">{transaction.amount}</p>
            : <p className="font-semibold text-red-700 text-sm md:text-base">{transaction.amount}</p> }
          </div>
        ))}
      </div>
    </div>
  );
};


export default RecentTransactionCard;