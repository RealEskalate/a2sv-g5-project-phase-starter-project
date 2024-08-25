"use client";
import React, { useState } from "react";

const Security: React.FC = () => {
  const [enableAuthentication, setenableAuthentication] = useState(false);

  return (
    <div className="flex flex-col lg:w-full h-auto space-y-5">
      <div className="flex flex-col space-y-3">
        <label className="font-semibold text-[#333B69]">Two-factor Authentication</label>

        <div className="flex items-center space-x-3">
          <label htmlFor="enableAuthentication" className="flex items-center cursor-pointer">
            <div className="relative">
              <input
                type="checkbox"
                id="enableAuthentication"
                className="sr-only"
                checked={enableAuthentication}
                onChange={() => setenableAuthentication(!enableAuthentication)}
              />
              <div
                className={`block ${
                  enableAuthentication ? "bg-[#16DBCC]"  : "bg-gray-300"
                } w-14 h-8 rounded-full`}
              ></div>
              <div
                className={`dot absolute left-1 top-1 bg-white w-6 h-6 rounded-full transition ${
                  enableAuthentication ? "transform translate-x-6" : ""
                }`}
              ></div>
            </div>
            <div className="ml-3 text-gray-700">Enable or disable two-factor authentication</div>
          </label>
        </div>
      </div>

      <div className="flex flex-col space-y-4">
        <label className="font-semibold text-[#333B69]">Change Password</label>
        <div className="flex flex-col space-y-3">
          <label htmlFor="currentPassword">Current Password</label>
          <input
            placeholder="**********"
            type="password"
            id="currentPassword"
            name="currentPassword"
            className="w-1/2 p-3 border rounded-2xl"
          />
        </div>

        <div className="flex flex-col space-y-3 mt-3">
          <label htmlFor="newPassword">New</label>
          <input
            placeholder="**********"
            type="password"
            id="newPassword"
            name="newPassword"
            className="w-1/2 p-3 border rounded-2xl"
          />
        </div>
      </div>

      <div className="flex justify-end mt-5">
        <button
          type="submit"
          className="w-full lg:w-52 h-14 rounded-xl bg-[#1814f3] lg:ml-auto mt-8 text-white"
        >
          Save
        </button>
      </div>
    </div>
  );
};

export default Security;
