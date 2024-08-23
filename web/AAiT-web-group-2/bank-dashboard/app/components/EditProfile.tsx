"use client"
import { useForm, SubmitHandler } from 'react-hook-form';
import pencil from '../../public/pencil.svg';
import Image from 'next/image';

interface FormData {
  name: string;
  email: string;
  dob: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
}

const EditProfile = () => {
  const { register, handleSubmit, formState: { errors } } = useForm<FormData>({
    defaultValues: {
      name: '',
      email: '',
      dob: '',
      permanentAddress: '',
      postalCode: '',
      username: '',
      password: '',
      presentAddress: '',
      city: '',
      country: ''
    }
  });

  const onSubmit: SubmitHandler<FormData> = (data) => {
    console.log(data);
  };

  return (
    <div className="mx-auto bg-white p-10">
      <form onSubmit={handleSubmit(onSubmit)} className="grid grid-cols-1 lg:grid-cols-5 gap-8">
        <div className="flex flex-col space-y-1 items-center">
          <div className=" w-36 h-36 rounded-lg flex items-center justify-center relative">
            <input
              type="file"
              accept="image/*"
              className="hidden"
              id="profilePic"
            />
            <label htmlFor="profilePic" className="cursor-pointer absolute right-3 bottom-3">
              <Image src={pencil} alt="pencil" />
            </label>
            <div className="flex justify-center items-center">
              <img className="bg-yellow-50 w-36 h-36 rounded-full" src="/christina.svg" alt="profile" />
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
            <label htmlFor="dob" className="block text-custom-light-dark mb-2 text-base font-normal">Date of Birth</label>
            <input
              id="dob"
              type="date"
              {...register('dob')}
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

        <div className="space-y-2  md:col-span-2">
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
      </form>
    </div>
  );
};

export default EditProfile;

