import { items } from "./lastTransactionItems";
import Image from "next/image";

const LastTransaction = () => {
  return (
    // edit the width of display width and margin while using it and the background color is white
    <div className="w-full ml-[300px] flex flex-col gap-5 p-5 rounded-3xl">
      {items.map((item, index) => (
        <div key={index} className="flex items-center justify-between ">
          <div className="flex w-[100px] justify-start pl-0">
            <Image
              src={item.image}
              width={40}
              height={40}
              alt={`${item.title}-image`}
            />
          </div>

          <div className="w-[180px]">
            <p className="text-[12px] text-[#333B69]">{item.title}</p>
            <span className="text-[12px] text-[#718EBF]">{item.date}</span>
          </div>
          <div className="w-[100px] flex justify-center items-center text-[12px] text-[#718EBF]">
            {item.category}
          </div>
          <div className="w-[100px] flex justify-center items-center text-[12px] text-[#718EBF]">
            {item.pass}
          </div>
          <div className="w-[100px] flex justify-start items-center text-[12px] text-[#718EBF]">
            {item.status}
          </div>
          {item.amount[0] === "-" ? (
            <div className="w-[100px] flex justify-end items-center text-[12px] text-[#FE5C73]">
              {item.amount}
            </div>
          ) : (
            <div className="w-[100px] flex justify-end items-center text-[12px] text-[#16DBAA]">
              {item.amount}
            </div>
          )}
        </div>
      ))}
    </div>
  );
};

export default LastTransaction;
