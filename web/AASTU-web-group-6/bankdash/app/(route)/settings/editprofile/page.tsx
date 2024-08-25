import EditProfileForm from "@/app/components/Forms/EditProfileForm";
import Navigation from "@/app/components/Settings/Navigation";
import React from "react";

const ProfilePage: React.FC = () => {
  return (
    <div className="w-full flex justify-center mt-15  md:max-h-[717px] ">
      <div className="w-full max-w-[1110px] bg-white px-3 mt-2 flex flex-col justify-center items-center rounded-xl  ">
        <Navigation />
        <EditProfileForm />
      </div>
    </div>
  );
};

export default ProfilePage;
