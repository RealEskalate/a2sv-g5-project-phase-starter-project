'use client';
import React, { useState } from 'react';
import { useForm, FormProvider } from 'react-hook-form';
import Step1 from './Step1';
import Step2 from './Step2';
import Step3 from './Step3';
import { useSignUpMutation } from '@/lib/redux/api/authApi';

type FormData = {
    name: string;
    email: string;
    dateOfBirth: string;
    permanentAddress: string;
    postalCode: string;
    presentAddress: string;
    city: string;
    country: string;
    profilePicture: string;
    username: string;
    password: string;
    preference: {
      currency: string;
      sentOrReceiveDigitalCurrency: boolean;
      receiveMerchantOrder: boolean;
      accountRecommendations: boolean;
      timeZone:string;
      twoFactorAuthentication: boolean;
    };
};

const SignUpForm: React.FC = () => {
  const methods = useForm<FormData>({
    defaultValues: {
      name: '',
      email: '',
      dateOfBirth: '',
      permanentAddress: '',
      postalCode: '',
      presentAddress: '',
      city: '',
      country: '',
      profilePicture:"/",
      username: '',
      password: '',
      preference: {
        currency: '',
        sentOrReceiveDigitalCurrency: false,
        receiveMerchantOrder: false,
        accountRecommendations: false,
        timeZone: 'GMT + 3',
        twoFactorAuthentication: false,
      },
    },
  });

  
  const [step, setStep] = useState(1);
  const [putSignUp ] = useSignUpMutation()
  
  const onSubmit = async (data: FormData) => {
    console.log(data);
    try {
      await putSignUp(data ).unwrap();
      window.location.reload()

    } catch (err) {
      console.error(err);
    }

  };

  const nextStep = () => {
    methods.trigger().then((valid) => {
      if (valid) {
        setStep((prev) => prev + 1);
      }
    });
  };

  const prevStep = () => setStep((prev) => prev - 1);

  const renderStep = () => {
    switch (step) {
      case 1:
        return <Step1 step={step} />;
      case 2:
        return <Step2 step={step} />;
      case 3:
        return <Step3 step={step} />;
      default:
        return <Step1 step={step} />;
    }
  };

  return (
    <FormProvider {...methods}>
      <form onSubmit={methods.handleSubmit(onSubmit)} className="max-w-lg mx-auto p-8 border rounded-md shadow-md h-5/6">
        <div className="flex flex-col justify-between h-full">
          <div className="flex-1">{renderStep()}</div>
          <div className="flex items-center justify-between mt-8">
            {step > 1 && (
              <button
                type="button"
                onClick={prevStep}
                className="bg-gray-300 text-gray-700 w-10 h-10 flex items-center justify-center rounded-full"
              >
                ←
              </button>
            )}
            <div className="flex-1 text-center">
              <p className="text-sm font-semibold">
                Step {step} of 3
              </p>
            </div>
            {step < 3 ? (
              <button
                type="button"
                onClick={nextStep}
                className="bg-blue-500 text-white w-10 h-10 flex items-center justify-center rounded-full"
              >
                →
              </button>
            ) : (
              <button
                type="submit"
                className="bg-blue-500 text-white w-10 h-10 flex items-center justify-center rounded-full"
              >
                ✓
              </button>
            )}
          </div>
        </div>
      </form>
    </FormProvider>
  );
};

export default SignUpForm;