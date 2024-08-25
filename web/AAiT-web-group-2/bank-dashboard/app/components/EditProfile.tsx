
"use client"
import { useForm, SubmitHandler } from 'react-hook-form';
import pencil from '../../public/pencil.svg';
import Image from 'next/image';
import { useGetCurrentUserQuery, useUpdateUserProfileMutation } from '@/lib/redux/api/bankApi';
import Spinner from './Spinner';
import Error from './Error';
import { useSession } from 'next-auth/react';
import { User } from '@/types/User';
import { useEffect } from 'react';
import { CustomSerializedError } from '@/lib/redux/types/CustomSerializedError';
import Alert from './Alert';

interface FormData {
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
}

const EditProfile = () => {
  const session = useSession();
  const access_token = session.data?.access_token;

  const { isLoading, isError, error, data } = useGetCurrentUserQuery(access_token as string);
  const [updateUserProfile, {isLoading: isUpdateLoading, isSuccess: isUpdateSuccess, isError: isUpdateError, error: updateError, data: updateData}] = useUpdateUserProfileMutation()

  const { register, handleSubmit, formState: { errors }, reset, setValue } = useForm<FormData>();
  const errorUpdate = updateError as CustomSerializedError

  useEffect(() => {
    if (data?.data) {
      reset({
        name: data.data.name || '',
        email: data.data.email || '',
        dateOfBirth: new Date(data.data.dateOfBirth).toISOString().split('T')[0],
        permanentAddress: data.data.permanentAddress || '',
        postalCode: data.data.postalCode || '',
        username: data.data.username || '',
        password: '', 
        presentAddress: data.data.presentAddress || '',
        city: data.data.city || '',
        country: data.data.country || '',
        profilePicture: data.data.profilePicture || ''
      });
    }
  }, [data, reset]);

  const onSubmit: SubmitHandler<FormData> = (data) => {
    console.log(data);
    updateUserProfile({token: access_token as string, userUpdate: data})
  };
  const onFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        const base64String = reader.result as string;
        setValue("profilePicture", base64String); 
      };
      reader.readAsDataURL(file); 
    }
  };

  if (isLoading)
    return <Spinner />

  if (isError)
    return <Error message={"Error loading the profile"} />
    
  if (!data)
    return <Error message={"No profile"} />

  const name_part = data.data.name.split(' ')
  let initials = name_part[0][0].toUpperCase()
  if(name_part.length >=2){

    initials = initials.concat(name_part[1][0].toUpperCase())
}

  return (
    <div className="mx-auto bg-white p-10">
      <form onSubmit={handleSubmit(onSubmit)} className="grid grid-cols-1 lg:grid-cols-5 gap-8">
        <div className="flex flex-col space-y-1 items-center">
          <div className="w-36 h-36 rounded-lg flex items-center justify-center relative">
            <input
              type="file"
              accept="image/*"
              className="hidden"
              id="profilePicture"
              onChange={onFileChange}
            />
            <label htmlFor="profilePicture" className="cursor-pointer absolute right-3 bottom-3">
              <Image src={pencil} alt="pencil" />
            </label>
            <div className="flex justify-center items-center">
              { data.data.profilePicture &&
                <img className="bg-yellow-50 w-36 h-36 rounded-full" src={data.data.profilePicture} alt="profile" />
              }
              {data.data.name && !data.data.profilePicture &&
                <div className='w-28 h-28 text-3xl font-medium rounded-full bg-gray-300 flex justify-center items-center'>
                  <p>{initials}</p>
                </div>  
              }
            </div>
          </div>
        </div>

        <div className="space-y-2 md:col-span-2">
          <div>
            <label htmlFor="name" className="block text-custom-light-dark mb-2 text-base font-normal">Your Name</label>
            <input
              id="name"
              type="text"
              {...register('name')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
            {errors.name && <p className="text-red-500">{errors.name.message}</p>}
          </div>

          <div>
            <label htmlFor="email" className="block text-custom-light-dark mb-2 text-base font-normal">Email</label>
            <input
              id="email"
              type="email"
              {...register('email')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
            {errors.email && <p className="text-red-500">{errors.email.message}</p>}
          </div>

          <div>
            <label htmlFor="dateOfBirth" className="block text-custom-light-dark mb-2 text-base font-normal">Date of Birth</label>
            <input
              id="dateOfBirth"
              type="date"
              {...register('dateOfBirth')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>

          <div>
            <label htmlFor="permanentAddress" className="block text-custom-light-dark mb-2 text-base font-normal">Permanent Address</label>
            <input
              id="permanentAddress"
              type="text"
              {...register('permanentAddress')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>

          <div>
            <label htmlFor="postalCode" className="block text-custom-light-dark mb-2 text-base font-normal">Postal Code</label>
            <input
              id="postalCode"
              type="text"
              {...register('postalCode')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>
        </div>

        <div className="space-y-2 md:col-span-2">
          <div>
            <label htmlFor="username" className="block text-custom-light-dark mb-2 text-base font-normal">User Name</label>
            <input
              id="username"
              type="text"
              {...register('username')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>

          <div>
            <label htmlFor="password" className="block text-custom-light-dark mb-2 text-base font-normal">Password</label>
            <input
              id="password"
              type="password"
              {...register('password')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>

          <div>
            <label htmlFor="presentAddress" className="block text-custom-light-dark mb-2 text-base font-normal">Present Address</label>
            <input
              id="presentAddress"
              type="text"
              {...register('presentAddress')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>

          <div>
            <label htmlFor="city" className="block text-custom-light-dark mb-2 text-base font-normal">City</label>
            <input
              id="city"
              type="text"
              {...register('city')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>

          <div>
            <label htmlFor="country" className="block text-custom-light-dark mb-2 text-base font-normal">Country</label>
            <input
              id="country"
              type="text"
              {...register('country')}
              className="w-full p-2 border border-custom-light-grey rounded-xl text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
            />
          </div>

          <div className="flex justify-end">
            <button
              type="submit"
              className="bg-custom-bright-purple w-full sm:w-2/4 text-white px-4 py-2 hover:shadow-md font-body font-medium text-md rounded-xl mt-6"
            >
              Save
            </button>
          </div>
        </div>
        {isUpdateError && <Alert type="error" message={errorUpdate.data.message} duration={2000} />}
        {isUpdateSuccess && <Alert type="success" message="Successfully updated! Refresh to see the changes." duration={2000} />}

      </form>
    </div>
  );
};

export default EditProfile;
