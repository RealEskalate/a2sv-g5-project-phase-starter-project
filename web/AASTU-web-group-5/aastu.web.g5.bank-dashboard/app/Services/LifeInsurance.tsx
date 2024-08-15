'use client'
import React from 'react'
import Image from 'next/image'

const LifeInsurance = () => {
  return (
    <div className="w-full flex  items-center gap-6 pr-4 pl-4 bg-gray-100">
    <div  className="flex   items-center bg-white shadow rounded-lg p-4 gap-4 ">

    
      
    <div className=" flex">
        <div className='p-4'>
        <div className='bg-white p-2 py- pr-5 rounded-xl'>
            <div className='flex items-center gap-2 '>
                <div className='flex items-center justify-center bg-blue-300 rounded-full w-12 h-12'>
                    <Image
                    src="/images/heart.png"
                    alt="heart Icon"
                    // layout="fill"
                    objectFit="cover"
                    width={25}
                    height={25}
                    />
                </div>
                <div >
                    <h3 className="font-sans text-black text-lg font-semibold ">Life Insurance</h3>
                    <p className="font-sans text-blue-900 opacity-60 text-sm">Unlimited Protection</p>
                </div>

            </div>

        </div>
        </div>
        <div className='p-4'>
        <div className='bg-white p-2 py- pr-5 rounded-xl'>
            <div className='flex items-center gap-2 '>
                <div className='flex items-center justify-center bg-blue-300 rounded-full w-12 h-12'>
                    <Image
                    src="/images/heart.png"
                    alt="heart Icon"
                    // layout="fill"
                    objectFit="cover"
                    width={25}
                    height={25}
                    />
                </div>
                <div >
                    <h3 className="font-sans text-black text-lg font-semibold ">Shopping</h3>
                    <p className="font-sans text-blue-900 opacity-60 text-sm">Buy. Think. Grow</p>
                </div>

            </div>

        </div>
        </div>
        <div className='p-4'>
        <div className='bg-white p-2 py- pr-5 rounded-xl'>
            <div className='flex items-center gap-2 '>
                <div className='flex items-center justify-center bg-blue-300 rounded-full w-12 h-12'>
                    <Image
                    src="/images/heart.png"
                    alt="heart Icon"
                    // layout="fill"
                    objectFit="cover"
                    width={25}
                    height={25}
                    />
                </div>
                <div >
                    <h3 className="font-sans text-black text-lg font-semibold ">Saftey</h3>
                    <p className="font-sans text-blue-900 opacity-60 text-sm">We are your allies</p>
                </div>

            </div>

        </div>
        </div>
        
       
        </div>

        </div>
  </div>
  )
}

export default LifeInsurance