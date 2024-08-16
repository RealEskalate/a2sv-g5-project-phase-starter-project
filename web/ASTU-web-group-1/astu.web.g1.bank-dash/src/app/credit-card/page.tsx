import AddNewCard from "@/components/AddNewCard/AddNewCard";
import CardList from "@/components/CardList/CardList";
import CardSettings from "@/components/CardSettings/CardSettings";
import MyCard from "@/components/MyCard/MyCard";
import React from "react";

export default function page() {
  return (
    <div className="space-y-5 ">
      <div>
        <h1 className="text-[#333B69] py-2 font-semibold">My Cards</h1>
        <div className="flex overflow-x-scroll space-x-2 scroll whitespace-nowrap scroll-smooth lg:flex lg:space-x-4  ">
          <MyCard />
          <MyCard />
          <MyCard />
        </div>
      </div>

      <div className="space-y-5 md:flex md:gap-6">
        <div className=" md:w-1/3 ">
          <p className="text-[#333B69] pb-2 font-semibold">
            Card Expense Statistics
          </p>
          <div className="bg-white rounded-3xl ">
            shvshs dfhbsfs hhhhhhhhhhhhhhhhhhfds hhhhhhhhhhh hsjdffff
            ffffffffffajcvdjhbc havvvvvvvabervhdrvb b chxvjzbvhvchv vh
            dfahvbhvash v dshvshvshddsh
          </div>
        </div>
        <div className="w-full md:w-2/3 ">
          <CardList />
        </div>
      </div>
      <div className="space-y-5 md:flex md:gap-6">
        <AddNewCard />
        <CardSettings />
      </div>
    </div>
  );
}
