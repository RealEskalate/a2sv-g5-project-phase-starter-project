import React from 'react';
const recentlistitems = [
  {   
      transactionName: "Deposit from my",
      date: "28 January 2021", 
      amount: "-$880",
      isDeposited: false,
      icons: "/assets/recentTransaction/icon1.svg"
  },
  {
      transactionName: "Deposit Paypal",
      date: "28 January 2021",
      amount: "+$2500",
      isDeposited: true,
      icons: "/assets/recentTransaction/icon2.svg"
  },
  {
      transactionName: "Jemi Wilson",
      date: "28 January 2021",
      amount: "+$5,400",
      isDeposited: true,
      icons: "/assets/recentTransaction/icon3.svg"
  }
];

const RecentTransaction = () => {
  return (
    // <div className='max-w-[350px] mx-auto h-auto'>
      // <h1 className='mb-[20px] text-[22px] text-[#343C6A] font-semibold'>Recent Transactions</h1>
      <div className='flex  flex-col flex-initial flex-wrap gap-[10px] bg-white drop-shadow-xl font-medium rounded-[25px] p-[25px]'>
          {recentlistitems.map((value, index) => (
            <div key={index} className='flex items-center gap-3'>
              <img src={value.icons} alt='Icon'  />
              <div className='flex flex-col  gap-1'>
                <p className='text-[16px] text-[#232323]  leading-[19.36px]'>
                  {value.transactionName}
                </p>
                <p className='text-[15px] leading-[18.36px] text-[#718EBF]'>{value.date}</p>
              </div>
              <p className='text-lg ml-auto' style={{ color: value.isDeposited ? '#41D4A8' : '#FF4B4A' }}>
                {value.amount}
              </p>
            </div>
          ))}
        </div>
      // </div>


  );
}

export default RecentTransaction;
