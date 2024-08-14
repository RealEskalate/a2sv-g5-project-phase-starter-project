"use client";
import React from "react";
import InputGroup from "../Form/InputGroup";
import Select from "react-select";

const options = [
  { value: "chocolate", label: "Chocolate" },
  { value: "strawberry", label: "Strawberry" },
  { value: "vanilla", label: "Vanilla" },
];

const Preferences = () => {
  return (
    <div>
      <form action="">
        <div className="flex flex-col md:flex-row md:space-x-5">
          <div className=" w-full lg:w-6/12 space-y-3 my-3">
            <label htmlFor="select" className="gray-dark text-16px">
              {"Select"} <br />
            </label>
            <Select
              id="select"
              options={options}
              className="w-full border-2 border-[#DFEAF2] p-5 py-3 rounded-xl placeholder:text-blue-steel focus:border-blue-steel outline-none"
            />
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

        <label className="gray-dark text-16px">
            Notification
        </label>

        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-blue-bright text-white px-10 py-2 rounded-lg"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default Preferences;
