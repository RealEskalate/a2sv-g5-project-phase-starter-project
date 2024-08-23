import Image from "next/image";
import React from "react";

const Service = ({ image, title }: { image: string; title: string }) => {
  return (
    <div className="flex items-center justify-between rounded-xl p-3 px-4 bg-white hover:shadow-md">
      <div className="bg-white">
        <div className="flex gap-6">
          <div className="bg-[#E7EDFF] rounded-2xl flex items-center p-3">
            <Image src={image} width={26} height={35} alt="" />
          </div>
          <div>
            <p className="font-semibold">{title}</p>
            <p className="text-[15px] text-[#718EBF]">
              It is a long established{" "}
            </p>
          </div>
        </div>
      </div>
      <div>
        <p className="font-semibold">Lorem Ipsum</p>
        <p className="text-[15px] text-[#718EBF]">Many publishing</p>
      </div>
      <div>
        <p className="font-semibold">Lorem Ipsum</p>
        <p className="text-[15px] text-[#718EBF]">Many publishing</p>
      </div>
      <div>
        <p className="font-semibold">Lorem Ipsum</p>
        <p className="text-[15px] text-[#718EBF]">Many publishing</p>
      </div>
      <button className="border transition hover:text-white hover:bg-blue-500 ease-in-out duration-200 border-[#718EBF] text-[#718EBF] px-4 py-1 text-[15px] flex items-center rounded-3xl">
        View Details
      </button>
    </div>
  );
};

export default Service;
