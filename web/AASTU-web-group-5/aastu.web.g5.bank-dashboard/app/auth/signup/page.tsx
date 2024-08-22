"use client";
import { useForm } from 'react-hook-form';
import axios from 'axios';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import Background from '@/public/assests/images/register_background.jpg';

interface Preference {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

interface RegistrationFormData {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string;
  preference: Preference;
}

const Register = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<RegistrationFormData>();
  const router = useRouter();
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  const onSubmit = async (data: RegistrationFormData) => {
    try {
      const response = await axios.post('https://bank-dashboard-1tst.onrender.com/auth/register', data);
      console.log(response.data); // Handle response data if needed
      router.push('/auth/signin'); // Redirect to login page after successful registration
    } catch (error: any) {
      setErrorMessage(error.response?.data?.message || 'Registration failed');
    }
  };

  return (
    <div className="flex flex-col items-center min-h-screen bg-gray-100 p-4">
      <div className='pt-2 bg-green-100 w-[90%] text-center mt-0'>
        <h1 className="text-4xl font-bold mb-6 text-purple-700">Create Your Account</h1>
      </div>
      <div className="flex justify-center w-[90%] bg-white p-8 rounded-lg shadow-md">
        <form onSubmit={handleSubmit(onSubmit)} className="grid grid-cols-1 gap-6 w-full">
          {errorMessage && <p className="text-red-500 mb-4 text-center col-span-1">{errorMessage}</p>}
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div className="flex flex-col items-center">
              <input {...register('name', { required: true })} placeholder="Name" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.name && <p className="error-text text-red-500 mt-1">Name is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('email', { required: true })} type="email" placeholder="Email" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.email && <p className="error-text text-red-500 mt-1">Email is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('dateOfBirth', { required: true })} type="date" placeholder="Date of Birth" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.dateOfBirth && <p className="error-text text-red-500 mt-1">Date of Birth is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('permanentAddress', { required: true })} placeholder="Permanent Address" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.permanentAddress && <p className="error-text text-red-500 mt-1">Permanent Address is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('postalCode', { required: true })} placeholder="Postal Code" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.postalCode && <p className="error-text text-red-500 mt-1">Postal Code is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('username', { required: true })} placeholder="Username" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.username && <p className="error-text text-red-500 mt-1">Username is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('password', { required: true })} type="password" placeholder="Password" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.password && <p className="error-text text-red-500 mt-1">Password is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('presentAddress', { required: true })} placeholder="Present Address" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.presentAddress && <p className="error-text text-red-500 mt-1">Present Address is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('city', { required: true })} placeholder="City" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.city && <p className="error-text text-red-500 mt-1">City is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('country', { required: true })} placeholder="Country" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.country && <p className="error-text text-red-500 mt-1">Country is required</p>}
            </div>

            <div className="flex flex-col items-center">
              <input {...register('profilePicture', { required: true })} type="text" 
              placeholder="Profile Picture URL" className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100" />
              {errors.profilePicture && <p className="error-text text-red-500 mt-1">Profile Picture is required</p>}
            </div>
          </div>
          <div className="flex justify-center">
            <h2 className="text-xl font-semibold text-gray-700">Preferences</h2>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="flex flex-col items-center">
            <select 
              {...register('preference.currency', { required: true })} 
              className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100"
            >
              <option value="">Select Currency</option>
              <option value="USD">USD</option>
              <option value="EUR">EUR</option>
              <option value="GBP">GBP</option>
              <option value="JPY">JPY</option>
            </select>
            {errors.preference?.currency && <p className="error-text text-red-500 mt-1">Currency is required</p>}
          </div>

          <div className="flex flex-col items-center">
            <select 
              {...register('preference.timeZone', { required: true })} 
              className="input-field w-[80%] bg-blue-50 border border-blue-300 rounded-md p-2 transition duration-300 ease-in-out focus:bg-blue-100"
            >
              <option value="">Select Time Zone</option>
              <option value="UTC-12:00">UTC-12:00</option>
              <option value="UTC-11:00">UTC-11:00</option>
              <option value="UTC-10:00">UTC-10:00</option>
              <option value="UTC-09:00">UTC-09:00</option>
              <option value="UTC-08:00">UTC-08:00</option>
              <option value="UTC-07:00">UTC-07:00</option>
              <option value="UTC-06:00">UTC-06:00</option>
              <option value="UTC-05:00">UTC-05:00</option>
              <option value="UTC-04:00">UTC-04:00</option>
              <option value="UTC-03:00">UTC-03:00</option>
              <option value="UTC-02:00">UTC-02:00</option>
              <option value="UTC-01:00">UTC-01:00</option>
              <option value="UTC+00:00">UTC+00:00</option>
              <option value="UTC+12:00">UTC+12:00</option>
            </select>
            {errors.preference?.timeZone && <p className="error-text text-red-500 mt-1">Time Zone is required</p>}
          </div>
          </div>

          <div className="flex flex-wrap justify-center space-x-4 mt-4">
            <label className="flex items-center">
              <input {...register('preference.sentOrReceiveDigitalCurrency')} type="checkbox" className="mr-2" />
              Send or Receive Digital Currency
            </label>

            <label className="flex items-center">
              <input {...register('preference.receiveMerchantOrder')} type="checkbox" className="mr-2" />
              Receive Merchant Order
            </label>

            <label className="flex items-center">
              <input {...register('preference.accountRecommendations')} type="checkbox" className="mr-2" />
              Account Recommendations
            </label>

            <label className="flex items-center">
              <input {...register('preference.twoFactorAuthentication')} type="checkbox" className="mr-2" />
              Two-Factor Authentication
            </label>
          </div>

          <div className="flex justify-center mt-6">
            <button type="submit" className="w-full md:w-1/2 py-3 bg-purple-700 text-white font-semibold rounded-full transition duration-300">
              Register
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Register;