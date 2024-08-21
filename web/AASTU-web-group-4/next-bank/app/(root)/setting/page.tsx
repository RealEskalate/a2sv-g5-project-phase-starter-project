// SettingsPage.tsx
'use client';
import React, { useState } from 'react';
import EditProfile from '../../../components/updateuser';
import Preference from '../../../components/updateprefrences';
import Security from '../../../components/securityForm';

const SettingsPage = () => {
  const [activeTab, setActiveTab] = useState<'editProfile' | 'preference' | 'security'>('editProfile');

  return (
    <div className="lg:grid lg:grid-cols-[250px_1fr]">
      <aside className="hidden lg:block bg-gray-100 h-screen sticky top-0">
        {/* Sidebar content */}
      </aside>

      <div className="p-4 lg:p-10">
        <div className="flex flex-wrap justify-center lg:justify-start mb-4 space-x-4">
          <button className={`py-2 px-4 ${activeTab === 'editProfile' ? 'bg-blue-800 text-white' : 'bg-gray-200'}`} onClick={() => setActiveTab('editProfile')}>Edit Profile</button>
          <button className={`py-2 px-4 ${activeTab === 'preference' ? 'bg-blue-800 text-white' : 'bg-gray-200'}`} onClick={() => setActiveTab('preference')}>Preference</button>
          <button className={`py-2 px-4 ${activeTab === 'security' ? 'bg-blue-800 text-white' : 'bg-gray-200'}`} onClick={() => setActiveTab('security')}>Security</button>
        </div>

        <div>
          {activeTab === 'editProfile' && <EditProfile />}
          {activeTab === 'preference' && <Preference />}
          {activeTab === 'security' && <Security />}
        </div>
      </div>
    </div>
  );
};

export default SettingsPage;
