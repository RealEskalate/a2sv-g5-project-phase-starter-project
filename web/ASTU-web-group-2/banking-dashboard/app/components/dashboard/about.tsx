import React from 'react';

const About = () => {
  return (
    <div className='flex flex-col sm:flex-row justify-center items-start p-8 gap-32 min-h-screen'>
      <div className='flex flex-col items-center justify-center w-full md:w-1/3  rounded-lg p-6 '>
        <h1 className='text-blue-800 text-center mb-4 text-2xl md:text-3xl font-extrabold'>About Us</h1>
        <div className='mb-4'>
          <img src='/about.svg' alt='About Us' className='w-full h-auto max-w-xs md:max-w-md' />
        </div>
        <p className='text-md text-gray-700 text-center'>
          Bank  - Your trusted financial partner for loans. Quick approvals, competitive rates,
          and personalized solutions to meet your unique needs. Empowering you to achieve your financial goals.
          Apply online today!
        </p>
      </div>

      <div className='flex flex-col items-center justify-center w-full sm:w-1/3 rounded-lg p-6 m-4'>
        <h1 className='text-blue-800 text-center mb-4 text-2xl md:text-3xl font-extrabold'>Contact Us</h1>
        <input
          className='w-full text-center p-4 h-12 mb-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
          name='name'
          placeholder='Full Name'
          type='text'
        />
        <input
          className='w-full text-center p-4 h-12 mb-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
          name='email'
          placeholder='Email Address'
          type='email'
        />
        <textarea
          className='w-full h-48 text-center p-4 mb-4 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
          name='message'
          placeholder='Message'
        
        />
        <button className='text-center h-12 rounded-md text-white w-1/3 bg-blue-800 hover:bg-blue-700'>
          Send
        </button>
      </div>
    </div>
  );
}

export default About;
