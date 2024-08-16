import Image from "next/image";
import React from "react";
import { ProfileForm } from "./Form";

const EditProfile = () => {
  return (
    <div className="md:flex gap-12 md:px-12">
      <div>
        <Image
          alt="Profile"
          className="rounded-full ml-auto mr-auto mt-5 lg:mt-0"
          src="https://github.com/shadcn.png"
          width={100}
          height={100}
        />
      </div>

      <div className="flex-grow">
        <ProfileForm />
      </div>
    </div>
  );
};

export default EditProfile;
