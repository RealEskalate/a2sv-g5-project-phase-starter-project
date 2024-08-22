'use client';

import React from "react";
import DepositeFromCard from '../../public/DepositFromCard';
import DepositWithPaypal from '../../public/Paypal';
import JemiWilson from "../../public/JemiWilson";

const RecentTransaction = () => {
  return (
    <div className="flex flex-col justify-center my-6 mx-auto w-my-card-width-2 h-my-card-height-3 rounded-my-card-radius drop-shadow-sm shadow-md">
      <div className="flex justify-start items-center">
        <div className="flex justify-center items-center bg-custom-light-orange mx-6 w-14 h-14 rounded-full">
          <DepositeFromCard />
        </div>
        <div className="flex flex-col justify-start items-start">
          <span className="font-medium text-md text-custom-light-dark">
            {" "}
            Deposit from my Card
          </span>
          <span className="text-custom-light-purple">28 January 2021</span>
        </div>
        <span className="text-custom-red ml-auto mr-6">-$850</span>
      </div>
      <div className="flex justify-start my-3 items-center">
        <div className="flex justify-center items-center bg-custom-light-blue mx-6 w-14 h-14 rounded-full">
          <DepositWithPaypal />
        </div>
        <div className="flex flex-col justify-start items-start">
          <span className="font-medium text-md text-custom-light-dark">
            Deposit Paypal
          </span>
          <span className="text-custom-light-purple">25 January 2021</span>
        </div>
        <span className="text-my-color-8 ml-auto mr-6">+$2,500</span>
      </div>
      <div className="flex justify-start items-center">
        <div className="flex justify-center items-center bg-custom-light-teal mx-6 w-14 h-14 rounded-full">
          <JemiWilson />
        </div>
        <div className="flex flex-col justify-start items-start">
          <span className="font-medium text-md text-custom-light-dark">
            Jemi Wilson
          </span>
          <span className="text-custom-light-purple">21 January 2021</span>
        </div>
        <span className="text-my-color-8 ml-auto mr-6">+$5,400</span>
      </div>
    </div>
  );
};

export default RecentTransaction;
