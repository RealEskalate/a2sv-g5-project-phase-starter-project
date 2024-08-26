import { useState, useEffect } from "react";
import { currentuser } from "@/services/userupdate";
import { BalanceCard } from "./smallCard";
import { TbFileSad } from "react-icons/tb";

interface BalanceCardProps {
  iconSrc: string;
  altText: string;
  title: string;
  amount: string;
  index: number;
}

const App: React.FC = () => {
  const [status, setStatus] = useState<'loading' | 'error' | 'success'>('loading');
  const [balances, setBalances] = useState<BalanceCardProps[]>([]);

  useEffect(() => {
    const fetchBalance = async () => {
      try {
        const balanceInfo = await currentuser();
        const fetchedBalance = `$${balanceInfo.data.accountBalance}`;

        const updatedBalances: BalanceCardProps[] = [
          {
            iconSrc: "/Images/1.png",
            index: 1,
            altText: "Money Bag Icon",
            title: "My Balance",
            amount: fetchedBalance,
          },
          {
            iconSrc: "/Images/2.png",
            index: 2,
            altText: "Income Icon",
            title: "Income",
            amount: "$5,600",
          },
          {
            iconSrc: "/Images/3.png",
            index: 3,
            altText: "Expense Icon",
            title: "Expense",
            amount: "$3,460",
          },
          {
            iconSrc: "/Images/4.png",
            index: 4,
            altText: "Savings Icon",
            title: "Total Saving",
            amount: "$7,920",
          },
        ];

        setBalances(updatedBalances);
        setStatus('success');
      } catch (error) {
        console.error("Failed to fetch balance", error);
        setStatus('error');
      }
    };

    fetchBalance();
  }, []);

  if (status === 'loading') {
    return (
      <div className="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-4 gap-4 w-full">
        {[...Array(4)].map((_, index) => (
          <div key={index} className="w-full h-20 bg-gray-300 rounded-md animate-pulse"></div>
        ))}
      </div>
    );
  }

  if (status === 'error') {
    return (
      <div className="w-full h-20 flex justify-center  flex-col items-center text-red-500 text-center">
        <TbFileSad
          className={`text-gray-300 dark:text-[#993d4b] w-[400px] h-[70px] pb-2 block mx-auto`}
          strokeWidth={1}
        />
        <p>Failed to fetch</p>
      </div>
    );
  }

  if (status === 'success' && balances.length === 0) {
    return <div className="text-center text-gray-500">No data to display.</div>;
  }

  return (
    <div className="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-4 gap-4 w-full">
      {balances.map((balance) => (
        <BalanceCard key={balance.index} balance={balance} />
      ))}
    </div>
  );
};

export default App;


// import { useState, useEffect, useContext } from "react";
// import { currentuser } from "@/services/userupdate";
// import { BalanceCard } from "./smallCard";
// import { TbFileSad } from "react-icons/tb";
// import { CurrencyContext } from "../context/CurrencyContext";

// interface BalanceCardProps {
//   iconSrc: string;
//   altText: string;
//   title: string;
//   amount: string;
//   index: number;
// }

// const BalanceCards: React.FC = () => {
//   const [status, setStatus] = useState<'loading' | 'error' | 'success'>('loading');
//   const [balances, setBalances] = useState<BalanceCardProps[]>([]);
//   const { exchangeRate } = useContext(CurrencyContext);

//   useEffect(() => {
//     const fetchBalance = async () => {
//       try {
//         const balanceInfo = await currentuser();
//         const fetchedBalance = (parseFloat(balanceInfo.data.accountBalance) * exchangeRate).toFixed(2);       
//         // const fetchedBalance = `$${balanceInfo.data.accountBalance}`;

//         const updatedBalances: BalanceCardProps[] = [
//           {
//             iconSrc: "/Images/1.png",
//             index: 1,
//             altText: "Money Bag Icon",
//             title: "My Balance",
//             amount: `$${fetchedBalance}`,
//             // amount: fetchedBalance,
//           },
//           {
//             iconSrc: "/Images/2.png",
//             index: 2,
//             altText: "Income Icon",
//             title: "Income",
//             amount: `$${(5600 * exchangeRate).toFixed(2)}`,
//             // amount: "$5,600",
//           },
//           {
//             iconSrc: "/Images/3.png",
//             index: 3,
//             altText: "Expense Icon",
//             title: "Expense",
//             amount: `$${(3460 * exchangeRate).toFixed(2)}`,
//             // amount: "$3,460",
//           },
//           {
//             iconSrc: "/Images/4.png",
//             index: 4,
//             altText: "Savings Icon",
//             title: "Total Saving",
//             amount: `$${(7920 * exchangeRate).toFixed(2)}`,
//             // amount: "$7,920",
//           },
//         ];

//         setBalances(updatedBalances);
//         setStatus('success');
//       } catch (error) {
//         console.error("Failed to fetch balance", error);
//         setStatus('error');
//       }
//     };

//     fetchBalance();
//   }, []);

//   if (status === 'loading') {
//     return (
//       <div className="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-4 gap-4 w-full">
//         {[...Array(4)].map((_, index) => (
//           <div key={index} className="w-full h-20 bg-gray-300 rounded-md animate-pulse"></div>
//         ))}
//       </div>
//     );
//   }

//   if (status === 'error') {
//     return (
//       <div className="w-full h-20 flex justify-center  flex-col items-center text-red-500 text-center">
//         <TbFileSad
//           className={`text-gray-300 dark:text-white w-[400px] h-[70px]`}
//         />
//         <p>Failed to fetch</p>
//       </div>
//     );
//   }

//   if (status === 'success' && balances.length === 0) {
//     return <div className="text-center text-gray-500">No data to display.</div>;
//   }

//   return (
//     <div className="grid grid-cols-2 sm:grid-cols-2 lg:grid-cols-4 gap-4 w-full">
//       {balances.map((balance) => (
//         <BalanceCard key={balance.index} balance={balance} />
//       ))}
//     </div>
//   );
// };

// export default BalanceCards;
