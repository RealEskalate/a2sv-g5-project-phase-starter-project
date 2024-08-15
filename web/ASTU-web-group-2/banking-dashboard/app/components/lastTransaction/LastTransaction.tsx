import { items } from "./lastTransactionItems";
import Image from "next/image";

const LastTransaction = () => {
  return (
    // The width depends on the width of container
    <div className="w-[325px] sm:w-[487px] md:w-[730px]  flex flex-col bg-white gap-5 p-5 rounded-3xl">
      {items.map((item, index) => (
        <div key={index} className="flex items-center justify-between ">
          <div className="flex w-[45px] sm:w-[55px] justify-center mr-5">
            <Image
              src={item.image}
              width={55}
              height={55}
              alt={`${item.title}-image`}
            />
          </div>

          <div className="w-[137px] sm:w-[117px] md:w-[156px] ">
            <p className="text-[14px] sm:text-[16px] text-[#333B69] font-medium">
              {item.title}
            </p>
            <span className="text-[12px] sm:text-[15px] text-[#718EBF]">
              {item.date}
            </span>
          </div>
          <div className="hidden  sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.category}
          </div>
          <div className="hidden sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.pass}
          </div>
          <div className="hidden  sm:w-[100px] sm:flex justify-start items-center text-[16px] text-[#718EBF]">
            {item.status}
          </div>
          {item.amount[0] === "-" ? (
            <div className="w-[100px] flex justify-end items-center text-[16px] text-[#FE5C73]">
              {item.amount}
            </div>
          ) : (
            <div className="w-[100px] flex justify-end items-center text-[16px] text-[#16DBAA]">
              {item.amount}
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

export default LastTransaction;
