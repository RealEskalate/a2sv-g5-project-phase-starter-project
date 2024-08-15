import React from "react";
import Avatar from "../Avatar/Avatar";

const ProfileCard = () => {
  return (
    <div className="flex flex-col gap-4 justify-center items-center w-[60px] lg:w-[85px] text-12px lg:text-16px">
      <Avatar />
      <div className="flex flex-col gap-1 justify-center w-full">
        <h1 className="text-gray-dark truncate whitespace-nowrap overflow-hidden text-overflow-ellipsis">Randy sdfdss Press</h1>
        <p className="text-center text-blue-steel">CEO</p>
      </div>
    </div>
  );
};

export default ProfileCard;
