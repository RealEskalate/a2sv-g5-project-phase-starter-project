'use client';

import React from 'react';

export default function ErrorPage() {
  return (
    <div className='flex flex-col items-center justify-center min-h-screen bg-gray-100'>
      <h1 className='text-6xl font-bold text-red-600 mb-4'>Something went wrong</h1>
      <p className='text-xl text-gray-600 mb-8'>
        We encountered an unexpected error. Please try again later.
      </p>
      <button
        onClick={() => window.location.reload()}
        className='px-6 py-3 bg-blue-600 text-white font-semibold rounded-lg shadow-md hover:bg-blue-700 transition-colors duration-300'
      >
        Reload Page
      </button>
    </div>
  );
}
