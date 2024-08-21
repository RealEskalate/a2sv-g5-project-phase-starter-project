import React from "react";

const AddNewCard = () => {
  return (
    <div className="bg-white rounded-3xl grid grid-cols-1 gap-6 p-6">
      <p className="text-[#718EBF] font-normal text-base leading-6">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>
      <form className="grid gap-6">
        <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Card Type
            </label>
            <input
              type="text"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] text-[#718EBF] pl-4"
              placeholder="Classic"
            />
          </div>

          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Name On Card
            </label>
            <input
              type="text"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] pl-4"
              placeholder="My Cards"
            />
          </div>
        </div>

        <div className="grid grid-cols-1 sm:grid-cols-2 gap-6">
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Balance
            </label>
            <input
              type="number"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] pl-4"
              placeholder="27,000$"
            />
          </div>
          <div className="flex flex-col gap-2">
            <label className="text-base font-normal text-[#232323]">
              Expiration Date
            </label>
            <input
              type="date"
              className="w-full h-12 rounded-lg border border-[#E2E8F0] text-[#718EBF] pl-4"
            />
          </div>
        </div>
        <div className="flex justify-start">
          <button
            type="submit"
            className="w-auto h-12 rounded-lg bg-[#1814F3] text-white px-6 py-2"
          >
            Add Card
          </button>
        </div>
      </form>
    </div>
  );
};

export default AddNewCard;
