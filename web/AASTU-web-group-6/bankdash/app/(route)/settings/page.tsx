import EditProfileForm from "@/app/components/Forms/EditProfileForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const ProfilePage: React.FC = () => {
  return (
    <div className="flex flex-col px-6">
    <Navigation/>
    <EditProfileForm/>
    </div>
  );
};

export default ProfilePage;
