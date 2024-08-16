'use client'

import React from 'react'
import InputBox from './InputBox'

const page = () => {
  return (
    <div>
      <div className='flex'>
        <div className='w-[20%] bg-green-500'>

        </div>
        <div className='w-[80%]'>
          <div className='flex'>
            <div className='w-[50%]'>
              <InputBox/>
            </div>
            <div className='W-[50%]'>
              <InputBox/>
            </div>

          </div>
          
        </div>

      </div>
        
        
    </div>
  )
}

export default page