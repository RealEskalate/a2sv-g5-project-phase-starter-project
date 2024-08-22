import { useState, useEffect } from "react";
import { currentuser } from "@/services/userupdate";
import { BalanceCard } from './smallCard';

interface BalanceCardProps {
  iconSrc: string;
  altText: string;
  title: string;
  amount: string;
  index: number;
}

const App: React.FC = () => {
  const [balances, setBalances] = useState<BalanceCardProps[]>([]);

  useEffect(() => {
    const fetchBalance = async () => {
      try {
        const balanceInfo = await currentuser();
        const fetchedBalance = `$${balanceInfo.data.accountBalance}`;

        const updatedBalances: BalanceCardProps[] = [
          {
            iconSrc: '/Images/1.png',
            index: 1,
            altText: "Money Bag Icon",
            title: "My Balance",
            amount: fetchedBalance, // Using fetched balance
          },
          {
            iconSrc: '/Images/2.png',
            index: 2,
            altText: "Income Icon",
            title: "Income",
            amount: "$5,600",
          },
          {
            iconSrc: '/Images/3.png',
            index: 3,
            altText: "Expense Icon",
            title: "Expense",
            amount: "$3,460",
          },
          {
            iconSrc: '/Images/4.png',
            index: 4,
            altText: "Savings Icon",
            title: "Total Saving",
            amount: "$7,920",
          },
        ];

        setBalances(updatedBalances);
      } catch (error) {
        console.error("Failed to fetch balance", error);
      }
    };

    fetchBalance();
  }, []);

  return (
    <div className="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-4 gap-4 w-full">
      {balances.map((balance) => (
        <BalanceCard key={balance.index} balance={balance} />
      ))}
    </div>
  );
};

export default App;




// import Image from 'next/image';  
// import { useState, useEffect } from "react";
// import { getAllBankServices, getBankServiceById } from "@/services/bankseervice";

// interface BalanceCardProps {
//   iconSrc: string; // Path to the image source
//   altText: string;
//   title: string;
//   balance: string;
//   index: number;
// }

// const BalanceCard: React.FC<{ accountData: BalanceCardProps }> = ({ accountData }) => {
//   const { iconSrc, altText, title, balance } = accountData;

//   return (
//     <div className="bg-white rounded-lg shadow-md p-3 flex items-center space-x-3 w-full">
//       {/* Icon */}
//       <div className="rounded-full p-2">
//         <Image 
//           src={iconSrc} 
//           alt={altText} 
//           width= {35}
//           height={35}
//           // className="w-10 h-10" 
//         />
//       </div>

//       {/* Balance Details */}
//       <div>
//         <p className="text-blue-500 text-xs">{title}</p>
//         <p className="text-black text-lg font-bold">{balance}</p>
//       </div>
//     </div>
//   );
// };


// const App: React.FC = () => {
//   return (
//     <div className="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-4 gap-4 w-full">
//       {accountDatas.map((accountData, index) => (
//         <BalanceCard key={index} accountData={accountData} />
//       ))}
//     </div>
//   );
// };

// export default App;



// const accountDatas: BalanceCardProps[] = [
//   {
//     iconSrc: '/Images/1.png',
//     index: 1,
//     altText: "Money Bag Icon",
//     title: "My Balance",
//     balance: "$12,750",
//   },
//   {
//     iconSrc: '/Images/2.png',
//     index: 2,
//     altText: "Income Icon",
//     title: "Income",
//     balance: "$5,600",
//   },
//   {
//     iconSrc: '/Images/3.png',
//     index: 3,
//     altText: "Expense Icon",
//     title: "Expense",
//     balance: "$3,460",
//   },
//   {
//     iconSrc: '/Images/4.png',
//     index: 4,
//     altText: "Savings Icon",
//     title: "Total Saving",
//     balance: "$7,920",
//   },
// ];


