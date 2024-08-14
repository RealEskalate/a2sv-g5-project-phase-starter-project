import React from 'react'

const MobileCreditCard = () => {
  return (
    <div className='bg-blue-700 w-[265px] h-[170px] rounded-xl relative'>
      <div className="flex justify-between w-[95%]">
        <div className='mt-1 ml-3 p-2'>
            <span className='text-[11px]'>Balance</span>
            <span className='block text-[16px] font-bold'>$5,756</span>
        </div>

        <img src="icons/chip.png" alt="chip card" className='w-[29px] h-[29px] mt-4 mr-2' />
      </div>

      <div className="flex justify-between w-[90%]">
        <div className='ml-3 pl-1.5'>
            <span className='text-[10px] text-gray-300'>CARD HOLDER</span>
            <span className='block text-[13px] font-bold'>Tekola Chane</span>
        </div>

        <div className='mr-3'>
            <span className='text-[10px] text-gray-300'>VALID THRU</span>
            <span className='block text-[13px] font-bold'>12/22</span>
        </div>
      </div>
      
      <div className="flex justify-between child-div absolute bottom-0 left-0 right-0 bg-gradient-to-b from-blue-600">
        <span className='p-3 ml-2 text-[15px]'>3778 **** **** 1234</span>
        <img src="icons/masterCard.png" alt="master card icon" className='w-[35px] h-[33px] mt-1.5 mr-3' />
      </div>

    </div>
  )
}

export default MobileCreditCard
