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
    width: "clamp(45px, 6vw + 15px, 55px)",
    height: "clamp(45px, 6vw + 15px, 55px)",
  };

  const imageSize = {
    width: "clamp(25px, 2.5vw + 2px, 30px)",
    height: "clamp(25px, 2.5vw + 2px, 30px)",
  };

  return (
    <div className=" flex items-center justify-center w-full py-4 p-2 rounded-[25px] bg-white sm:min-w-[200px]">
      <div
        style={ImageContainerSize}
        className={`flex ${color} rounded-full items-center justify-center flex-shrink-0`}
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
            fontSize: "clamp(0.95rem, 0.9rem + 0.5vw, .8rem)",
          }}
          className="md:text-[16px]"
        >
          {name}
        </p>
        <h1
          // style={{ fontSize: "clamp(16px, 2vw + 4px, 30px)" }}
          className="text-xl md:text-2xl font-medium"
        >
          ${balance}
        </h1>
      </div>
    </div>
  );
};

export default AccountInformationCard;
