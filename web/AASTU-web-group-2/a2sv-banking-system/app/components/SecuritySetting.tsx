"use client";
import React from "react";

const HeadingLabel = ({ label }: { label: string }) => {
  return (
    <h1 className="text-sm font-medium lg:text-lg text-[#333B69]">{label}</h1>
  );
};

const InputLabel = ({ label, htmlFor }: { label: string; htmlFor: string }) => {
  return (
    <label htmlFor={htmlFor} className="text-xs text-[#232323] lg:text-base">
      {label}
    </label>
  );
};

import { useState } from "react";

const ToggleSwitch = () => {
  const [enabled, setEnabled] = useState(false);

  const toggleSwitch = () => {
    setEnabled(!enabled);
  };

  return (
    <div className="flex items-center gap-3">
      <input
        type="checkbox"
        id="two-factor-toggle"
        className="peer hidden"
        checked={enabled}
        onChange={toggleSwitch}
      />
      <label
        htmlFor="two-factor-toggle"
        className={`cursor-pointer rounded-full w-12 h-6 flex items-center relative transition-colors duration-300 ${
          enabled ? "bg-[#16DBCC]" : "bg-gray-200"
        }`}
      >
        <span
          className={`bg-white w-6 h-6 rounded-full transition-transform duration-300 ${
            enabled ? "translate-x-6" : ""
          }`}
        ></span>
      </label>

      <InputLabel
        label="Enable or disable two factor authentication"
        htmlFor="two-factor-toggle"
      />
    </div>
  );
};

const SecuritySetting = () => {
  return (
    <div className="">
      <form action="" className="flex flex-col gap-3">
        <HeadingLabel label="Two-factor Authentication" />
        <div className="flex items-center">
          <ToggleSwitch />
        </div>

        <HeadingLabel label="Change Password" />
        <div className="space-y-2">
          <InputLabel label="Current Password" htmlFor="old_password" />
          <input
            type="password"
            name=""
            id="old_password"
            placeholder="**********"
            className="border rounded-lg px-3 py-2 text-xs flex items-end w-full md:w-64"
          />
        </div>
        <div className="space-y-2">
          <InputLabel label="New Password" htmlFor="new_password" />
          <input
            type="password"
            name=""
            id="new_password"
            placeholder="**********"
            className="border rounded-lg px-3 py-2 text-xs flex items-end w-full md:w-64"
          />
        </div>
        <div className="w-full flex justify-end">
          <button
            type="submit"
            className="bg-[#1814F3] hover:bg-[#423fef] text-white px-5 py-3 rounded-xl md:w-2/12 text-sm w-full  self-end"
          >
            Save
          </button>
        </div>
      </form>
    </div>
  );
};

export default SecuritySetting;
