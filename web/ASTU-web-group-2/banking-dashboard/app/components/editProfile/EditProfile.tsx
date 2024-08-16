'use client';
import React, { useState } from 'react';
import FormComponent from './FormComponent'; // Adjust the import path as necessary
import YourFormComponent from './Preference';
import ProfileSecurity from '../profileSecurity/profileSecurity';

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

// Provide a default value that matches the MainData type
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

const EditProfile = () => {
  const [activeTab, setActiveTab] = useState<number>(0);
  const tabs = ['Edit Profile', 'Preferences', 'Security'];
  const [mainData, setMainData] = useState<MainData>(defaultMainData); 

  return (
    <div className="flex flex-col items-center p-6 ">
      <div className="bg-white rounded-[25px] p-[1rem] grid gap-[2rem] w-full">
        <div className="mb-6 w-[30%] max-md:w-[40%] relative">
          <ul className="flex justify-around list-none p-0">
            {tabs.map((tab, index) => (
              <li
                key={index}
                className={`cursor-pointer p-2 font-[500] text-[16px] ${activeTab === index ? 'text-blue-700 font-bold' : 'text-[#718EBF]'}`}
                onClick={() => setActiveTab(index)}
              >
                {tab}
              </li>
            ))}
          </ul>
          <div
            className="absolute bottom-0 left-2 w-1/3 h-[3px] rounded-t-[1rem] bg-[#1814F3] transition-all duration-300"
            style={{ transform: `translateX(${activeTab * 100}%)` }}
          />
        </div>
        {activeTab == 0&&<FormComponent 
        setMainData = {setMainData}
        mainData = {mainData}
        />}

          
          {activeTab == 1 && <YourFormComponent/>}
          {activeTab == 2 && <ProfileSecurity/>}
        
      </div>
    </div>
  );
};

export default EditProfile;
export type {MainData}