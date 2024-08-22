import React from "react";
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
  return (
    <div className="flex flex-col gap-6 bg-white drop-shadow-xl font-medium rounded-[25px] p-[25px]">
      {recentlistitems.map((value, index) => (
        <div key={index} className="flex items-center gap-3">
          <img src={value.icons} alt="Icon" />
          <div className="flex flex-col flex-initial gap-1">
            <p className="text-base text-[#B1B1B1] leading-[19.36px]">
              {value.transactionName}
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
