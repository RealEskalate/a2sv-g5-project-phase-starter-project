import React from 'react'
interface props{
    data: string [][]
}

const TrendingStock = ({data}:props) => {
  return (
    <div className='border border-solid rounded-3xl overflow-hidden w-[100%]  bg-white'>
    <table className='w-full' >
        <thead>
            <tr className='text-left '>
                <th className='px-6 py-6 font-medium text-base text-[#718EBF]'>SL No</th>
                <th className='px-6 py-6 font-medium text-base text-[#718EBF]'>Name</th>
                <th className='px-6 py-6 font-medium text-base text-[#718EBF]'>Price</th>
                <th className='px-6 py-6 font-medium text-base text-[#718EBF]'>Return</th>

            </tr>
        </thead>
        <tbody>
        {data.map((row, rowIndex) => (
          <tr key={rowIndex} >
            {row.map((cell, cellIndex) => {
              const isLastCell = cellIndex === row.length - 1;
              const cellColor = isLastCell && typeof cell === 'string' && cell.startsWith('-')
                ? 'text-[#FE5C73]'
                : 'text-[#16DBAA]';
              const cellClasses = isLastCell
                ? `px-6 py-4 whitespace-nowrap font-inter font-normal text-base ${cellColor}`
                : 'px-6 py-4 whitespace-nowrap font-inter font-normal text-base text-[#232323]';

              return (
                <td key={cellIndex} className={cellClasses}>
                  {cell}
                </td>
              );
            })}
          </tr>
        ))}
        </tbody>
    </table>
    </div>
  )
}

export default TrendingStock
