// "use client"

// import ToggleButton from "./Toggle"
// import { SubmitHandler, useForm } from 'react-hook-form';
// import { useEffect, useState } from 'react';
// import { useSession } from "next-auth/react";
// import { useGetCurrentUserQuery } from "@/lib/redux/api/bankApi";

// interface FormData {
//     password: string;
//     newPassword: string;
// }

// const Security = () => {
//     const { register, handleSubmit, formState: { errors } } = useForm<FormData>({
//         defaultValues: {
//           password: '',
//           newPassword: "",
//         }
//       });

//       const session = useSession();
//       const access_token = session.data?.access_token;
    
//       const { isLoading, isError, error, data } = useGetCurrentUserQuery(access_token as string);

//       const pref = data?.data.preference
//       useEffect(() => {
//         if (pref) {
          
//           setTwoFactor(pref.twoFactorAuthentication)
//         }
//       }, [data]);
//       const onSubmit: SubmitHandler<FormData> = (data) => {
//         console.log(data);
//       };
//     const [twoFactor, setTwoFactor] = useState(true)
//   return (
//     <div className="p-6 flex flex-col gap-4">
//         <h2 className="text-[#333B69]">Two-factor Authentication</h2>
//         <div className='flex gap-2 items-center'><ToggleButton enabled={twoFactor} onEnable={() => setTwoFactor(!twoFactor)} /> <p className='text-sm font-normal'>Enable or disable two factor authentication</p></div>
//         <h2 className="mt-4 text-[#333B69]">Change Password</h2>
//         <form onSubmit={handleSubmit(onSubmit)} className=''>
//             <div className='flex flex-col sm:w-2/4 w-full sm:gap-4 '>
//                 <div className=''>
//                     <label htmlFor="password" className="text-custom-light-dark mb-4 text-base font-normal">Current Password</label>
//                     <input
//                     id="password"
//                     type="password"                  
//                     {...register('password')}
//                     className="w-full p-2 border border-custom-light-grey rounded-xl mt-2 text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
//                     />
//                     {errors.password && <p className="text-red-500">{errors.password.message}</p>}
//                 </div>
//                 <div className=''>
//                     <label htmlFor="newPassword" className="text-custom-light-dark mb-4 text-base font-normal">New Password</label>
//                     <input
//                     id="newPassword"
//                     type="password"
//                     {...register('newPassword')}
//                     className="w-full p-2 border border-custom-light-grey rounded-xl mt-2 text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
//                     />
//                     {errors.newPassword && <p className="text-red-500">{errors.newPassword.message}</p>}
//                 </div>
//             </div>
//             <div className="flex justify-end">
//             <button
//               type="submit"
//               className="bg-custom-bright-purple w-full sm:w-1/4 text-white px-4 py-2 hover:shadow-md font-body font-medium text-md rounded-xl mt-6"
//             >
//               Save
//             </button>
//           </div>
//         </form>
//     </div>
//   )
// }

// export default Security


"use client";

import ToggleButton from "./Toggle";
import { SubmitHandler, useForm } from 'react-hook-form';
import { useEffect, useState } from 'react';
import { useSession } from "next-auth/react";
import { useChangePasswordMutation, useGetCurrentUserQuery, useUpdateUserPreferenceMutation } from "@/lib/redux/api/bankApi";
import Alert from "./Alert";
import { CustomSerializedError } from "@/lib/redux/types/CustomSerializedError";

interface PasswordFormData {
    password: string;
    newPassword: string;
}

interface TwoFactorFormData {
    twoFactorAuthentication: boolean;
}

