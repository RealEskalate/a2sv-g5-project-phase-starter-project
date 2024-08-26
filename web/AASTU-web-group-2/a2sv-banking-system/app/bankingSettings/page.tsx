"use client";
import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { getSession } from "next-auth/react";
import Tabs from '../components/Tabs';
import NotificationToggle from '../components/NotificationToggle';
import EditProfile from '../components/EditProfile';
import SecuritySetting from '../components/SecuritySetting';
import { getCurrentUser, userUpdatePreference } from '../../lib/api/userControl';
import User, { Preference } from '../../types/userInterface';
import Refresh from '@/app/api/auth/[...nextauth]/token/RefreshToken';

type Data = {
  access_token: string;
  data: string;
  refresh_token: string;
};

type SessionDataType = {
  user: Data;
};

const SettingsPage: React.FC = () => {
  const router = useRouter();
  const [loading, setLoading] = useState(true);
  const [session, setSession] = useState<SessionDataType | null>(null);
  const [activeTab, setActiveTab] = useState<string>('Edit Profile');
  const [user, setUser] = useState<User | null>(null);
  const [notifications, setNotifications] = useState<Preference>({
    currency: "",
    sentOrReceiveDigitalCurrency: true,
    receiveMerchantOrder: false,
    accountRecommendations: true,
    timeZone: "",
    twoFactorAuthentication: false,
  });
  const [message, setMessage] = useState<string | null>(null);

  useEffect(() => {
    const fetchSessionAndUser = async () => {
      setLoading(true);

      const sessionData = (await getSession()) as SessionDataType | null;

      if (sessionData && sessionData.user) {
        setSession(sessionData);
        try {
          const userData = await getCurrentUser(sessionData.user.access_token);
          setUser(userData);
          setNotifications(userData.preference);
        } catch (error) {
          console.error("Error fetching user data:", error);
        }
      } else {
        router.push(`./api/auth/signin?callbackUrl=${encodeURIComponent("/accounts")}`);
      }
      setLoading(false);
    };

    fetchSessionAndUser();
  }, [router]);

  const handleTabChange = (tab: string) => setActiveTab(tab);

  const handleNotificationChange = (key: keyof Preference, value: boolean | string) => {
    setNotifications(prev => ({ ...prev, [key]: value }));
  };

  const handleTextInputChange = (key: keyof Preference, value: string) => {
    setNotifications(prev => ({ ...prev, [key]: value }));
  };

  const handlePreferencesUpdate = async (event: React.FormEvent) => {
    event.preventDefault();
    setMessage(null); 
    try {
      if (session?.user?.access_token) {
        const accessToken = await Refresh();
        await userUpdatePreference(notifications, accessToken);
        setMessage('Preferences updated successfully!');
      }
    } catch (error) {
      console.error('Error updating preferences:', error);
      setMessage('Failed to update preferences. Please try again.');
    }
  };

  if (loading) {
    return (
      <div className="flex min-h-screen bg-gray-50 dark:bg-gray-900">
        <main className="flex-1 p-4 md:p-8">
          <div className="bg-white dark:bg-gray-800 p-6 rounded-2xl shadow-md">
            <div className="animate-pulse space-y-6">
              <div className="flex space-x-4">
                <div className="h-10 w-32 bg-gray-200 dark:bg-gray-700 rounded-lg"></div>
                <div className="h-10 w-32 bg-gray-200 dark:bg-gray-700 rounded-lg"></div>
                <div className="h-10 w-32 bg-gray-200 dark:bg-gray-700 rounded-lg"></div>
              </div>

              <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div className="space-y-2">
                  <div className="h-4 bg-gray-200 dark:bg-gray-700 rounded w-1/4"></div>
                  <div className="h-10 bg-gray-200 dark:bg-gray-700 rounded-lg"></div>
                </div>
                <div className="space-y-2">
                  <div className="h-4 bg-gray-200 dark:bg-gray-700 rounded w-1/4"></div>
                  <div className="h-10 bg-gray-200 dark:bg-gray-700 rounded-lg"></div>
                </div>
              </div>

              <div className="space-y-4">
                <div className="h-4 bg-gray-200 dark:bg-gray-700 rounded w-1/6"></div>
                <div className="h-10 bg-gray-200 dark:bg-gray-700 rounded-lg w-1/3"></div>
                <div className="h-10 bg-gray-200 dark:bg-gray-700 rounded-lg w-1/3"></div>
                <div className="h-10 bg-gray-200 dark:bg-gray-700 rounded-lg w-1/3"></div>
              </div>

              <div className="flex justify-end">
                <div className="h-12 w-32 bg-gray-200 dark:bg-gray-700 rounded-lg"></div>
              </div>
            </div>
          </div>
        </main>
      </div>
    );
  }

  return (
    <div className="flex min-h-screen bg-gray-50 dark:bg-[#090b0e]">
      <main className="flex-1 p-4 md:p-8">
        <div className="bg-white dark:bg-[#020817] p-6 rounded-2xl shadow-md">
          <Tabs
            tabs={['Edit Profile', 'Preferences', 'Security']}
            activeTab={activeTab}
            onTabChange={handleTabChange}
          />

          <div className="mt-6">
            {activeTab === 'Preferences' && (
              <form onSubmit={handlePreferencesUpdate}>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                  <div>
                    <label htmlFor="currency-input" className="block text-sm font-medium text-[#232323] dark:text-[#9faaeb]">Currency</label>
                    <input
                      type="text"
                      id="currency-input"
                      title="Currency"
                      className="mt-1 block w-full border border-[#DFEAF2] dark:bg-[#020817] dark:border dark:border-[#333B69] rounded-full shadow-sm px-4 py-2 text-[#232323] dark:text-[#9faaeb] bg-white placeholder-[#718EBF] dark:placeholder-[#9faaeb]"
                      placeholder="Enter your currency"
                      value={notifications.currency || ""}  
                      onChange={(e) => handleTextInputChange('currency', e.target.value)}
                    />
                  </div>
                  <div>
                    <label htmlFor="timezone-input" className="block text-sm font-medium text-[#232323] dark:text-[#9faaeb]">Time Zone</label>
                    <input
                      type="text"
                      id="timezone-input"
                      title="Time Zone"
                      className="mt-1 block w-full border border-[#DFEAF2] dark:bg-[#020817] dark:border dark:border-[#333B69] rounded-full shadow-sm px-4 py-2 text-[#718EBF] dark:text-[#9faaeb] bg-white placeholder-[#718EBF] dark:placeholder-[#9faaeb]"
                      placeholder="Enter your time zone"
                      value={notifications.timeZone || ""}  
                      onChange={(e) => handleTextInputChange('timeZone', e.target.value)}
                    />
                  </div>
                </div>

                <div className="mt-6">
                  <label className="block text-sm font-medium text-[#333B69] dark:text-[#9faaeb]">Notifications</label>
                  <div className="flex flex-col mt-4 gap-4">
                  <NotificationToggle
                    id="notification1"
                    label="Send/Receive Digital Currency"
                    checked={notifications.sentOrReceiveDigitalCurrency}
                    onChange={(checked) => handleNotificationChange('sentOrReceiveDigitalCurrency', checked)}
                  />
                  <NotificationToggle
                    id="notification2"
                    label="Receive Merchant Orders"
                    checked={notifications.receiveMerchantOrder}
                    onChange={(checked) => handleNotificationChange('receiveMerchantOrder', checked)}

                    />
                  <NotificationToggle
                    id="notification3"
                    label="Account Recommendations"
                    checked={notifications.accountRecommendations}
                    onChange={(checked) => handleNotificationChange('accountRecommendations', checked)}
                   
                  />

                  </div>
                </div>

                {message && (
                  <div className="mt-6 text-center">
                    <span className="text-sm text-green-600 dark:text-[#9faaeb]">{message}</span>
                  </div>
                )}

                <div className="mt-6 flex justify-end">
                  <button
                    type="submit"
                    className="bg-[#1814F3] hover:bg-[#423fef]   dark:bg-[#1814F3] dark:hover:bg-[#423fef] text-white px-5 py-3 rounded-xl md:w-2/12 text-sm w-full self-end"
                  >
                    Save
                  </button>
                </div>
              </form>
            )}

            {activeTab === 'Edit Profile' && <EditProfile />}
            {activeTab === 'Security' && <SecuritySetting />}
          </div>
        </div>
      </main>
    </div>
  );
};

export default SettingsPage;
