import React, { useEffect, useRef, useState } from "react";
import { useForm } from "react-hook-form";
import Image from "next/image";
import {
  getStorage,
  ref,
  uploadBytesResumable,
  getDownloadURL,
} from "firebase/storage";
import app from "@/app/firebase";
import { circleWithPen, profilepic } from "@/../../public/Icons";
import { useDispatch, useSelector } from "react-redux";
import { usePutSettingMutation } from "@/lib/redux/api/settingApi";
import { RootState } from "@/lib/redux/store";
import { settingPutUserResponse } from "@/lib/redux/types/setting";

interface FormInput{
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture: string; // URL of the uploaded image
}
interface EditProfileProps {
  userData: settingPutUserResponse[];
}

const EditProfile = ({ userData }: EditProfileProps ) => {
  const [currentUser, setCurrentUser] = useState<settingPutUserResponse>(userData[0]);

  useEffect(() => {
    if (userData && userData.length > 0) {
      setCurrentUser(userData[0]);
    }
  }, [userData]);


  const { register, handleSubmit, setValue,formState } = useForm<FormInput>({
    defaultValues: {
      name: currentUser?.data.name || "",
      email:currentUser?.data.email || "",
      dateOfBirth:"",
      permanentAddress:currentUser?.data.permanentAddress || "",
      postalCode: currentUser?.data.postalCode || "",
      username: currentUser?.data.username || "",
      presentAddress: currentUser?.data.presentAddress || "",
      city: currentUser?.data.city || "",
      country:currentUser?.data.country || "",
      profilePicture:currentUser?.data.profilePicture || "",
    },
  });


  const {isSubmitSuccessful,isSubmitting} = formState
  const dispatch = useDispatch();
  const { loading, error } = useSelector((state: RootState) => state.setting);

  const [putSetting] = usePutSettingMutation();

  const fileInputRef = useRef<HTMLInputElement>(null);

  const onSubmit = async (data: FormInput) => {
    try {
      await putSetting(data ).unwrap();
      window.location.reload()

    } catch (err) {
      console.error(err);
    }
  };
  

  const handleImageClick = () => {
    if (fileInputRef.current) {
      fileInputRef.current.click();
    }
  };

  const handleFileChange = async (event: React.ChangeEvent<HTMLInputElement>) => {
    const file = event.target.files?.[0];
    if (file) {
      const imageUrl = await uploadImageToCloud(file);
      setValue("profilePicture", imageUrl);
    } else {
      setValue("profilePicture", "");
    }
  };

  // Mock upload function (replace this with actual implementation)
  const uploadImageToCloud = async (file: File): Promise<string> => {
    const storage = getStorage(app);
    const storageRef = ref(storage, file.name);
  
    const uploadTask = uploadBytesResumable(storageRef, file);
  
    return new Promise((resolve, reject) => {
      uploadTask.on(
        'state_changed',
        (snapshot) => {
          const progress = (snapshot.bytesTransferred / snapshot.totalBytes) * 100;
          if(progress == 100){
            alert(`Upload is ${progress}% done`);
          }
        },
        (error) => {
          console.error("Error during upload:", error);
          reject(error);
        },
        async () => {
          try {
            const url = await getDownloadURL(uploadTask.snapshot.ref);
            console.log("File available at", url);
            resolve(url);
          } catch (err) {
            console.error("Error getting download URL:", err);
            reject(err);
          }
        }
      );
    });
  };

  console.log("current user",currentUser,currentUser?.data.name)
  const formFields = [
    { id: "name", label: "Your Name", value: currentUser?.data.name || "Charlene Reed", type: "text" },
    { id: "username", label: "User Name", value: currentUser?.data.username || "Charlene Reed", type: "text" },
    { id: "email", label: "Email", value: currentUser?.data.email || "charlenereed@gmail.com", type: "email" },
    { id: "dateOfBirth", label: "Date of Birth", value:  "Enter Date of Birth", type: "date" },
    { id: "presentAddress", label: "Present Address", value: currentUser?.data.presentAddress || "San Jose, California, USA", type: "text" },
    { id: "permanentAddress", label: "Permanent Address", value: currentUser?.data.permanentAddress || "San Jose, California, USA", type: "text" },
    { id: "city", label: "City", value: currentUser?.data.city || "San Jose", type: "text" },
    { id: "postalCode", label: "Postal Code", value: currentUser?.data.postalCode || "45962", type: "text" },
    { id: "country", label: "Country", value: currentUser?.data.country || "USA", type: "text" },
  ];


  return (
    <div className="p-4 flex flex-col md:flex-row gap-8 dark:bg-[#172941] ">
      <div className="relative rounded-full w-64 h-64 mb-5 md:w-40 md:h-40">
        <Image
          src={currentUser.data.profilePicture}
          width={256}
          height={256}
          alt="profilepic"
          className="w-64 h-64 md:w-40 md:h-40 object-cover rounded-full"
        />
        <Image
          src={circleWithPen}
          alt="edit icon"
          width={64}
          height={64}
          className="absolute z-30 right-1 bottom-10 object-cover md:w-10 md:h-10 md:bottom-5 cursor-pointer"
          onClick={handleImageClick}
        />
        <input
          type="file"
          ref={fileInputRef}
          className="hidden"
          onChange={handleFileChange}
        />
      </div>

      <form
        onSubmit={handleSubmit(onSubmit)}
        className="flex flex-wrap items-end justify-between w-full md:w-4/5"
      >
        {formFields.map((field) => (
          <div key={field.id} className="mb-3 w-full md:w-[45%]">
            <label className="block text-black text-sm mb-2 dark:text-white" htmlFor={field.id}>
              {field.label}
            </label>
            <input
              className="w-full p-3 md:p-2 text-[#718EBF] border-2 text-sm border-[#DFEAF2] rounded-lg focus:outline-none bg-white dark:border-white dark:bg-[#172941]"
              type={field.type}
              id={field.id}
              defaultValue={field.value}
              {...register(field.id as keyof FormInput, {
                required: `${field.label} is required`,
              })}
            />
          </div>
        ))}
        <div className="flex justify-end w-full">
          <button
            className="w-full md:w-1/5 bg-[#1814F3] text-white font-semibold py-2 px-4 rounded-lg focus:outline-none"
            type="submit"
          >{
            !isSubmitting ? "Save":"saving ... " }
          </button>
        </div>
      </form>
    </div>
  );
};

export default EditProfile;
 