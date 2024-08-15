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
    <div className={`flex border w-full lg:${width}  justify-center items-center rounded-3xl py-2 gap-7`}>
        <div className='border  flex justify-center items-center rounded-full w-[70px] h-[70px]' style={{ backgroundColor: color , borderColor:color}}>
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
