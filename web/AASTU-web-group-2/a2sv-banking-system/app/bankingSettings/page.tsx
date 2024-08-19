"use client";
import React, { useState } from 'react';
import Tabs from "../components/Tabs";
import TextInput from '../components/TextInput';
import NotificationToggle from '../components/NotificationToggle';
import EditProfile from '../components/EditProfile';
import SecuritySetting from '../components/SecuritySetting';

const SettingsPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<string>('Preferences');
  const [notifications, setNotifications] = useState({
    digitalCurrency: true,
    merchantOrders: false,
    recommendations: true
  });

  const handleTabChange = (tab: string) => {
    setActiveTab(tab);
  };

  const handleNotificationChange = (key: keyof typeof notifications, checked: boolean) => {
    setNotifications(prev => ({ ...prev, [key]: checked }));
  };

  return (
    <div className="flex min-h-screen bg-gray-50">
      <main className="flex-1 p-4 md:p-8">
        <div className="bg-white p-6 rounded-2xl shadow-md">
          <Tabs
            tabs={['Edit Profile', 'Preferences', 'Security']}
            activeTab={activeTab}
            onTabChange={handleTabChange}
          />

          <div className="mt-6">
            {activeTab === 'Preferences' && (
              <form>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <TextInput id="currency-input" label="Currency" value="USD" readOnly />
                  <TextInput id="timezone-input" label="Time Zone" value="(GMT-12:00) International Date Line West" readOnly />
                </div>

                <div className="mt-6">
                  <label className="block text-sm font-medium text-[#333B69]">Notification</label>
                  <div className="flex flex-col mt-4 gap-4">
                    <NotificationToggle
                      id="notification1"
                      label="I send or receive digital currency"
                      checked={notifications.digitalCurrency}
                      onChange={(checked) => handleNotificationChange('digitalCurrency', checked)}
                    />
                    <NotificationToggle
                      id="notification2"
                      label="I receive merchant orders"
                      checked={notifications.merchantOrders}
                      onChange={(checked) => handleNotificationChange('merchantOrders', checked)}
                    />
                    <NotificationToggle
                      id="notification3"
                      label="There are recommendations for my account"
                      checked={notifications.recommendations}
                      onChange={(checked) => handleNotificationChange('recommendations', checked)}
                    />
                  </div>
                </div>

                <div className="mt-6 flex justify-end">
                  <button
                    type="submit"
                    className="bg-[#1814F3] border border-[#1814F3] rounded-xl text-white px-6 py-3 font-semibold text-xl md:w-1/4"
                  >
                    Save
                  </button>
                </div>
              </form>
            )}

            {activeTab === 'Edit Profile' && (
              <div>
                <EditProfile />
              </div>
            )}

            {activeTab === 'Security' && (
              <div>
                <SecuritySetting />
              </div>
            )}
          </div>
        </div>
      </main>
    </div>
  );
};

export default SettingsPage;
