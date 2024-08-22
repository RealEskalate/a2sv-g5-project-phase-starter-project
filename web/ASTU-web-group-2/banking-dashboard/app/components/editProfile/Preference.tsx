'use client';
import React from 'react';
import { useForm, Controller } from 'react-hook-form';
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import ToggleSwitch from './ToggleSwitch'; // Import the ToggleSwitch component
import { useSelector, useDispatch } from 'react-redux';
import { RootState } from '@/lib/store';
import { setUser } from '@/lib/features/userSlice/userSlice'; // Adjust the import path as needed
import { useSession } from 'next-auth/react';
import { useUpdatePreferenceMutation } from '@/lib/service/UserService';
import notify from '@/utils/notify';
// Define the Yup schema for validation
const schema = Yup.object().shape({
  currency: Yup.string().required('Currency is required'),
  timeZone: Yup.string().required('TimeZone is required'),
  sentOrReceiveDigitalCurrency: Yup.boolean().required(),
  receiveMerchantOrder: Yup.boolean().required(),
  accountRecommendations: Yup.boolean().required(),
});

interface FormValues {
  currency: string;
  timeZone: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
}

const YourFormComponent = () => {
  const dispatch = useDispatch();

  // Get the user preferences from the Redux store
  const user = useSelector((state: RootState) => state.user.user);

  // Set up form with default values from Redux store
  const { control, handleSubmit, register, formState: { errors } } = useForm<FormValues>({
    resolver: yupResolver(schema),
    defaultValues: {
      currency: user?.preference?.currency || '',
      timeZone: user?.preference?.timeZone || '',
      sentOrReceiveDigitalCurrency: user?.preference?.sentOrReceiveDigitalCurrency || false,
      receiveMerchantOrder: user?.preference?.receiveMerchantOrder || false,
      accountRecommendations: user?.preference?.accountRecommendations || false,
     
    },
  });
  const [updatePreference,{isLoading}] = useUpdatePreferenceMutation() // Use the mutation hook
  const { data: session, status } = useSession();

  // Handle form submission
  const onSubmit = async(data: FormValues) => {
    if (user && user.id) {
      const updatedUser = {
        ...user,
        preference: {
          ...data,
          twoFactorAuthentication: user.preference?.twoFactorAuthentication ?? true,
          accountRecommendations:false
        },
      };
      
      try {
        if (session?.user?.accessToken) {
          console.log("Updated User:",  session?.user?.accessToken);
          const response = await updatePreference({
            accessToken: session?.user?.accessToken,
            updatedPreference: updatedUser.preference,
          }).unwrap();
          console.log("Updated User:", response);
          dispatch(setUser(updatedUser));
          notify.success("Profile updated successfully");
          // Show success message or handle successful update
        } else {
          notify.error("Access token is missing");
          
          throw new Error("Access token is missing");
  
        }
      } catch (error) {
        console.error("Failed to update user:", error);
        // Show error message or handle error
      }
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='space-y-4'>
      <div className='grid grid-cols-1 md:grid-cols-2 gap-4'>
      <div className='w-full'>
      <label className='block mb-1 font-400 text-[16px] text-[#232323] capitalize'>
      Currency
          </label>
          <input
            placeholder={`USD`}
            className='w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200'
            {...register('currency')}
          />
          {errors.currency && <p className='text-red-500'>{errors.currency.message}</p>}
        </div>
        <div className='w-full'>
          <label className='block mb-1 font-400 text-[16px] text-[#232323] capitalize'>
            TimeZone
          </label>
          <input
            placeholder={`{GMT-12:00} International Data Line West`}
            className='w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200'
            {...register('timeZone')}
          />
          {errors.timeZone && <p className='text-red-500'>{errors.timeZone.message}</p>}
        </div>
      </div>

      <div className='flex flex-col gap-[0.4rem]'>
        <label className=' mb-1 font-400 text-[16px] text-[#232323] capitalize'>
          Notification
        </label>
        <Controller
          name='sentOrReceiveDigitalCurrency'
          control={control}
          render={({ field }) => (
            <ToggleSwitch
              id='sentOrReceiveDigitalCurrency'
              label='I send or receive digital currency'
              checked={field.value}
              onChange={field.onChange}
            />
          )}
        />
        <Controller
          name='receiveMerchantOrder'
          control={control}
          render={({ field }) => (
            <ToggleSwitch
              id='receiveMerchantOrder'
              label='I receive merchant order'
              checked={field.value}
              onChange={field.onChange}
            />
          )}
        />
        <Controller
          name='accountRecommendations'
          control={control}
          render={({ field }) => (
            <ToggleSwitch
              id='accountRecommendations'
              label='There are recommendations for my account'
              checked={field.value}
              onChange={field.onChange}
            />
          )}
        />
      </div>
      <div className='flex justify-end'>
        <button
          type='submit'
          className='w-full md:w-auto px-4 py-2 bg-[#1814F3] text-white rounded-lg'
          disabled = {isLoading}
        >
          Save Changes
        </button>
      </div>
    </form>
  );
};

export default YourFormComponent;
