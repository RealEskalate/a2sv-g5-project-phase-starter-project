'use client';
import React from 'react';
import { useFormContext } from 'react-hook-form';

interface StepProps {
  step: number;
}

const Step3: React.FC<StepProps> = ({ step }) => {
  const { register, formState: { errors }, watch } = useFormContext();

  const isDigitalCurrencyEnabled = watch('preference.sentOrReceiveDigitalCurrency');
  const isMerchantOrderEnabled = watch('preference.receiveMerchantOrder');
  const isAccountRecommendationsEnabled = watch('preference.accountRecommendations');
  const isTwoFactorAuthEnabled = watch('preference.twoFactorAuthentication');

  const switchBaseClasses = 'w-14 h-8 flex items-center rounded-full p-1 duration-300 ease-in-out';
  const switchCircleClasses = 'h-6 w-6 bg-white rounded-full shadow-md transform duration-300 ease-in-out';

  return (
    <div className="space-y-6 min-h-[350px] flex flex-col justify-between">
      <h2 className="text-2xl font-semibold">Step {step}: Account Setup & Preferences</h2>
      <div>
        <label className="block text-sm font-medium">Username</label>
        <input
          {...register('username', { required: 'Username is required' })}
          className="mt-1 p-2 block w-full border rounded-md"
        />
        {errors.username && <p className="text-red-500 text-sm mt-1">{String(errors.username.message)}</p>}
      </div>
      <div>
        <label className="block text-sm font-medium">Password</label>
        <input
          type="password"
          {...register('password', { required: 'Password is required' })}
          className="mt-1 p-2 block w-full border rounded-md"
        />
        {errors.password && <p className="text-red-500 text-sm mt-1">{String(errors.password.message)}</p>}
      </div>
      <div>
        <label className="block text-sm font-medium">Currency</label>
        <select
          {...register('preference.currency', { required: 'Currency is required' })}
          className="mt-1 p-2 block w-full border rounded-md"
        >
          <option value="USD">US Dollar (USD)</option>
          <option value="EUR">Euro (EUR)</option>
          <option value="ETB">Ethiopian Birr (ETB)</option>
          <option value="GBP">British Pound (GBP)</option>
        </select>
        {(errors.preference as any).currency && <p className="text-red-500 text-sm mt-1">{String((errors.preference as any).currency?.message)}</p>}
      </div>
      <div className="flex items-center justify-between">
        <label className="block text-sm font-medium">Send or Receive Digital Currency</label>
        <label className="relative inline-flex items-center cursor-pointer">
          <input 
            type="checkbox" 
            {...register('preference.sentOrReceiveDigitalCurrency')} 
            className="sr-only peer" 
          />
          <div className={`${switchBaseClasses} ${isDigitalCurrencyEnabled ? 'bg-green-500' : 'bg-gray-300'}`}>
            <div className={`${switchCircleClasses} ${isDigitalCurrencyEnabled ? 'translate-x-6' : ''}`}></div>
          </div>
        </label>
      </div>
      <div className="flex items-center justify-between">
        <label className="block text-sm font-medium">Receive Merchant Order</label>
        <label className="relative inline-flex items-center cursor-pointer">
          <input 
            type="checkbox" 
            {...register('preference.receiveMerchantOrder')} 
            className="sr-only peer" 
          />
          <div className={`${switchBaseClasses} ${isMerchantOrderEnabled ? 'bg-green-500' : 'bg-gray-300'}`}>
            <div className={`${switchCircleClasses} ${isMerchantOrderEnabled ? 'translate-x-6' : ''}`}></div>
          </div>
        </label>
      </div>
      <div className="flex items-center justify-between">
        <label className="block text-sm font-medium">Account Recommendations</label>
        <label className="relative inline-flex items-center cursor-pointer">
          <input 
            type="checkbox" 
            {...register('preference.accountRecommendations')} 
            className="sr-only peer" 
          />
          <div className={`${switchBaseClasses} ${isAccountRecommendationsEnabled ? 'bg-green-500' : 'bg-gray-300'}`}>
            <div className={`${switchCircleClasses} ${isAccountRecommendationsEnabled ? 'translate-x-6' : ''}`}></div>
          </div>
        </label>
      </div>
      <div className="flex items-center justify-between">
        <label className="block text-sm font-medium">Two-Factor Authentication</label>
        <label className="relative inline-flex items-center cursor-pointer">
          <input 
            type="checkbox" 
            {...register('preference.twoFactorAuthentication')} 
            className="sr-only peer" 
          />
          <div className={`${switchBaseClasses} ${isTwoFactorAuthEnabled ? 'bg-green-500' : 'bg-gray-300'}`}>
            <div className={`${switchCircleClasses} ${isTwoFactorAuthEnabled ? 'translate-x-6' : ''}`}></div>
          </div>
        </label>
      </div>
    </div>
  );
};

export default Step3;
