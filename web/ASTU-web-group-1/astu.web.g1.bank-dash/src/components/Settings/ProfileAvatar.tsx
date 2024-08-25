'use client';
import { useAppSelector } from '@/hooks/hoooks';
import Image from 'next/image';

const ProfileAvatar = () => {
  const getData = useAppSelector((state) => state.profile);

  return (
    <div className='mr-10 flex justify-center md:block py-3 h-auto'>
      <label htmlFor='profilePicture'>
        <Image
          alt='Profile Image'
          src={getData.profilePicture || '/assets/default-user.svg'}
          width={170}
          height={170}
          className='rounded-full aspect-square border-2'
        />
      </label>
    </div>
  );
};

export default ProfileAvatar;
