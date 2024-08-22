import React from 'react'
import { IconType } from 'react-icons'

interface Cardprops{
  icon: IconType
  title:string 
  sub_title:string
}

const card = ({icon:Icon,title,sub_title}:Cardprops) => {
  return (
    <div className=' inline-flex   gap-4  items-center rounded-[25px] bg-white px-16 py-7'>
        <Icon className='w-[70px] h-[70px]'/>
        <div >
            <p className='font-inter font-semibold text-[20px] leading-[24.2px] text-[#232323]'>{title}</p>
            <p className='font-normal text-[#718EBF] text-[16px] leading-[19.36px] mt-2'>{sub_title}</p>
        </div>
    </div>
  )
}

export default card