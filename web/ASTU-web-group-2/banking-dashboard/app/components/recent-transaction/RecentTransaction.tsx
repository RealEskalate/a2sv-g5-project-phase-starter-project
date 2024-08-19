'use client'
import React, { useActionState } from 'react';
import { useSession } from 'next-auth/react';
import { useGetAllTransactionQuery } from '@/lib/service/TransactionService';
interface Props{
  description : string,
  date:string,
  amount : number,
  type : string,
  icon : string,
  receiverUserName : string,


}
const icons = [
  "/assets/recentTransaction/icon1.svg",
 "/assets/recentTransaction/icon2.svg",
 "/assets/recentTransaction/icon3.svg",
 ]
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
  const session = useSession()
  const accessToken = session.data?.user.accessToken || ""
  const { data, isLoading, error ,isSuccess } = useGetAllTransactionQuery(accessToken);
  if (isLoading) {
    return <div>Loading transactions...</div>;
  }

  if (error) {
    return <div>Error fetching transactions</div>;
  }
  let fetcheddata: Props[] = data?.data || recentlistitems;
  if (fetcheddata.length > 3){fetcheddata = fetcheddata.slice(0,3);}
  console.log(fetcheddata)


  return (
      <div className='flex  flex-col flex-initial flex-wrap gap-[10px] bg-white drop-shadow-xl font-medium rounded-[25px] p-[25px]'>
          {fetcheddata.map((value, index) => (
            <div key={index} className='flex items-center gap-3'>
              <img src={icons[index]} alt='Icon'  />
              <div className='flex flex-col  gap-1'>
                <p className='text-[16px] text-[#232323]  leading-[19.36px]'>
                {value.receiverUserName? value.receiverUserName : recentlistitems[index].transactionName}
                </p>
                <p className='text-[15px] leading-[18.36px] text-[#718EBF]'>{value.date}</p>
              </div>
              {value.amount > 0 ? ( <p className='text-lg ml-auto text-[#41D4A8]'>
                +${value.amount}
              </p>) : ( <p className='text-lg ml-auto text-[#FF4B4A]'>
                -${value.amount}
              </p>)}
        
            </div>
          ))}
        </div>
    


  );
}

export default RecentTransaction;
