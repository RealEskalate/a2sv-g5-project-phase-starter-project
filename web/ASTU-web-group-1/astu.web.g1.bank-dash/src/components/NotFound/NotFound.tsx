'use client'
import { useRouter } from 'next/navigation';
import React from 'react';

export default function NotFound() {
  const router = useRouter();
  const handleBack = () => {
    router.back();
  };

  return (
    <div className='flex items-center justify-around w-screen min-h-screen bg-white'>
      <div>
        <h1 className='text-[60px] font-bold text-blue-bright'>Oops...</h1>
        <p className='text-3xl text-slate-600 font-[600]'>Page not found</p>
        <h3 className='my-2 font-[400] text-slate-500'>Message</h3>
        <button
          className='border bg-blue-steel border-blue-600 text-white font-[550] px-3 py-1 rounded-xl hover:scale-105 hover:bg-blue-bright transition-all duration-300'
          onClick={handleBack}
        >
          Go Back
        </button>
      </div>
      <div className='rounded-[36px] shadow-lg max-w-md'>
        <img
          src='https://www.digitalmesh.com/blog/wp-content/uploads/2020/05/404-error.jpg'
          alt='404 Error'
        />
      </div>
    </div>
  );
}
