import Image from 'next/image';  

interface BalanceCardProps {
    iconSrc: string; // Path to the image source
    altText: string;
    title: string;
    amount: string;
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