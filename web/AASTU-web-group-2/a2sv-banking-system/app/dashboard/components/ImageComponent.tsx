import React from 'react';
import Image from 'next/image';

interface ImageComponentProps {
  src: string; // URL of the image
  alt: string; // Alt text for the image
  width: number; // Width of the image
  height: number; // Height of the image
  name: string; // Name to display
  role: string; // Role to display
}

const ImageComponent: React.FC<ImageComponentProps> = ({ src, alt, width, height, name, role }) => {
  return (
    <div className='flex'>
      <div className='flex flex-col gap-4'>
        <div className='flex justify-center'>
          <Image
            width={width}
            height={height}
            className='rounded-full'
            alt={alt}
            src={src}
          />
        </div>
        <div className='flex flex-col gap-2'>
          <p className='text-xs text-center text-nowrap'>{name}</p>
          <p className='text-xs text-[#718EBF] text-center'>{role}</p>
        </div>
      </div>
    </div>
  );
};

export default ImageComponent;
