import React from 'react';
import Image from 'next/image'

interface Props {
    image: string,
    title: string,
    value: string
}

const Card = ({ image, title, value }: Props) => {
  return (
    <div className='flex items-center gap-4 p-6 bg-white rounded-xl'>
      <Image src={image} alt='title' className='w-16 h-16'/>
      <div className='flex flex-col items-start justify-center gap-2'>
        <p>{title}</p>
        <span className='font-bold'>${value}</span>
      </div>
    </div>
  )
}

export default Card
