'use client';
import React from 'react';
import { useFormContext } from 'react-hook-form';
import Image from 'next/image';
import { logo } from "@/../../public/Icons";
import Link from 'next/link';

interface StepProps {
  step: number;
}

const Step1: React.FC<StepProps> = ({ step }) => {
  const { register, formState: { errors }, watch } = useFormContext();

  // Calculate minimum age validation
  const validateAge = (value: string) => {
    const birthDate = new Date(value);
    const today = new Date();
    let age = today.getFullYear() - birthDate.getFullYear();
    const monthDiff = today.getMonth() - birthDate.getMonth();

    if (
      monthDiff < 0 ||
      (monthDiff === 0 && today.getDate() < birthDate.getDate())
    ) {
      age--;
    }

    return age >= 16 || 'You must be at least 16 years old';
  };

  return (
    <>
      <div className='flex justify-center gap-2 mb-2'>
        <Image src={logo} width={36} height={36} alt='bankdash-logo' />
        <h1 className='font-extrabold text-2xl text-[#343C6A]'>BankDash</h1>
      </div>
      <div className="space-y-4 min-h-[350px] flex flex-col justify-between">
        <h2 className="text-2xl font-semibold">Step {step}: Basic Information</h2>
        <div>
          <label className="block text-sm font-medium">Name <span className="text-red-500">*</span></label>
          <input
            {...register('name', { required: 'Name is required' })}
            className="mt-1 p-2 block w-full border rounded-md"
          />
          {errors.name && <p className="text-red-500 text-sm mt-1">{String(errors.name.message)}</p>}
        </div>
        <div>
          <label className="block text-sm font-medium">Email <span className="text-red-500">*</span></label>
          <input
            {...register('email', {
              required: 'Email is required',
              pattern: {
                value: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
                message: 'Invalid email address'
              }
            })}
            className="mt-1 p-2 block w-full border rounded-md"
          />
          {errors.email && <p className="text-red-500 text-sm mt-1">{String(errors.email.message)}</p>}
        </div>
        <div>
          <label className="block text-sm font-medium">Date of Birth <span className="text-red-500">*</span></label>
          <input
            type="date"
            {...register('dateOfBirth', {
              required: 'Date of Birth is required',
              validate: validateAge
            })}
            className="mt-1 p-2 block w-full border rounded-md"
          />
          {errors.dateOfBirth && <p className="text-red-500 text-sm mt-1">{String(errors.dateOfBirth.message)}</p>}
        </div>
        <div>
          <label className="block text-sm font-medium">Profile Picture <span className="text-red-500">*</span></label>
          <input
            type="file"
            {...register('profilePicture', { required: 'Profile picture is required' })}
            className="mt-1 block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-violet-50 file:text-violet-700 hover:file:bg-violet-100"
          />
          {errors.profilePicture && <p className="text-red-500 text-sm mt-1">{String(errors.profilePicture.message)}</p>}
        </div>
      </div>
      <div className="text-center mt-4 flex justify-end my-4">
        <p className="text-base">
          Already have an account?{' '}
          <Link href="/auth/signin" className="text-[#4640DE] hover:underline">Sign in here</Link>.
        </p>
      </div>
    </>
  );
};

export default Step1;
