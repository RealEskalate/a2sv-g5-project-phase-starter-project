'use client';
import React, { useState } from 'react';
import on from '../../../public/images/on.svg';
import off from '../../../public/images/of.svg';
import UserProfile from './setting';
interface FormData {
  currency: string;
  Time_Zone: string;
}

const dummyData: FormData = {
  currency: 'USD',
  Time_Zone: 'GMT-5',
};

const Preferences = () => {
  const [formData, setFormData] = useState<FormData>(dummyData);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });
  };

  const [activeTab, setActiveTab] = useState('editProfile');
  const [notifications, setNotifications] = useState({
    digitalCurrency: true,
    merchantOrder: false,
    recommendations: true,
  });

  const toggleNotification = (key: keyof typeof notifications) => {
    setNotifications(prevState => ({
      ...prevState,
      [key]: !prevState[key],
    }));
  };

  return (
    <div className='container md:mt-10 md:ml-10 space-y-6 rounded-[60px] shadow-md p-6 bg-white w-full md:w-[1110px] '>
      <div>
        <div className="flex flex-row space-x-6 mt-3">
          <h4 className={`cursor-pointer ${activeTab === 'Edit Profile' ? 'text-[#1814F3]' : 'text-[#718EBF]'}`}
            onClick={() => setActiveTab('Edit Profile')}  > Edit Profile </h4>
          <h4 className={`cursor-pointer ${activeTab === 'Preference' ? 'text-[#1814F3]' : 'text-[#718EBF]'}`}
            onClick={() => setActiveTab('Preference')}>Preference </h4>
          <h4 className={`cursor-pointer ${activeTab === 'Security' ? 'text-[#1814F3]' : 'text-[#718EBF]'}`}
            onClick={() => setActiveTab('Security')} > Security</h4>
        </div>

        <div className="relative">
          {activeTab === 'Edit Profile' && <div className="border-b-2 border-[#1814F3] w-24 absolute left-0"></div>}
          {activeTab === 'Preference' && <div className="border-b-2 border-[#1814F3] w-24 absolute left-[100px]"></div>}
          {activeTab === 'Security' && <div className="border-b-2 border-[#1814F3] w-24 absolute left-[200px]"></div>}
        </div>
      </div>
      {activeTab === 'Edit Profile' && <UserProfile/>}
      {activeTab === 'Preference' && (
        <div>
          <form action="">
            <div className='flex flex-col md:flex-row space-y-4 md:space-y-0 md:space-x-4 '>
              <div className='flex flex-col space-y-3'>
                <label htmlFor="currency">Currency</label>
                <input
                  type="text"
                  id="currency"
                  name="currency"
                  value={formData.currency}
                  onChange={handleChange}
                  className='border border-1 rounded-[15px] p-[10px] w-full md:w-[510px]  text-[#718EBF] text-[15px]'
                />
              </div>
              <div className='flex flex-col space-y-3'>
                <label htmlFor="Time_Zone">Time Zone</label>
                <input
                  type="text"
                  id="Time_Zone"
                  name="Time_Zone"
                  value={formData.Time_Zone}
                  onChange={handleChange}
                  className='border border-1 rounded-[15px] p-[10px] w-full md:w-[510px] text-[#718EBF] text-[15px]'
                />
              </div>
            </div>
          </form>

          <div className='flex flex-col gap-2'>
            <div className='font-medium mt-3 mb-3'>Notification</div>
            <div className='flex flex-row gap-5 items-center' onClick={() => toggleNotification('digitalCurrency')}>
              <img src={notifications.digitalCurrency ? on.src : off.src} alt="" />
              <h1 className='font-normal '>I send or receive digital currency</h1>
            </div>
            <div className='flex flex-row gap-5 items-center' onClick={() => toggleNotification('merchantOrder')}>
              <img src={notifications.merchantOrder ? on.src : off.src} alt="" />
              <h2>I receive merchant order</h2>
            </div>
            <div className='flex flex-row gap-5 items-center' onClick={() => toggleNotification('recommendations')}>
              <img src={notifications.recommendations ? on.src : off.src} alt="" />
              <h2>There are recommendations for my account</h2>
            </div>
            <button type="submit" className='mt-4 bg-[#1814F3] text-white p-2 rounded-[15px] w-full md:w-[130px] h-[40px] md:ml-[900px]'>Save</button>
          </div>
        </div>
      )}
    </div>
  );
};

export default Preferences;
