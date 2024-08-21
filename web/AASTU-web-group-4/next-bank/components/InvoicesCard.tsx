import { useEffect, useState } from 'react';
import { BanknotesIcon } from "@heroicons/react/24/outline";
import { getExpenses } from "@/services/transactionfetch";
import { formatDistanceToNowStrict } from 'date-fns';

interface Transaction {
  receiverUserName: string;
  type: string;
  amount: string;
  date: string;
}

const InvoicesCard: React.FC = () => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const transactionData = await getExpenses(0, 5); 

        if (Array.isArray(transactionData.data.content)) {
          setTransactions(transactionData.data.content);
        } else {
          console.error("Transaction data is not an array");
        }
      } catch (error) {
        console.error("Failed to fetch transactions", error);
      }
    };

    fetchTransactions();
  }, []);

  const formatTimeSince = (dateString: string) => {
    const date = new Date(dateString);
    return formatDistanceToNowStrict(date, { addSuffix: true });
  };

  return (
    <div className="flex-1 flex flex-col justify-between bg-white rounded-lg shadow-md p-4 space-y-4">
        {transactions.map((transaction, index) => (
        <div key={index} className="flex items-center justify-between">
          <div className="w-10 h-10 flex items-center justify-center rounded-full">
            <BanknotesIcon className="h-8 w-8 text-green-500"/>
          </div>
          <div className="flex-1 px-4">
            <div className="text-gray-800 font-medium">{transaction.receiverUserName}</div>
            <div className="text-gray-400 text-sm">{formatTimeSince(transaction.date)}</div>
          </div>
          <div className="text-gray-800 font-semibold">${transaction.amount}</div>
        </div>
        ))} 
    </div>
  );
};

export default InvoicesCard;
