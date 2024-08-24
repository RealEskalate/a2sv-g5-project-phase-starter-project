"use client";

import React from "react";
import DepositeFromCard from "../../public/DepositFromCard";
import DepositWithPaypal from "../../public/Paypal";
import JemiWilson from "../../public/JemiWilson";

const RecentTransaction = () => {
  return (
    <div className="flex flex-col justify-start py-5 w-full min-w-[231px] min-h-[170px] max-w-[338px] bg-white lg:h-full rounded-my-card-radius drop-shadow-sm shadow-md">
      <div className="flex justify-start items-center">
        <div className="flex justify-center items-center bg-custom-light-orange mx-6 min-w-[50px] min-h-[50px] max-w-[60px] max-h-[60px] rounded-full">
          <DepositeFromCard />
        </div>
        <div className="flex flex-col justify-start items-start">
          <span className="font-medium text-[clamp(13px,2vw,16px)] text-custom-light-dark">
            Deposit from my Card
          </span>
          <span className="text-custom-light-purple text-[clamp(10px,2vw,14px)] ">
            28 January 2021
          </span>
        </div>
        <span className="text-custom-red ml-auto mr-6 text-[clamp(11px,2vw,16px)]">
          -$850
        </span>
      </div>
      <div className="flex justify-start my-3 items-center">
        <div className="flex justify-center items-center bg-custom-light-blue mx-6 min-w-[50px] min-h-[50px] max-w-[60px] max-h-[60px] rounded-full">
          <DepositWithPaypal />
        </div>
        <div className="flex flex-col justify-start items-start">
          <span className="font-medium text-md text-custom-light-dark text-[clamp(13px,2vw,16px)]">
            Deposit Paypal
          </span>
          <span className="text-custom-light-purple text-[clamp(10px,2vw,14px)]">
            25 January 2021
          </span>
        </div>
        <span className="text-my-color-8 ml-auto mr-6 text-[clamp(11px,2vw,16px)]">
          +$2,500
        </span>
      </div>
      <div className="flex justify-start items-center">
        <div className="flex justify-center items-center bg-custom-light-teal mx-6 min-w-[50px] min-h-[50px] max-w-[60px] max-h-[60px] rounded-full">
          <JemiWilson />
        </div>
        <div className="flex flex-col justify-start items-start">
          <span className="font-medium text-md text-custom-light-dark text-[clamp(13px,2vw,16px)]">
            Jemi Wilson
          </span>
          <span className="text-custom-light-purple text-[clamp(10px,2vw,14px)]">
            21 January 2021
          </span>
        </div>
        <span className="text-my-color-8 ml-auto mr-6 text-[clamp(11px,2vw,16px)]">
          +$5,400
        </span>
      </div>
    </div>
  );
};

export default RecentTransaction;
