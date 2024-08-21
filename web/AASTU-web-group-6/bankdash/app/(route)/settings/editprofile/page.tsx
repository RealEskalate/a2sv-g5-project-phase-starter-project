import EditProfileForm from "@/app/components/Forms/EditProfileForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const ProfilePage: React.FC = () => {
  return (
    <div className="w-full mt-2 px-10 py-7 flex flex-col ">
      <div className="bg-white dark:bg-[#232328] rounded-xl min-h-screen">
        <Navigation />
        <EditProfileForm />
      </div>
    </div>
  );
};

export default ProfilePage;
