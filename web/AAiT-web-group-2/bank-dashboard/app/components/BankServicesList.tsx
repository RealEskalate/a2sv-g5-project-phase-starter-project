import React from "react";
import BankService from "./BankService";

const BankServicesList = () => {
  return (
    <>
      <p className="text-[#343C6A] text-[22px]">Bank Services List</p>
      {[...Array(6)].map((_, index) => (
        <BankService key={index} />
      ))}
    </>
  );
};

export default BankServicesList;
