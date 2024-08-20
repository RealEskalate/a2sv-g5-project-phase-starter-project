import SignUpForm from '@/components/Form/AuthForm/SignUpForm';
import Image from 'next/image';

export default function page() {
  return (
    <div className='flex justify-around items-center min-h-screen'>
      <div className='md:w-1/2 h-[60vh] relative flex items-center justify-center'>
        <Image
          src='/assets/icons/signup.svg'
          alt='hello'
          width={300}
          height={300}
          className='object-cover'
        />
      </div>
      <div className='h-[100vh] bg-slate-200 md:w-1/2 p-6 flex items-center justify-center '>
        <SignUpForm />
      </div>
    </div>
  );
}
