"use client";
import { SubmitHandler, useForm } from 'react-hook-form';
import ToggleButton from './Toggle';
import { useEffect } from 'react';
import { useSession } from 'next-auth/react';
import { useGetCurrentUserQuery, useUpdateUserPreferenceMutation } from '@/lib/redux/api/bankApi';
import Alert from './Alert';
import { CustomSerializedError } from '@/lib/redux/types/CustomSerializedError';

interface FormData {
  currency: string;
  timeZone: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  twoFactorAuthentication: boolean;
}

const Preferences = () => {
  const { register, handleSubmit, formState: { errors }, reset, setValue, getValues } = useForm<FormData>({
    defaultValues: {
      currency: '',
      timeZone: '',
      sentOrReceiveDigitalCurrency: true,
      receiveMerchantOrder: false,
      accountRecommendations: true,
    }
  });

  const session = useSession();
  const access_token = session.data?.access_token;

  const { isLoading, isError, error, data } = useGetCurrentUserQuery(access_token as string);
  const [updateUserPreference, {isLoading: isUpdateLoading, isSuccess: isUpdateSuccess, isError: isUpdateError, error: updateError, data: updateData}] = useUpdateUserPreferenceMutation()


  const pref = data?.data.preference;
  const errorUpdate = updateError as CustomSerializedError

  useEffect(() => {
    if (pref) {
      reset({
        currency: pref.currency || '',
        timeZone: pref.timeZone || '',
        sentOrReceiveDigitalCurrency: pref.sentOrReceiveDigitalCurrency || true,
        receiveMerchantOrder: pref.receiveMerchantOrder || false,
        accountRecommendations: pref.accountRecommendations || true,
        twoFactorAuthentication: pref.twoFactorAuthentication || true
      });
    }
  }, [pref, reset]);

  const onSubmit: SubmitHandler<FormData> = (data) => {
    console.log(data);
    updateUserPreference({userUpdate: data, token: access_token as string})
    console.log(updateError, 'errr ')

  };

  const handleToggleChange = (field: keyof FormData) => {
    setValue(field, !getValues(field), { shouldDirty: true });
  };

  return (
    <div>
      <form onSubmit={handleSubmit(onSubmit)} className='p-4'>
        <div className='flex sm:flex-nowrap flex-wrap sm:gap-5 '>
          <div className='flex-grow'>
            <label htmlFor="currency" className="text-custom-light-dark mb-4 text-base font-normal">Currency</label>
            <input
              id="currency"
              type="text"
              {...register('currency')}
              className="w-full p-2 border border-custom-light-grey rounded-xl mt-2 text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
            {errors.currency && <p className="text-red-500">{errors.currency.message}</p>}
          </div>
          <div className='flex-grow'>
            <label htmlFor="timezone" className="text-custom-light-dark mb-4 text-base font-normal">Time Zone</label>
            <input
              id="timeZone"
              type="text"
              {...register('timeZone')}
              className="w-full p-2 border border-custom-light-grey rounded-xl mt-2 text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
            {errors.timeZone && <p className="text-red-500">{errors.timeZone.message}</p>}
          </div>
        </div>

        <div className='flex flex-col gap-4'>
          <h2 className='font-medium my-3 text-base text-[#333B69]'>Notification</h2>
          <div className='flex gap-2 items-center'>
            <ToggleButton
              enabled={getValues('sentOrReceiveDigitalCurrency')}
              onEnable={() => handleToggleChange('sentOrReceiveDigitalCurrency')}
            />
            <p className='text-sm font-normal'>I send or receive digital currency</p>
          </div>
          <div className='flex gap-2 items-center'>
            <ToggleButton
              enabled={getValues('receiveMerchantOrder')}
              onEnable={() => handleToggleChange('receiveMerchantOrder')}
            />
            <p className='text-sm font-normal'>I receive merchant order</p>
          </div>
          <div className='flex gap-2 items-center'>
            <ToggleButton
              enabled={getValues('accountRecommendations')}
              onEnable={() => handleToggleChange('accountRecommendations')}
            />
            <p className='text-sm font-normal'>There are recommendations for my account</p>
          </div>
        </div>

        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-custom-bright-purple w-full sm:w-1/4 text-white px-4 py-2 hover:shadow-md font-body font-medium text-md rounded-xl mt-6"
          >
            Save
          </button>
        </div>
      </form>
      {isUpdateError && <Alert type="error" message={errorUpdate.data.message} duration={2000} />}
      {isUpdateSuccess && <Alert type="success" message="Successfully updated! Refresh to see the changes." duration={2000} />}

    </div>
  );
}

export default Preferences;
