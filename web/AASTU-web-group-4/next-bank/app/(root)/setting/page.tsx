'use client';
import React, { useState } from 'react';
import { useForm, Controller } from 'react-hook-form';
import { FaPencilAlt } from 'react-icons/fa'; // Importing the pencil icon from react-icons
import Input from '@/components/ui/Input';
import Toggle from '@/components/ui/Toggle';
import Image from 'next/image';

const SettingsPage = () => {
  const [activeTab, setActiveTab] = useState<'editProfile' | 'preference' | 'security'>('editProfile');
  const { control, register, handleSubmit } = useForm();

  const onSubmit = (data: any) => {
    console.log(data);
  };

  return (
    <div className="lg:grid lg:grid-cols-[250px_1fr]"> {/* CSS Grid: Sidebar and Content */}
      {/* Sidebar */}
      <aside className="hidden lg:block bg-gray-100 h-screen sticky top-0">
        {/* Your sidebar content */}
      </aside>

      {/* Main Content */}
      <div className="p-4 lg:p-10">
        {/* Tab Buttons */}
        <div className="flex flex-wrap justify-around gap-4 border-b border-gray-300">
          {['editProfile', 'preference', 'security'].map((tab) => (
            <button
              key={tab}
              onClick={() => setActiveTab(tab as 'editProfile' | 'preference' | 'security')}
              className={`pb-2 ${activeTab === tab ? 'border-b-2 border-blue-800' : ''} text-sm lg:text-base`}
            >
              {tab === 'editProfile' ? 'Edit Profile' : tab.charAt(0).toUpperCase() + tab.slice(1)}
            </button>
          ))}
        </div>

        {/* Tab Content */}
        <div className="mt-4 mx-auto max-w-4xl"> {/* Centered and max-width constraint */}
          {activeTab === 'editProfile' && (
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4 md:grid md:grid-cols-3 md:gap-6">
              {/* First Column: Profile Picture */}
              <div className="flex justify-center md:justify-start md:col-span-1">
                <div className="relative">
                  <Image
                    src="/Images/profilepic.jpeg"
                    alt="Profile"
                    width={150}
                    height={150}
                    className="rounded-full aspect-square object-cover"
                  />
                  <span className="absolute bottom-4 right-0 md:bottom-[22rem] p-2 bg-blue-800 rounded-full cursor-pointer">
                    <FaPencilAlt className="text-white" /> {/* White pencil icon */}
                  </span>
                </div>
              </div>

              {/* Second Column: Name, Email, Date of Birth, Permanent Address, Postal Code */}
              <div className="md:col-span-1 space-y-4">
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Your Name" placeholder="John Doe" {...register('name')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Email" type="email" placeholder="john@example.com" {...register('email')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Date of Birth" type="date" placeholder="YYYY-MM-DD" {...register('dob')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Permanent Address" placeholder="123 Main St" {...register('permanentAddress')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Postal Code" placeholder="12345" {...register('postalCode')} />
                </div>
              </div>

              {/* Third Column: Username, Password, Present Address, City, Country, and Save Button */}
              <div className="md:col-span-1 space-y-4">
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Username" placeholder="john_doe" {...register('username')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Password" type="password" placeholder="******" {...register('password')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Present Address" placeholder="456 Another St" {...register('presentAddress')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="City" placeholder="Cityname" {...register('city')} />
                </div>
                <div className="w-full max-w-xs mx-auto">
                  <Input label="Country" placeholder="Countryname" {...register('country')} />
                </div>

                <div className='md:pt-5'>
                  <button type="submit" className="w-full max-w-xs mx-auto bg-blue-800 text-white py-2 rounded-md">
                    Save
                  </button>
                </div>
              </div>
            </form>
          )}

          {activeTab === 'preference' && (
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
              {/* Wrapper Grid for Large Screens */}
              <div className="md:grid md:grid-cols-2 md:gap-6">
                {/* First Column: Currency and Time Zone in a Row */}
                <div className="md:col-span-2 space-y-4 md:flex md:space-y-0 md:space-x-6">
                  <div className="w-full max-w-xs">
                    <Input label="Currency" placeholder="USD" {...register('currency')} />
                  </div>
                  <div className="w-full max-w-xs">
                    <Input label="Time Zone" placeholder="GMT-5" {...register('timeZone')} />
                  </div>
                </div>

                {/* Second Column: Notifications */}
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

              {/* Save Button */}
              <div className="mt-6 flex justify-center md:pt-32">
                <button type="submit" className="w-full max-w-xs mx-auto bg-blue-800 text-white py-2 rounded-md">
                  Save
                </button>
              </div>
            </form>
          )}

          {activeTab === 'security' && (
            <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
              {/* Security Section */}
              <h3 className="font-semibold">Two-factor Authentication</h3>
              <Controller
                control={control}
                name="twoFactorAuth"
                render={({ field }) => (
                  <Toggle label="Enable or disable two-factor authentication" {...field} />
                )}
              />

              <h3 className="font-semibold">Change Password</h3>
              <div className="w-full max-w-xs">
                <Input label="Current Password" type="password" placeholder="******" {...register('currentPassword')} />
              </div>
              <div className="w-full max-w-xs">
                <Input label="New Password" type="password" placeholder="******" {...register('newPassword')} />
              </div>
              <div className="flex justify-center md:pt-20">
                <button type="submit" className="w-full max-w-xs mx-auto bg-blue-800 text-white py-2 rounded-md">
                  Save
                </button>
              </div>
            </form>
          )}
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;
