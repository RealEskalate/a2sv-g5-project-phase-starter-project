import React from "react";
import AddCardForm from "../Forms/AddCardForm";

const AddCard = () => {
  return (
    <div className="bg-white dark:bg-[#232328] rounded-3xl p-6 sm:p-10 shadow-sm">
      <p className="text-[#718EBF] text-sm sm:text-base pb-4 sm:pb-5 dark:text-gray-400">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>
      <AddCardForm />
    </div>
  );
};

export default AddCard;
