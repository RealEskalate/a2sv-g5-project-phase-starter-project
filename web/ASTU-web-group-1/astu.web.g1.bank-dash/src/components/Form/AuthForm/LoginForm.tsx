'use client';
import React, { useState } from 'react';
import InputGroup from '../InputGroup';

import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { signIn, useSession } from 'next-auth/react';
import Link from 'next/link';
import { useRouter } from 'next/navigation';
import { toastError, toastSuccess } from '@/components/Toastify/Toastify';
import { Loader } from 'lucide-react';

const LoginFormSchema = z.object({
  username: z
    .string({
      required_error: 'Username is required',
      message: 'Username is required',
    })
    .min(4, {
      message: 'Username is too short',
    }),
  password: z.string().min(6),
});

type FormData = z.infer<typeof LoginFormSchema>;

const LoginForm = () => {
  const [isLoading, setLoading] = useState(false);
  const session = useSession();
  const router = useRouter();

  if (session.data) {
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
    // console.log("data is", data);
    setLoading(true);
    try {
      const result = await signIn('credentials', {
        redirect: false,
        username: data.username,
        password: data.password,
      });
      // console.log('result from signIn', result);
      if (!result?.ok) {
        setLoading(false);
        toastError('invalid credentials');
        // throw new Error("invalid credentials");
      }

      if (result?.ok) {
        setLoading(false);
        toastSuccess('Login Successful');
        // console.log('redirecting to ', result?.url);
        const parsedUrl = new URL(result?.url || '/');
        const callbackUrl = parsedUrl.searchParams.get('callbackUrl');
        // console.log('callbackUrl is ', callbackUrl);
        router.push(callbackUrl || '/bank-dash');
      }
    } catch (error) {
      toastError('invalid credentials');
    }
  };

  return (
    <form
      className='flex flex-col items-center w-full lg:w-3/4 justify-center py-6 p-4 lg:p-6 rounded-2xl bg-white'
      onSubmit={handleSubmit(onSubmit)}
    >
      <p className='text-[#333B69] pb-3 text-28px text-left font-semibold w-full'>Login</p>
      <div className='w-full flex flex-col'>
        <InputGroup
          id='username'
          label='Username'
          inputType='text'
          registerName='username'
          register={register}
          placeholder='Enter Username'
          errorMessage={errors?.username?.message as string}
        />
        <InputGroup
          id='password'
          label='Password'
          inputType='password'
          registerName='password'
          register={register}
          placeholder='Enter password'
          errorMessage={errors?.password?.message as string}
        />
      </div>

      <button
        type='submit'
        className='bg-[#1814f3] text-white text-center px-10 py-3 font-Lato font-bold rounded-lg w-full mt-4'
        id='login-btn'
      >
        {isLoading ? <Loader className='animate-spin w-full' /> : 'Login'}
      </button>
      <div className='w-full mt-5'>
        {`Already have an account?`}
        <Link href='/api/auth/signup'>
          <span className='text-indigo-800 font-[700] ml-1'>SignUp</span>
        </Link>
      </div>
    </form>
  );
};

export default LoginForm;
