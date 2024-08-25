import React from 'react';
import Image from 'next/image'
import { RecentTransactionProps } from '@/types/index.';

const RecentTransaction: React.FC<RecentTransactionProps> = ({ title, date, amount, type, imageSrc }) => (
  <div className="flex justify-between items-center">
    <div className="flex items-center">
      <Image src={imageSrc} alt={title} className="w-8 h-8 mr-4 rounded-full" /> {/* Image element */}
      <div>
        <p className="font-medium">{title}</p>
        <p className="text-gray-500">{date}</p>
      </div>
    </div>
    <p className={`font-medium ${type === 'income' ? 'text-green-500' : 'text-red-500'}`}>
      {amount}
    </p>
  </div>
);

export default RecentTransaction;