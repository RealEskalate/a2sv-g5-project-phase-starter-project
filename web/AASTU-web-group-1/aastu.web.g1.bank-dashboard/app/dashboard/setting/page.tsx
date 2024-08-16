import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import EditProfile from "./component/EditProfile";
import Preferences from "./component/Preferences";
import Security from "./component/Security";

const Setting = () => {
  return (
    <div className="flex w-full mt-5 md:ml-5">
      <Tabs
        defaultValue="editProfile"
        className="w-full max-w-full md:max-w-[900px] bg-white p-4 rounded-2xl"
      >
        <TabsList className="bg-transparent md:p-12">
          <TabsTrigger value="editProfile" className="tabs">
            Edit Profile
          </TabsTrigger>
          <TabsTrigger value="Preferences" className="tabs">
            Preferences
          </TabsTrigger>
          <TabsTrigger value="security" className="tabs">
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
