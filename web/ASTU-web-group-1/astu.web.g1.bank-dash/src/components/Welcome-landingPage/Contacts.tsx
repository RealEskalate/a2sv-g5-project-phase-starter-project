import Image from "next/image";
import React from "react";

const Contacts = () => {
  return (
    <>
      <div
        className=" mt-10 flex flex-wrap space-y-5 justify-center md:space-y-0 md:space-x-1 w-full p-5 text-white md:px-24 shadow-[5px_5px_10px_10px_rgb(250,250,250)]"
        id="contacts"
      >
        <div className="w-64 pt-4 space-y-4">
          <p className="font-Inter text-lg font-bold text-gray-400">Bank</p>
          <ul className="list-disc space-y-2 text-deepNavy">
            <li className="">About us</li>
            <li className="">Contact us</li>
            <li className="">Careers</li>
          </ul>
        </div>
        <div className="w-64 pt-4 space-y-4">
          <p className="font-Inter text-lg font-bold text-gray-400">Contacts</p>
          <ul className="space-y-2 list-disc text-deepNavy">
            <li className=" flex space-x-3">
              <Image
                src="/assets/images/email.png"
                alt=""
                width={1000}
                height={1000}
                className="w-6 h-6"
              />
              <p>support@brainwave.io</p>
            </li>
            <li className="flex space-x-3">
              <Image
                src="/assets/images/phone.png"
                alt=""
                width={1000}
                height={1000}
                className="w-6 h-6"
              />
              <p>+133-394-3439-1435</p>
            </li>
            <li className="flex space-x-3">
              <Image
                src="/assets/images/sms.png"
                alt=""
                width={1000}
                height={1000}
                className="w-6 h-6"
              />
              <p>3439</p>
            </li>
          </ul>
        </div>
        <div className="w-64 pt-4 space-y-4">
          <p className="font-Inter text-lg font-bold text-gray-400">Services</p>
          <ul className="space-y-2 list-disc text-deepNavy">
            <li className="">Accounts</li>
            <li className="">Credit-card</li>
            <li className="">Investments</li>
            <li className="">Loans</li>
            <li className="">Setting</li>
            <li className="">Transactions</li>
          </ul>
        </div>
        <div className="w-64 pt-4 space-y-4">
          <p className="font-Inter text-lg font-bold text-gray-400">
            Social Media
          </p>
          <ul className="flex items-center space-x-10">
            <li className="hover:text-green-200">
              <Image
                src="/assets/images/youtube.png"
                alt=""
                width={1000}
                height={1000}
                className="w-9 h-9"
              />
            </li>
            <li className="hover:text-green-200 ">
              <Image
                src="/assets/images/instagram.png"
                alt=""
                width={1000}
                height={1000}
                className="w-9 h-9"
              />
            </li>
            <li className="hover:text-green-200 relative">
              <Image
                src="/assets/images/facebook.png"
                alt=""
                width={1000}
                height={1000}
                className="w-9 h-9"
              />
            </li>
          </ul>
        </div>
      </div>
    </>
  );
};

export default Contacts;
