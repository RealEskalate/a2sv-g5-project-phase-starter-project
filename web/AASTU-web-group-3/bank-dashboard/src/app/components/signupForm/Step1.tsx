'use client';
import React, { useRef, useState } from 'react';
import { useFormContext } from 'react-hook-form';
import Image from 'next/image';
import Link from 'next/link';
import {
  getStorage,
  ref,
  uploadBytesResumable,
  getDownloadURL,
} from 'firebase/storage';
import app from '@/app/firebase';
import { logo } from '@/../../public/Icons';

interface StepProps {
  step: number;
}

const Step1: React.FC<StepProps> = ({ step }) => {
  const { register, formState: { errors }, setValue } = useFormContext();
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [uploadProgress, setUploadProgress] = useState<number | null>(null);

  const handleImageClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      const imageUrl = await uploadImageToCloud(file);
      setValue("profilePicture", imageUrl);
    } else {
      setValue("profilePicture", "");
    }
  };

  const uploadImageToCloud = async (file: File): Promise<string> => {
    const storage = getStorage(app);
    const storageRef = ref(storage, file.name);
    const uploadTask = uploadBytesResumable(storageRef, file);

    return new Promise((resolve, reject) => {
      uploadTask.on(
        'state_changed',
        (snapshot) => {
          const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
          setUploadProgress(progress); // Set the progress state
        },
        (error) => {
          console.error('Error during upload:', error);
          reject(error);
        },
        async () => {
          try {
            const url = await getDownloadURL(uploadTask.snapshot.ref);
            console.log('File available at', url);
            setUploadProgress(null); // Reset progress after completion
            resolve(url);
          } catch (err) {
            console.error('Error getting download URL:', err);
            reject(err);
          }
        }
      );
    });
  };

  const validateAge = (value: string) => {
    const birthDate = new Date(value);
    const today = new Date();
    let age = today.getFullYear() - birthDate.getFullYear();
    const monthDiff = today.getMonth() - birthDate.getMonth();

    if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
      age--;
    }

    return age >= 16 || 'You must be at least 16 years old';
  };

  return (
    <>
      {/* <div className='flex justify-center gap-2 mb-2'>
        <Image src={logo} width={36} height={36} alt='bankdash-logo' />
        <h1 className='font-extrabold text-2xl text-[#343C6A]'>BankDash</h1>
      </div> */}
      <div className="space-y-4 min-h-[350px] flex flex-col justify-between">
        <h2 className="text-2xl font-semibold">Step {step}: Basic Information</h2>
        <div>
          <label className="block text-sm font-medium">Name <span className="text-red-500">*</span></label>
          <input
            {...register('name', { required: 'Name is required' })}
            className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"

          />
          {errors.name && <p className="text-red-500 text-sm mt-1">{String(errors.name?.message)}</p>}
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
            className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"

          />
          {errors.email && <p className="text-red-500 text-sm mt-1">{String(errors.email?.message)}</p>}
        </div>
        <div>
          <label className="block text-sm font-medium">Date of Birth <span className="text-red-500">*</span></label>
          <input
            type="date"
            {...register('dateOfBirth', {
              required: 'Date of Birth is required',
              validate: validateAge
            })}
            className="border rounded-md w-full h-12 border-[#CCCCF5] dark:border-darkComponent p-2 bg-white dark:bg-darkPage dark:text-darkText"

          />
          {errors.dateOfBirth && <p className="text-red-500 text-sm mt-1">{String(errors.dateOfBirth?.message)}</p>}
        </div>
        <div>
          <label className="block text-sm font-medium">Profile Picture <span className="text-red-500">*</span></label>
          <input
            type="file"
            ref={fileInputRef}
            className="mt-1 block w-full text-sm text-gray-500 file:mr-4 file:py-2 file:px-4 file:rounded-full file:border-0 file:text-sm file:font-semibold file:bg-violet-50 file:text-violet-700 hover:file:bg-violet-100"
            onChange={handleFileChange}
          />
          {uploadProgress !== null && (
            <div className="w-full bg-gray-200 rounded-full h-2.5 mt-2">
              <div
                className="bg-blue-600 h-2.5 rounded-full"
                style={{ width: `${uploadProgress}%` }}
              ></div>
              <p className="text-sm mt-1">{Math.round(uploadProgress)}% uploaded</p>
            </div>
          )}
          {errors.profilePictureFile && <p className="text-red-500 text-sm mt-1">{String(errors.profilePictureFile?.message)}</p>}
        </div>
      </div>
      <div className="text-center mt-4 flex justify-end my-4">
        <p className="text-base">
          Already have an account?{' '}
          <Link href="/auth/signin" className="text-[#4640DE] hover:underline">Sign in here</Link>
        </p>
      </div>
    </>
  );
};

export default Step1;
