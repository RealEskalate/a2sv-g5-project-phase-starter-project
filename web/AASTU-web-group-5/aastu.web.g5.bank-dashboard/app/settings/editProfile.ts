'use client'

import React, { useState } from 'react';
import Image from 'next/image';

const EditProfile = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [dateOfBirth, setDateOfBirth] = useState('');
  const [permanentAddress, setPermanentAddress] = useState('');
  const [postalCode, setPostalCode] = useState('');
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [presentAddress, setPresentAddress] = useState('');
  const [city, setCity] = useState('');
  const [country, setCountry] = useState('');

  const [message, setMessage] = useState('');
  const [isSuccess, setIsSuccess] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    console.log('Form submitted',(dateOfBirth.toString()));
    const formattedDateOfBirth = new Date(dateOfBirth).toISOString();


    const token = "eyJhbGciOiJIUzM4NCJ9.eyJzdWIiOiJmZXZlbiIsImlhdCI6MTcyNDA1ODMzMywiZXhwIjoxNzI0MTQ0NzMzfQ.Xu_fc3v5wg1bTTtrjmqCoZRUWPzujw8OOrtn9YHSJg8CheLg197moMs5zMwgie-w"; // Replace with your actual token retrieval method

    try {
      console.log(111)
      const response = await fetch('https://bank-dashboard-6acc.onrender.com/user/update', {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`, // Include the token in the headers
        },
        body: JSON.stringify({
          name,
          email,
          dateOfBirth: formattedDateOfBirth,
          permanentAddress,
          postalCode,
          username,
          password,
          presentAddress,
          city,
          country
        }),
      });
      console.log(response)

      if (response.ok) {
        setMessage('Registered successfully!'); // Success message
        setIsSuccess(true);
      } else {
        const errorData = await response.json();
        setMessage(`Error: ${errorData.message}`); // Error message from server
        setIsSuccess(false);
      }
    } catch (error) {
      console.log(error)
      setMessage('An unexpected error occurred. Please try again later.'); // Generic error message
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
              {/* Name */}
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


              {/* Email */}
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

              {/* Date of Birth */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Date of Birth</p>
                <input 
                  type='date' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='Select your date of birth'
                  value={dateOfBirth}
                  onChange={(e) => setDateOfBirth(e.target.value)}
                />
              </div>

              {/* Permanent Address */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Permanent Address</p>
                <input 
                  type="text" 
                  className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="San Jose, California, USA"
                  value={permanentAddress}
                  onChange={(e) => setPermanentAddress(e.target.value)}
                />
              </div>

              {/* Postal Code */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Postal Code</p>
                <input 
                  type="string" 
                  className="border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="45962"
                  value={postalCode}
                  onChange={(e) => setPostalCode(e.target.value)}
                />
              </div>
            </div>

            {/* Right Column */}
            <div className='w-full md:w-[50%] space-y-6 mt-8 md:mt-0'>
              {/* Username */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>User Name</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='Charlene Reed'
                  value={username}
                  onChange={(e) => setUsername(e.target.value)}
                />
              </div>

              {/* Password */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Password</p>
                <input 
                  type='password' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='********'
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                />
              </div>

              {/* Present Address */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Present Address</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='San Jose, California, USA'
                  value={presentAddress}
                  onChange={(e) => setPresentAddress(e.target.value)}
                />
              </div>


              {/* City */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>City</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='San Jose'
                  value={city}
                  onChange={(e) => setCity(e.target.value)}
                />
              </div>

              {/* Country */}
              <div className='bg-white'>
                <p className='text-black font-sans text-lg mb-2'>Country</p>
                <input 
                  type='text' 
                  className='border border-[#DFEAF2] p-2 w-full rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500'
                  placeholder='USA'
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
