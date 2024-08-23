import Image from "next/image";
import React from "react";

const Desc = () => {
  return (
    <div className="mt-4 px-32 py-16">
      <h1 className="mb-8 w-[30%] text-4xl text-primaryBlack">
        Let Banking Power your financial operations
      </h1>

      <div className="flex gap-5 items-center justify-between">
        <div className="flex gap-4">
          <Image src="/icons/apple.svg" height={42} width={42} alt="Icon" />
          <div>
            <h1 className="text-primaryBlack">Accounts Payable</h1>
            <p className="text-sm">Manage Pay and Recouncil business</p>
          </div>
        </div>
        <div className="flex gap-4">
          <Image src="/icons/apple.svg" height={42} width={42} alt="Icon" />
          <div>
            <h1 className="text-primaryBlack">Loans</h1>
            <p className="text-sm">Actively Track Loans</p>
          </div>
        </div>
        <div className="flex gap-4">
          <Image src="/icons/apple.svg" height={42} width={42} alt="Icon" />
          <div>
            <h1 className="text-primaryBlack">Investments</h1>
            <p className="text-sm">Focus on your Investments</p>
          </div>
        </div>
        <div className="flex gap-4">
          <Image src="/icons/apple.svg" height={42} width={42} alt="Icon" />
          <div>
            <h1 className="text-primaryBlack">Credit Cards</h1>
            <p className="text-sm">Unlimited Credit cards at fingertip</p>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Desc;
