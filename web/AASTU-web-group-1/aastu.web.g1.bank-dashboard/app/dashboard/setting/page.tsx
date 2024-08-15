import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import Link from "next/link";
import React from "react";
import EditProfile from "./component/EditProfile";
import Preferences from "./component/Preferences";
import Security from "./component/Security";

const Setting = () => {
  return (
    <div className="flex items-center justify-center h-full">
      <Tabs defaultValue="editProfile" className="w-[400px] bg-white">
        <TabsList className="bg-transparent p-12">
          <TabsTrigger
            value="editProfile"
            className="data-[state=active]:border-b-primaryBlue data-[state=active]:border-b-4 data-[state=active]:text-primaryBlue bg-gray-100 text-gray-800 bg-transparent pb-4 px-5"
          >
            Edit Profile
          </TabsTrigger>
          <TabsTrigger
            value="Preferences"
            className="data-[state=active]:border-b-primaryBlue data-[state=active]:border-b-4 data-[state=active]:text-primaryBlue bg-gray-100 text-gray-800 bg-transparent pb-4 px-5"
          >
            Preferences
          </TabsTrigger>
          <TabsTrigger
            value="security"
            className="data-[state=active]:border-b-primaryBlue data-[state=active]:border-b-4 data-[state=active]:text-primaryBlue bg-gray-100 text-gray-800 bg-transparent pb-4 px-5"
          >
            Security
          </TabsTrigger>
          <hr />
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
