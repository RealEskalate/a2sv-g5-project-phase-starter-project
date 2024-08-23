'use client';
import { useState } from 'react';
import { useRouter } from 'next/navigation';
import img from '../../../public/images/image.png';
import icon from '../../../public/images/icon.svg';
import { dummyUserProfile } from './duumydata';

const UserProfile = () => {
  const router = useRouter();

  const [formData, setFormData] = useState(dummyUserProfile);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    // Handle form submission logic here
  };

  return (
    <div className='flex flex-col md:flex-row gap-10 mt-10'>
      <div className='flex justify-center items-center h-full'>
        <div className='relative w-[130px] h-[130px]'>
          <img src={icon.src} alt='icon' className='absolute top-20 right-0 w-8 h-8' />
          <img src={img.src} alt='' className='w-[130px] h-[130px] rounded-full' />
        </div>
      </div>
      <div>
        <form onSubmit={handleSubmit}>
          <div className='flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4'>
            <div className='flex flex-col'>
              <label htmlFor="name">Name</label>
              <input 
                type="text" 
                id="name" 
                name="name" 
                value={formData.name} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
            <div className='flex flex-col'>
              <label htmlFor="username">Username</label>
              <input 
                type="text" 
                id="username" 
                name="username" 
                value={formData.username} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
          </div>

          <div className='flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4 mt-4'>
            <div className='flex flex-col'>
              <label htmlFor="email">Email</label>
              <input 
                type="email" 
                id="email" 
                name="email" 
                value={formData.email} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
            <div className='flex flex-col'>
              <label htmlFor="password">Password</label>
              <input 
                type="text" 
                id="password" 
                name="password" 
                value={formData.password} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
          </div>

          <div className='flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4 mt-4'>
            <div className='flex flex-col mt-4'>
              <label htmlFor="dob">Date of Birth</label>
              <input 
                type="date" 
                id="dob" 
                name="dob" 
                value={formData.dob} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
            <div className='flex flex-col mt-4'>
              <label htmlFor="presentAddress">Present Address</label>
              <input 
                type="text" 
                id="presentAddress" 
                name="presentAddress" 
                value={formData.presentAddress} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
          </div>

          <div className='flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4 mt-4'>
            <div className='flex flex-col mt-4'>
              <label htmlFor="permanentAddress">Permanent Address</label>
              <input 
                type="text" 
                id="permanentAddress" 
                name="permanentAddress" 
                value={formData.permanentAddress} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
            <div className='flex flex-col mt-4'>
              <label htmlFor="city">City</label>
              <input 
                type="text" 
                id="city" 
                name="city" 
                value={formData.city} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
          </div>

          <div className='flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4 mt-4'>
            <div className='flex flex-col mt-4'>
              <label htmlFor="postalCode">Postal Code</label>
              <input 
                type="text" 
                id="postalCode" 
                name="postalCode" 
                value={formData.postalCode} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
            <div className='flex flex-col mt-4'>
              <label htmlFor="country">Country</label>
              <input 
                type="text" 
                id="country" 
                name="country" 
                value={formData.country} 
                onChange={handleChange} 
                className='border border-1 rounded-[15px] p-[10px] w-full md:w-[418px] text-[#718EBF] text-[15px]' />
            </div>
          </div>
          <button type="submit" className='mt-4 bg-[#1814F3] text-white p-2 rounded-[15px] w-full md:w-[150px] h-[60px] md:ml-[700px]'>Save</button>
        </form>
      </div>
    </div>
  );
};

export default UserProfile;