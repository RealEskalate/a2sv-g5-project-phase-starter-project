import AddNewCard from "@/components/AddNewCard/AddNewCard";
import CardSettings from "@/components/CardSettings/CardSettings";
import MyCard from "@/components/MyCard/MyCard";
import React from "react";

export default function page() {
  return (
    <div>
      <div className="flex overflow-x-scroll space-x-2 scroll whitespace-nowrap scroll-smooth lg:flex lg:space-x-4  ">
        <MyCard />
        <MyCard />
        <MyCard />
      </div>
      <div className="md:flex md:gap-6">
        <AddNewCard />
        <CardSettings />
      </div>
    </div>
  );
}
