'use client'

import React, { useEffect, useState } from 'react';
import { colors } from '@/constants';
import { getAllTransactionsss } from '@/services/transactionfetch';
import LifeInsuranceIcon from '@/public/icons/LifeInsuranceIcon';
import ShoppingIcon from '@/public/icons/ShoppingIcon';
import SavingAccountsIcon from '@/public/icons/SavingAccountsIcon';
import BusinessLoans from '@/public/icons/BusinessLoans';

const RecentTransaction = () => {
  const [recentTransaction, setRecentTransaction] = useState<any[]>([]);

  useEffect(() => {
    const fetchRecentTransaction = async () => {
      try {
        const response = await getAllTransactionsss();
        setRecentTransaction(response.data.content);
      } catch (error) {
        console.error("Error fetching the recent transactions: ", error);
      }
    };
    fetchRecentTransaction();
  }, []);

  // Get the last 3 transactions
  const lastThreeTransactions = recentTransaction.slice(-3).reverse();

  return (
    <div className="max-w-auto">
      {lastThreeTransactions.map((transaction) => {
        // Determine icon and amount color based on transaction type
        let IconComponent = LifeInsuranceIcon; // Default icon for other types
        let amountTextColor = 'text-green-500';
        let formattedAmount = `$${transaction.amount.toLocaleString()}`;

        // Check if the amount exceeds one million
        if (transaction.amount > 1000000) {
          formattedAmount = `${formattedAmount.slice(0, formattedAmount.indexOf(',') + 6)}...`;
        }

        switch (transaction.type) {
          case 'transfer':
            IconComponent = LifeInsuranceIcon;
            amountTextColor = 'text-red-500';
            formattedAmount = `-${formattedAmount}`;
            break;
          case 'shopping':
            IconComponent = ShoppingIcon;
            amountTextColor = 'text-red-500';
            formattedAmount = `-${formattedAmount}`;
            break;
          case 'service':
            IconComponent = SavingAccountsIcon;
            amountTextColor = 'text-red-500';
            formattedAmount = `-${formattedAmount}`;
            break;
          case 'deposit':
            IconComponent = BusinessLoans;
            amountTextColor = 'text-green-500';
            formattedAmount = `+${formattedAmount}`;
            break;
          default:
            break;
        }

        return (
          <div key={transaction.transactionId} className={`${colors.white} max-w-auto`}>
            <div className='flex p-2 justify-between w-auto md:w-auto'>
              <div className='flex gap-4'>
                <div className={`${colors.white} rounded-full flex items-center justify-center p-4 h-20 w-20`}>
                  <IconComponent className="w-12 h-12" /> {/* Adjust size as needed */}
                </div>
                <div className='my-2'>
                  <p className='text-lg font-medium'>{transaction.senderUserName}</p>
                  <p className={`${colors.textgray} text-sm text-start`}>{transaction.date}</p>
                </div>
              </div>
              <div className='my-5'>
                <p className={`${amountTextColor}`}>{formattedAmount}</p>
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
};

export default RecentTransaction;
