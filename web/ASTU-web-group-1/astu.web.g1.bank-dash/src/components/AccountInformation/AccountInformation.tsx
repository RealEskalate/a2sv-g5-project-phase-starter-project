import React from "react";

interface AccountInformationProps {
  image: string;
  name: string;
  balance: string;
  color: string;
}

const AccountInformation = ({
  image,
  name,
  balance,
  color,
}: AccountInformationProps) => {
  return (
    <div className="flex shadow-2xl items-center justify-center max-w-[255px] h-[85px] md:h-[120px] bg-white rounded-[25px] px-[36px] py-[25px]">
      <div
        className={`w-[45px]  md:w-[70px] ${color} relative right-4 h-[45px] md:h-[70px] rounded-full items-center justify-center`}
      >
        <img
          src={image}
          alt="image"
          className="mx-auto my-[15px] md:my-[20px] w-[15px] md:w-[30px] h-[15px] md:h-[30px] object-cover"
        />
      </div>
      <div className="flex flex-col gap-1">
        <p className="text-[#718EBF] text-[12px] md:text-[16px]">{name}</p>
        <h1 className="text-[16px] md:text-[25px] font-medium">${balance}</h1>
      </div>
    </div>
  );
};

export default AccountInformation;