const Security = () => {
    // Password form setup
    const { register: registerPassword, handleSubmit: handleSubmitPassword, formState: { errors: errorsPassword } } = useForm<PasswordFormData>({
        defaultValues: {
            password: '',
            newPassword: '',
        }
    });

    // Two-factor form setup
    const { register: registerTwoFactor, handleSubmit: handleSubmitTwoFactor, formState: { errors: errorsTwoFactor }, setValue, getValues } = useForm<TwoFactorFormData>({
        defaultValues: {
            twoFactorAuthentication: true,
        }
    });

    const session = useSession();
    const access_token = session.data?.access_token;

    const { isLoading, isError, error, data } = useGetCurrentUserQuery(access_token as string);
    const [updateUserPreference, {isLoading: isUpdateLoading, isError: isUpdateError, error: updateError, data: updateData}] = useUpdateUserPreferenceMutation()
    const [changePassword, {isLoading: isChangePasswordLoading, isSuccess: isChangePasswordSuccess, isError: isChangePasswordError, error: changePasswordError, data: changePasswordData}] = useChangePasswordMutation()

    const errorChangePassword = changePasswordError as CustomSerializedError

    const pref = data?.data.preference;
    
    useEffect(() => {
        if (pref) {
            setValue('twoFactorAuthentication', pref.twoFactorAuthentication);
        }
    }, [data, setValue, pref]);

    const onSubmitPassword: SubmitHandler<PasswordFormData> = (data) => {
        console.log("Password Data:", data);
        changePassword({credentials: data, token: access_token as string})

    };

    const onSubmitTwoFactor: SubmitHandler<TwoFactorFormData> = (data) => {
        console.log("Two-Factor Data:", data);
        
    };

    if(isChangePasswordSuccess){
      console.log("suc")
    }
    if(isChangePasswordError){
      console.log("err")
    }
    return (
        <div className="p-6 flex flex-col gap-4">
            <h2 className="text-[#333B69]">Two-factor Authentication</h2>
            <form onSubmit={handleSubmitTwoFactor(onSubmitTwoFactor)} className='flex flex-col gap-2'>
                <div className='flex gap-2 items-center'>
                    <ToggleButton
                        enabled={getValues('twoFactorAuthentication')}
                        onEnable={() => setValue('twoFactorAuthentication', !getValues('twoFactorAuthentication'), { shouldDirty: true })}
                    />
                    <p className='text-sm font-normal'>Enable or disable two-factor authentication</p>
                </div>
                {/* <div className="flex justify-end w-full mt-4">
                    <button
                        type="submit"
                        className="bg-custom-bright-purple w-full sm:w-1/4 text-white px-4 py-2 hover:shadow-md font-body font-medium text-md rounded-xl"
                    >
                        Save
                    </button>
                </div> */}
            </form>

            <h2 className="mt-4 text-[#333B69]">Change Password</h2>
            <form onSubmit={handleSubmitPassword(onSubmitPassword)} className=''>
                <div className='flex flex-col sm:w-2/4 w-full sm:gap-4 '>
                    <div className=''>
                        <label htmlFor="password" className="text-custom-light-dark mb-4 text-base font-normal">Current Password</label>
                        <input
                            id="password"
                            type="password"
                            {...registerPassword('password')}
                            className="w-full p-2 border border-custom-light-grey rounded-xl mt-2 text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
                        />
                        {errorsPassword.password && <p className="text-red-500">{errorsPassword.password.message}</p>}
                    </div>
                    <div className=''>
                        <label htmlFor="newPassword" className="text-custom-light-dark mb-4 text-base font-normal">New Password</label>
                        <input
                            id="newPassword"
                            type="password"
                            {...registerPassword('newPassword')}
                            className="w-full p-2 border border-custom-light-grey rounded-xl mt-2 text-custom-light-purple text-base font-normal focus:outline-none focus:ring-2 focus:ring-custom-bright-purple focus:border-transparent"
                        />
                        {errorsPassword.newPassword && <p className="text-red-500">{errorsPassword.newPassword.message}</p>}
                    </div>
                </div>
                <div className="flex justify-end">
                    <button
                        type="submit"
                        className="bg-custom-bright-purple w-full sm:w-1/4 text-white px-4 py-2 hover:shadow-md font-body font-medium text-md rounded-xl mt-6"
                    >
                        Save
                    </button>
                </div>
            </form>

            {isChangePasswordError && <Alert type="error" message={errorChangePassword.data.message} duration={2000} />}
            {isChangePasswordSuccess && <Alert type="success" message="Password changed successfully!" duration={2000} />}

        </div>
    );
}

export default Security;
