"use client"
import React from 'react'
import { Switch } from "@/components/ui/switch"


const PrefPage = () => {
  
  return (
    <div className='flex flex-col mt-10 text-sm space-y-10'>
        <div className='flex justify-between'>
            <div className='flex flex-col gap-3'>
                <div className='text-[#232323]'>Currency</div>
                <input placeholder='currency'  className='text-[#DFEAF2] rounded-xl w-[510px] border border-[#DFEAF2] py-3 px-5'/>
            </div>
            <div className='flex flex-col gap-3'>
                <div className='text-[#232323]'>Time Zone</div>
                <input placeholder='TimeZone' className='text-[#DFEAF2] rounded-xl w-[510px] border border-[#DFEAF2] py-3 px-5'/>
            </div>
        </div>
        <div className='flex flex-col gap-5'>
            <div className='font-semibold'>Notification</div>
            <div className='flex items-center gap-4'>
                <Switch />
                <div>I send or receive digita currency</div>
            </div>
            <div className='flex items-center gap-4'>
                <Switch />
                <div>I receive merchant order</div>
            </div>
            <div className='flex items-center gap-4'>
                <Switch className='bg-[#16DBCC]'/>
                <div>There are recommendation for my account</div>
            </div>

        </div>
        <div className="flex w-full justify-end mt-10  px-[30px] ">
          <button className="px-10 py-3 text-white rounded-xl bg-[#1814F3]">
            {" "}
            Save
          </button>
        </div>
    </div>
  )
}

export default PrefPage