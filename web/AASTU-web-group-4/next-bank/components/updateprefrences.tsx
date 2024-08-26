// PreferenceForm.tsx
import React, { useEffect, useState } from 'react';
import { useForm, Controller } from 'react-hook-form';
import Input from '@/components/ui/Input';
import Toggle from '@/components/ui/Toggle';
import { updatePreference, currentuser} from '@/services/userupdate';

const PreferenceForm = () => {
  const { control, register, handleSubmit, setValue } = useForm();
  const [loading, setLoading] = useState(true); // To handle loading state

  useEffect(() => {
    const fetchData = async () => {
      try {
        const Userdata = await currentuser();
        const data = Userdata.data.preference // Fetch preference data from backend
        console.log('Fetched Data:', data);

        // Prefill the form with fetched data
        setValue('currency', data.currency || '');
        setValue('timeZone', data.timeZone || '');
        setValue('digitalCurrencyNotification', data.sentOrReceiveDigitalCurrency || false);
        setValue('merchantOrderNotification', data.receiveMerchantOrder || false);
        setValue('accountRecommendationNotification', data.accountRecommendations || false);

        setLoading(false); // Set loading to false after data is fetched
      } catch (error) {
        console.error('Error fetching preferences:', error);
        setLoading(false);
      }
    };

    fetchData();
  }, [setValue]);

  const onSubmit = async (data: any) => {
    try {
      console.log('Form Data:', data);
      const response = await updatePreference(data); // Update preferences
      console.log('Update Preference Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  if (loading) return <p>Loading...</p>; // Optional: Add a loading state

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
      <div className="md:grid md:grid-cols-2 md:gap-6">
        <div className="md:col-span-2 space-y-4 md:flex md:space-y-0 md:space-x-6">
          <div className="w-full max-w-xs">
            <input
              className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
              placeholder="USD"
              {...register('currency')}
            />
          </div>
          <div className="w-full max-w-xs">
            <input
              className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
              placeholder="GMT-5"
              {...register('timeZone')}
            />
          </div>
        </div>

        <div className="md:col-span-2 space-y-4">
          <h3 className="font-semibold">Notification</h3>
          <div className="space-y-4 flex flex-col">
            <Controller
              control={control}
              name="digitalCurrencyNotification"
              render={({ field }) => (
                <Toggle label="I send or receive digital currency" {...field} />
              )}
            />
            <Controller
              control={control}
              name="merchantOrderNotification"
              render={({ field }) => (
                <Toggle label="I receive merchant order" {...field} />
              )}
            />
            <Controller
              control={control}
              name="accountRecommendationNotification"
              render={({ field }) => (
                <Toggle label="There are recommendations for my account" {...field} />
              )}
            />
          </div>
        </div>
      </div>

      <div className="mt-6 flex justify-center md:pt-32">
        <button type="submit" className="w-full max-w-xs mx-auto bg-blue-800 text-white py-2 rounded-md">
          Save
        </button>
      </div>
    </form>
  );
};

export default PreferenceForm;
