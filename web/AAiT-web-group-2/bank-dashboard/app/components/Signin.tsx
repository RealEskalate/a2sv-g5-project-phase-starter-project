'use client';

import { useForm, SubmitHandler } from 'react-hook-form';
import { signIn } from 'next-auth/react';
import { useRouter } from 'next/navigation';
import Link from 'next/link';
import { useEffect, useState } from 'react';
import Alert from './Alert';

interface SignInFormData {
  username: string;
  password: string;
}

export default function SignIn() {
  const { register, handleSubmit, formState: { errors } } = useForm<SignInFormData>();
  const router = useRouter();
  const [signinStatus, setSigninStatus] = useState('')
  const [isError, setIsError] = useState(false)
  const [isSuccess, setIsSuccess] = useState(false)

  const reset = () => {
    setIsSuccess(false)
    setIsError(false)
  }

  
  const onSubmit: SubmitHandler<SignInFormData> = async (data) => {
    console.log("data", data)
    const result = await signIn('sign-in', {
      redirect: false,
      userName: data.username,
      password: data.password,
    });
    console.log("result", result?.error)
    if(result?.error){
      setSigninStatus(result?.error as string)
      setIsError(true)      
    } else{
      setSigninStatus('Successfully logged in.')
      setIsSuccess(true)

    }
    

   
    
    if (result?.ok) {
      router.push('/dashboard');
    } else {
      console.error(result?.error);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center ">
      <form onSubmit={handleSubmit(onSubmit)} className="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
         <h1 className='font-bold mb-4 text-custom-purple text-3xl text-center '>Welcome to BankDash</h1>
        
        <h2 className="text-2xl font-bold mb-6 text-center">Sign In</h2>

        <div className="flex flex-col mb-4">
          <label htmlFor="username" className="mb-2 text-base font-semibold text-gray-700">Username</label>
          <input
            id="username"
            {...register('username', { required: 'Username is required' })}
            placeholder="Username"
            className="p-3 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          {errors.username && <span className="text-red-500 text-sm mt-1">{errors.username.message}</span>}
        </div>

        <div className="flex flex-col mb-6">
          <label htmlFor="password" className="mb-2 text-base font-semibold text-gray-700">Password</label>
          <input
            id="password"
            type="password"
            {...register('password', { required: 'Password is required' })}
            placeholder="Password"
            className="p-3 border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          {errors.password && <span className="text-red-500 text-sm mt-1">{errors.password.message}</span>}
        </div>

        <button type="submit" className="w-full py-3 bg-custom-purple text-white font-bold rounded-lg hover:shadow-md transition duration-300">
          Sign In
        </button>
        <p className='text-gray-500 font-body font-normal text-base mt-6 mb-6'>Don't have an account? <Link className='text-custom-purple font-semibold text-base' href='/auth/signup'>Sign Up</Link></p>
        <p className='text-gray-500 font-body font-normal text-sm'>By clicking 'Signin', you acknowledge that you have read and accepted our terms of <Link className='text-purple-tag'  href={'/terms-of-service'}>Terms of Service</Link> and <Link className='text-purple-tag '  href={'/privacy-policy'}>Privacy Policy</Link></p>

      </form>
        {isSuccess && <Alert type='success' message={signinStatus} duration={2000} onClose={reset} />}
        {isError && <Alert type='error' message={signinStatus} duration={2000} onClose={reset} />}

    </div>
  );
}
