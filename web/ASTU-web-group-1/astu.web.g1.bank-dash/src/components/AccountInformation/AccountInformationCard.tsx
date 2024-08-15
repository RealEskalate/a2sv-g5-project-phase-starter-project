import React from "react";

interface AccountInformationProps {
  image: string;
  name: string;
  balance: string;
  color: string;
}

const AccountInformationCard = ({
  image,
  name,
  balance,
  color,
}: AccountInformationProps) => {
  const ImageContainerSize = {
    width: "clamp(45px, 8vw + 15px, 70px)",
    height: "clamp(45px, 8vw + 15px, 70px)",
  };

  const imageSize = {
    width: "clamp(15px, 3vw + 5px, 30px)",
    height: "clamp(15px, 3vw + 5px, 30px)",
  };

  return (
    <div className=" flex items-center justify-center w-full py-4 p-2 rounded-[25px] bg-white sm:min-w-[200px]">
      <div
        style={ImageContainerSize}
        className={`flex ${color} rounded-full items-center justify-center`}
      >
        <img
          src={image}
          alt="image"
          style={imageSize}
          className="object-cover"
        />
      </div>
      <div className="flex flex-col ml-4">
        <p
          style={{
            color: "#718EBF",
            fontSize: "clamp(0.75rem, 0.875rem + 0.4vw, .5rem)",
          }}
          className="md:text-[16px]"
        >
          {name}
        </p>
        <h1
          style={{ fontSize: "clamp(16px, 2vw + 4px, 30px)" }}
          className="text-[16px] md:text-[25px] font-medium"
        >
          ${balance}
        </h1>
      </div>
    </div>
  );
};

export default AccountInformationCard;
