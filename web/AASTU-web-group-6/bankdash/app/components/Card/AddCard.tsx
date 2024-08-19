import React from "react";
import AddCardForm from "../Forms/AddCardForm";

const AddCard = () => {
  return (
    <div className="border bg-white rounded-3xl p-10 ">
      <p className="text-[#718EBF] text-base pb-5">
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
