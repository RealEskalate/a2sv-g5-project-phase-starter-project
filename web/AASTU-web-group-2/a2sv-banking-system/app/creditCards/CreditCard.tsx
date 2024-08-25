import Link from "next/link";
import React, { ReactNode } from "react";

interface Props {
  icon: ReactNode;
  data: Array<[string, string]>;
  handleDetail: Function;
  card: Card;
  initialValue: boolean;
}

import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import { Card } from "@/types/cardController.Interface";

const DetailCard = ({
  handleClick,
  initialValue,
  card,
}: {
  handleClick: Function;
  initialValue: boolean;
  card: Card;
}) => {
  return (
    <Dialog
      isOpen={initialValue}
      onClose={() => {
        handleClick(false);
      }}
    >
      <DialogContent>
        <div className="text-center w-full">
          <h3 className="text-xl font-semibold mb-4 text-gray-800 dark:text-[#9faaeb]">
            Card Details
          </h3>
          <div className="flex gap-4">
            <div
              className="bg-gray-100 p-4 rounded-lg mb-4 w-full dark:bg-[#050914] dark:border dark:border-[#333B69] 
"
            >
              <p className="text-sm text-gray-500 dark:text-[#9faaeb]">
                Card Holder
              </p>
              <p className="text-lg font-medium text-gray-900 dark:text-[#dde1f8]">
                {card.cardHolder}
              </p>
            </div>
            <div
              className="bg-gray-100 p-4 rounded-lg mb-4 w-full dark:bg-[#050914] dark:border dark:border-[#333B69] 
"
            >
              <p className="text-sm text-gray-500 dark:text-[#9faaeb]">
                Balance
              </p>
              <p className="text-lg font-medium text-gray-900 dark:text-[#dde1f8]">
                ${card.balance.toFixed(2)}
              </p>
            </div>
          </div>
          <div className="flex gap-4">
            <div
              className="bg-gray-100 p-4 rounded-lg mb-4 w-full dark:bg-[#050914] dark:border dark:border-[#333B69] 
"
            >
              <p className="text-sm text-gray-500 dark:text-[#9faaeb]">
                Card Number
              </p>
              <p className="text-lg font-medium text-gray-900 dark:text-[#dde1f8]">
                {card.expiryDate}
              </p>
            </div>
            <div
              className="bg-gray-100 p-4 rounded-lg mb-4 w-full dark:bg-[#050914] dark:border dark:border-[#333B69] 
"
            >
              <p className="text-sm text-gray-500 dark:text-[#9faaeb]">
                Card Type
              </p>
              <p className="text-lg font-medium text-gray-900 dark:text-[#dde1f8]">
                **** **** **** {card.semiCardNumber}
              </p>
            </div>
          </div>
          <button
            onClick={() => handleClick(false)}
            className="mt-6 bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700 transition"
          >
            Close
          </button>
        </div>
      </DialogContent>
    </Dialog>
  );
};

const CreditCard = ({
  icon,
  data,
  handleDetail,
  card,
  initialValue,
}: Props) => {
  return (
    <>
      <DetailCard
        handleClick={handleDetail}
        initialValue={initialValue}
        card={card}
      />
      <div className="flex justify-around items-center bg-white p-3 lg:p-4 rounded-2xl shadow-sm dark:bg-[#050914] dark:border dark:border-[#333B69]">
        {icon}
        <div className="flex gap-9">
          {data.map((data, index) => {
            return (
              <div className={`${index > 1 && "hidden"}  lg:block`} key={index}>
                <p className="text-[#232323] font-medium text-sm lg:text-base dark:text-[#d2d6ef]">
                  {data[0]}
                </p>
                <p className="text-[#8297c0] text-xs lg:text-sm ">{data[1]}</p>
              </div>
            );
          })}
        </div>
        <button
          className="text-[#1814F3] font-medium text-xs lg:text-base dark:text-[#1492f3]"
          onClick={() => {
            handleDetail(true);
          }}
        >
          View Detail
        </button>
      </div>
    </>
  );
};

export default CreditCard;
