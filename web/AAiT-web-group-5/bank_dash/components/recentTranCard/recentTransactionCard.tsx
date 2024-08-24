import React from 'react';
import PaypalIcon from '../icons/PaypalIcon';
import CreditIcon from '../icons/CreditCardIcon';
import CoinIcon from '../icons/CoinIcon';

function TransactionCard() {
  return (
    <section className='flex w-5/12 px-10 flex-col gap-5 mt-5'>
      <h2 className="flex items-center text-2xl font-semibold text-primary-200">
        Recent Transactions
      </h2>
      <div className="bg-white w-full p-5 overflow-hidden rounded-3xl space-y-5">
        <ul className="text-primary-200 space-y-4">
          <li className="flex items-center gap-3">
            <div className="w-[55px] h-[55px] flex items-center justify-center bg-[#FFF5D9] text-[#FFBB38] rounded-full">
              <CreditIcon />
            </div>
            <div className="flex-1 flex items-center justify-between space-x-3">
              <div>
                <p className="text-black text-lg">Deposit from my Card</p>
                <p className="text-[#718EBF]">25 January 2021</p>
              </div>
              <p className="text-[#FF4B4A] font-semibold">-$500</p>
            </div>
          </li>
          <li className="flex items-center gap-3">
            <div className="w-[55px] h-[55px] flex items-center justify-center bg-[#E7EDFF] text-[#396AFF] rounded-full">
              <PaypalIcon />
            </div>
            <div className="flex-1 flex items-center justify-between space-x-3">
              <div>
                <p className="text-black text-lg">Deposit Paypal</p>
                <p className="text-[#718EBF]">25 January 2021</p>
              </div>
              <p className="text-[#41D4A8] font-semibold">+$500</p>
            </div>
          </li>
          <li className="flex items-center gap-3">
            <div className="w-[55px] h-[55px] flex items-center justify-center bg-[#DCFAF8] text-[#16DBCC] rounded-full">
              <CoinIcon />
            </div>
            <div className="flex-1 flex items-center justify-between space-x-3">
              <div>
                <p className="text-black text-lg">Jemi Wilson</p>
                <p className="text-[#718EBF]">25 January 2021</p>
              </div>
              <p className="text-[#41D4A8] font-semibold">+$500</p>
            </div>
          </li>
        </ul>
      </div>
    </section>
  );
}

export default TransactionCard;
