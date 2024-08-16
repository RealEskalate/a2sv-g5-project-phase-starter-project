'use client'
import Image from 'next/image'
import React from 'react'

const BankServiceList = () => {
  return (
    <div className="bg-gray-100 p-3"> {/* Gray container with padding */}
      <div className="flex flex-col sm:flex-row items-center bg-white shadow rounded-lg p-3 gap-4 w-full"> {/* White inner container */}
        <div className="bg-pink-100 rounded-xl flex-col items-center justify-center w-full sm:w-auto p-4">
          <div className="flex items-center justify-center">
            <Image
              src="/images/dollar.png"
              alt="dollar Icon"
              width={20}
              height={20}
            />
          </div>
          <div className="flex items-center justify-center ml-4">
            <Image
              src="/images/hand.png"
              alt="hand Icon"
              width={30}
              height={30}
            />
          </div>
        </div>

        <div className="flex flex-col sm:flex-row justify-between gap-10 flex-grow w-full">
          <div className="flex flex-col">
            <h3 className="font-sans text-black text-lg mb-1">Business Loans</h3>
            <p className="font-sans text-blue-900 opacity-60 text-sm">It is a long established</p>
          </div>
          <div className="flex flex-col">
            <h3 className="font-sans text-black text-lg mb-1">Lorem Ipsum</h3>
            <p className="font-sans text-blue-900 opacity-60 text-sm">Many Publishing</p>
          </div>
          <div className="flex flex-col">
            <h3 className="font-sans text-black text-lg mb-1">Lorem Ipsum</h3>
            <p className="font-sans text-blue-900 opacity-60 text-sm">Many Publishing</p>
          </div>
          <div className="flex flex-col">
            <h3 className="font-sans text-black text-lg mb-1">Lorem Ipsum</h3>
            <p className="font-sans text-blue-900 opacity-60 text-sm">Many Publishing</p>
          </div>
          <div className="flex items-center py-4">
            <button className="text-blue-900 opacity-60 border border-gray-400 rounded-full px-3 py-1 hover:bg-blue-50">
              View Details
            </button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default BankServiceList
