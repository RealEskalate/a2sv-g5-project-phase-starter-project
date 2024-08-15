import React from 'react';

type RecentTransactionCardProps = {
  TransactionName: string;
  calender: string;
  amount: number;
  imageUrl: string;
  moneyColor: string;
  sign: string;
};

const RecentTransactionCard: React.FC<RecentTransactionCardProps> = ({
  TransactionName,
  calender,
  amount,
  imageUrl,
  moneyColor,
  sign,
}) => {
  return (
    <div className='flex items-center'>
      <div className='flex-shrink-0'>
        <img
          className=' w-[40px] h-[40px] rounded-full'
          src={imageUrl}
          alt={`${TransactionName} image`}
        />
      </div>
      <div className='flex-1 min-w-0 ms-4'>
        <p className='text-sm text-[#232323] truncate'>{TransactionName}</p>
        <p className='text-sm text-[#718EBF]'>{calender}</p>
      </div>
      <div className='inline-flex items-center text-sm font-medium' style={{ color: moneyColor }}>
        {sign}${amount}
      </div>
    </div>
  );
};

export default RecentTransactionCard;
