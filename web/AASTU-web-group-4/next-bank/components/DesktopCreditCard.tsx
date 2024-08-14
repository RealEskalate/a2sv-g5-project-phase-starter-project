import React from 'react'

const DesktopCreditCart = () => {
  return (
    <div className='bg-blue-700 w-[350px] h-[235px] rounded-xl relative'>
      <div className="flex justify-between w-[95%]">
        <div className='mt-1 ml-3 p-2'>
            <span className='text-[12px]'>Balance</span>
            <span className='block text-[20px] font-bold'>$5,756</span>
        </div>

        <img src="icons/chip.png" alt="chip card" className='w-[29px] h-[29px] mt-4 mr-2' />
      </div>

      <div className="flex justify-between w-[90%]">
        <div className='mt-1 ml-3 p-2'>
            <span className='text-[12px] text-gray-300'>CARD HOLDER</span>
            <span className='block text-[15px] font-bold'>Tekola Chane</span>
        </div>

        <div className='mt-1 mr-9 p-2'>
            <span className='text-[12px] text-gray-300'>VALID THRU</span>
            <span className='block text-[15px] font-bold'>12/22</span>
        </div>
      </div>
      
      <div className="flex justify-between child-div absolute bottom-0 left-0 right-0 bg-gradient-to-b from-blue-600">
        <span className='p-3 ml-2 text-[22px]'>3778 **** **** 1234</span>
        <img src="icons/masterCard.png" alt="master card icon" className='w-[44px] h-[42px] mt-1.5 mr-3' />
      </div>

    </div>
  )
}

export default DesktopCreditCart
