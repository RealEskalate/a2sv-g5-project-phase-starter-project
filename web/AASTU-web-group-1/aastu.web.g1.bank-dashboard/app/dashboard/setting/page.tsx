"use client";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import EditProfile from "./component/EditProfile";
import Preferences from "./component/Preferences";
import Security from "./component/Security";
import { useUser } from "@/contexts/UserContext";
import Refresh from "@/app/api/auth/[...nextauth]/token/RefreshToken";
import { useEffect } from "react";

const Setting = () => {
  const { isDarkMode } = useUser();
  return (
    <div className="flex justify-center w-full mt-5 md:ml-5">
      <Tabs
        defaultValue="editProfile"
        className={`w-full max-w-full md:max-w-[1100px] p-4 rounded-2xl ${
          isDarkMode ? "bg-gray-800 text-gray-200" : "bg-white text-black"
        }`}
      >
        <TabsList
          className={`bg-transparent md:p-12 ${
            isDarkMode ? "text-gray-300" : "text-black"
          }`}
        >
          <TabsTrigger
            value="editProfile"
            className={`tabs ${
              isDarkMode
                ? "text-gray-300 hover:bg-gray-700"
                : "text-black hover:bg-gray-100"
            }`}
          >
            Edit Profile
          </TabsTrigger>
          <TabsTrigger
            value="Preferences"
            className={`tabs ${
              isDarkMode
                ? "text-gray-300 hover:bg-gray-700"
                : "text-black hover:bg-gray-100"
            }`}
          >
            Preferences
          </TabsTrigger>
          <TabsTrigger
            value="security"
            className={`tabs ${
              isDarkMode
                ? "text-gray-300 hover:bg-gray-700"
                : "text-black hover:bg-gray-100"
            }`}
          >
            Security
          </TabsTrigger>
        </TabsList>
        <TabsContent value="editProfile">
          <EditProfile />
        </TabsContent>
        <TabsContent value="Preferences">
          <Preferences />
        </TabsContent>
        <TabsContent value="security">
          <Security />
        </TabsContent>
      </Tabs>
    </div>
  );
};

export default Setting;
