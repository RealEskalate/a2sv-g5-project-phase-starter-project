"use client"
import { useState } from "react";
import { useRouter } from "next/navigation";


import { useForm, SubmitHandler } from "react-hook-form";
import Link from "next/link";
import { signIn } from "next-auth/react";
import Alert from "./Alert";

interface Preference {
  currency: string;
  sentOrReceiveDigitalCurrency: boolean;
  receiveMerchantOrder: boolean;
  accountRecommendations: boolean;
  timeZone: string;
  twoFactorAuthentication: boolean;
}

interface FormData {
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

}

const Signup = () => {
    const [step, setStep] = useState(1);
    const router = useRouter();
    const [signupStatus, setSignupStatus] = useState('')
    const [isError, setIsError] = useState(false)
    const [isSuccess, setIsSuccess] = useState(false)
  
    const {
      register,
      trigger,
      handleSubmit,
      formState: { errors, isValid },
    } = useForm<FormData>({
        mode: "onChange", 
        shouldUnregister: false, 
    });

    const reset = () => {
      setIsSuccess(false)
      setIsError(false)
    }
  
    
    const onSubmit: SubmitHandler<FormData> = async (data) => {
      console.log(data);
      const result = await signIn("sign-up", {...data, redirect: false, callbackUrl: '/settings'})
      if(result?.error){
        setSignupStatus(result?.error as string)
        setIsError(true)      
      } 
      if(result?.ok){
        setSignupStatus('Successfully registered.')
        setIsSuccess(true)
  
      }

      if (result?.ok) {
        router.push('/dashboard');
      }
      
    };

    const handleNext = () => {
        trigger()
        console.log(errors)
        if(isValid)
            setStep(step + 1)
    }

    if(step === 3){
        console.log('err', errors)
    }
    const validatePassword = (value: string) => {
        if(value.length < 6){
            return "Password must be atleast six characters long"
        }
        const hasLetter = /[a-zA-Z]/.test(value);
        if (!hasLetter) {
            return "Password must include at least one letter";
        }
        
        const hasNumber = /[0-9]/.test(value);
        if (!hasNumber) {
            return "Password must include at least one number";
        }
        
        return true;
        
    }
  
    return (
      <div className="max-w-lg mx-auto bg-white p-8 rounded-lg shadow-md mt-10">
        <div>
          <h1 className="font-bold text-custom-purple text-3xl text-center">Signup to BankDash</h1>
          <button
              className='flex gap-4 justify-center items-center mb-8 border font-body font-bold text-xl h-12 w-full border-custom-light-grey text-custom-light-purple rounded-md mt-6 p-4 hover:shadow-md hover:shadow-gray-100'
              // onClick={() => signIn('google', { callbackUrl: '/landing' })}
          >
              <span>  
                  <svg width="21" height="20" viewBox="0 0 21 20" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M18.6712 8.36788H18V8.33329H10.5V11.6666H15.2096C14.5225 13.607 12.6762 15 10.5 15C7.73874 15 5.49999 12.7612 5.49999 9.99996C5.49999 7.23871 7.73874 4.99996 10.5 4.99996C11.7746 4.99996 12.9342 5.48079 13.8171 6.26621L16.1742 3.90913C14.6858 2.52204 12.695 1.66663 10.5 1.66663C5.89791 1.66663 2.16666 5.39788 2.16666 9.99996C2.16666 14.602 5.89791 18.3333 10.5 18.3333C15.1021 18.3333 18.8333 14.602 18.8333 9.99996C18.8333 9.44121 18.7758 8.89579 18.6712 8.36788Z" fill="#FFC107"/>
                      <path d="M3.12749 6.12121L5.8654 8.12913C6.60624 6.29496 8.4004 4.99996 10.5 4.99996C11.7746 4.99996 12.9342 5.48079 13.8171 6.26621L16.1742 3.90913C14.6858 2.52204 12.695 1.66663 10.5 1.66663C7.29915 1.66663 4.52332 3.47371 3.12749 6.12121Z" fill="#FF3D00"/>
                      <path d="M10.5 18.3333C12.6525 18.3333 14.6083 17.5095 16.0871 16.17L13.5079 13.9875C12.6432 14.6451 11.5865 15.0008 10.5 15C8.33251 15 6.49209 13.6179 5.79876 11.6891L3.08126 13.7829C4.46043 16.4816 7.26126 18.3333 10.5 18.3333Z" fill="#4CAF50"/>
                      <path d="M18.6713 8.36796H18V8.33337H10.5V11.6667H15.2096C14.8809 12.5902 14.2889 13.3972 13.5067 13.988L13.5079 13.9871L16.0871 16.1696C15.9046 16.3355 18.8333 14.1667 18.8333 10C18.8333 9.44129 18.7758 8.89587 18.6713 8.36796Z" fill="#1976D2"/>
                  </svg>
              </span>
              Sign Up with Google
          </button>
  
        </div>
        <h2 className="text-2xl text-gray-400 font-semibold mb-6">{`Step ${step} / 3`}</h2>
        <form onSubmit={handleSubmit(onSubmit)} noValidate>
          {step === 1 && (
            <>
              <div className="mb-4">
                <label htmlFor="name" className="block text-gray-700">
                  Name
                </label>
                <input
                  id="name"
                  type="text"
                  {...register("name", { required: "Name is required" })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.name && (
                  <div className="text-red-600 mt-1">{errors.name.message}</div>
                )}
              </div>
  
              <div className="mb-4">
                <label htmlFor="email" className="block text-gray-700">
                  Email
                </label>
                <input
                  id="email"
                  type="email"
                  {...register("email", {
                    required: "Email is required",
                    pattern: {
                      value: /^\S+@\S+$/i,
                      message: "Enter a valid email",
                    },
                  })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.email && (
                  <div className="text-red-600 mt-1">{errors.email.message}</div>
                )}
              </div>

              <div className="mb-4">
                <label htmlFor="username" className="block text-gray-700">
                  Username
                </label>
                <input
                  id="username"
                  type="text"
                  {...register("username", {
                    required: "Username is required",
                  })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.username && (
                  <div className="text-red-600 mt-1">
                    {errors.username.message}
                  </div>
                )}
              </div>
  
             
            </>
          )}
  
          {step === 2 && (
            <>
            <div className="mb-4">
                <label htmlFor="password" className="block text-gray-700">
                  Password
                </label>
                <input
                  type="password"
                  id='password'
                  {...register('password', { 
                      required: {
                          value: true,
                          message: 'Password is required'
                      },
                      validate: (value) => validatePassword(value)
                  })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.password && (
                  <div className="text-red-600 mt-1">
                    {errors.password.message}
                  </div>
                )}
              </div>

              <div className="mb-4">
                <label
                  htmlFor="permanentAddress"
                  className="block text-gray-700"
                >
                  Permanent Address
                </label>
                <input
                  id="permanentAddress"
                  type="text"
                  {...register("permanentAddress", {
                    required: "Permanent Address is required",
                  })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.permanentAddress && (
                  <div className="text-red-600 mt-1">
                    {errors.permanentAddress.message}
                  </div>
                )}
              </div>
  
              <div className="mb-4">
                <label htmlFor="postalCode" className="block text-gray-700">
                  Postal Code
                </label>
                <input
                  id="postalCode"
                  type="text"
                  {...register("postalCode", {
                    required: "Postal Code is required",
                  })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.postalCode && (
                  <div className="text-red-600 mt-1">
                    {errors.postalCode.message}
                  </div>
                )}
              </div>

              
  
              
            </>
          )}
  
          {step === 3 && (
            <>
                <div className="mb-4">
                    <label htmlFor="dateOfBirth" className="block text-gray-700">
                    Date of Birth
                    </label>
                    <input
                    id="dateOfBirth"
                    type="date"
                    {...register("dateOfBirth", {
                        required: "Date of Birth is required",
                    })}
                    className="w-full p-2 border border-gray-300 rounded"
                    />
                    {errors.dateOfBirth && (
                    <div className="text-red-600 mt-1">
                        {errors.dateOfBirth.message}
                    </div>
                    )}
              </div>
              <div className="mb-4">
                <label htmlFor="presentAddress" className="block text-gray-700">
                  Present Address
                </label>
                <input
                  id="presentAddress"
                  type="text"
                  {...register("presentAddress", {
                    required: "Present Address is required",
                  })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.presentAddress && (
                  <div className="text-red-600 mt-1">
                    {errors.presentAddress.message}
                  </div>
                )}
              </div>
  
              <div className="mb-4">
                <label htmlFor="city" className="block text-gray-700">
                  City
                </label>
                <input
                  id="city"
                  type="text"
                  {...register("city", { required: "City is required" })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.city && (
                  <div className="text-red-600 mt-1">{errors.city.message}</div>
                )}
              </div>
  
              <div className="mb-4">
                <label htmlFor="country" className="block text-gray-700">
                  Country
                </label>
                <input
                  id="country"
                  type="text"
                  {...register("country", { required: "Country is required" })}
                  className="w-full p-2 border border-gray-300 rounded"
                />
                {errors.country && (
                  <div className="text-red-600 mt-1">
                    {errors.country.message}
                  </div>
                )}
              </div>
            </>
          )}
  
          <div className="flex justify-between mt-6">
            {step > 1 && (
              <button
                type="button"
                onClick={() => setStep(step - 1)}
                className="bg-gray-500 text-white px-4 py-2 rounded"
              >
                Previous
              </button>
            )}
            {step < 3 ? (
              <button
                type="button"
                onClick={handleNext}
                className="bg-blue-500 text-white px-4 py-2 rounded"
              >
                Next
              </button>
            ) : (
              <button
                type="submit"
                className="bg-green-500 text-white px-4 py-2 rounded"
              >
                Sign Up
              </button>
            )}
          </div>
        </form>
        <p className='text-gray-600 font-normal text-base mt-6 mb-6'>Already have an account? <Link className='text-custom-purple font-semibold text-base' href='/auth/signin'>Login</Link></p>
        <p className='text-gray-600 font-normal text-sm'>By clicking 'Signup', you acknowledge that you have read and accepted our terms of <Link className='text-purple-tag'  href={'/terms-of-service'}>Terms of Service</Link> and <Link className='text-purple-tag '  href={'/privacy-policy'}>Privacy Policy</Link></p>
        {isSuccess && <Alert type='success' message={signupStatus} duration={2000} onClose={reset} />}
        {isError && <Alert type='error' message={signupStatus} duration={2000} onClose={reset} />}

      </div>
    );
}

export default Signup
