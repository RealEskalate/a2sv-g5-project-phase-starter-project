'use client'
import React, { useState } from "react";

function Toggle() {
  const [isChecked, setIsChecked] = useState(false);

  const handleToggle = () => {
    setIsChecked(!isChecked);
  };

  return (
    <label className="inline-flex items-center cursor-pointer">
      <input
        type="checkbox"
        value=""
        className="sr-only peer"
        checked={isChecked}
        onChange={handleToggle}
      />
      <div className="relative w-11 h-6 bg-gray-200 rounded-full peer peer-focus:ring-4 peer-focus:ring-blue-300 dark:peer-focus:bg-teal-400 dark:bg-gray-700 peer-checked:after:translate-x-full rtl:peer-checked:after:-translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:start-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-teal-400"></div>
    </label>
  );
}
const activeColor = 'text-blue-700';
const disabledColor = 'text-slate-400';

function Settings() {
  const [enabled, setEnabled] = useState('Edit Profile');

  return (
    <div className="bg-white rounded-2xl text-neutral-800 text-sm md:text-base mx-auto w-full max-w-max mt-6 md:mt-8 px-[25px] sm:px-[40px]">
      <div className="relative mt-8 md:mt-10">
        <div className="flex gap-[42px] md:flex-row md:gap-16 lg:gap-[73px]">
          <div
            className={`cursor-pointer ${enabled === 'Edit Profile' ? activeColor : disabledColor}`}
            onClick={() => setEnabled('Edit Profile')}
          >
            Edit Profile
          </div>
          <div
            className={`cursor-pointer ${enabled === 'Preference' ? activeColor : disabledColor}`}
            onClick={() => setEnabled('Preference')}
          >
            Preference
          </div>
          <div
            className={`cursor-pointer ${enabled === 'Security' ? activeColor : disabledColor}`}
            onClick={() => setEnabled('Security')}
          >
            Security
          </div>
        </div>
        <div className="flex flex-col md:flex-row md:gap-10 lg:gap-[30px] mt-10 md:mt-12 mx-4 md:mx-8 lg:mx-16">
          <div>
            <div>Currency</div>
            <input
              type="text"
              className="border-slate-200 border-[1px] w-full h-[40px] mt-3 rounded-3xl md:w-[334px] lg:w-[510px]"
              style={{ paddingLeft: '20px' }}
              placeholder="USD"
            />
          </div>
          <div>
            <div>Time Zone</div>
            <input
              type="text"
              className="border-slate-200 border-[1px] w-full h-[40px] mt-3 rounded-2xl md:w-[334px] lg:w-[510px]"
              placeholder="(GMT-12:00) International Date Line West"
              style={{ paddingLeft: '20px' }}
            />
          </div>
        </div>
        <div className="mt-6 md:mt-8 text-slate-700 text-sm md:text-base lg:text-[17px]">
          Notification
          <div className="flex flex-col gap-4 mt-5 md:mt-6">
            <div className="flex gap-5 md:gap-[20px]">
              <Toggle />
              <div>I send or receive digital currency</div>
            </div>
            <div className="flex gap-5 md:gap-[20px]">
              <Toggle />
              <div>I receive merchant order</div>
            </div>
            <div className="flex gap-5 md:gap-[20px]">
              <Toggle />
              <div>There are recommendations for my account</div>
            </div>
            
            
           
          </div>
        </div>
        
        <div className="flex justify-end mt-16 md:mt-18 mb-5">
          <button className="border-none bg-blue-700 text-white w-full h-[50px] rounded-full md:w-[190px] md:ml-[90px] text-[13px] md:text-base">
            Save
          </button>
        </div>
      </div>
    </div>
  );
}

export default Settings;