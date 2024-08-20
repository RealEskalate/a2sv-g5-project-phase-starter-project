'use client';

import { useForm } from 'react-hook-form';
import Email from './Email';
import Password from './Password';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { signIn, useSession } from 'next-auth/react';
import Link from 'next/link';
import { redirect, useRouter } from 'next/navigation';

const LoginFormSchema = z.object({
  username: z
    .string({
      required_error: 'Email is required',
      message: 'Email is required',
    })
    .min(4, {
      message: 'Email is too short',
    }),
  password: z.string().min(8),
});

type FormData = z.infer<typeof LoginFormSchema>;

function LoginForm(req: any) {
  const session = useSession();
  const router = useRouter();

  if (session.data) {
    // console.log('redirect user to posts', session);
    router.push('/bank-dash');
  }

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<FormData>({
    resolver: zodResolver(LoginFormSchema),
  });

  const onSubmit = async (data: FormData) => {
    console.log('data is', data);
    try {
      const result = await signIn('credentials', {
        redirect: false,
        username: data.username,
        password: data.password,
      });
      // console.log('result from signIn', result);
      if (!result?.ok) {
        throw new Error('invalid credentials');
      }
      if (result?.ok) {
        // console.log('redirecting to ', result?.url);
        const parsedUrl = new URL(result?.url || '/');
        const callbackUrl = parsedUrl.searchParams.get('callbackUrl');
        // console.log('callbackUrl is ', callbackUrl);
        router.push(callbackUrl || '/bank-dash');
      }
    } catch (error) {
      alert('invalid credentials');
    }
  };

  return (
    <div className=' sm:min-w-[450px] max-w-lg shadow-lg rounded-xl p-3 mb-8 pb-14 space-y-5'>
      <h1 className='text-indigo-900 text-4xl m-5 text-center font-poppins font-[800]'>
        Welcome Back
      </h1>
      <div className='w-full flex justify-between px-5 my-4 mb-5'>
        <p className='min-w-32 border border-indigo-500'></p>
        <p className='min-w-32 border border-indigo-500'></p>
      </div>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className='my-3'>
          <label
            htmlFor='Email'
            className='block mb-1 font-epilogue text-sm font-[600] text-indigo-900'
          >
            Email Address
          </label>
          <input
            type='text'
            id='email'
            {...register('username')}
            className='w-full font-[600] font-epilogue outline-none rounded-lg p-2 text-indigo-950 text-sm border border-slate-400'
          />
          <p className='text-red-500 text-xs mt-1 font-poppins font-[550] md:max-w-[400px]'>
            {errors.username?.message}
          </p>
        </div>
        <div className='my-3'>
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
          <p className='text-red-500 text-xs mt-1 font-poppins font-[550] md:max-w-[400px]'>
            {errors.password?.message}
          </p>
        </div>
        <button
          type='submit'
          id='submit'
          className='w-full bg-indigo-900 text-white rounded-3xl py-2 font-epilogue font-[700] mt-5 hover:bg-indigo-800 transition-all duration-500'
        >
          Login
        </button>
        <p className='text-sm font-epilogue font-medium text-slate-400 mt-2 mx-2'>
          {!session ? (
            <>
              {`Already have an account?`}
              <Link href='/api/auth/signin'>
                <span className='text-indigo-800 font-[700] ml-1'>Sign In</span>
              </Link>
            </>
          ) : (
            <>
              {`Don't have an account?`}
              <Link href='/api/auth/signup'>
                <span className='text-indigo-800 font-[700] ml-1'>Sign Up</span>
              </Link>
            </>
          )}
        </p>
      </form>
    </div>
  );
}

export default LoginForm;
