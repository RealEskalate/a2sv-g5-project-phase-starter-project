
'use client'

import React from "react";
import Toggle from './toogle';

function Preference() {
  return (
    <div>
      <div className="flex flex-wrap flex-col md:flex-row md:gap-10 lg:gap-12 mt-10 md:mt-12 mx-4">
        <div>
          <div>Currency</div>
          <input
            type="text"
            className="border-slate-200 border-[1px] w-full h-10 mt-3 rounded-3xl md:w-[20rem] lg:w-[30rem]"
            style={{ paddingLeft: '1.25rem' }}
            placeholder="USD"
          />
        </div>
        <div>
          <div>Time Zone</div>
          <input
            type="text"
            className="border-slate-200 border-[1px] w-full h-10 mt-3 rounded-2xl md:w-[20rem] lg:w-[30rem]"
            placeholder="(GMT-12:00) International Date Line West"
            style={{ paddingLeft: '1.25rem' }}
          />
        </div>
      </div>
      <div className="mt-6 md:mt-8 text-slate-700 text-sm md:text-base lg:text-[17px]">
        Notification
        <div className="flex flex-col gap-4 mt-5 md:mt-6">
          <div className="flex gap-5 md:gap-6">
            <Toggle />
            <div>I send or receive digital currency</div>
          </div>
          <div className="flex gap-5 md:gap-6">
            <Toggle />
            <div>I receive merchant order</div>
          </div>
          <div className="flex gap-5 md:gap-6">
            <Toggle />
            <div>There are recommendations for my account</div>
          </div>
        </div>
      </div>
      <div className="flex justify-end mt-16 md:mt-18">
        <button className="border-none bg-blue-700 text-white w-full h-12 rounded-full md:w-[12rem] text-[13px] md:text-base">
          Save
        </button>
      </div>
    </div>
  );
}

export default Preference;