import React from "react";

const AddCardForm = () => {
  return (
    <div>
      <form className="w-[50%]" action="">
        <div className="flex gap-10">
          <div>
            <div className="flex flex-col">
              <label className="text-[#232323] text-base font-semibold pb-2">
                Card Type
              </label>
              <input
                className="border border-[#DFEAF2] p-2 rounded-2xl mb-5"
                type="text"
                placeholder="Classic"
              />
            </div>
            <div className="flex flex-col">
              <label className="text-[#232323] text-base font-semibold pb-2">
                Balance
              </label>
              <input
                className="border border-[#DFEAF2] p-2 rounded-2xl"
                type="text"
                placeholder="Classic"
              />
            </div>
          </div>
          <div>
            <div className="flex flex-col">
              <label className="text-[#232323] text-base font-semibold pb-2">
                Name on Card
              </label>
              <input
                className="border border-[#DFEAF2] p-2 rounded-2xl mb-5"
                type="text"
                placeholder="Classic"
              />
            </div>
            <div className="flex flex-col">
              <label className="text-[#232323] text-base font-semibold pb-2">
                Expiration Date
              </label>
              <input
                className="border border-[#DFEAF2] p-2 rounded-2xl"
                type="date"
                placeholder="Classic"
              />
            </div>
          </div>
        </div>
        <button
          type="submit"
          className="mt-4 px-8 py-3 bg-[#1814F3] text-white text-lg font-semibold rounded-[9px]"
        >
          Add Card
        </button>
      </form>
    </div>
  );
};

export default AddCardForm;
