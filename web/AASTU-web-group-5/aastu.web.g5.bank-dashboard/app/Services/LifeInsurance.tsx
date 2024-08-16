'use client'
import React from 'react'
import Image from 'next/image'

const LifeInsurance = () => {
  return (
    <div className="w-full flex items-center bg-gray-100 shadow rounded-lg p-4 gap-5"> {/* Increased padding for height */}
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4 w-full">
        <div className="bg-white p-5 pr-6 rounded-xl flex-grow"> {/* Increased padding for height */}
          <div className="flex items-center gap-2">
            <div className="flex items-center justify-center bg-blue-300 rounded-full w-12 h-12">
              <Image
                src="/images/heart.png"
                alt="heart Icon"
                objectFit="cover"
                width={25}
                height={25}
              />
            </div>
            <div>
              <h3 className="font-sans text-black text-2xl font-semibold">Life Insurance</h3>
              <p className="font-sans text-blue-900 opacity-60 text-sm">Unlimited Protection</p>
            </div>
          </div>
        </div>
        <div className="bg-white p-6 rounded-xl flex-grow"> {/* Increased padding for height */}
          <div className="flex items-center gap-2">
            <div className="flex items-center justify-center bg-blue-300 rounded-full w-12 h-12">
              <Image
                src="/images/heart.png"
                alt="heart Icon"
                objectFit="cover"
                width={25}
                height={25}
              />
            </div>
            <div>
              <h3 className="font-sans text-black text-2xl font-semibold">Shopping</h3>
              <p className="font-sans text-blue-900 opacity-60 text-sm">Buy. Think. Grow</p>
            </div>
          </div>
        </div>
        <div className="bg-white p-6 rounded-xl flex-grow"> {/* Increased padding for height */}
          <div className="flex items-center gap-2">
            <div className="flex items-center justify-center bg-blue-300 rounded-full w-12 h-12">
              <Image
                src="/images/heart.png"
                alt="heart Icon"
                objectFit="cover"
                width={25}
                height={25}
              />
            </div>
            <div>
              <h3 className="font-sans text-black text-2xl font-semibold">Safety</h3>
              <p className="font-sans text-blue-900 opacity-60 text-sm">We are your allies</p>
            </div>
          </div>
        </div>
        {/* Additional items if needed */}
      </div>
    </div>
  )
}

export default LifeInsurance
