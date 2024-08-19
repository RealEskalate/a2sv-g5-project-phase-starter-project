import React from "react";
import Image from "next/image";
const CardSetting = () => {
  return (
    <div className="max-sm:rounded-[15px] bg-white rounded-[25px] flex flex-col gap-[40px] max-sm:gap-[15px]">
      <div className="flex pl-[30px] pt-[30px] gap-[20px] items-center max-sm:pt-[12px] max-sm:gap-[12px] ">
        <Image
          alt="icon"
          src="/assets/CardSetting/BlockCard.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="text-[16px] leading-[19.36px] text-[#232323] font-medium   max-sm:text-[14px]">
            Block Card
          </p>
          <p className="font-normal text-[15px] leading-[18.15px] text-[#718EBF] pt-[7px] max-sm:text-[12px]">
            Instantly block your card
          </p>
        </div>
      </div>
      <div className="flex pl-[30px]  gap-[20px] items-center max-sm:gap-[12px]">
        <Image
          alt="icon"
          src="/assets/CardSetting/ChangePinCode.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="font-medium text-[16px] leading-[19.36px] text-[#232323] max-sm:text-[14px]">
            Change Pin Code
          </p>
          <p className="font-normal text-[15px] leading-[18.15px] text-[#718EBF] pt-[7px] max-sm:text-[12px]">
            Choose another pin code
          </p>
        </div>
      </div>
      <div className="flex pl-[30px]  gap-[20px] items-center max-sm:gap-[12px]">
        <Image
          alt="icon"
          src="/assets/CardSetting/AddToGooglePay.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="font-medium text-[16px] leading-[19.36px] text-[#232323] max-sm:text-[14px]">
            Add to Google Pay
          </p>
          <p className="font-normal text-[15px] leading-[18.15px] text-[#718EBF] pt-[7px] max-sm:text-[12px]">
            Withdraw without any card
          </p>
        </div>
      </div>
      <div className="flex pl-[30px]   gap-[20px] items-center max-sm:gap-[12px]">
        <Image
          alt="icon"
          src="/assets/CardSetting/AddToApplePay.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="font-medium text-[16px] leading-[19.36px] text-[#232323] max-sm:text-[14px]">
            Add to Apple Pay
          </p>
          <p className="font-normal text-[15px] leading-[18.15px] text-[#718EBF] pt-[7px] max-sm:text-[12px]">
            Withdraw without any card
          </p>
        </div>
      </div>
      <div className="flex pl-[30px]  gap-[20px] items-center max-sm:gap-[12px]">
        <Image
          alt="icon"
          src="/assets/CardSetting/AddToAppleStore.svg"
          width={45}
          height={45}
        />
        <div>
          <p className="font-medium text-[16px] leading-[19.36px] text-[#232323] max-sm:text-[14px]">
            Add to Apple Store
          </p>
          <p className="font-normal text-[15px] leading-[18.15px] text-[#718EBF] pt-[7px] max-sm:text-[12px]">
            Withdraw without any card
          </p>
        </div>
      </div>
    </div>
  );
};

export default CardSetting;
