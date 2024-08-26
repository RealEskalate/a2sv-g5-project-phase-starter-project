import { useEffect, useState } from 'react';
import { CurrencyDollarIcon } from "@heroicons/react/24/outline";
import { getExpenses } from "@/services/transactionfetch";
import { formatDistanceToNowStrict } from 'date-fns';

import { TbFileSad } from "react-icons/tb";
interface Transaction {
  id: number;
  receiverUserName: string;
  date: string;
  amount: number;
  description: string;
}

const TransactionList: React.FC = () => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [status, setStatus] = useState<'loading' | 'error' | 'success'>('loading');

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const transactionData = await getExpenses(0, 5);

        if (Array.isArray(transactionData.data.content)) {
          setTransactions(transactionData.data.content);
          setStatus('success');
        } else {
          console.error("Transaction data is not an array");
          setStatus('error');
        }
      } catch (error) {
        console.error("Failed to fetch transactions", error);
        setStatus('error');
      }
    };

    fetchTransactions();
  }, []);

  const formatTimeSince = (date: string) => {
    return formatDistanceToNowStrict(new Date(date));
  };

  if (status === 'loading') {
    return (
      <div className="flex-1 py-3 flex flex-col justify-between  dark:bg-dark rounded-lg shadow-md p-4 space-y-4">
        {[...Array(5)].map((_, index) => (
          <div key={index} className="flex items-center dark:bg-dark justify-between animate-pulse">
            <div className="w-10 h-10 rounded-full bg-gray-200"></div>
            <div className="flex-1 px-3">
              <div className="h-4 bg-gray-200 rounded w-3/4 mb-2"></div>
              <div className="h-4 bg-gray-200 rounded w-1/2"></div>
            </div>
            <div className="h-4 bg-gray-200 rounded w-12"></div>
          </div>
        ))}
      </div>
    );
  }

  if (status === 'error') {
    return (
      <div className="p-3 gap-4  flex flex-col justify-center items-center h-auto  dark:bg-dark   text-center ">
        <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
        <p className="text-red-500" >Failed to fetch</p>
      </div>
    );
  }

  if (transactions.length === 0) {
    return <div className="text-center text-gray-500">No transactions to display.</div>;
  }

  return (
    <div className="flex-1 py-3 flex flex-col justify-between bg-white rounded-lg shadow-md p-4 space-y-4">
      {transactions.map((transaction) => (
        <div key={transaction.id} className="flex items-center justify-between">
          <div className="w-10 h-10 flex items-center justify-center rounded-full bg-blue-100">
            <CurrencyDollarIcon className="h-8 w-8 text-blue-700"/>
          </div>
          <div className="flex-1 px-3">
            <div className="text-gray-800 font-medium">{transaction.receiverUserName}</div>
            <div className="text-gray-400 text-sm">{formatTimeSince(transaction.date)}</div>
          </div>
          <div className="text-gray-800 font-semibold">${transaction.amount}</div>
        </div>
      ))}
    </div>
  );
};

export default TransactionList;
