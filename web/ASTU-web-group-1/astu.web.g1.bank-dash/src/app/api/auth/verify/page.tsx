import OTP from '@/components/OTP/OTP';

export default function page() {
  // take parameter from url

  return (
    <div className='flex justify-center min-h-screen items-center'>
      <div className='max-w-md p-5 shadow-lg rounded-xl text-center'>
        <h1 className='mb-8 mt-3 text-indigo-900 text-4xl m-5 text-center font- font-[1000]'>
          Verify Email
        </h1>
        <p className='my-8 text-[14.5px] font-epilogue text-indigo-300 font-[500] text-justify'>{`We've sent a verification code to the email address you provided. To complete the verification process, please enter the code here.`}</p>
        <div className='my-5'>
          <OTP />
        </div>
        <p className='text-xs mb-5'>
          You can request to <span className='text-indigo-900 font-[600]'>Resend code in </span>
          <span className='block text-indigo-900 font-[600]'>0:30</span>
        </p>

        <button
          type='submit'
          className='w-full bg-indigo-900 text-white rounded-3xl py-2 font-epilogue font-[700] mt-5 hover:bg-indigo-800 transition-all duration-500 mb-6'
        >
          Continue
        </button>
      </div>
    </div>
  );
}
