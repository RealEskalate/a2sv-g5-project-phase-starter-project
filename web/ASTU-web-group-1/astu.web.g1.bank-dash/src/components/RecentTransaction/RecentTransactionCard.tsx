import React from 'react';

type RecentTransactionCardProps = {
  TransactionName: string;
  calender: string;
  amount: number;
  imageUrl: string;
  sign: string;
};

const RecentTransactionCard: React.FC<RecentTransactionCardProps> = ({
  TransactionName,
  calender,
  amount,
  imageUrl,
  sign,
}) => {
  const dateObj = new Date(calender);
  const sign2 = sign === 'deposit';

  const options = { day: '2-digit', month: 'short', year: 'numeric' };
  // const formattedDate = dateObj.toLocaleDateString('en-GB', options);
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
        <p className='font-semibold text-[#232323] truncate'>{TransactionName}</p>
        <p className='text-sm text-[#718EBF] font-medium'>
          {dateObj.toLocaleString('en-GB', { day: '2-digit', month: 'short', year: 'numeric' })}
        </p>
      </div>
      <div
        className={`inline-flex items-center font-medium ${
          sign2 ? 'text-[#41D4A8]' : 'text-red-600'
        }`}
      >
        {sign2 ? '+' : '-'}${amount}
      </div>
    </div>
  );
};

export default RecentTransactionCard;
