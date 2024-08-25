import React from 'react'
import Image from 'next/image'

interface ListCardProps {
    imgpath: string,
    heading: textProps,
    description1: textProps,
    description2: textProps,
    description3: textProps,
}
interface textProps {
    title: string,
    description: string
}

const colors={

}


const ListCard = ({imgpath, heading, description1, description2, description3} : ListCardProps) => {
  return (
    <div className='flex bg-white py-3 pl-3 lg:p-4 rounded-2xl mr-4 mb-2 sm:mb-5'>
        <div className='mr-3 lg:mr-8 w-14 lg:w-16 my-auto'>
            <Image src={imgpath} width={60} height={60} alt='saving icon'/>
        </div>
        <div className='grid grid-cols-5 sm:grid-cols-11 w-full'>
            <div className='my-auto col-span-3 sm:col-span-3'>
                <div className='font-medium text-sm sm:text-base lg:text-lg '>{heading.title}</div>
                <div className='text-[#718EBF] text-xs lg:text-lg inline-block'>{heading.description}</div>
            </div>
            <div className='my-auto col-span-2 hidden sm:block'>
                <div className='font-medium text-sm sm:text-base lg:text-lg'>Lorem Ipsum</div>
                <div className='text-[#718EBF] text-xs lg:text-lg inline-block'>Many publishing</div>
            </div>
            <div className='my-auto col-span-2 hidden sm:block'>
                <div className='font-medium text-sm sm:text-base lg:text-lg'>Lorem Ipsum</div>
                <div className='text-[#718EBF] text-xs lg:text-lg inline-block'>Many publishing</div>
            </div>
            <div className='my-auto col-span-2 hidden sm:block'>
                <div className='font-medium text-sm sm:text-base lg:text-lg'>Lorem Ipsum</div>
                <div className='text-[#718EBF] text-xs lg:text-lg inline-block'>Many publishing</div>
            </div>
            <button className='rounded-full sm:border-2 lg:border-2 sm:border-barTextGray w-fit h-fit my-auto py-1 sm:px-2 lg:px-6  flex align-middle sm:hover:border-[#1814F3] col-span-2'>
                <span className='text-[#1814F3] sm:text-[#718EBF] text-xs lg:text-lg font-medium sm:text-barTextGray my-auto sm:hover:text-[#1814F3]'>View Details</span>
            </button>
        </div>
    </div>
  )
}

export default ListCard