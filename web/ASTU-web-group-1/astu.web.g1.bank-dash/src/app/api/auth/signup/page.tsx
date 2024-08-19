import SignUpForm from '@/components/Formx/SignUpForm';
import Image from 'next/image';

export default function page() {
  return (
    <div className='flex justify-around items-center min-h-screen'>
      <div className='w-[30vw] h-[60vh] relative'>
        <Image
          src='/assets/new-user.png'
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
