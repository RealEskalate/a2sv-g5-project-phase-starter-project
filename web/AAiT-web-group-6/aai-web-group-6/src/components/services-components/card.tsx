import React from "react";

interface CardProps {
  text1: string;
  text2: string;
  imageSrc: string;
  imageBackground: string;
}

export default function ServiceCard({
  text1,
  text2,
  imageSrc,
  imageBackground,
}: CardProps) {
  return (
    <div className="card py-6 px-12 space-x-5 flex bg-white rounded-3xl">
      <div
        className={`bg-[rgba(${imageBackground})] rounded-full w-16 h-16 flex justify-center items-center`}
      >
        <img src={imageSrc} alt="Card Image" />
      </div>
      <div className="flex flex-col space-y-2 justify-center ">
        <p className="text-[rgba(35,35,35,1)] font-semibold text-sm">{text1}</p>
        <p className="text-[rgba(113,142,191,1)] text-xs font-normal">
          {text2}
        </p>
      </div>
    </div>
  );
}
