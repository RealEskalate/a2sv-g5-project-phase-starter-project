import React, { ReactNode } from 'react'
import { FaUser } from 'react-icons/fa'
interface LoanProps {
  name:string
  amount:number
  color:string
  icon:ReactNode
}

const Card = () => {
  return (
    
    <div className='flex w-64 border-0 rounded-xl bg-white min-h-32 gap-3 items-center '>
      <div className="icons border-1 rounded-full ml-4 bg-gray-100 h-16 w-16 flex items-center justify-center">
          <FaUser color='#396AFF' size={30}/>
      </div>
      <div className="info">
        <p className='tex-[#718EBF] mt-2'>Personal Loans</p>
        <p className='font-semibold text-xl'>$500</p>
      </div>
    </div>
  )
}

export default Card
