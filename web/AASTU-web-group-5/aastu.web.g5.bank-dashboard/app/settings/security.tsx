'use client';
import React from 'react';
import Toggle from './toogle';

export default function Security() {
  return (
    <div className='text-[16px]'>
      <div className="text-slate-700 text-sm md:text-base lg:text-[17px]">Two-factor Authentication</div>
      <div className="flex gap-5 md:gap-6 mt-4">
        <Toggle />
        <div className="text-slate-700 text-sm md:text-base lg:text-[17px]">
          Enable or disable two-factor authentication
        </div>
      </div>
      
      <div className="mt-10 text-slate-700 text-sm md:text-base lg:text-[17px] text-[17px]">Change Password</div>
      
      <div className="mt-4">
        <div className="text-slate-700 text-sm md:text-base lg:text-[17px]">Current Password</div>
        <input
          type="password"
          className="border-slate-200 border-[1px] w-full h-10 mt-3 rounded-3xl md:w-[20rem] lg:w-[30rem]"
          style={{ paddingLeft: '1.25rem' }}
          placeholder="************"
        />
      </div>
      
      <div className="mt-4">
        <div className="text-slate-700 text-sm md:text-base lg:text-[17px]">New Password</div>
        <input
          type="password"
          className="border-slate-200 border-[1px] w-full h-10 mt-3 rounded-3xl md:w-[20rem] lg:w-[30rem]"
          style={{ paddingLeft: '1.25rem' }}
          placeholder="************"
        />
      </div>
      
      <div className="flex justify-end mt-16 md:mt-18">
        <button className="border-none bg-blue-700 text-white w-full h-12 rounded-full md:w-[12rem] text-[13px] md:text-base">
          Save
        </button>
      </div>
    </div>
  );
}
