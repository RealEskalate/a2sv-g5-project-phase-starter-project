import Image from "next/image";
import React from "react";

interface titleProp {
  title: string;
}

const Navbar = ({ title }: titleProp) => {
  return (
    <header className="flex flex-col justify-center items-center">
      <div className="flex w-full">
        <div className="logo flex items-center justify-between sm:mr-14 pl-[27px]">
          <Image
            src={"assets/navbar/hamburger.svg"}
            width={25}
            height={25}
            alt="hamburger"
            className="mr-3 sm:hidden block"
          />
          <Image
            src={"/assets/navbar/credit-card.svg"}
            width={36}
            height={36}
            alt="bankDash logo"
            className="mr-3 sm:block hidden"
          />
          <p className="font-black text-[25px] text-[#343C6A] sm:flex hidden">
            BankDash.
          </p>
        </div>

        <div className="w-full m-5 flex sm:justify-between items-center">
          <div className="flex w-full justify-center sm:w-auto">
            <p className="font-semibold text-[20px] sm:text-[25px] text-[#343C6A] sm:ml-7">
              {title}
            </p>
          </div>

          <div className="sm:flex position:absolute right-0 gap-5">
            <div className="search-div hidden sm:flex bg-[#F5F7FA] items-center rounded-full h-[50px] pl-5 pr-5 pt-3 pb-3">
              <Image
                src={"/assets/navbar/magnifying-glass.svg"}
                width={20}
                height={20}
                alt="magnifying-glass"
                className="mr-5 "
              />
              <input
                type="text"
                placeholder="Search for something"
                className="text-[15px] bg-[#F5F7FA]"
              />
            </div>

            <div className="bg-[#F5F7FA] hidden sm:flex justify-center rounded-full items-center">
              <Image
                src={"/assets/navbar/settings.svg"}
                width={50}
                height={50}
                alt="settings"
                className="flex-shrink-0"
              />
            </div>

            <div className="bg-[#F5F7FA] hidden sm:flex justify-center rounded-full items-center">
              <Image
                src={"/assets/navbar/notification.svg"}
                width={50}
                height={50}
                alt="notification"
                className="flex-shrink-0"
              />
            </div>

            <Image
              src={"/assets/navbar/default-image.svg"}
              width={50}
              height={50}
              alt="profile-picture"
              className="object-fill rounded-full "
            />
          </div>
        </div>
      </div>

      <div className="search-div flex w-4/5 sm:hidden bg-[#F5F7FA] items-center rounded-full pl-5 pr-5 pt-3 pb-3">
        <Image
          src={"/assets/navbar/magnifying-glass.svg"}
          width={20}
          height={20}
          alt="magnifying-glass"
          className="mr-5"
        />
        <input
          type="text"
          placeholder="Search for something"
          className="text-md bg-[#F5F7FA] w-full"
        />
      </div>
    </header>
  );
};

export default Navbar;
