import Image from "next/image";
import React from "react";

interface titleProp {
  title: string;
}

const Navbar = ({ title }: titleProp) => {
  return (
    <header>
      <div className="flex justify-between p-5">
        <div className="left-content flex justify-between items-center m-5">
          <div className="logo flex items-center justify-between mr-14">
            <Image
              src={"/credit-card.png"}
              width={50}
              height={50}
              alt="bankDash logo"
              className="mr-3"
            />
            <p className="font-extrabold text-3xl text-blue-950">BankDash.</p>
          </div>

          <div className="Title">
            <p className="font-bold text-3xl text-blue-950">{title}</p>
          </div>
        </div>

        <div className="right-contnet m-5 flex items-center">
          <div className="search-div flex bg-gray-100 items-center rounded-full pl-5 pr-5 pt-3 pb-3 mr-7">
            <Image
              src={"/search.png"}
              width={17}
              height={17}
              alt="magnifying-glass"
              className="mr-5"
            />
            <input
              type="text"
              placeholder="Search for something"
              className="text-md bg-gray-100"
            />
          </div>

          <div className="bg-gray-100 p-2 rounded-full items-center mr-7">
            <Image
              src={"/settings.png"}
              width={30}
              height={30}
              alt="magnifying-glass"
            />
          </div>

          <div className="bg-gray-100 p-2 rounded-full items-center mr-7">
            <Image
              src={"/notification.png"}
              width={30}
              height={30}
              alt="magnifying-glass"
            />
          </div>

          <Image
            src={"/default-image.png"}
            width={60}
            height={60}
            alt="profile-picture"
            className="object-cover rounded-full"
          />
        </div>
      </div>
    </header>
  );
};

export default Navbar;
