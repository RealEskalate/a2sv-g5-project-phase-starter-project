"use client"
import React from 'react'
import { Switch } from "@/components/ui/switch"


const PrefPage = () => {
  
  return (
    <div className=' flex flex-col mt-[43px]  h-[400px]'>
        <div className='flex flex-row gap-[30px]'>
            <div className='w-[510px] h-[80px]  flex flex-col gap-3'>
                <div className='text-[#232323] w-fit h-[19px]'>Currency</div>
                <input placeholder='currency'  className='text-[#DFEAF2] rounded-[15px] w-[510px] h-[50px] border border-[#DFEAF2] py-4 px-5'/>
            </div>
            <div className='w-[510px] h-[80px]  flex flex-col gap-3'>
                <div className='text-[#232323] w-fit h-[19px]'>Time Zone</div>
                <input placeholder='TimeZone' className='text-[#DFEAF2] rounded-[15px] w-[510px] h-[50px] border border-[#DFEAF2] py-4 px-5'/>
            </div>
        </div>
        <div className='flex flex-col mt-[27px] gap-5'>
            <div className='font-semibold'>Notification</div>
            <div className='flex flex-row gap-[15px]'>
                <Switch />
                <div>I send or receive digita currency</div>
            </div>
            <div className='flex flex-row gap-[15px]'>
                <Switch />
                <div>I receive merchant order</div>
            </div>
            <div className='flex flex-row gap-[15px]'>
                <Switch className='bg-[#16DBCC]'/>
                <div>There are recommendation for my account</div>
            </div>

        </div>
        <div className="flex w-full justify-end mt-10  px-[30px] ">
          <button className=" w-[190px] h-[50px] text-white py-[14px] px-[74px] rounded-[15px] bg-[#1814F3]">
            {" "}
            Save
          </button>
        </div>
    </div>
  )
}

export default PrefPage