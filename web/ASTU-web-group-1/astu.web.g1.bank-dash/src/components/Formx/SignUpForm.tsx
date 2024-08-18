'use client';

import { zodResolver } from '@hookform/resolvers/zod';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import Image from 'next/image';
import { signIn } from 'next-auth/react';

const SignUpFormSchema = z
  .object({
    fullName: z.string().min(3, {
      message: 'Full name must be at least 3 characters',
    }),
    email: z.string().min(1, { message: 'Email is required' }).email('Invalid email address'),
    password: z.string().min(6, { message: 'Password must be at least 6 characters' }),
    confirmPassword: z.string().min(6, { message: 'Password must be at least 6 characters' }),
  })
  .refine((value) => value.password === value.confirmPassword, {
    message: 'Passwords do not match',
    path: ['confirmPassword'],
  });

type FormData = z.infer<typeof SignUpFormSchema>;

export default function SignUpForm() {
  const router = useRouter();

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(SignUpFormSchema),
  });

  const onSubmit = async (data: FormData) => {
    const response = await fetch(`https://akil-backend.onrender.com/signup`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name: data.fullName,
        email: data.email,
        password: data.password,
        confirmPassword: data.confirmPassword,
        role: 'user',
      }),
    });
    const res = await response.json();
    if (!res.success) {
      alert(res.message);
    } else {
      router.push(`/api/auth/verify?email=${data.email}`);
      alert(res.message);
    }
  };

  const handleGoogleSignUp = async () => {
    const response = await signIn('google');
  };

  return (
    <div className='px-5 pb-10 bt-5 shadow-2xl rounded-xl relative'>
      <h1 className='text-indigo-900 text-4xl m-5 text-center font-poppins font-[1000]'>
        Sign Up Today!
      </h1>
      <div
        className='flex w-full justify-center border-2 py-1.5 items-center text-indigo-950 border-indigo-300 rounded-md'
        onClick={handleGoogleSignUp}
      >
        <Image src={'/assets/google.png'} width={18} height={18} alt='google' />
        <span className='font-poppins mx-1 font-[600] text-sm'>Sign Up With Google</span>
      </div>
      <p className='text-[14px] text-indigo-950 font-[600] font-poppins min-w-44 text-center flex justify-center mx-20 left-1/2 before:left-0 before:-z-10 bg-white before:bg-indigo-900 before:h-0.5 overflow-hidden before:mt-1 before:w-full before: before:absolute before:text-start before:items-start mt-5'>
        Or Sign Up with Google
      </p>
      <form onSubmit={handleSubmit(onSubmit)} className='md:min-w-[350px] sm:min-w-[300px]'>
        <div className='my-2'>
          <label
            htmlFor='fullName'
            className='block mb-1 font-epilogue text-sm font-[600] text-indigo-900'
          >
            Full Name
          </label>
          <input
            type='text'
            id='fullName'
            {...register('fullName')}
            className='w-full font-[600] font-epilogue outline-none rounded-lg p-2 text-indigo-950 text-sm border border-slate-400'
          />
          {errors.fullName && (
            <p className='text-red-500 text-xs mt-1 font-poppins font-[550] md:max-w-[400px]'>
              {errors.fullName.message}
            </p>
          )}
        </div>

        <div className='my-2'>
          <label
            htmlFor='email'
            className='block mb-1 font-epilogue text-sm font-[600] text-indigo-900'
          >
            Email Address
          </label>
          <input
            type='email'
            id='email'
            {...register('email')}
            className='w-full font-[600] font-epilogue outline-none rounded-lg p-2 text-indigo-950 text-sm border border-slate-400'
          />
          {errors.email && (
            <p className='text-red-500 text-xs mt-1 font-poppins font-[550] md:max-w-[400px]'>
              {errors.email.message}
            </p>
          )}
        </div>

        <div className='my-2'>
          <label
            htmlFor='password'
            className='block mb-1 font-epilogue text-sm font-[600] text-indigo-900'
          >
            Password
          </label>
          <input
            type='password'
            id='password'
            {...register('password')}
            className='w-full font-[600] font-epilogue outline-none rounded-lg p-2 text-indigo-950 text-sm border border-slate-400'
          />
          {errors.password && (
            <p className='text-red-500 text-xs mt-1 font-poppins font-[550] md:max-w-[400px]'>
              {errors.password.message}
            </p>
          )}
        </div>

        <div className='my-2'>
          <label
            htmlFor='confirmPassword'
            className='block mb-1 font-epilogue text-sm font-[600] text-indigo-900'
          >
            Confirm Password
          </label>
          <input
            type='password'
            id='confirmPassword'
            {...register('confirmPassword')}
            className='w-full font-[600] font-epilogue outline-none rounded-lg p-2 text-indigo-950 text-sm border border-slate-400'
          />
          {errors.confirmPassword && (
            <p className='text-red-500 text-xs mt-1 font-poppins font-[550] md:max-w-[400px]'>
              {errors.confirmPassword.message}
            </p>
          )}
        </div>

        <button
          type='submit'
          className='w-full bg-indigo-900 text-white rounded-3xl py-2 font-epilogue font-[700] mt-5 hover:bg-indigo-800 transition-all duration-500'
        >
          Continue
        </button>
        <p className='text-sm font-epilogue font-medium text-slate-400 mt-2 mx-2'>
          {`Already have an account?`}{' '}
          <Link href='/api/auth/signin'>
            <span className='text-indigo-800 font-[700] ml-1'>Log In</span>
          </Link>
        </p>
      </form>
    </div>
  );
}
