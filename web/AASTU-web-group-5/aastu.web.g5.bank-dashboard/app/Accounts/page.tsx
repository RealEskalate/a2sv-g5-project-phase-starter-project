'use client';
import React, { useEffect, useState } from 'react';
import Image from 'next/image';
import axios from 'axios';
import Card from '../components/common/card';
import ChartCard from './ChartCard';
import { balance as mockBalance, invoicesData as mockInvoicesData } from './mockData';
import Balance_img from '@/public/assests/icon/Accounts/Group494.png';
import Income_img from '@/public/assests/icon/Accounts/Group400.png';
import Expense_img from '@/public/assests/icon/Accounts/Group402.png';
import TotalSaving_img from '@/public/assests/icon/Accounts/Group401.png';
import Shopping from '@/public/assests/icon/Accounts/Group328.png';
import Service from '@/public/assests/icon/Accounts/Group327.png';
import Transfer from '@/public/assests/icon/Accounts/user2.png';
import { useSession } from 'next-auth/react';

const Accounts = () => {
  const [transactions, setTransactions] = useState<any[]>([]);
  const [balance, setBalance] = useState<number>(0);
  const [income, setIncome] = useState<number>(0);
  const [expense, setExpense] = useState<number>(0);
  const [netBalance, setNetBalance] = useState<number>(0);
  const [invoicesData, setInvoicesData] = useState<any[]>(mockInvoicesData); // Use mock data
  const { data: session } = useSession();

  
  const token: string =  ` Bearer ${session?.user?.accessToken} `;
  console.log(session,11111111,token)

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions?page=0&size=10', {
          headers: {
            Authorization: token,
          },
        });
        
        if (response.data.success) {
          setTransactions(response.data.data.reverse().slice(0,3)); // Store only the 3 most recent transactions
        }
      } catch (error) {
        console.error('Error fetching transactions:', error);
      }
    };

    const fetchBalanceData = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions/random-balance-history?monthsBeforeFirstTransaction=2', {
          headers: {
            Authorization: token,
          },
        });
        
        if (response.data.success) {
          const latestBalance = response.data.data[response.data.data.length - 1]?.value;
          setBalance(latestBalance);
          setNetBalance(latestBalance + income - expense); // Adjust according to your logic
        }
      } catch (error) {
        console.error('Error fetching balance data:', error);
      }
    };

    const fetchIncomes = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions/incomes?page=5&size=3', {
          headers: {
            Authorization: token,
          },
        });

        if (response.data.success) {
          setIncome(response.data.data.reduce((total: number, income: any) => total + income.amount, 0)); // Calculate total income
        }
      } catch (error) {
        console.error('Error fetching incomes:', error);
      }
    };

    const fetchExpenses = async () => {
      try {
        const response = await axios.get('https://bank-dashboard-6acc.onrender.com/transactions/expenses?page=0&size=3', {
          headers: {
            Authorization: token,
          },
        });

        if (response.data.success) {
          setExpense(response.data.data.reduce((total: number, expense: any) => total + expense.amount, 0)); // Calculate total expense
        }
      } catch (error) {
        console.error('Error fetching expenses:', error);
      }
    };

    fetchTransactions();
    fetchBalanceData();
    fetchIncomes();
    fetchExpenses();
  }, [income, expense]);

  return (
    <div className="bg-[#F5F7FA] space-y-8 mx-auto pt-3 px-4 md:px-8 lg:px-16 max-w-full overflow-hidden">
      {/* Balance and Overview */}
      <div className="grid grid-cols-2 sm:grid-cols-2 md:grid-cols-4 gap-4">
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Balance_img} alt="balance" />
          <div>
            <p className="text-sm md:text-lg font-semibold">My balance</p>
            <p className="text-base md:text-xl">${balance.toFixed(2)}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Income_img} alt="income" />
          <div>
            <p className="text-sm md:text-lg font-semibold">Income</p>
            <p className="text-base md:text-xl">${income.toFixed(2)}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={Expense_img} alt="expense" />
          <div>
            <p className="text-sm md:text-lg font-semibold">Expense</p>
            <p className="text-base md:text-xl">${expense.toFixed(2)}</p>
          </div>
        </div>
        <div className="p-4 bg-white rounded-lg flex items-center justify-center space-x-4">
          <Image height={44} width={44} src={TotalSaving_img} alt="total saving" />
          <div>
            <p className="text-sm md:text-lg font-semibold">Total saving</p>
            <p className="text-base md:text-xl">${netBalance.toFixed(2)}</p>
          </div>
        </div>
      </div>

      {/* Last Transactions and Card */}
      <div className="grid grid-cols-1 lg:grid-cols-3 gap-4">
      <div className="p-4 bg-white rounded-lg lg:col-span-2 space-y-7">
          <p className="text-lg font-semibold">Last transactions</p>
          {transactions.map((transaction, index) => (
            <div key={index} className="flex items-center pr-4 space-x-4 mb-4">
              <Image
                height={44}
                width={44}
                src={
                  transaction.type === 'shopping'
                    ? Shopping
                    : transaction.type === 'service'
                    ? Service
                    : Transfer
                }
                alt="transaction"
                className="rounded-full object-cover"
              />
              <div className="flex-1">
                <p className="font-semibold text-sm md:text-base">{transaction.description}</p>
                <p className="text-xs md:text-sm text-gray-500">{transaction.date}</p>
              </div>
              <p className="flex-1 text-xs md:text-sm break-words">{transaction.type}</p>
              <p className="flex-1 text-xs md:text-sm break-words">
                {transaction.transactionId ? `${transaction.transactionId.slice(0, 4)}***` : 'N/A'}
              </p>
              <p className="flex-1 text-xs md:text-sm break-words">{transaction.receiverUserName}</p>
              <p className="flex-1 text-xs md:text-sm break-words">
                {transaction.amount < 0 ? (
                  <span className="text-red-500">-${Math.abs(transaction.amount)}</span>
                ) : (
                  <span className="text-green-500">+${transaction.amount}</span>
                )}
              </p>
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
          <div className="p-1 bg-white rounded-lg shadow">
            <ChartCard />
          </div>
        </div>

        {/* Invoices Sent */}
        <div className="p-4 bg-gray-100 rounded-lg">
          <p className="text-lg font-semibold">Invoices sent</p>
          <div className="p-4 rounded-lg shadow h-[364px] bg-white">
            {invoicesData.slice(0, 4).map((data, index) => (
              <div key={index} className="flex items-center justify-between space-x-8 mb-4">
                <div className="flex items-center justify-between space-x-8 mb-4">
                  <Image height={44} width={44} src={data.image} alt="invoice" className="rounded-full object-cover" />
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
};

export default Accounts;