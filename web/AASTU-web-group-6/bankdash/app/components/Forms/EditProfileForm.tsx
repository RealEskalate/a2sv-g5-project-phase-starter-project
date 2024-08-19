import React from 'react';
import Image from '@/node_modules/next/image';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPencilAlt } from '@fortawesome/free-solid-svg-icons';

const EditProfileForm = () => {
  return (
    <div className="content-center max-h-[1100px] w-full mt-2 flex justify-between flex-wrap px-4 bg-white rounded-lg shadow-md">
      <div className="flex justify-center w-full sm:w-auto md:w-auto">
        <div className="relative h-fit flex">
          <Image
            src="/assets/profile-1.png"
            width={128}
            height={128}
            alt="Profile"
            className="rounded-full mr-4"
          />
          <button className="absolute bottom-3 right-1 px-2 py-1 bg-[#1814F3] text-white rounded-full">
            <FontAwesomeIcon icon={faPencilAlt} className="text-lg" />
          </button>
        </div>
      </div>

      <form className="w-full max-w-[848px] mt-8 space-y-6 flex flex-col">
        <div className="flex flex-col gap-6">
          <div className='flex flex-wrap gap-x-6'>
            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Your Name</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>

            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">User Name</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>
          </div>

          <div className='flex flex-wrap gap-x-6'>
            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Email</label>
              <input
                type="email"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>

            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Password</label>
              <input
                type="password"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>
          </div>

          <div className='flex flex-wrap gap-x-6'>
            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Date of Birth</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>

            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Present Address</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>
          </div>

          <div className='flex flex-wrap gap-x-6'>
            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Permanent Address</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>

            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">City</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>
          </div>

          <div className='flex flex-wrap gap-x-6'>
            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Postal Code</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>

            <div className='w-full max-w-[318px] md:w-[265px]'>
              <label className="block text-sm font-medium text-gray-700">Country</label>
              <input
                type="text"
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm"
              />
            </div>
          </div>
        </div>

        <button
          type="submit"
          className="w-full py-2 px-4 border border-transparent rounded-md shadow-sm text-sm font-medium text-white bg-blue-600 hover:bg-blue-700"
        >
          Save
        </button>
      </form>
    </div>
  )
}

export default EditProfileForm;
