
import { useEffect, useState } from "react";
import { getAllTransactions } from "@/services/transactionfetch";
import { currentuser } from "@/services/userupdate";
import TransactionCard from "./TransactionCard";
import { TbFileSad } from "react-icons/tb";
import TransactionCardShimmer from './TransactionCardShimmer';

// App Component
const App: React.FC = () => {
  const [status, setStatus] = useState<'loading' | 'error' | 'success'>('loading');
  const [transactions, setTransactions] = useState<[]>([]);
  const [currentUser, setCurrentUser] = useState("");

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const transactionData = await getAllTransactions(0, 5); 
        const current = await currentuser();
        setCurrentUser(current.data.name);

        console.log("on the card user:", currentuser); 
        
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

  if (status === 'loading') {
    return (
      <div className="flex flex-col gap-4">
      {[...Array(3)].map((_, index) => (
        <TransactionCardShimmer key={index} />
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

  if (status === 'success' && transactions.length === 0) {
    return (
      <div className="p-3 gap-4 flex-1 h-auto bg-gray-50 dark:bg-dark dark:text-white text-center text-gray-500">
        No transactions to display.
      </div>
    );
  }

  return (
    <div className="p-3  flex-1 h-auto bg-gray-50 dark:bg-dark text-gray-900 dark:text-white grid grid-cols-1  gap-4">
      {transactions.map((transaction, index) => (
        <TransactionCard key={index} transaction={transaction} currentname={currentUser} />
      ))}
    </div>
  );
};

export default App;


// import React, { useEffect, useState } from 'react';
// import { getAllTransactions } from '@/services/transactionfetch';
// import { currentuser } from '@/services/userupdate';
// import TransactionCard from './TransactionCard';
// import { TbFileSad } from 'react-icons/tb';
// import TransactionCardShimmer from './TransactionCardShimmer';
// import { useCurrency } from '../context/CurrencyContext';

// const App: React.FC = () => {
//   const [status, setStatus] = useState<'loading' | 'error' | 'success'>('loading');
//   const [transactions, setTransactions] = useState([]);
//   const [currentUser, setCurrentUser] = useState("");
//   const { currency } = useCurrency();

//   useEffect(() => {
//     const fetchTransactions = async () => {
//       try {
//         const transactionData = await getAllTransactions(0, 5); 
//         const current = await currentuser();
//         setCurrentUser(current.data.name);

//         if (Array.isArray(transactionData.data.content)) {
//           setTransactions(transactionData.data.content);
//           setStatus('success');
//         } else {
//           setStatus('error');
//         }
//       } catch (error) {
//         setStatus('error');
//       }
//     };

//     fetchTransactions();
//   }, []);

//   const formatAmount = (amount: number) => {
//     switch (currency) {
//       case 'EUR':
//         return `€${(amount * 0.85).toFixed(2)}`;
//       case 'GBP':
//         return `£${(amount * 0.75).toFixed(2)}`;
//       default:
//         return `$${amount.toFixed(2)}`;
//     }
//   };

//   return (
//     <div>
//       {/* Your render logic for loading, error, and success */}
//       {status === 'success' && transactions.map((transaction, index) => (
//         <TransactionCard key={index} transaction={transaction} formatAmount={formatAmount} />
//       ))}
//     </div>
//   );
// };

// export default App;

