"use client";
import React, { useEffect, useState } from 'react';
import FormComponent from './FormComponent'; // Adjust the import path as necessary
import YourFormComponent from './Preference';

interface Preference {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

interface MainData {
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

const defaultMainData: MainData = {
  name: '',
  email: '',
  dateOfBirth: new Date().toISOString(),
  permanentAddress: '',
  postalCode: '',
  username: '',
  password: '',
  presentAddress: '',
  city: '',
  country: '',
  profilePicture: '',
  preference: {
    currency: '',
    sentOrReceiveDigitalCurrency: true,
    receiveMerchantOrder: true,
    accountRecommendations: true,
    timeZone: '',
    twoFactorAuthentication: true,
  }
};

const Signup = () => {
  const [activeTab, setActiveTab] = useState<number>(0);
  const tabs = ['User Information', 'Preferences'];
  const [mainData, setMainData] = useState<MainData>(defaultMainData);

  useEffect(() => {
    console.log(mainData);
  }, [mainData]);

  return (
    <div className="flex flex-col items-center p-6">
      <h2 className="text-2xl font-bold text-blue-700 mb-6">Sign Up</h2>

      <div className="bg-white rounded-[25px] p-[1rem] grid gap-[2rem] w-full relative">
        <ul className="flex justify-between list-none p-0 relative w-[25%]">
          {tabs.map((tab, index) => (
            <li
              key={index}
              className={`cursor-pointer p-2 font-[500] text-[16px] ${activeTab === index ? 'text-blue-700 font-bold' : 'text-[#718EBF]'} relative`}
              onClick={() => setActiveTab(index)}
            >
              {tab}
              {activeTab === index && (
                <div className="absolute bottom-0 left-0 w-full h-[3px] rounded-t-[1rem] bg-[#1814F3] transition-all duration-300" ></div>
              )}
            </li>
          ))}
        </ul>
        {activeTab === 0 && (
          <FormComponent 
            setMainData={setMainData}
            mainData={mainData}
            setActiveTab={setActiveTab}
          />
        )}
        {activeTab === 1 && (
          <YourFormComponent
            setMainData={setMainData}
            mainData={mainData}
          />
        )}
      </div>
      <div className="mt-6 text-center">
        <p className="text-sm text-gray-600">
          Already have an account?{' '}
          <a href="/login" className="text-blue-700 font-semibold hover:underline">
            Log In
          </a>
        </p>
      </div>
    </div>
  );
};

export default Signup;
export type { MainData };
