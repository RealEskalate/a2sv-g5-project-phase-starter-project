'use client'
import Image from 'next/image'
import React from 'react'

const BankServiceList = () => {
  return (
   

    
    <div className="p-4 pr-60 pl- bg-gray-100">
       <div className="flex   items-center bg-white shadow rounded-lg p-4 gap-4 ">
        

        
          <div className=" bg-pink-100 rounded-xl">
            <div className="flex-row items-center justify-center  w-12 h-12">
              <Image
                src="/images/dollar.png"
                alt="dollar Icon"
                // layout="fill"
                // objectFit="cover"
                width={20}
                height={20}
                className="flex items-center justify-center"
              />
              <Image
                src="/images/hand.png"
                alt="hand Icon"
                // layout="fill"
                // objectFit="contain"
                width={30}
                height={30}
                className="flex items-center justify-center"
              />
            </div>
          </div>

          
          <div className="flex gap-6">
            <div>
                <h3 className="font-sans text-black text-lg mb-1">Business Loans</h3>
                <p className="font-sans text-blue-900 opacity-60 text-sm">It is a long established</p>
            </div>
            <div>
                <h3 className="font-sans text-black text-lg mb-1">Lorem Ipsum</h3>
                <p className="font-sans text-blue-900 opacity-60 text-sm">Many Publishing</p>
            </div>
            <div>
                <h3 className="font-sans text-black text-lg mb-1">Lorem Ipsum</h3>
                <p className="font-sans text-blue-900 opacity-60 text-sm">Many Publishing</p>
            </div>
            <div>
                <h3 className="font-sans text-black text-lg mb-1">Lorem Ipsum</h3>
                <p className="font-sans text-blue-900 opacity-60 text-sm">Many Publishing</p>
            </div>
            <div className=' py-4'>
                <button className="text-blue-900 opacity-60 border border-gray-400 rounded-full px-3 py-1  hover:bg-blue-50">
                    View Details
                </button>
          </div>
            
          </div>
          </div>
        
        {/* Repeat the above structure for other services */}
      
    </div>
    
  )
}

export default BankServiceList
