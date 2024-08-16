import AccountInformation from "@/components/AccountInformation/AccountInformation";
import DebiteAndCredit from "@/components/Charts/DebiteAndCredit";
import InvoiceSent from "@/components/InvoiceSent/InvoiceSent";
import MyCard from "@/components/MyCard/MyCard";
import LastTransaction from "@/components/Transaction/LastTransaction";
import React from "react";

export default function page() {
  return (
    <>
      <AccountInformation />
      <div className=" min-[890px]:flex min-[890px]:space-x-4 lg:space-x-10 ">
        <LastTransaction />
        <div className="mb-5">
          <div className="flex justify-between">
            <h1 className="text-[#333B69] pb-2 font-semibold">My Card</h1>
            <p className="text-[#333B69] pb-2 font-semibold">See All</p>
          </div>
          <div>
            <MyCard />
          </div>
        </div>
      </div>
      <div className="min-[890px]:flex min-[890px]:space-x-4 lg:space-x-10">
        <DebiteAndCredit />
        <InvoiceSent />
      </div>
    </>
  );
}
