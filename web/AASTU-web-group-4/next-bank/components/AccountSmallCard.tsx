// import Image from 'next/image';
import { TbMoneybag } from "react-icons/tb";

import Image from 'next/image';

const BalanceCard = () => {
  return (
    <div className="bg-white rounded-lg shadow-md p-3 flex items-center space-x-3 w-48">
      {/* Icon */}
      <div className="bg-yellow-100 rounded-full p-2">
        <Image 
          src="/path-to-your-icon/money-bag-icon.svg" // Replace with your icon path
          alt="Balance Icon"
          width={20}
          height={20}
        />
      </div>

      {/* Balance Details */}
      <div>
        <p className="text-blue-500 text-xs">My Balance</p>
        <p className="text-black text-lg font-bold">$12,750</p>
      </div>
    </div>
  );
}

export default BalanceCard;
