import Image from "next/image";
import React from "react";

const BankService = ({ iconUrl }: { iconUrl: string }) => {
  return (
    <div className=" bg-[#F5F7FA] sm:p-0">
      {" "}
      {/* Added padding for spacing */}
      <div className="bg-white rounded-2xl h-[90px] flex sm:justify-around items-center mt-5 w-full max-w-full">
        {" "}
        {/* Responsive width */}
        <div className="flex gap-4 pl-2 sm:pl-0 items-center">
          <Image
            className="sm:w-[60px] sm:h-[60px]"
            width={45}
            height={45}
            src={iconUrl}
            alt=""
          />
          <div>
            <p className="font-medium text-[#232323] text-sm sm:text-base sm:mb-auto mb-0.5">
              Bussiness loans
            </p>
            <p className="font-light text-[#718EBF] text-xs sm:text-[15px]">
              it is a long established
            </p>
          </div>
        </div>
        <div className="hidden sm:block">
          {" "}
          {/* Hidden on mobile, visible on small screens and up */}
          <p className="font-medium"> Lorem Ipsum</p>
          <p className="font-light text-[#718EBF] sm:text-[15px]">
            Many publishing
          </p>
        </div>
        <div className="hidden sm:block">
          {" "}
          {/* Hidden on mobile, visible on small screens and up */}
          <p className="font-medium"> Lorem Ipsum</p>
          <p className="font-light text-[#718EBF] sm:text-[15px]">
            Many publishing
          </p>
        </div>
        <div className="hidden sm:block">
          {" "}
          {/* Hidden on mobile, visible on small screens and up */}
          <p className="font-medium"> Lorem Ipsum</p>
          <p className="font-light text-[#718EBF] sm:text-[15px]">
            Many publishing
          </p>
        </div>
        <div className="rounded-3xl text-xs sm:text-sm font-medium pl-10 text-[#1814F3] sm:border sm:px-3 py-1 hover:text-[#1814F3] hover:border-blue-800">
          view details
        </div>
      </div>
    </div>
  );
};

export default BankService;
