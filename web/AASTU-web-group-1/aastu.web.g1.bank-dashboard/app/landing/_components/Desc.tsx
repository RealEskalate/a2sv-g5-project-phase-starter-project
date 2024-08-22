import Image from "next/image";
import React from "react";

const Desc = () => {
  return (
    <div>
      <h1>Let Banking Power your financial operations</h1>

      <div>
        <Image src="/icons/apple.svg" height={24} width={24} alt="Icon" />
        <h1>Accounts Payable</h1>
        <p>Manage Pay and Recouncil business bills</p>
      </div>
    </div>
  );
};

export default Desc;
