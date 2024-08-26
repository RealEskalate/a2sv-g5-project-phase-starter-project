import React from 'react'

const CardShimmer = () => {
  return (
    <div className="flex flex-col md:flex-row gap-3 md:gap-10 justify-center pt-4 w-full">
    {Array(3).fill("").map((_, index) => (<div key={index} className="w-full h-[80px] md:w-[30%] bg-gray-300 animate-pulse rounded-2xl"></div>))}
  </div>
  )
}

export default CardShimmer
