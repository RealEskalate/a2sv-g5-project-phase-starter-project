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

      <div className="p-4 lg:p-10 w-full">
        {/* Tabs */}
        <div className="flex justify-start gap-4 space-x-8 border-b-2 border-gray-200 mb-4">
          <button
            className={`py-2 text-lg ${activeTab === 'editProfile' ? 'border-b-4 border-blue-800 text-black font-semibold' : 'text-gray-500'}`}
            onClick={() => setActiveTab('editProfile')}
          >
            Edit Profile
          </button>
          <button
            className={`py-2 text-lg ${activeTab === 'preference' ? 'border-b-4 border-blue-800 text-black font-semibold' : 'text-gray-500'}`}
            onClick={() => setActiveTab('preference')}
          >
            Preference
          </button>
          <button
            className={`py-2 text-lg ${activeTab === 'security' ? 'border-b-4 border-blue-800 text-black font-semibold' : 'text-gray-500'}`}
            onClick={() => setActiveTab('security')}
          >
            Security
          </button>
        </div>

        {/* Tab Content */}
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
