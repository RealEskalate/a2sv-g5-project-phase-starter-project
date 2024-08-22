import React, { useState } from 'react';
import { useForm, Controller } from 'react-hook-form';
import { yupResolver } from '@hookform/resolvers/yup';
import * as yup from 'yup';
import { FaPencilAlt } from 'react-icons/fa';
import { MainData } from './Signup'; // Adjust the import path as necessary

interface FormValues {
  name: string;
  email: string;
  dateOfBirth: string; // Change to string to match the input type
  permanentAddress: string;
  postalCode: string;
  username: string;
  password: string;
  presentAddress: string;
  city: string;
  country: string;
  profilePicture?: string | null;
}

interface FormComponentProps {
  mainData: MainData;
  setMainData: React.Dispatch<React.SetStateAction<MainData>>;
  setActiveTab: React.Dispatch<React.SetStateAction<number>>;
}

// Define the schema
const schema = yup.object().shape({
  name: yup.string().required('Name is required'),
  email: yup.string().email('Invalid email format').required('Email is required'),
  dateOfBirth: yup.string().required('Date of Birth is required'),
  permanentAddress: yup.string().required('Permanent Address is required'),
  postalCode: yup.string().required('Postal Code is required'),
  username: yup.string().required('Username is required'),
  password: yup.string().min(6, 'Password must be at least 6 characters').required('Password is required'),
  presentAddress: yup.string().required('Present Address is required'),
  city: yup.string().required('City is required'),
  country: yup.string().required('Country is required'),
});

const FormComponent: React.FC<FormComponentProps> = ({ mainData, setMainData, setActiveTab }) => {
  const [selectedImage, setSelectedImage] = useState<string | null>(mainData.profilePicture || null);

  const { control, formState: { errors } } = useForm<FormValues>({
    resolver: yupResolver(schema),
    defaultValues: {
      name: mainData.name,
      email: mainData.email,
      dateOfBirth: mainData.dateOfBirth.split('T')[0], // Adjust for date format
      permanentAddress: mainData.permanentAddress,
      postalCode: mainData.postalCode,
      username: mainData.username,
      password: mainData.password,
      presentAddress: mainData.presentAddress,
      city: mainData.city,
      country: mainData.country,
      profilePicture: mainData.profilePicture,
    }
  });

  // Function to update mainData when input changes
  const handleInputChange = (field: keyof FormValues, value: string) => {
    setMainData((prev) => ({ ...prev, [field]: value }));
  };

  const handleImageChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    if (event.target.files && event.target.files[0]) {
      const reader = new FileReader();
      reader.onload = (e) => {
        const imageResult = e.target?.result as string;
        setSelectedImage(imageResult);
        handleInputChange('profilePicture', imageResult);
      };
      reader.readAsDataURL(event.target.files[0]);
    }
  };

  return (
    <form className=''>
      <div className="flex max-md:flex-col justify-between gap-[2rem]">
        <div className="flex flex-col w-[30rem] max-md:w-full">
          <div className="mb-6 flex justify-center">
            <label className="relative cursor-pointer">
              <input
                type="file"
                className="hidden"
                onChange={handleImageChange}
              />
              <div className="w-[10rem] h-[10rem]  rounded-full bg-gray-200 flex items-center justify-center text-gray-500 overflow-hidden">
                {selectedImage ? (
                  <img
                    src={selectedImage}
                    alt="Profile"
                    className="object-cover w-full h-full"
                  />
                ) : (
                  <img
                    src="/public/assets/auth/avatardefault.png"
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
          {(['name', 'email', 'dateOfBirth', 'permanentAddress', 'postalCode'] as const).map((field) => (
            <div className="mb-4" key={field}>
              <Controller
                name={field}
                control={control}
                render={({ field: { value, onChange, ...restField } }) => (
                  <div className="mb-4">
                    <label className="block mb-1 font-400 text-[16px] text-[#232323] capitalize">
                      {field.charAt(0).toUpperCase() + field.slice(1).replace(/([A-Z])/g, ' $1')}
                    </label>
                    <input
                      {...restField}
                      type={field === 'dateOfBirth' ? 'date' : 'text'}
                      placeholder={`Enter your ${field}`}
                      className="w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200"
                      value={value}
                      onChange={(e) => {
                        onChange(e);
                        handleInputChange(field, e.target.value);
                      }}
                    />
                    {errors[field] && <p className="text-red-500 text-sm">{errors[field]?.message}</p>}
                  </div>
                )}
              />
            </div>
          ))}
        </div>
        <div className="flex flex-col w-[100%]">
          {(['username', 'password', 'presentAddress', 'city', 'country'] as const).map((field) => (
            <div className="mb-4" key={field}>
              <Controller
                name={field}
                control={control}
                render={({ field: { value, onChange, ...restField } }) => (
                  <div className="mb-4">
                    <label className="block mb-1 font-400 text-[16px] text-[#232323] capitalize">
                      {field.charAt(0).toUpperCase() + field.slice(1).replace(/([A-Z])/g, ' $1')}
                    </label>
                    <input
                      {...restField}
                      type={field === 'password' ? 'password' : 'text'}
                      placeholder={`Enter your ${field}`}
                      className="w-full p-2 border border-[#DFEAF2] rounded-[15px] focus:outline-none focus:ring-2 focus:ring-blue-200"
                      value={value}
                      onChange={(e) => {
                        onChange(e);
                        handleInputChange(field, e.target.value);
                      }}
                    />
                    {errors[field] && <p className="text-red-500 text-sm">{errors[field]?.message}</p>}
                  </div>
                )}
              />
            </div>
          ))}
        </div>
      </div>
      <div className="flex justify-end mt-6 ">
        <button type="button" onClick={() => setActiveTab(1)} className="bg-blue-700 md:w-[190px] text-white px-6 py-2 rounded-lg hover:bg-blue-600">
          Next
        </button>
      </div>
    </form>
  );
};

export default FormComponent;
