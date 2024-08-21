// import Image from 'next/image';  

// interface BalanceCardProps {
//     iconSrc: string; // Path to the image source
//     altText: string;
//     title: string;
//     balance: string;
//     index: number;
//   }
  

// export const BalanceCard: React.FC<{ accountData: BalanceCardProps }> = ({ accountData }) => {
//     const { iconSrc, altText, title, balance } = accountData;
  
//     return (
//       <div className="bg-white rounded-lg shadow-md p-3 flex items-center space-x-3 w-full">
//         {/* Icon */}
//         <div className="rounded-full p-2">
//           <Image 
//             src={iconSrc} 
//             alt={altText} 
//             width= {35}
//             height={35}
//             // className="w-10 h-10" 
//           />
//         </div>
  
//         {/* Balance Details */}
//         <div>
//           <p className="text-blue-500 text-xs">{title}</p>
//           <p className="text-black text-lg font-bold">{balance}</p>
//         </div>
//       </div>
//     );
//   };
  

//   const accountDatas: BalanceCardProps[] = [
//   {
//     iconSrc: '/Images/1.png',
//     index: 1,
//     altText: "Money Bag Icon",
//     title: "My Balance",
//     balance: ,
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

