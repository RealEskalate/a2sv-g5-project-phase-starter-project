import { useForm, Controller } from 'react-hook-form';
import * as Yup from 'yup';
import { yupResolver } from '@hookform/resolvers/yup';
import ToggleSwitch from './ToggleSwitch'; // Import the ToggleSwitch component

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

  const { control, handleSubmit, register, formState: { errors } } = useForm<FormValues>({
    resolver: yupResolver(schema),
  });

  // Handle form submission
  const onSubmit = (data: FormValues) => {
    const formattedData = {
      preference: {
        currency: data.currency,
        sentOrReceiveDigitalCurrency: data.sentOrReceiveDigitalCurrency,
        receiveMerchantOrder: data.receiveMerchantOrder,
        accountRecommendations: data.accountRecommendations,
        timeZone: data.timeZone,
        twoFactorAuthentication: true, // You can set this based on your needs
      },
    };
    console.log(formattedData);
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
        Submit
      </button>
    </form>
  );
};

export default YourFormComponent;
