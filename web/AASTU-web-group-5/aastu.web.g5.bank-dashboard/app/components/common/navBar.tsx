import React from "react";
import { FaSearch } from "react-icons/fa";
import seetings from "/public/assets/icons/group417.png";
import notification from "/public/assets/icons/group418.png";
import person from "/public/assets/icons/MaskGroup.png";
import magnifying from "/public/assets/icons/magnifying-glass.png";
import { FaBars } from "react-icons/fa";

const primary_2 = "rgba(52, 60, 106, 1)";
const primary_3 = "rgba(45, 96, 255, 1)";
const sidecolor = "#B1B1B1";

import Image from "next/image";
const SearchInput = () => {
  return (
    <div className="relative">
      <input
        type="text"
        placeholder="Search..."
        className="pl-12 pr-4 py-2 rounded-full w-full"
        style={{
          backgroundColor: "#F5F7FA",
          borderRadius: "40px",
          border: "none",
        }}
      />
      <div className="absolute left-4 top-1/2 transform -translate-y-1/2 text-gray-500">
        <Image src={magnifying} alt="Search Icon" />
      </div>
    </div>
  );
};

function Desktop() {
  return (
    <div className="flex justify-between h-[101px]   items-center pl-[46px] pr-[40px]">
      <div className="font-semibold text-[22px]" style={{ color: primary_2 }}>
        Overview
      </div>
      <div className="flex gap-[20px]">
        <SearchInput />
        <Image
          src={seetings}
          alt="Transfer Icon"
          className="h-[40px] w-[40px]"
        />
        <Image
          src={notification}
          alt="Transfer Icon"
          className="h-[40px] w-[40px]"
        />
        <Image src={person} alt="Transfer Icon" className="h-[45px] w-[45px]" />
      </div>
    </div>
  );
}
function Mobile() {
  return (
    <div className="pl-[25px] pr-[25px] w-full h-[140px] pt-[25px]">
    <div className="flex justify-between items-center">
      <FaBars />
      <div className="font-semibold text-[22px]" style={{ color: primary_2 }}>
        Overview
      </div>
      <Image src={person} alt="Transfer Icon" className="h-[45px] w-[45px]" />
    </div>
    <div className="mt-4 w-full flex">
      <SearchInput  />
    </div>
  </div>
  );
}

function NavBar() {
  return (
    <>
      <div className="hidden sm:block">
        <Desktop />
      </div>
      <div className="block sm:hidden">
        <Mobile />
      </div>
    </>
  );
}

export default NavBar;
