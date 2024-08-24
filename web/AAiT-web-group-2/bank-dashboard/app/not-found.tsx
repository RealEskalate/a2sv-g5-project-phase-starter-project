import React from 'react'

const NotFound = () => {
  return (
    <div className='flex flex-col gap-4 justify-center items-center h-full'>
        <p className='text-custom-purple font-body font-semibold text-4xl'>404</p>
        <h1 className='font-medium text-3xl' > The Page you are looking for does not exist.</h1>
    </div>
  )
}

export default NotFound