'use client'
import React from 'react'

const Input = () => {
  return (
    <div className='flex-col bg-white p-4'>
      <div className='mb-2'>
        <p className='text-black font-sans text-lg'>Your Name</p>
      </div>
      <div>
        <input 
          type='text' 
          className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
          placeholder='Charlene Reed' 
        />
      </div>
    </div>
  )
}

export default Input
