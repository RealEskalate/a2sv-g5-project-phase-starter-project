import React from "react";

const AddNewCard = () => {
  return (
    <div className="ml-290px mt-849px bg-white rounded-[25px]">
      <p className="pl-[30px] pt-[27px] text-[#718EBF] font-normal text-[16px] leading-[26px]">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>
      <div className="pl-[30px] pt-[29px]">
        <form className="contact-form ">
          <div className="flex gap-[30px] pb-[22px] max-sm:flex-col">
            <div className="flex flex-col gap-[11px] ">
              <label className="text-[16px] leading-[19.26px] font-normal text-[#232323]">
                Card Type
              </label>
              <input
                type="text"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] text-[#718EBF] pl-[20px]"
                placeholder="Classic"
              />
            </div>

            <div className="flex flex-col gap-[11px] ">
              <label className="text-[16px] leading-[19.26px] font-normal text-[#232323]">
                Name On Card
              </label>
              <input
                type="text"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] pl-[20px]"
                placeholder="My Cards"
              />
            </div>
          </div>
          <div className="flex gap-[30px] pb-[30px] max-sm:flex-col">
            <div className="flex flex-col gap-[11px] ">
              <label className="text-[16px] leading-[19.26px] font-normal text-[#232323]">
                Balance
              </label>
              <input
                type="number"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] pl-[20px]"
                placeholder="27,000$"
              />
            </div>
            <div className="flex flex-col gap-[11px]">
              <label
                className="text-[16px] leading-[19.26px] font-normal text-[#232323]"
                htmlFor=""
              >
                Expiration Date
              </label>
              <input
                type="date"
                className="w-[320px] h-[50px] rounded-[15px] border-[1px] text-[#718EBF] pl-[20px]"
              />
            </div>
          </div>
          <button
            type="submit"
            className="w-[160px] h-[50px] rounded-[9px] bg-[#1814F3] text-[#ffffff]"
          >
            Add Card
          </button>
        </form>
      </div>
    </div>
  );
};

export default AddNewCard;
