import React from 'react';
import RecentTransactionCard from './RecentTransactionCard';

const RecentTransaction = () => {
  return (
    <div className='w-full'>
      <h1 className='text-[#333B69] pb-2 font-semibold'>Recent Transaction</h1>
      <div className=' max-w-md  bg-white border border-gray-200 rounded-[15px] px-4 py-3'>
        <div className='flow-root'>
          <ul role='list' className=' '>
            <li className='py-1'>
              <RecentTransactionCard
                TransactionName='Deposit from my'
                calender='28 January 2021'
                amount={850}
                imageUrl='/assets/images/deposit.png'
                moneyColor='#FF4B4A'
                sign='-'
              />
            </li>
            <li className='py-1'>
              <RecentTransactionCard
                TransactionName='Depoist Paypal'
                calender='25 January 2021'
                amount={2500}
                imageUrl='/assets/images/paypal.png'
                moneyColor='#41D4A8'
                sign='+'
              />
            </li>
            <li className='py-1'>
              <RecentTransactionCard
                TransactionName='Jemi Wilson'
                calender='21 January 2021'
                amount={5400}
                imageUrl='/assets/images/dollarCoin.png'
                moneyColor='#41D4A8'
                sign='+'
              />
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default RecentTransaction;
