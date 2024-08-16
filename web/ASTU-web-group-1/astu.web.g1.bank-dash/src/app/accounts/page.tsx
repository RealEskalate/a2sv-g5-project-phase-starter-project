import AccountInformation from "@/components/AccountInformation/AccountInformation";
import InvoiceSent from "@/components/InvoiceSent/InvoiceSent";
import MyCard from "@/components/MyCard/MyCard";
import LastTransaction from "@/components/Transaction/LastTransaction";
import React from "react";

export default function page() {
  return (
    <div>
      <AccountInformation />
      <div className="min-[890px]:flex min-[890px]:space-x-4 lg:space-x-10 mb-5">
        <LastTransaction />
        <div>
          <div className="flex justify-between">
            <h1 className="text-[#333B69] py-2 font-semibold">My Card</h1>
            <p className="text-[#333B69] py-2 text-sm">See All</p>
          </div>
          <div>
            <MyCard />
          </div>
        </div>
      </div>
      <div className="min-[890px]:flex min-[890px]:space-x-4 lg:space-x-10">
        <LastTransaction />
        <InvoiceSent />
      </div>
    </div>
  );
}
