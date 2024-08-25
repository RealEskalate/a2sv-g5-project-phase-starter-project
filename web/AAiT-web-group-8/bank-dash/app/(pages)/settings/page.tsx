'use client';

import EditProfile from '@/components/settings/EditProfile';
import Preference from '@/components/settings/Preference';
import Security from '@/components/settings/Security';
import { useState } from 'react';


const Tabs = () => {
  const [activeTab, setActiveTab] = useState('editProfile');

  const renderTabContent = () => {
    switch (activeTab) {
      case 'editProfile':
        return <EditProfile isActive={false} />;
      case 'preferences':
        return <Preference />;
      case 'security':
        return <Security />;
      default:
        return <EditProfile isActive={false} />;
    }
  };

  return (
    <div className="bg-white w-full lg:w-[1030px] rounded-2xl">
      <div className="ml-4 lg:ml-14 flex justify-start gap-4 lg:gap-10 lg:w-[950px] border-b bg-white pt-5 pb-0 text-[#718EBF]">
        <button
          className={`${
            activeTab === 'editProfile'
              ? 'border-b-2 border-[#1814f3] text-[#1814f3] font-bold'
              : ''
          }`}
          onClick={() => setActiveTab('editProfile')}
        >
          Edit Profile
        </button>
        <button
          className={`${
            activeTab === 'preferences'
              ? 'border-b-2 border-[#1814f3] text-[#1814f3] font-bold'
              : ''
          }`}
          onClick={() => setActiveTab('preferences')}
        >
          Preferences
        </button>
        <button
          className={`${
            activeTab === 'security'
              ? 'border-b-2 border-[#1814f3] text-[#1814f3] font-bold'
              : ''
          }`}
          onClick={() => setActiveTab('security')}
        >
          Security
        </button>
      </div>
      <div className="p-6 w-full lg:w-[1030px]">{renderTabContent()}</div>
    </div>
  );
};

export default Tabs;