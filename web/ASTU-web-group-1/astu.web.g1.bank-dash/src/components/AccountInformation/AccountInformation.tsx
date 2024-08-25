import React from "react";
import AccountInformationCard from "./AccountInformationCard";

const AccountInformation = () => {
  return (
    <div className="w-full grid grid-cols-2 grid-rows-2 min-[1030px]:grid-cols-4 md:grid-rows-1 grid-flow-row gap-4 mb-5">
      <AccountInformationCard
        image="/assets/images/balance.png"
        name="My Balance"
        balance="12,7000"
        color="bg-[#FFF5D9]"
      />
      <AccountInformationCard
        image="/assets/images/income.png"
        name="Income"
        balance="5,600"
        color="bg-[#E7EDFF]"
      />

      <AccountInformationCard
        image="/assets/images/expense.png"
        name="Expense"
        balance="3,460"
        color="bg-[#FFE0EB]"
      />

      <AccountInformationCard
        image="/assets/images/saving.png"
        name="Total Saving"
        balance="7,920"
        color="bg-[#DCFAF8]"
      />
    </div>
  );
};

export default AccountInformation;
