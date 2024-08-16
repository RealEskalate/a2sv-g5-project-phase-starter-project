import { Tabs, TabsContent, TabsList, TabsTrigger } from '@/components/ui/tabs';

import EditProfile from './EditProfile';
import Preferences from './Preferences';
import Security from './Security';

const Settings = () => {
  return (
    <>
      <div className='bg-white p-8 rounded-3xl'>
        <Tabs defaultValue='account' className='w-full'>
          <TabsList className='w-full mb-6 md:justify-start bg-white'>
            <TabsTrigger
              value='edit_profile'
              className='data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]'
            >
              Edit Profile
            </TabsTrigger>
            <TabsTrigger
              value='preferences'
              className=' data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]'
            >
              Preferences
            </TabsTrigger>
            <TabsTrigger
              value='security'
              className=' data-[state=active]:border-b-transparent border-b-4 border-b-transparent data-[state=active]:text-[#1814f3] data-[state=active]:border-b-[#1814f3]'
            >
              Security
            </TabsTrigger>
          </TabsList>

          <TabsContent value='edit_profile'>
            <EditProfile />
          </TabsContent>
          <TabsContent value='preferences'>
            <Preferences />
          </TabsContent>
          <TabsContent value='security'>
            <Security />
          </TabsContent>
        </Tabs>
      </div>
    </>
  );
};

export default Settings;
