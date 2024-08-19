'use client'
import React from "react";
import { useGetAllTransactionQuery } from "../recent-transaction/transcation";
import { useSession } from "next-auth/react";
interface Props{
  description : string,
  date:string,
  amount : number,
  type : string,
  icon : string,
  receiverUserName: string;


}
const icons = [
  "/assets/invoicesSent/icon1.svg",
 "/assets/invoicesSent/icon2.svg",
 "/assets/invoicesSent/icon3.svg",
  "/assets/invoicesSent/icon4.svg",
 
 ]
const recentlistitems = [
  {
    transactionName: "Apple Store",
    date: "5h ago",
    amount: "$450",
    icons: "/assets/invoicesSent/icon1.svg",
  },
  {
    transactionName: "Michael",
    date: "5 days ago",
    amount: "$160",
    icons: "/assets/invoicesSent/icon2.svg",
  },
  {
    transactionName: "Play Station",
    date: "5 days ago",
    amount: "$700",
    icons: "/assets/invoicesSent/icon3.svg",
  },
  {
    transactionName: "William",
    date: "10 days ago",
    amount: "$700",
    icons: "/assets/invoicesSent/icon4.svg",
  },

];

const InvoicesSent = () => {
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
  if (fetcheddata.length > 4){fetcheddata = fetcheddata.slice(0,4);}
  console.log(fetcheddata)

  return (
    <div className="flex flex-col gap-6 bg-white drop-shadow-xl font-medium rounded-[25px] p-[25px]">
      {fetcheddata.map((value, index) => (
        <div key={index} className="flex items-center gap-3">
          <img src={icons[index]} alt="Icon" />
          <div className="flex flex-col flex-initial gap-1">
            <p className="text-base text-[#B1B1B1] leading-[19.36px]">
              {value.receiverUserName? value.receiverUserName : recentlistitems[index].transactionName}
            </p>
            <p className="text-[15px] leading-[18.36px] text-[#718EBF]">
              {value.date}
            </p>
          </div>
          <p className="text-lg ml-auto text-[#718EBF]">{value.amount}</p>
        </div>
      ))}
    </div>
  );
};


export default InvoicesSent;
