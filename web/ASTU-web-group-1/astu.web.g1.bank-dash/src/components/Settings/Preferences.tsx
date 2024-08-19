"use client";
import React from "react";
import InputGroup from "../Form/InputGroup";
import ToggleInput from "../Form/ToggleInput";

const Preferences = () => {
  return (
    <div>
      <form action="">
        <div className="flex flex-col md:flex-row md:space-x-5">
          <div className=" w-full lg:w-6/12 space-y-3 my-3">
            <label htmlFor="select" className="gray-dark text-16px">
              Country
            </label>
            <select
              id="select"
              className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
            >
              <option selected>USD</option>
              <option value="US">Birr</option>
              <option value="CA">Birr</option>
              <option value="FR">Birr</option>
            </select>
          </div>

          <InputGroup
            id="timeZone"
            label="Time Zone"
            inputType="text"
            registerName="timeZone"
            register={undefined}
            placeholder="(GMT-1200) International Date Line West"
          />
        </div>

        <label className="gray-dark text-16px">Notification</label>

        <ToggleInput
          label="I send or receive digital currency"
          inputType="checkbox"
          id="email"
          registerName="email"
          register={undefined}
          placeholder="Email"
          currentState={true}
        />
        <ToggleInput
          label="I recieve merchant order"
          inputType="checkbox"
          id="email"
          registerName="email"
          register={undefined}
          placeholder="Email"
          currentState={false}
        />
        <ToggleInput
          label="There are recommendation for my account"
          inputType="checkbox"
          id="email"
          registerName="email"
          register={undefined}
          placeholder="Email"
          currentState={true}
        />

        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-[#1814f3] text-white px-10 py-2 rounded-lg w-full md:w-auto mt-4"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default Preferences;
