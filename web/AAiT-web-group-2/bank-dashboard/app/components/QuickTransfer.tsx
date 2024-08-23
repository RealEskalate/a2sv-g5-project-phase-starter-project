import React from 'react'

const QuickTransfer = () => {
  return (
    <div className='bg-[#F5F7FA]'>
      <div className="bg-white px-5 py-5 w-[445px] rounded-xl">
      <div className=' flex items-center gap-5 '>
        <div className='font-bold flex flex-col items-center text-center'>
            <img  src="/profile1.png" alt="" />
            <p className='pt-4'>Livia Bator</p>
            <p className='text-[#718EBF]'>CEO</p>
        </div>
        <div className='flex flex-col items-center text-center'>
            <img src="/profile2.png" alt="" />
            <p className='pt-4'>Randy Press</p>
            <p className='text-[#718EBF]'>Director</p>
        </div>
        <div className='flex flex-col items-center text-center'>
            <img src="/profile3.png" alt="" />
            <p className='pt-4'>Workman</p>
            <p className='text-[#718EBF]'>Designer</p>
        </div>

      <img src="/vector2.png" alt="" />        
      </div>

      <div className="flex items-center pt-5 gap-8">
        <p className="text-[#718EBF]">Write amount</p>
        <div className=" h-[50px] relative">
          <input className='focus:outline-none px-5 bg-[#EDF1F7] h-[50px] rounded-3xl text-[#718EBF]'type="text" />
          <div className='absolute right-0 top-0 rounded-3xl px-5 w-[140px] cursor-pointer flex justify-center items-center gap-2 h-[50px] bg-[#1814F3] text-white'>
            <p>send</p> 
            <img src="/send.png" alt="" />
          </div>
        </div>
      </div>
      </div>
      
    </div>
  )
}

export default QuickTransfer
