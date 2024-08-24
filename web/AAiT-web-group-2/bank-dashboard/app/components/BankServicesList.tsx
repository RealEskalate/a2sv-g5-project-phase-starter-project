import React from "react";
import BankService from "./BankService";
import styles from "../../src/styles/hide-scroll.module.css";
import BankServiceCard from "./BankServiceCard";

const BankServicesList = () => {
  const icons = [
    "./icons/insurance.svg",
    "./icons/shopping.svg",
    "./icons/safety.svg",
  ];
  const icons2 = [
    "./icons/loans.svg",
    "./icons/checking.svg",
    "./icons/saving.svg",
    "./icons/debit.svg",
    "./icons/life_insurance.svg",
    "./icons/loans.svg",
  ];
  return (
    <>
      <div
        className={`flex  gap-10 mb-3 overflow-x-auto ${styles.hide_scroll}`}
      >
        {[...Array(3)].map((_, index) => (
          <BankServiceCard iconUrl={icons[index]} key={index} />
        ))}
      </div>

      <p className="text-[#343C6A] text-[22px] font-semibold my-5">
        Bank Services List
      </p>
      {[...Array(6)].map((_, index) => (
        <BankService iconUrl={icons2[index]} key={index} />
      ))}
    </>
  );
};

export default BankServicesList;
