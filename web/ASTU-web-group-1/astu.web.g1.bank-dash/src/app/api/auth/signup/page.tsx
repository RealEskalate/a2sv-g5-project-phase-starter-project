import SignUpForm from '@/components/Form/AuthForm/SignUpForm';
import Image from 'next/image';

export default function page() {
  return (
    <div className='flex justify-around items-center min-h-screen'>
      <div className='md:w-1/2 h-[60vh] relative'>
        <Image
          src='/assets/images/apple.png'
          alt='hello'
          layout='fill'
          objectFit='cover'
          sizes='100vw'
          className='object-cover'
        />
      </div>
      <SignUpForm />
    </div>
  );
}
