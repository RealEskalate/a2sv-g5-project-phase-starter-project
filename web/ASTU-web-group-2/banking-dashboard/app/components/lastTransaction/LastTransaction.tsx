"use client";
import { useGetAllTransactionQuery } from "@/lib/service/TransactionService";
import Image from "next/image";

export type Item = {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
};

const LastTransaction = () => {
  const access =
    "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJzYW1pdGVzdCIsImlhdCI6MTcyMzgxMDA0NiwiZXhwIjoxNzIzODk2NDQ2fQ.6W4MhXDf2fYFVvHonw5Bs597XszBLvNnB71B6pqrbhkMU3IFkF7NLwHLLJPotY51";

  const { data, isError, isLoading } = useGetAllTransactionQuery(access);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error fetching transactions</div>;
  }

  let items: Item[] = data?.data || [];
  if (items.length > 3) {
    items = items.slice(0, 3);
  }

  return (
    // The width depends on the width of container
    <div className="w-[325px] sm:w-[487px] md:w-[730px]  flex flex-col bg-white gap-5 p-5 rounded-3xl">
      {items.map((item, index) => (
        <div key={index} className="flex items-center justify-between ">
          <div className="flex w-[45px] sm:w-[55px] justify-center mr-5">
            {item.type === "shopping" ? (
              <Image
                src={"/assets/lastTransaction/spot-sub.svg"}
                width={55}
                height={55}
                alt={`${item.receiverUserName}-image`}
              />
            ) : item.type === "service" ? (
              <Image
                src={"/assets/lastTransaction/settings.svg"}
                width={55}
                height={55}
                alt={`${item.receiverUserName}-image`}
              />
            ) : (
              <Image
                src={"/assets/lastTransaction/user.svg"}
                width={55}
                height={55}
                alt={`${item.receiverUserName}-image`}
              />
            )}
          </div>

          <div className="w-[137px] sm:w-[117px] md:w-[156px] ">
            <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
              {item.receiverUserName}
            </p>
            <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
              {item.date}
            </span>
          </div>
          <div className="hidden  sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.type}
          </div>
          <div className="hidden sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.transactionId.slice(0, 4)}****
          </div>
          <div className="hidden  sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.description}
          </div>
          {item.amount < 0 ? (
            <div className="w-[100px] flex justify-end items-center text-[16px] text-[#FE5C73]">
              ${item.amount}
            </div>
          ) : (
            <div className="w-[100px] flex justify-end items-center text-[16px] text-[#16DBAA]">
              ${item.amount}
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

export default LastTransaction;
