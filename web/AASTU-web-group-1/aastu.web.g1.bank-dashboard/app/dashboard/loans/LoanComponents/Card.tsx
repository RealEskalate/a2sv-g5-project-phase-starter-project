import React from 'react'
import { IoPersonCircleSharp } from "react-icons/io5";
export default function Card(){
  return (
    <div className='h-[10rem] w-[25rem] bg-white rounded-3xl m-2 flex justify-around'>
            <IoPersonCircleSharp size='100' className='my-auto w-1/3 pl-3' color='blue'/>
            <div className='pl-3 my-auto w-2/3'>
              <h1 className='text-bg-gray text-xl'>Personal Loans</h1>
              <p className='font-[500] text-xl'>$50,000</p>
            </div>
    </div>
  )
}

