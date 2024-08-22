import LoginForm from '@/components/Formx/LoginForm';
import React from 'react';
import Image from 'next/image';
import AuthProvider from '@/components/Formx/AuthProvider';

export default function Page() {
  return (
    <>
      <div className='bg-white w-screen h-screen flex justify-around items-center'>
        <div className='text-center hidden min-[800px]:block'>
          <h1 className='text-indigo-900 text-4xl font-poppins font-[900]'>Welcome To Bank-Dash</h1>
          <div className='w-[35vw] h-[60vh] relative'>
            <Image
              src='/assets/images/welcome-page.png'
              alt='hello'
              layout='fill'
              objectFit='cover'
              sizes='100vw'
              className='object-cover'
            />
          </div>
        </div>
        <AuthProvider>
          <LoginForm />
        </AuthProvider>
      </div>
    </>
  );
}
