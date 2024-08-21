import React from 'react';
import { useForm } from 'react-hook-form';
import Image from 'next/image';
import Input from '@/components/ui/Input';  // Ensure this Input component is set up properly
import { updateUserDetails } from '@/services/userupdate';


interface EditProfileFormData {
  "name": "string",
  "email": "string",
  "dateOfBirth": "2024-08-20T10:36:26.481Z",
  "permanentAddress": "string",
  "postalCode": "string",
  "username": "string",
  "presentAddress": "string",
  "city": "string",
  "country": "string",
  "profilePicture": "string"
}

const EditProfileForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const onSubmit = async (data: any) => {
    console.log('Form Data:', data);
    try {
      const response = await updateUserDetails(data);
      console.log('Update User Details Response:', response);
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="space-y-4 md:grid md:grid-cols-3 md:gap-6">
      <div className="flex justify-center md:justify-start md:col-span-1">
        <div className="relative">
          <Image
            src="/Images/profilepic.jpeg"
            alt="Profile"
            width={150}
            height={150}
            className="rounded-full aspect-square object-cover"
          />
        </div>
      </div>

      <div className="md:col-span-1 space-y-4">
        <div className="w-full max-w-xs mx-auto">
          <input className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
placeholder="John Doe" {...register('name', { required: true })} />
          {errors.name && <p className="text-red-500">Name is required</p>}
        </div>
        <div className="w-full max-w-xs mx-auto">
          <input className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
   type="email" placeholder="john@example.com" {...register('email', { required: true })} />
          {errors.email && <p className="text-red-500">Email is required</p>}
        </div>
        <div className="w-full max-w-xs mx-auto">
          <input      className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
 type="date" placeholder="YYYY-MM-DD" {...register('dateOfBirth')} />
        </div>    ``
        <div className="w-full max-w-xs mx-auto">
          <input      className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
 placeholder="123 Main St" {...register('permanentAddress')} />
        </div>
        <div className="w-full max-w-xs mx-auto">
          <input      className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
 placeholder="12345" {...register('postalCode')} />
        </div>
      </div>

      <div className="md:col-span-1 space-y-4">
        <div className="w-full max-w-xs mx-auto">
          <input      className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
 placeholder="john_doe" {...register('username', { required: true })} />
          {errors.username && <p className="text-red-500">Username is required</p>}
        </div>
        <div className="w-full max-w-xs mx-auto">
          <input      className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
  placeholder="456 Another St" {...register('presentAddress')} />
        </div>
        <div className="w-full max-w-xs mx-auto">
          <input      className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
 placeholder="Cityname" {...register('city')} />
        </div>
        <div className="w-full max-w-xs mx-auto">
          <input className="mt-1 p-2 border border-gray-300 rounded-xl focus:outline-none focus:border-blue-800"
 placeholder="Countryname" {...register('country')} />
        </div>
        <div className='md:pt-5'>
          <button type="submit" className="w-full max-w-xs mx-auto bg-blue-800 text-white py-2 rounded-md">
            Save
          </button>
        </div>
      </div>
    </form>
  );
};

export default EditProfileForm;
