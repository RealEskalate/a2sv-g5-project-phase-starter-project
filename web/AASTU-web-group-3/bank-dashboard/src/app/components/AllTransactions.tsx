"use client"
import React from "react";
import { FaRegArrowAltCircleUp, FaRegArrowAltCircleDown } from "react-icons/fa";
import { useGetAllTransactionsQuery } from "@/lib/redux/api/transactionsApi"; // Adjust the import path based on your folder structure

const formatDate = (dateString: string) => {
  const date = new Date(dateString);

  const formattedDate = date.toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
  });

  const formattedTime = date.toLocaleTimeString('en-US', {
    hour: '2-digit',
    minute: '2-digit',
    hour12: true,
  });

  return `${formattedDate}, ${formattedTime}`;
};

const AllTransactions = () => {
  const { data, error, isLoading } = useGetAllTransactionsQuery({ size: 10, page: 1 }); // Adjust size and page based on your needs

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>Error loading transactions</div>;

  return (
    <div className="">
      <h1 className="m-2 text-xl font-semibold">Recent Transactions</h1>
      <section className="border-0 rounded-xl bg-white shadow-md mx-2 p-2">
        <div className="hidden lg:grid lg:grid-cols-7 xl:grid-cols-9 font-medium text-sky-300 text-xs h-7 items-center border-b mt-2">
          <div className="lg:col-span-2">Description</div>
          <div className="hidden xl:block xl:col-span-2">Transaction Id</div>
          <div>Type</div>
          <div>Sender</div>
          <div>Date</div>
          <div>Amount</div>
          <div className="justify-self-center">Receipt</div>
        </div>

        {data?.data.map((transaction, index) => (
          <div
            key={index}
            className="grid grid-cols-7 xl:grid-cols-9 border-b min-h-12 items-center text-xs lg:font-medium xl:text-[18px]"
          >
            <div className="flex items-center gap-2 col-span-5 lg:col-span-2 lg:font-medium ">
              {transaction.amount < 0 ? (
                <FaRegArrowAltCircleUp
                  color="#718EBF"
                  className="text-4xl md:text-xl xl:text-3xl"
                />
              ) : (
                <FaRegArrowAltCircleDown
                  color="#718EBF"
                  className="text-4xl md:text-xl xl:text-3xl"
                />
              )}
              <span>{transaction.description}</span>
            </div>
            <div className="hidden xl:block xl:col-span-2">{transaction.transactionId}</div>
            <div className="hidden lg:block">{transaction.type}</div>
            <div className="hidden lg:block">{transaction.senderUserName}</div>
            <div className="hidden lg:block">{formatDate(transaction.date)}</div>
            <div
              className={`col-span-2 lg:col-span-1 justify-self-end lg:justify-self-auto ${
                transaction.amount < 0 ? "text-red-500" : "text-green-500"
              }`}
            >
              {transaction.amount < 0 ? "-" : "+"}${Math.abs(transaction.amount)}
            </div>
            {/* <div className="hidden lg:block">{transaction.receiverUserName}</div> */}
            <div className="hidden lg:block border p-1 rounded-lg w-auto justify-self-center hover:border-blue-400 hover:cursor-pointer">Download</div>
          </div>
        ))}
      </section>
    </div>
  );
};

export default AllTransactions;
