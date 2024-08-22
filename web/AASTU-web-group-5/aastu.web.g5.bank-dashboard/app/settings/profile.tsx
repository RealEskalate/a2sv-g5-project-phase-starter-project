'use client';

import React, { useState } from 'react';
import Image from 'next/image';

import { useSession } from 'next-auth/react';
import axios from 'axios';
interface ExtendedUser {
  name?: string;
  email?: string;
  image?: string;
  accessToken?: string;
  }
const EditProfile = () => {

  const { data: session } = useSession();
  
  const user = session?.user as ExtendedUser;
  
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [dateOfBirth, setDateOfBirth] = useState('');
  const [permanentAddress, setPermanentAddress] = useState('');
  const [postalCode, setPostalCode] = useState('');
  const [username, setUsername] = useState('');
  const [presentAddress, setPresentAddress] = useState('');
  const [city, setCity] = useState('');
  const [country, setCountry] = useState('');
  const [message, setMessage] = useState('');
  const [isSuccess, setIsSuccess] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const formattedDateOfBirth = new Date(dateOfBirth).toISOString();

    try {
      const response = await axios.put(
        'https://bank-dashboard-1tst.onrender.com/user/update',
        {
          name,
          email,
          dateOfBirth: formattedDateOfBirth,
          permanentAddress,
          postalCode,
          username: user.name,
          presentAddress,
          profilePicture: "/",
          city,
          country
        },
        {
          headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${user.accessToken}`,
          },
        }
      );

      if (response.status === 200) {
        setMessage('Profile updated successfully!');
        setIsSuccess(true);
      } else {
        const errorData = response.data;
        setMessage(`Error: ${errorData.message}`);
        setIsSuccess(false);
      }
    } catch (error) {
      console.log(error, 11);
      setMessage('An unexpected error occurred. Please try again later.');
      setIsSuccess(false);
    }
  };


  return (
    <div className='bg-white p-4 md:p-8'>
      <form onSubmit={handleSubmit}>
        <div className='flex flex-col md:flex-row md:space-x-8'>  
          {/* Profile Image Section */}
          <div className='w-full md:w-[20%] flex justify-center mb-8 md:mb-0'>
            <div className='w-56 h-56 md:w-40 md:h-40 rounded-full overflow-hidden flex items-center justify-center'>
              <Image
                src="/images/christina.png"
                alt="Profile Picture"
                width={224}
                height={224}
                className="object-cover"
              />
            </div>
          </div>

          {/* Form Section */}
          <div className='w-full md:w-[80%] flex flex-col md:flex-row gap-6'>
            {/* Left Column */}
            <div className='w-full md:w-[50%] space-y-6'>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Your Name</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='Charlene Reed'
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                />
              </div>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Email</p>
                <input 
                  type='email' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='charlenereed@gmail.com'
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                />
              </div>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Date of Birth</p>
                <input 
                  type='date' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  value={dateOfBirth}
                  onChange={(e) => setDateOfBirth(e.target.value)}
                />
              </div>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Permanent Address</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='San Jose, California, USA'
                  value={permanentAddress}
                  onChange={(e) => setPermanentAddress(e.target.value)}
                />
              </div>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Postal Code</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='45962'
                  value={postalCode}
                  onChange={(e) => setPostalCode(e.target.value)}
                />
              </div>
            </div>

            {/* Right Column */}
            <div className='w-full md:w-[50%] space-y-6 mt-8 md:mt-0'>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Username</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </div>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Present Address</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  value={presentAddress}
                  onChange={(e) => setPresentAddress(e.target.value)}
                />
              </div>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>City</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  value={city}
                  onChange={(e) => setCity(e.target.value)}
                />
              </div>
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Country</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  value={country}
                  onChange={(e) => setCountry(e.target.value)}
                />
              </div>
            </div>
          </div>
        </div>

        {/* Save Button */}
        <div className="flex justify-center md:justify-end mt-8 md:mt-12">
          <button type="submit" className="bg-blue-800 text-white w-full h-12 rounded-full md:w-[12rem] text-[13px] md:text-base">
            Save
          </button>
        </div>
      </form>

      {/* Display success or error message */}
      {message && (
        <div className="mt-4 text-center">
          <p className={isSuccess ? 'text-green-600' : 'text-red-600'}>{message}</p>
        </div>
      )}
    </div>
  )
}

export default EditProfile;
