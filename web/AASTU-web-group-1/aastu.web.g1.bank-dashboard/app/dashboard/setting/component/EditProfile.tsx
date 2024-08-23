"use client";
import ky from "ky";
import { getSession } from "next-auth/react";
import Image from "next/image";
import { useEffect, useState } from "react";
import { ProfileForm } from "./Form";

interface FormData {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  presentAddress: string;
  city: string;
  country: string;
  username: string;
  password: string;
}

const EditProfile = () => {
  const [profilePic, setProfilePic] = useState("");

  const [formData, setFormData] = useState<FormData>({
    name: "",
    email: "",
    dateOfBirth: "",
    permanentAddress: "",
    postalCode: "",
    presentAddress: "",
    city: "",
    country: "",
    username: "",
    password: "",
  });

  useEffect(() => {
    const fetchUser = async () => {
      const session = await getSession();
      const accessToken = session?.user.accessToken;
      console.log(accessToken);
      if (!accessToken) {
        throw new Error("No access token found");
      }

      try {
        const res: any = await ky(
          `${process.env.NEXT_PUBLIC_BASE_URL}/user/current`,
          {
            headers: {
              Authorization: `Bearer ${accessToken}`,
            },
            timeout: 10000,
          }
        ).json();
        setProfilePic(res.data.profilePicture);

        setFormData((prevFormData) => ({
          ...prevFormData,
          name: res.data.name,
          email: res.data.email,
          dateOfBirth: res.data.dateOfBirth,
          permanentAddress: res.data.permanentAddress,
          postalCode: res.data.postalCode,
          presentAddress: res.data.presentAddress,
          city: res.data.city,
          country: res.data.country,
          username: res.data.username,
          password: res.data.password,
        }));
      } catch (error) {
        console.error("Failed to fetch user:", error);
      }
    };

    fetchUser();
  }, []);
  return (
    <div className="md:flex gap-12 md:px-12">
      <div>
        <Image
          alt="Profile"
          className="rounded-full ml-auto mr-auto mt-5 lg:mt-0"
          src={profilePic}
          width={100}
          height={100}
        />
      </div>

      <div className="flex-grow">
        <ProfileForm formData={formData} />
      </div>
    </div>
  );
};

export default EditProfile;
