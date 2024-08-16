import React from 'react';
import { useForm, Controller } from 'react-hook-form';
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import ToggleSwitch from './ToggleSwitch'; // Import the ToggleSwitch component
import { MainData } from './Signup'; // Adjust the import path as necessary
import { useSignUpMutation } from '@/lib/serice/TransactionService';
import { useRouter } from 'next/navigation'; // Import useRouter from next/router

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

interface YourFormComponentProps {
  mainData: MainData;
  setMainData: React.Dispatch<React.SetStateAction<MainData>>;
}

const YourFormComponent: React.FC<YourFormComponentProps> = ({ setMainData, mainData }) => {
  const router = useRouter(); // Initialize the useRouter hook
  const { control, handleSubmit, register, formState: { errors } } = useForm<FormValues>({
    resolver: yupResolver(schema),
    defaultValues: {
      currency: mainData.preference?.currency || '',
      timeZone: mainData.preference?.timeZone || '',
      sentOrReceiveDigitalCurrency: mainData.preference?.sentOrReceiveDigitalCurrency || false,
      receiveMerchantOrder: mainData.preference?.receiveMerchantOrder || false,
      accountRecommendations: mainData.preference?.accountRecommendations || false,
    },
  });
  const [signUp] = useSignUpMutation();
  
  const onSubmit = async (data: FormValues) => {
    const formattedData = {
      ...mainData,
      preference: {
        currency: data.currency,
        sentOrReceiveDigitalCurrency: data.sentOrReceiveDigitalCurrency,
        receiveMerchantOrder: data.receiveMerchantOrder,
        accountRecommendations: data.accountRecommendations,
        timeZone: data.timeZone,
        twoFactorAuthentication: true, // You can set this based on your needs
      },
    };

    setMainData(formattedData);
    
    const res = await signUp(formattedData); // Make sure to send formattedData
    
    if (res.data) {
      router.push('/login'); // Navigate to the login page after successful signup
    } else {
      // Handle error cases here
      console.error('Signup failed', res.error);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className='space-y-4'>
      <div className='flex max-md:grid items-center justify-center space-x-4'>
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

      <button
        type='submit'
        className='px-4 py-2 bg-blue-500 text-white rounded-lg'
      >
        Register
      </button>
    </form>
  );
};

export default YourFormComponent;
