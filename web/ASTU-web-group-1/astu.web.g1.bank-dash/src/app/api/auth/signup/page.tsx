import SignUpForm from '@/components/Form/AuthForm/SignUpForm';
import Image from 'next/image';

export default function page() {
  return (
    <div className='flex justify-around items-center min-h-screen'>
      <div className='hidden md:w-1/2 minrelative lg:flex items-center justify-center'>
        <Image
          src='/assets/icons/signup.svg'
          alt='hello'
          width={300}
          height={300}
          className='object-cover'
        />
      </div>
      <div className='min-h-[100vh] lg:bg-slate-200 w-full lg:w-1/2 lg:p-6 flex items-center justify-center '>
      <SignUpForm />
      </div>
    </div>
  );
}
