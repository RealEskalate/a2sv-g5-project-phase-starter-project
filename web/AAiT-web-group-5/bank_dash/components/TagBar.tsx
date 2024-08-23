import React from "react";
import lockIcon from "../public/svgs/Group.svg";
import Image from "next/image";

const TagBar = ({
  image,
  title,
  subtitle,
}: {
  image: string;
  title: string;
  subtitle: string;
}) => {
  return (
    <div className="flex gap-4 justify-center items-center bg-white rounded-2xl px-10 py-5 hover:shadow-md">
      <div className="flex justify-center items-center  bg-[#E7EDFF] rounded-full size-16">
        <Image src={image} width={28} height={34} alt={title} />
      </div>
      <div className="flex flex-col">
        <p className="text-xl">{title}</p>
        <p className="text-[#718EBF] ">{subtitle}</p>
      </div>
    </div>
  );
};

export default TagBar;
