import Image from "next/image";
import React from "react";

const MyCard = () => {
  return (
    <div className="w-[350px] h-[235px] bg-grad-end text-white  border rounded-3xl flex flex-col justify-between">
      <div className="flex flex-col  px-8  pt-6 h-full">
        <div className="flex justify-between ">
          <div className="flex flex-col">
            <span className="text-12px">Balance</span>
            <span className="text-20px  font-semibold">$5,756</span>
          </div>
          <div className="">
            <Image
              src="/chip-card-white.svg"
              alt="chip_card"
              width={35}
              height={35}
            />
          </div>
        </div>
        <div className="flex h-full pb-5">
          <div className="flex flex-1 flex-col justify-center">
            <span className="text-12px">CARD HOLDER</span>
            <span className="text-18px font-semibold">Eddy Cusuma</span>
          </div>
          <div className="flex flex-1 flex-col justify-around items-center">
            <div className="flex flex-col">
              <span className="text-12px">VALID THRU</span>
              <span className="text-18px  font-semibold">12/22</span>
            </div>
          </div>
        </div>
      </div>
      <div className="rounded-b-3xl flex justify-between px-8 py-5 bg-gradient-to-b from-grad-start to-grad-end">
        <div className="flex">
          <span className="text-22px font-text-navy">3778 **** **** 1234</span>
        </div>
        <div className="flex pr-[15px]">
          <div className="flex bg-[#9199AF] h-8 w-8 rounded-full opacity-50"></div>
          <div className="flex bg-[#9199AF] h-8 w-8 rounded-full opacity-50 -mx-[15px]"></div>
        </div>
      </div>
    </div>
  );
};

export default MyCard;
