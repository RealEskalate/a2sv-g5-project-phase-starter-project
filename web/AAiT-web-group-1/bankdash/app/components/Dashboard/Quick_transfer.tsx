"use client";
import Image from "next/image";
import React, { useState } from "react";
import send from "../../../public/images/send.svg";
import p2 from "../../../public/images/profile1.svg";
import p3 from "../../../public/images/profile2.svg";
import p4 from "../../../public/images/profile3.svg";
import arrow from "../../../public/images/arrow.svg";

interface Type {
  name: string;
  position: string;
  profilePicture: string;
}

interface Props {
  profiles: Type[];
}

const Quick_transfer = () => {
  const profiles = [
    { name: "Livia Bator", position: "CEO", profilePicture: p2 },
    { name: "Randy press", position: "Director", profilePicture: p3 },
    { name: "Workman", position: "Designer", profilePicture: p4 },
    { name: "Livia Bator", position: "CEO", profilePicture: p2 }
  ];
  const [currIndex, setCurrIndex] = useState(0);

  const threeProfiles = profiles.slice(currIndex, currIndex + 3);

  const handleNext = () => {
    if (profiles.length - currIndex > 3) {
      setCurrIndex(currIndex + 1);
    }
  };

  return (
    <div className="flex flex-col h-full ">
      <div>
        <h1 className="text-[22px] font-semibold text-[#343C6A] ml-2 mb-3">Quick Transfer</h1>
      </div>
      <div className="mt-2 md:mt-3 py-5 bg-white rounded-3xl space-y-10 pl-2 h-full">
        <div className="flex ">
          <div className="flex space-x-7">
            {threeProfiles.map((item, index) => (
              <div key={index} className="flex flex-col items-center space-y-1">
                <Image
                  className="w-10 h-10"
                  src={item.profilePicture}
                  alt="profile picture"
                />
                <h3 className="text-xs font-semibold">{item.name}</h3>
                <h3 className="text-xs">{item.position}</h3>
              </div>
            ))}
          </div>
          <div>
            <Image
              className="w-16 h-16 mt-4 ml-4"
              onClick={handleNext}
              src={arrow}
              alt="arrow"
            ></Image>
          </div>
        </div>

        <div className="flex w-full items-center">
          <div className="w-1/2">
            <h2 className="text-sm text-center font-semibold text-[#718EBF]">
              Write Amount
            </h2>
          </div>
          <div className="flex 1/2 relative">
            <input
              type="text"
              className="w-7/12 p-2 bg-[#EDF1F7] rounded-l-full outline-none"
              
            />
            <button className="bg-[#1814F3] absolute right-4 bottom-0 flex p-2 space-x-4 w-1/2 rounded-full">
              <h3 className="text-white">send</h3>
              <Image className="w-6 h-6" src={send} alt="send" />
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Quick_transfer;
