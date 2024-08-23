import React from "react";
import Image from "next/image";

const CardSetting = () => {
  return (
    <div className="bg-white rounded-3xl grid gap-6 p-6">
      <div className="flex items-center gap-4">
        <Image
          alt="icon"
          src="/assets/CardSetting/BlockCard.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="text-lg font-medium text-[#232323]">
            Block Card
          </p>
          <p className="text-sm font-normal text-[#718EBF] mt-1">
            Instantly block your card
          </p>
        </div>
      </div>
      <div className="flex items-center gap-4">
        <Image
          alt="icon"
          src="/assets/CardSetting/ChangePinCode.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="text-lg font-medium text-[#232323]">
            Change Pin Code
          </p>
          <p className="text-sm font-normal text-[#718EBF] mt-1">
            Choose another pin code
          </p>
        </div>
      </div>
      <div className="flex items-center gap-4">
        <Image
          alt="icon"
          src="/assets/CardSetting/AddToGooglePay.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="text-lg font-medium text-[#232323]">
            Add to Google Pay
          </p>
          <p className="text-sm font-normal text-[#718EBF] mt-1">
            Withdraw without any card
          </p>
        </div>
      </div>
      <div className="flex items-center gap-4">
        <Image
          alt="icon"
          src="/assets/CardSetting/AddToApplePay.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="text-lg font-medium text-[#232323]">
            Add to Apple Pay
          </p>
          <p className="text-sm font-normal text-[#718EBF] mt-1">
            Withdraw without any card
          </p>
        </div>
      </div>
      <div className="flex items-center gap-4">
        <Image
          alt="icon"
          src="/assets/CardSetting/AddToAppleStore.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="text-lg font-medium text-[#232323]">
            Add to Apple Store
          </p>
          <p className="text-sm font-normal text-[#718EBF] mt-1">
            Withdraw without any card
          </p>
        </div>
      </div>
    </div>
  );
};

export default CardSetting;
