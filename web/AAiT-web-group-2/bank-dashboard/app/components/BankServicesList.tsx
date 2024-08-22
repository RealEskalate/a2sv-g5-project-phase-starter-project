import React from 'react'

const BankServicesList = () => {
  return (
    <div className='w-[1000px] bg-[#F5F7FA]'>
        <p className='text-[#343C6A] text-[22px]'>Bank Services List</p>
        <div className="bg-white rounded-2xl h-[90px] flex justify-around items-center mt-5">
        <img src="/loans.png" alt="" />
        <div>
            <p> Business loans</p>
            <p className="font-light text-[#718EBF]">it is a long established</p>
        </div>
        <div>
            <p> Lorem Ipsum</p>
            <p className="font-light text-[#718EBF]">Many publishing</p>
        </div>
        <div>
            <p> Lorem Ipsum</p>
            <p className="font-light text-[#718EBF]">Many publishing</p>
        </div>
        <div>
            <p> Lorem Ipsum</p>
            <p className="font-light text-[#718EBF]">Many publishing</p>
        </div>
        <div className="rounded-3xl text-[#718EBF] border px-9 py-1 hover:text-[#1814F3] hover:border-blue-800">
            view details
        </div>
      </div>
    </div>
  )
}

export default BankServicesList
