import Image from 'next/image';  

interface BalanceCardProps {
    iconSrc: string; // Path to the image source
    altText: string;
    title: string;
    amount: string;
    index: number;
  }
  

export const BalanceCard: React.FC<{ balance: BalanceCardProps }> = ({ balance }) => {
    const { iconSrc, altText, title, amount } = balance;
  
    return (
      <div className="bg-white rounded-lg shadow-md p-3 flex items-center space-x-3 w-full">
        {/* Icon */}
        <div className="rounded-full p-2">
          <Image 
            src={iconSrc} 
            alt={altText} 
            width= {35}
            height={35}
            // className="w-10 h-10" 
          />
        </div>
  
        {/* amount Details */}
        <div>
          <p className="text-blue-500 text-xs">{title}</p>
          <p className="text-black text-lg font-bold">{amount}</p>
        </div>
      </div>
    );
  };
  

  const balances: BalanceCardProps[] = [
  {
    iconSrc: '/Images/1.png',
    index: 1,
    altText: "Money Bag Icon",
    title: "My Balance",
    amount: balance.accountBalance,
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

