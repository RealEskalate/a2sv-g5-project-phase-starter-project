import React from 'react';
import Image from 'next/image';
import Card from '../components/common/card';
import ChartCard from './ChartCard';
import { balance, income, expense, netBalance, transaction, invoicesData } from './mockData';
import Balance_img from '@/public/assests/icon/Accounts/Group494.png';
import Income_img from '@/public/assests/icon/Accounts/Group400.png';
import Expense_img from '@/public/assests/icon/Accounts/Group402.png';
import TotalSaving_img from '@/public/assests/icon/Accounts/Group401.png';

const Accounts = () => {
  // Get the 3 most recent transactions
  const recentTransactions = transaction.slice(0, 3);

  // Get the 4 most recent invoices
  const recentInvoices = invoicesData.slice(0, 4);

  return (
    <div className=" bg-[#F5F7FA] space-y-8 mx-auto pt-3 px-4 md:px-8 lg:px-16 max-w-full overflow-hidden">
      {/* Balance and Overview */}
      <div className="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-4 gap-4">
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Balance_img} alt='balance' />
          <div>
            <p className="text-sm md:text-lg font-semibold">My balance</p>
            <p className="text-base md:text-xl">${balance}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Income_img} alt='income' />
          <div>
            <p className="text-sm md:text-lg font-semibold">Income</p>
            <p className="text-base md:text-xl">${income}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Expense_img} alt='expense' />
          <div>
            <p className="text-sm md:text-lg font-semibold">Expense</p>
            <p className="text-base md:text-xl">${expense}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={TotalSaving_img} alt='total saving' />
          <div>
            <p className="text-sm md:text-lg font-semibold">Total saving</p>
            <p className="text-base md:text-xl">${netBalance}</p>
          </div>
        </div>
      </div>

      {/* Last Transactions and Card */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
        <div className="p-4 bg-white rounded-lg lg:col-span-2">
          <p className="text-lg font-semibold">Last transactions</p>
          {recentTransactions.map((transaction, index) => (
            <div key={index} className="flex items-center space-x-4 mb-4">
              <Image height={44} width={44} src={transaction.image} alt='transaction' className="rounded-full object-cover" />
              <div className="flex-1">
                <p className="font-semibold text-sm md:text-base">{transaction.name}</p>
                <p className="text-xs md:text-sm text-gray-500">{transaction.date}</p>
              </div>
              <p className="flex-1 text-xs md:text-sm break-words">{transaction.type}</p>
              <p className="flex-1 text-xs md:text-sm break-words">{transaction.number}</p>
              <p className="flex-1 text-xs md:text-sm break-words">{transaction.status}</p>
              <p className="flex-1 text-xs md:text-sm break-words">{transaction.amount}</p>
            </div>
          ))}
        </div>

        {/* Hidden on small screens */}
        <div className="hidden lg:block bg-white rounded-lg">
          <div className="flex justify-between p-2">
            <p className="text-lg font-semibold">My card</p>
            <p>See all</p>
          </div>
          <Card />
        </div>
      </div>

      {/* Debit and Credit Overview */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
        <div className="p-4 bg-gray-100 rounded-lg lg:col-span-2">
          <p className="text-lg font-semibold">Debit and credit overview</p>
          <div className="p-1  bg-white rounded-lg shadow">
           
              {/* Implement the bar graph here */}
              <ChartCard />
            
          </div>
        </div>

        {/* Invoices Sent */}
        <div className="p-4 bg-gray-100 rounded-lg">
          <p className="text-lg font-semibold">Invoices sent</p>
          <div className="p-4 rounded-lg shadow h-[364px] bg-white">
            {recentInvoices.map((data, index) => (
              <div key={index} className="flex items-center justify-between space-x-8 mb-4">
                <div className='flex items-center justify-between space-x-8 mb-4'>
                <Image height={44} width={44} src={data.image} alt='invoice' className="rounded-full object-cover" />
                <div>
                  <p className="font-semibold text-sm md:text-base">{data.name}</p>
                  <p className="text-xs md:text-sm text-gray-500">{data.date}</p>
                </div>
                </div>
                <p className="font-semibold text-sm md:text-base">{data.amount}</p>
            </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}

export default Accounts;
