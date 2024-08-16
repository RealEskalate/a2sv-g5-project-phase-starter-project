import React from 'react'
interface props{
    title:string;
    amount: string;
    icon:string;
    color:string;
    width:string
}
const Card = ({title,amount,icon,color , width}:props) => {
  return (
    <div className={`flex border border-gray-300 ${width}  justify-center items-center rounded-3xl py-2 gap-2 lg:gap-7 bg-white shadow-lg min-w-[170px]`}>
        <div className='border  flex justify-center items-center rounded-full w-[45px] h-[45px] lg:w-[70px] lg:h-[70px]' style={{ backgroundColor: color , borderColor:color}}>
            <img src={icon}/>
        </div>

        <div>
            <p className='text-[#718EBF] font-normal text-base font-inter' >{title}</p>
            <p className='text-[#232323] font-semibold text-2xl font-inter'>{amount}</p>
        </div>

      
    </div>
  )
}

export default Card
