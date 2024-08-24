import React, { useState } from "react";
import { useForm, Controller } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import * as yup from "yup";
import { FaPencilAlt } from "react-icons/fa";
import { useDispatch, useSelector } from "react-redux";
import { RootState } from "@/lib/store";
import { setUser } from "@/lib/features/userSlice/userSlice";
import { useUpdateUserMutation } from "@/lib/service/UserService";
import { useSession } from "next-auth/react";
import notify from "@/utils/notify";

interface FormValues {
  name: string;
  email: string;
  dateOfBirth: string;
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture?: string | null;
}

const schema = yup.object().shape({
  name: yup.string().required("Name is required"),
  email: yup
    .string()
    .email("Invalid email format")
    .required("Email is required"),
  dateOfBirth: yup.string().required("Date of Birth is required"),
  permanentAddress: yup.string().required("Permanent Address is required"),
  postalCode: yup.string().required("Postal Code is required"),
  username: yup.string().required("Username is required"),
  password: yup
    .string()
    .min(6, "Password must be at least 6 characters")
    .required("Password is required"),
  presentAddress: yup.string().required("Present Address is required"),
  city: yup.string().required("City is required"),
  country: yup.string().required("Country is required"),
});

const FormComponent: React.FC = () => {
  const user = useSelector((state: RootState) => state.user.user);
  const dispatch = useDispatch();
  const [updateUser, { isLoading, isError }] = useUpdateUserMutation(); // Use the mutation hook

  const [selectedImage, setSelectedImage] = useState<string | null>(
    user?.profilePicture || null
  );

  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm<FormValues>({
    resolver: yupResolver(schema),
    defaultValues: {
      name: user?.name || "",
      email: user?.email || "",
      dateOfBirth: user?.dateOfBirth?.split("T")[0] || "",
      permanentAddress: user?.permanentAddress || "",
      postalCode: user?.postalCode || "",
      username: user?.username || "",
      password: "",
      presentAddress: user?.presentAddress || "",
      city: user?.city || "",
      country: user?.country || "",
      profilePicture: user?.profilePicture || "",
    },
  });
  const { data: session, status } = useSession();

  const onSubmit = async (data: FormValues) => {
    const updatedUser = {
      ...user,
      ...data,
      dateOfBirth: new Date(data.dateOfBirth).toISOString(),
      profilePicture: selectedImage,
    };

    try {
      if (session?.user?.accessToken) {
        const response = await updateUser({
          accessToken: session.user.accessToken,
          updatedUser: updatedUser,
        }).unwrap();
        console.log("Updated User:", response);
        dispatch(setUser(response.data));
        notify.success("Profile updated successfully");
        // Show success message or handle successful update
      } else {
        notify.error("Access token is missing");

        throw new Error("Access token is missing");
      }
    } catch (error) {
      console.error("Failed to update user:", error);
      // Show error message or handle error
    }
  };

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const reader = new FileReader();
      reader.onload = (e) => {
        setSelectedImage(e.target?.result as string);
      };
      reader.readAsDataURL(event.target.files[0]);
    }
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)} className="">
      <div className="flex max-md:flex-col justify-between gap-[2rem]">
        <div className="flex flex-col w-[30rem] max-md:w-full">
          <div className="mb-6 flex justify-center">
            <label className="relative cursor-pointer">
              <input
                type="file"
                className="hidden"
                onChange={handleImageChange}
              />
              <div className="w-[10rem] h-[10rem] rounded-full bg-gray-200 flex items-center justify-center text-gray-500 overflow-hidden">
                {selectedImage ? (
                  <img
                    src={selectedImage}
                    alt="Profile"
                    className="object-cover w-full h-full"
                  />
                ) : (
                  <img
                    src="/assets/auth/avatardefault.png"
                    alt="Profile"
                    className="object-cover w-full h-full"
                  />
                )}
                <div className="absolute bottom-2 right-6 bg-blue-700 p-1 rounded-full transform translate-x-1/4 translate-y-1/4">
                  <FaPencilAlt color="white" size={16} />
                </div>
              </div>
            </label>
          </div>
        </div>

        <div className="flex flex-col w-[100%]">
          {(
            [
              "name",
              "email",
              "dateOfBirth",
              "permanentAddress",
              "postalCode",
            ] as const
          ).map((field) => (
            <div className="mb-4" key={field}>
              <Controller
                name={field}
                control={control}
                render={({ field }) => (
                  <div className="mb-4">
                    <label className="block mb-1 font-400 text-[16px] text-[#232323] capitalize">
                      {field.name.charAt(0).toUpperCase() +
                        field.name.slice(1).replace(/([A-Z])/g, " $1")}
                    </label>
                    <input
                      {...field}
                      type={field.name === "dateOfBirth" ? "date" : "text"}
                      placeholder={field.value as string}
                      className="w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200"
                      onChange={(e) => field.onChange(e.target.value)}
                    />
                    {errors[field.name] && (
                      <p className="text-red-500 text-sm">
                        {errors[field.name]?.message}
                      </p>
                    )}
                  </div>
                )}
              />
            </div>
          ))}
        </div>
        <div className="flex flex-col w-[100%]">
          {(
            [
              "username",
              "password",
              "presentAddress",
              "city",
              "country",
            ] as const
          ).map((field) => (
            <div className="mb-4" key={field}>
              <Controller
                name={field}
                control={control}
                render={({ field }) => (
                  <div className="mb-4">
                    <label className="block mb-1 font-400 text-[16px] text-[#232323] capitalize">
                      {field.name.charAt(0).toUpperCase() +
                        field.name.slice(1).replace(/([A-Z])/g, " $1")}
                    </label>
                    <input
                      {...field}
                      type={field.name === "password" ? "password" : "text"}
                      placeholder={field.value as string}
                      className="w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200"
                      onChange={(e) => field.onChange(e.target.value)}
                    />
                    {errors[field.name] && (
                      <p className="text-red-500 text-sm">
                        {errors[field.name]?.message}
                      </p>
                    )}
                  </div>
                )}
              />
            </div>
          ))}
        </div>
      </div>
      <div className="flex justify-end mt-6">
        <button
          type="submit"
          className="w-full md:w-auto px-4 py-2 bg-[#1814F3] text-white rounded-lg"
          disabled={isLoading}
        >
          Save Changes
        </button>
      </div>
    </form>
  );
};

export default FormComponent;
