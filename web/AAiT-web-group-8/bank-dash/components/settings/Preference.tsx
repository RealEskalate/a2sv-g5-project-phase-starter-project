"use client";
import React, { useState } from "react";

const Preference: React.FC = () => {
  // State for each notification toggle
  const [sendReceive, setSendReceive] = useState(false);
  const [merchantOrder, setMerchantOrder] = useState(false);
  const [accountRecommendation, setAccountRecommendation] = useState(false);

  // Save function (you can update this to actually handle saving preferences)
  const handleSave = () => {
    const preferences = {
      sendReceive,
      merchantOrder,
      accountRecommendation,
    };
    console.log("Preferences saved:", preferences);
    
  };

  return (
    <div className="flex flex-col lg:w-full h-auto space-y-5">
      <div className="flex flex-col lg:flex-row items-start lg:space-x-5">
        <div className="flex flex-col w-full">
          <label htmlFor="currency">Currency</label>
          <input
            placeholder="USD"
            type="text"
            id="currency"
            name="currency"
            className="w-full p-3 mt-3 border rounded-2xl"
          />
        </div>

        <div className="flex flex-col w-full">
          <label htmlFor="timezone">Time Zone</label>
          <input
            placeholder="(GMT-12:00) International Date Line West"
            type="text"
            id="timezone"
            name="timezone"
            className="w-full p-3 mt-3 border rounded-2xl"
          />
        </div>
      </div>

      <div className="flex flex-col space-y-3">
        <label className="font-semibold">Notification</label>

        <div className="flex items-center space-x-3">
          <label htmlFor="sendReceive" className="flex items-center cursor-pointer">
            <div className="relative">
              <input
                type="checkbox"
                id="sendReceive"
                className="sr-only"
                checked={sendReceive}
                onChange={() => setSendReceive(!sendReceive)}
              />
              <div
                className={`block ${
                  sendReceive ? "bg-[#16DBCC]"  : "bg-gray-300"
                } w-14 h-8 rounded-full`}
              ></div>
              <div
                className={`dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition ${
                  sendReceive ? "transform translate-x-6" : ""
                }`}
              ></div>
            </div>
            <div className="ml-3 text-gray-700">I send or receive digital currency</div>
          </label>
        </div>

        <div className="flex items-center space-x-3">
          <label htmlFor="merchantOrder" className="flex items-center cursor-pointer">
            <div className="relative">
              <input
                type="checkbox"
                id="merchantOrder"
                className="sr-only"
                checked={merchantOrder}
                onChange={() => setMerchantOrder(!merchantOrder)}
              />
              <div
                className={`block ${
                  merchantOrder ? "bg-[#16DBCC]"  : "bg-gray-300"
                } w-14 h-8 rounded-full`}
              ></div>
              <div
                className={`dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition ${
                  merchantOrder ? "transform translate-x-6" : ""
                }`}
              ></div>
            </div>
            <div className="ml-3 text-gray-700">I receive merchant order</div>
          </label>
        </div>
        <div className="flex items-center space-x-3">
          <label
            htmlFor="accountRecommendation"
            className="flex items-center cursor-pointer"
          >
            <div className="relative">
              <input
                type="checkbox"
                id="accountRecommendation"
                className="sr-only"
                checked={accountRecommendation}
                onChange={() =>
                  setAccountRecommendation(!accountRecommendation)
                }
              />
              <div
                className={`block ${
                  accountRecommendation ? "bg-[#16DBCC]"  : "bg-gray-300"
                } w-14 h-8 rounded-full`}
              ></div>
              <div
                className={`dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition ${
                  accountRecommendation ? "transform translate-x-6" : ""
                }`}
              ></div>
            </div>
            <div className="ml-3 text-gray-700">
              There are recommendations for my account
            </div>
          </label>
        </div>
      </div>

      <div className="flex justify-end mt-5">
      <button
            type="submit"
            className="w-full lg:w-52 h-14 rounded-xl bg-[#1814f3] lg:ml-auto mt-12 text-white"
          >
            Save
          </button>
      </div>
    </div>
  );
};

export default Preference;
