import React from "react";

interface LastTransData {
  transactionId: string;
  type: string;
  senderUserName: string;
  description: string;
  date: string;
  amount: number;
  receiverUserName: string;
}

const LastTrans = ({
  description,
  date,
  type,
  amount,
  receiverUserName,
  transactionId,
  senderUserName
}: LastTransData) => {
  let amountStr = amount?.toLocaleString();
  let amount_str = ""
  if (amount < 0) {
     amount_str = `-$${amountStr.slice(1)}`; 
  } else {
     amount_str=  `+$${amountStr}`;
  }
  
  const textColor = amountStr?.startsWith("-")
    ? "text-[#FE5C73]"
    : "text-[#16DBAA]";
  let color = "";
  let icon = "";

  if (type === "shopping") {
    icon = "/assets/renew.svg";
    color = "#DCFAF8";
  } else if (type === "transfer") {
    color = "#FFE0EB";
    icon = "/assets/userr.svg";
  } else if (type === "service") {
    color = "#E7EDFF";
    icon = "/assets/settings.svg";
  } else {
    color = "#FFE0EB";
    icon = "/assets/userr.svg";
  }

  return (
    <div className="flex gap-10 items-center mb-5 ">
      <div
        className="border border-solid rounded-2xl w-[55px] h-[55px] flex justify-center items-center"
        style={{ borderColor: color, backgroundColor: color }}
      >
        <img src={icon} />
      </div>
      <div className="flex flex-col w-[25%]">
        <p className="font-inter font-medium text-base text-[#232323] ">
          {description}
        </p>
        <p className="font-inter font-normal text-[15px] text-[#718EBF]">
          {date}
        </p>
      </div>

      <p className="hidden lg:block font-inter font-normal text-base text-[#718EBF] w-[15%]">
        {type}
      </p>
      <p className="hidden lg:block font-inter font-normal text-base text-[#718EBF] w-[10%]">
        123***
      </p>
      <p className="hidden lg:block font-inter font-normal text-base text-[#718EBF] w-[10%]">
        completed
      </p>

      <p className={`font-inter font-medium text-base ${textColor}`}>
        {amount_str}
      </p>
    </div>
  );
};

export default LastTrans;
