import React from "react";

export const InputLabel = ({ label }: { label: string }) => {
  return (
    <label htmlFor="" className="text-sm lg:text-base">
      {label}
    </label>
  );
};

const AddCardForm = () => {
  return (
    <div className="bg-white p-5 flex-col gap-7 rounded-xl flex shadow h-full">
      <p className="text-xs text-[#718EBF] leading-5 lg:text-base">
        Credit Card generally means a plastic card issued by Scheduled
        Commercial Banks assigned to a Cardholder, with a credit limit, that can
        be used to purchase goods and services on credit or obtain cash
        advances.
      </p>

      <form action="" className="flex flex-col items-start gap-4">
        <div className="flex flex-col md:flex-row gap-4 w-full">
          <div className="flex flex-col gap-2">
            <InputLabel label="Classic" />
            <input
              type="text"
              placeholder="Classic"
              className="border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base"
            />
          </div>
          <div className="flex flex-col gap-2">
            <InputLabel label="Name On Card" />

            <input
              type="text"
              placeholder="My Cards"
              className="border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base"
            />
          </div>
        </div>
        <div className="flex flex-col md:flex-row gap-4 w-full">
          <div className="flex flex-col gap-2">
            <InputLabel label="Balance" />

            <input
              type="number"
              placeholder="27,000$"
              className="border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base"
            />
          </div>
          <div className="flex flex-col gap-2 w-full md:w-auto">
            <InputLabel label="Expiration Date" />

            <input
              type="date"
              placeholder="27,000$"
              className="border px-3 py-3 rounded-xl lg:w-80 text-xs lg:text-base"
            />
          </div>
        </div>

        <button
          type="submit"
          className="bg-[#1814F3] hover:bg-[#423fef] text-white px-5 py-3 rounded-xl w-full lg:w-auto text-sm"
        >
          Add Card
        </button>
      </form>
    </div>
  );
};

export default AddCardForm;
